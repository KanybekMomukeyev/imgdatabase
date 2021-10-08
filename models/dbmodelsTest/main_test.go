package dbmodelsTest

import (
	"os"
	"testing"

	models "github.com/KanybekMomukeyev/imgdatabase/v3/models"

	log "github.com/sirupsen/logrus"
)

var DbMng *models.DbManager

func init() {
	var providerLogs = log.New()
	providerLogs.SetFormatter(&log.JSONFormatter{})
	providerLogs.SetOutput(os.Stdout)
	providerLogs.SetLevel(log.InfoLevel)

	log.Printf("\nSTART MIGRATIONS\n")
	models.MigrateDatabaseDown(models.TestPath, "file://../migrations/")
	models.MigrateDatabaseUp(models.TestPath, "file://../migrations/")
	log.Printf("\nMIGRATIONS FINISHED\n")

	DbMng = models.NewDbManager(models.TestPath, providerLogs)
}

func TestMain(m *testing.M) {
	log.Printf("\n START TEST -V This gets run BEFORE any tests get run!\n")
	exitVal := m.Run()
	log.Printf("\n END TEST -V This gets run AFTER any tests get run!\n")
	os.Exit(exitVal)
}

func TestOne(t *testing.T) {
	log.Println("TestOne running")
}

func TestTwo(t *testing.T) {
	log.Println("TestTwo running")
}

// go test -v
