package dbmodelsTest

import (
	"os"
	"testing"

	models "github.com/KanybekMomukeyev/imgdatabase/models"

	log "github.com/sirupsen/logrus"
)

var DbMng *models.DbManager

func init() {
	var providerLogs = log.New()
	providerLogs.SetFormatter(&log.JSONFormatter{})
	providerLogs.SetOutput(os.Stdout)
	providerLogs.SetLevel(log.InfoLevel)

	log.Println("START MIGRATIONS")
	models.MigrateDatabaseDown(models.TestPath, "file://../migrations/")
	models.MigrateDatabaseUp(models.TestPath, "file://../migrations/")
	log.Println("MIGRATIONS FINISHED")
	DbMng = models.NewDbManager(models.TestPath, providerLogs)
}

func TestMain(m *testing.M) {
	log.Println("This gets run BEFORE any tests get run!")
	exitVal := m.Run()
	log.Println("This gets run AFTER any tests get run!")
	os.Exit(exitVal)
}

func TestOne(t *testing.T) {
	log.Println("TestOne running")
}

func TestTwo(t *testing.T) {
	log.Println("TestTwo running")
}

// go test -v
