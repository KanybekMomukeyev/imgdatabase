package models

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	// DefaultWritePermissions la
	DefaultWritePermissions = 0760

	// MaximumNewDirectoryAttempts la
	MaximumNewDirectoryAttempts = 1000
)

var (
	fileInfo os.FileInfo
	err      error
)

// ParseText la lala
func ParseText(path string) {
	//b, err := ioutil.ReadFile("./xlsparser/file.txt") // just pass the file name
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(b)   // print the content as 'bytes'
	str := string(b) // convert content to a 'string'
	fmt.Println(str) // print the content as a 'string'
}

// GetFileInfo la lala
func GetFileInfo(path string) {
	fileInfo, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

// FileExists credit
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// IsEmpty credit
func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

// CreateUniqueDirectory creates a new directory but if the combination of dir and name exists
// then append a number until a unique name is found
func CreateUniqueDirectory(dir string, name string, maximumAttempts int) (string, error) {
	for i := 0; i < maximumAttempts; i++ {
		n := name
		if i > 0 {
			n += strconv.Itoa(i)
		}
		p := filepath.Join(dir, n)
		exists, err := FileExists(p)
		if err != nil {
			return p, err
		}
		if !exists {
			err := os.MkdirAll(p, DefaultWritePermissions)
			if err != nil {
				return "", fmt.Errorf("Failed to create directory %s due to %s", p, err)
			}
			return p, nil
		}
	}
	return "", fmt.Errorf("Could not create a unique file in %s starting with %s after %d attempts", dir, name, maximumAttempts)
}

// RenameDir credit
func RenameDir(src string, dst string, force bool) (err error) {
	err = CopyDir(src, dst, force)
	if err != nil {
		return fmt.Errorf("failed to copy source dir %s to %s: %s", src, dst, err)
	}
	err = os.RemoveAll(src)
	if err != nil {
		return fmt.Errorf("failed to cleanup source dir %s: %s", src, err)
	}
	return nil
}

// RenameFile credit
func RenameFile(src string, dst string) (err error) {
	err = CopyFile(src, dst)
	if err != nil {
		return fmt.Errorf("failed to copy source file %s to %s: %s", src, dst, err)
	}
	err = os.RemoveAll(src)
	if err != nil {
		return fmt.Errorf("failed to cleanup source file %s: %s", src, err)
	}
	return nil
}

// CopyDir credit https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
func CopyDir(src string, dst string, force bool) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		if force {
			os.RemoveAll(dst)
		} else {
			return fmt.Errorf("destination already exists")
		}
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath, force)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}

// CopyFile credit https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

// LoadBytes loads a file
func LoadBytes(dir, name string) ([]byte, error) {
	path := filepath.Join(dir, name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error loading file %s in directory %s, %v", name, dir, err)
	}
	return bytes, nil
}

//mainnot NOT USED HERE ----------------------------------------------------------------------
func mainnot() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("received", msg1)
				return // return will exit from function/block
			case msg2 := <-c2:
				fmt.Println("received", msg2)
				// default:
				// 	// not waiting channels
				// 	fmt.Println("default")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("FINISHED")
}

// {"level":"info","msg":"total: 1728","time":"2018-11-03T04:17:58Z"}
// {"level":"info","msg":"megacomUsers: 419","time":"2018-11-03T04:17:58Z"}
// {"level":"info","msg":"beelineUsers: 398","time":"2018-11-03T04:17:58Z"}
// {"level":"info","msg":"nurtelecomUsers: 911","time":"2018-11-03T04:17:58Z"}
