package models

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	dbmodels "github.com/KanybekMomukeyev/imgdatabase/v3/models/dbmodels"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// ALTER USER username WITH ENCRYPTED PASSWORD 'password';
// CREATE USER kanybek WITH PASSWORD 'KanybekNazgul1984';
// CREATE DATABASE databasename;
// GRANT ALL PRIVILEGES ON DATABASE databasename TO  kanybek;

const (
	RDSPath         = "user=kanybek dbname=databasename password=nazgulum host=databasename.c9s8gyz9e38h.us-west-2.rds.amazonaws.com port=5432 sslmode=require"
	DevelopmentPath = "user=newuser dbname=postgres password=password host=localhost port=5432 sslmode=disable"
	TestPath        = "user=kanybek dbname=databasename password=KanybekNazgul1984 host=localhost port=5432 sslmode=disable"
)

type DatabaseInterface interface {
	followChannel()
}

type DbManager struct {
	DB              *sqlx.DB
	dbChannel       chan func()
	responseChannel chan uint64
	errorChannel    chan error
	providerLogs    *log.Logger
}

func NewDbManager(path string, logger *log.Logger) *DbManager {

	db, err := sqlx.Connect("postgres", path)
	if err != nil {
		log.Fatalln(err)
	}

	dbM := new(DbManager)
	dbM.DB = db
	dbM.providerLogs = logger
	dbM.dbChannel = make(chan func(), 1000)
	dbM.responseChannel = make(chan uint64, 1000)
	dbM.errorChannel = make(chan error, 1000)
	dbM.followChannel()
	return dbM
}

func (dbM *DbManager) StopDBChannel() {

}

//----------------------------- PRIVATE METHODS ------------------------------------------
func (dbM *DbManager) followChannel() {
	go func() {
		for f := range dbM.dbChannel {
			f()
		}
	}()
}

func (dbM *DbManager) sendLastInsertedId(lastInsertId uint64) {
	go func() {
		dbM.responseChannel <- lastInsertId
	}()
}

func (dbM *DbManager) sendError(err error) {
	if err != nil {
		dbM.providerLogs.WithFields(log.Fields{"DbManager error": err}).Warn("SendError")
		go func() {
			dbM.errorChannel <- err
		}()
	}
}

func (dbM *DbManager) waitForLastInsertedId() (uint64, error) {
	select {
	case lastInsertedId := <-dbM.responseChannel:
		return lastInsertedId, nil
	case err := <-dbM.errorChannel:
		return 0, err
	}
}

func (dbM *DbManager) CreateUser(user *dbmodels.UserRequest) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		lastInsertId, err := dbmodels.StoreUser(tx, user)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(lastInsertId)
	}

	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}
func (dbM *DbManager) UpdateUser(user *dbmodels.UserRequest) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		rowsAffected, err := dbmodels.UpdateUser(tx, user)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(rowsAffected)
	}

	dbM.dbChannel <- f

	return dbM.waitForLastInsertedId()
}

func (dbM *DbManager) AllUsers(companyId uint64) ([]*dbmodels.UserRequest, error) {
	users, err := dbmodels.AllUsersForCompany(dbM.DB, companyId)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dbM *DbManager) AllUsersForUpdate(filter *dbmodels.UserFilter, companyId uint64) ([]*dbmodels.UserRequest, error) {
	users, err := dbmodels.AllUsersForUpdate(dbM.DB, filter, companyId)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dbM *DbManager) UserForEmail(email string) (*dbmodels.UserRequest, error) {
	user, err := dbmodels.UserForEmail(dbM.DB, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dbM *DbManager) UserForUUID(uuid string) (*dbmodels.UserRequest, error) {
	user, err := dbmodels.UserForUUID(dbM.DB, uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dbM *DbManager) UserCompanyIdForUserUUID(uuid string) (uint64, error) {
	user, err := dbmodels.UserForUUID(dbM.DB, uuid)
	if err != nil {
		return 0, err
	}
	return user.CompanyId, nil
}

//---------------------------------- Customer -------------------------------------------
func (dbM *DbManager) CreateCuttedImage(cutted *dbmodels.CuttedImage, companyId uint64) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		lastInsertId, err := dbmodels.StoreCuttedImage(tx, cutted, companyId)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(lastInsertId)
	}
	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}

func (dbM *DbManager) UpdateCuttedImage(cutted *dbmodels.CuttedImage, companyId uint64) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		rowsAffected, err := dbmodels.UpdateCuttedImage(tx, cutted, companyId)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(rowsAffected)
	}
	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}

func (dbM *DbManager) AllCuttedImagesForCompany(companyId uint64) ([]*dbmodels.CuttedImage, error) {
	customers, err := dbmodels.AllCuttedImagesForCompany(dbM.DB, companyId)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (dbM *DbManager) StatOfMarkedCuttedImagesForTelegramId(companyId uint64, cuttedImageType uint32) ([]*dbmodels.CuttedImage, error) {
	customers, err := dbmodels.StatOfMarkedCuttedImagesForTelegramId(dbM.DB, companyId, cuttedImageType)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (dbM *DbManager) CuttedImageForID(ImageID uint64) (*dbmodels.CuttedImage, error) {
	cuttedImage, err := dbmodels.CuttedImageForID(dbM.DB, ImageID)
	if err != nil {
		return nil, err
	}
	return cuttedImage, nil
}

// STATISTICS FOR DOCUMENTS
func (dbM *DbManager) TotalCuttedImagesFor(docModelID uint64) (int, error) {
	translateCount, err := dbmodels.CountCuttedImagesForDoc(dbM.DB, docModelID)
	if err != nil {
		return 0, err
	}
	return translateCount, nil
}

func (dbM *DbManager) Cutted_image_not_translated_FOR(docModelID uint64) (int, error) {
	translateCount, err := dbmodels.StatsCuttedImageForDoc(dbM.DB, docModelID, dbmodels.CUTTED_IMAGE_NOT_TRANSLATED)
	if err != nil {
		return 0, err
	}
	return translateCount, nil
}

func (dbM *DbManager) Cutted_image_translated_FOR(docModelID uint64) (int, error) {
	translateCount, err := dbmodels.StatsCuttedImageForDoc(dbM.DB, docModelID, dbmodels.CUTTED_IMAGE_TRANSLATED)
	if err != nil {
		return 0, err
	}
	return translateCount, nil
}

func (dbM *DbManager) Cutted_image_misstaked_FOR(docModelID uint64) (int, error) {
	translateCount, err := dbmodels.StatsCuttedImageForDoc(dbM.DB, docModelID, dbmodels.CUTTED_IMAGE_MISSTAKED)
	if err != nil {
		return 0, err
	}
	return translateCount, nil
}

func (dbM *DbManager) Cutted_image_marked_as_invalid_FOR(docModelID uint64) (int, error) {
	translateCount, err := dbmodels.StatsCuttedImageForDoc(dbM.DB, docModelID, dbmodels.CUTTED_IMAGE_MARKED_INVALID)
	if err != nil {
		return 0, err
	}
	return translateCount, nil
}

// ------------------------------------------------------------ //

func (dbM *DbManager) CuttedImagesForFolder(folderID uint64) ([]*dbmodels.CuttedImage, error) {
	cuttedImages, err := dbmodels.CuttedImagesForFolder(dbM.DB, folderID)
	if err != nil {
		return nil, err
	}
	return cuttedImages, nil
}

func (dbM *DbManager) CuttedImagesForType(cuttedImageType uint32) ([]*dbmodels.CuttedImage, error) {
	cuttedImages, err := dbmodels.CuttedImagesForType(dbM.DB, cuttedImageType)
	if err != nil {
		return nil, err
	}
	return cuttedImages, nil
}

func (dbM *DbManager) RandomNotTranslatedCuttedImageNOT_IN_ARRAY(imageIDs []uint64) (*dbmodels.CuttedImage, error) {
	randomCuttedImage, err := dbmodels.RandomNotTranslatedCuttedImageNOT_IN_ARRAY(dbM.DB, imageIDs)
	if err != nil {
		return nil, err
	}
	return randomCuttedImage, nil
}

func (dbM *DbManager) RandomNotTranslatedCuttedImageNotInArray_TYPE_IS_HANDWRITTEN(imageIDs []uint64) (*dbmodels.CuttedImage, error) {
	randomCuttedImage, err := dbmodels.RandomNotTranslatedCuttedImageNotInArray_TYPE_IS_HANDWRITTEN(dbM.DB, imageIDs)
	if err != nil {
		return nil, err
	}
	return randomCuttedImage, nil
}

func (dbM *DbManager) RandomUnknownType() (*dbmodels.CuttedImage, error) {
	randomCuttedImage, err := dbmodels.RandomUnknownType(dbM.DB)
	if err != nil {
		return nil, err
	}
	return randomCuttedImage, nil
}

func (dbM *DbManager) AllUpdatedCuttedImages(custFilter *dbmodels.ImageFilter, companyId uint64) ([]*dbmodels.CuttedImage, error) {
	customers, err := dbmodels.AllUpdatedCuttedImages(dbM.DB, custFilter, companyId)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

//---------------------------------- FOLDERS -------------------------------------------
func (dbM *DbManager) CreateFolderModel(folder *dbmodels.FolderModel, companyId uint64) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		lastInsertId, err := dbmodels.StoreFolderModel(tx, folder, companyId)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(lastInsertId)
	}
	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}

func (dbM *DbManager) UpdateFolderModel(folder *dbmodels.FolderModel, companyId uint64) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		rowsAffected, err := dbmodels.UpdateFolderModel(tx, folder, companyId)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(rowsAffected)
	}
	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}

func (dbM *DbManager) AllFolderModelsForInitial(companyId uint64) ([]*dbmodels.FolderModel, error) {
	customers, err := dbmodels.AllFoldersForCompany(dbM.DB, companyId)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (dbM *DbManager) AllFoldersForDoc(docModelID uint64) ([]*dbmodels.FolderModel, error) {
	customers, err := dbmodels.AllFoldersForDoc(dbM.DB, docModelID)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (dbM *DbManager) FolderForFolderId(folderModelID uint64) (*dbmodels.FolderModel, error) {
	customers, err := dbmodels.FolderForFolderId(dbM.DB, folderModelID)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (dbM *DbManager) AllUpdatedSuppliers(suppFilter *dbmodels.FolderFilter, companyId uint64) ([]*dbmodels.FolderModel, error) {
	customers, err := dbmodels.AllFoldersForUpdate(dbM.DB, suppFilter, companyId)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

//---------------------------------- Cutted_Image_Translates -------------------------------------------
func (dbM *DbManager) TranslatesForSearchKeyword(searchKey string) ([]*dbmodels.CuttedImageTranslate, error) {
	translates, err := dbmodels.TranslatesForSearchKeyword(dbM.DB, searchKey)
	if err != nil {
		return nil, err
	}
	return translates, nil
}

func (dbM *DbManager) AutoMigrateTranslates() (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()
		translates, _ := dbmodels.AllTranslates(dbM.DB)
		for _, translate := range translates {
			_, err := dbmodels.UpdateTranslateFTS(tx, translate)
			if err != nil {
				tx.Rollback()
				dbM.sendError(err)
				return
			}
		}
		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}
		dbM.sendLastInsertedId(uint64(1000))
	}
	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}

func (dbM *DbManager) CreateCuttedImageTranslate(cuttedImageTranslate *dbmodels.CuttedImageTranslate) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		lastInsertId, err := dbmodels.StoreCuttedImageTranslate(tx, cuttedImageTranslate)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(lastInsertId)
	}
	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}

// STATISTICS FOR TELEGRAM USER
func (dbM *DbManager) Translated_STATUSES_forTgUser(telegramUserID uint64, translateStatus dbmodels.AcceptStatus) ([]*dbmodels.CuttedImageTranslate, error) {
	translates, err := dbmodels.Translated_STATUSES_forTgUser(dbM.DB, telegramUserID, translateStatus)
	if err != nil {
		return nil, err
	}
	return translates, nil
}

func (dbM *DbManager) LastActivityForTgUser(telegramUserID uint64) (*dbmodels.CuttedImageTranslate, error) {
	lastTranslate, err := dbmodels.LastActivityForTgUser(dbM.DB, telegramUserID)
	if err != nil {
		return nil, err
	}
	return lastTranslate, nil
}

func (dbM *DbManager) Count_WAITING_TranslatesForCuttedImage(cuttedImageID uint64) (int, error) {

	translateCount, err := dbmodels.Count_WAITING_TranslatesForCuttedImage(dbM.DB, cuttedImageID, dbmodels.TRANSLATE_WAITING)
	if err != nil {
		return 0, err
	}
	return translateCount, nil
}

func (dbM *DbManager) CountFor_WAITING_TranslatesFotTgUser(telegramUserID uint64) (int, error) {

	translateCount, err := dbmodels.CountFor_Translated_STATUSES_FotTgUser(dbM.DB, telegramUserID, dbmodels.TRANSLATE_WAITING)
	if err != nil {
		return 0, err
	}
	return translateCount, nil
}

func (dbM *DbManager) CountFor_ACCEPTED_TranslatesFotTgUser(telegramUserID uint64) (int, error) {
	translateCount, err := dbmodels.CountFor_Translated_STATUSES_FotTgUser(dbM.DB, telegramUserID, dbmodels.TRANSLATE_ACCEPTED)
	if err != nil {
		return 0, err
	}
	return translateCount, nil
}

func (dbM *DbManager) CountFor_REJECTED_TranslatesFotTgUser(telegramUserID uint64) (int, error) {
	translateCount, err := dbmodels.CountFor_Translated_STATUSES_FotTgUser(dbM.DB, telegramUserID, dbmodels.TRANSLATE_REJECTED)
	if err != nil {
		return 0, err
	}
	return translateCount, nil
}

func (dbM *DbManager) CountFor_USER_MARKED_INVALID_TranslatesFotTgUser(telegramUserID uint64) (int, error) {
	translateCount, err := dbmodels.CountFor_Translated_STATUSES_FotTgUser(dbM.DB, telegramUserID, dbmodels.TRANSLATE_MARKED_AS_INVALID)
	if err != nil {
		return 0, err
	}
	return translateCount, nil
}

func (dbM *DbManager) UpdateCuttedImageTranslate(cuttedImageTranslate *dbmodels.CuttedImageTranslate) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		rowsAffected, err := dbmodels.UpdateCuttedImageTranslate(tx, cuttedImageTranslate)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(rowsAffected)
	}
	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}

func (dbM *DbManager) CuttedImageIDsForTelegramUser(telegramUserID uint64) ([]*dbmodels.CuttedImageTranslate, error) {
	imageIDs, err := dbmodels.CuttedImageIDsForTelegramUser(dbM.DB, telegramUserID)
	if err != nil {
		return nil, err
	}
	return imageIDs, nil
}

func (dbM *DbManager) TranslatesForDocModelAndAcceptedStatus(docModelID uint64, translateStatus dbmodels.AcceptStatus) ([]*dbmodels.CustomTranslate, error) {
	custTranslates, err := dbmodels.TranslatesForDocModelAndAcceptedStatus(dbM.DB, docModelID, translateStatus)
	if err != nil {
		return nil, err
	}
	return custTranslates, nil
}

func (dbM *DbManager) TranslatesForDocModelAndTelegramUser(docModelID uint64, telegramUserID uint64) ([]*dbmodels.CustomTranslate, error) {
	custTranslates, err := dbmodels.TranslatesForDocModelAndTelegramUser(dbM.DB, docModelID, telegramUserID)
	if err != nil {
		return nil, err
	}
	return custTranslates, nil
}

func (dbM *DbManager) TelegramUserIdsDocModel(docModelID uint64) ([]*dbmodels.TelegramUserIDsSelectResponse, error) {
	telegramIDs, err := dbmodels.TelegramUserIdsDocModel(dbM.DB, docModelID)
	if err != nil {
		return nil, err
	}
	return telegramIDs, nil
}

func (dbM *DbManager) DocIdsForTelegramUser(telegramUserID uint64) ([]*dbmodels.DocIDsSelectResponse, error) {
	imageIDs, err := dbmodels.DocIdsForTelegramUser(dbM.DB, telegramUserID)
	if err != nil {
		return nil, err
	}
	return imageIDs, nil
}

func (dbM *DbManager) FolderIdsForCuttedImageTranslate(cuttedImageTranslateIDs []uint64) ([]uint64, error) {
	imageIDs, err := dbmodels.FolderIdsForCuttedImageTranslate(dbM.DB, cuttedImageTranslateIDs)
	if err != nil {
		return nil, err
	}
	return imageIDs, nil
}

func (dbM *DbManager) CuttedImageTranslatesFor(cuttedImageID uint64) ([]*dbmodels.CuttedImageTranslate, error) {
	customers, err := dbmodels.CuttedImageTranslatesFor(dbM.DB, cuttedImageID)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (dbM *DbManager) AllUpdatedCuttedImageTranslates(cutImgFilter *dbmodels.CuttedImageTranslateFilter, cuttedImageID uint64) ([]*dbmodels.CuttedImageTranslate, error) {
	customers, err := dbmodels.AllUpdatedCuttedImageTranslatesForImage(dbM.DB, cutImgFilter, cuttedImageID)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

//---------------------------------- DocModel -------------------------------------------
func (dbM *DbManager) CreateDocModel(doc *dbmodels.DocModel, companyId uint64) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		lastInsertId, err := dbmodels.StoreDocModel(tx, doc, companyId)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(lastInsertId)
	}
	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}

func (dbM *DbManager) UpdateDocModel(doc *dbmodels.DocModel, companyId uint64) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		rowsAffected, err := dbmodels.DocModelUpdate(tx, doc, companyId)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(rowsAffected)
	}
	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}

func (dbM *DbManager) AllDocModelsForCompany(companyId uint64) ([]*dbmodels.DocModel, error) {
	customers, err := dbmodels.AllDocModelsForCompany(dbM.DB, companyId)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (dbM *DbManager) DocForDocModelID(docModelID uint64) (*dbmodels.DocModel, error) {
	customers, err := dbmodels.DocForDocModelID(dbM.DB, docModelID)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

//---------------------------------- TelegramUserModel -------------------------------------------
func (dbM *DbManager) CreateSampleTelegramUsersIfNotExsist() {
	// 1. Kano // 120545053
	// 2. Azamat // 182220854
	// 3. Tariel // 1824029863
	kano, _ := dbM.TelegramUserForTelegramID(120545053) // kano
	if kano == nil {
		tgKano := new(dbmodels.TelegramUser)
		tgKano.UserID = uint64(1200)
		tgKano.CompanyID = uint64(1)
		tgKano.TelegramID = 120545053
		tgKano.FirstName = "Каныбек"
		tgKano.SecondName = "Момукеев Саадаевич"
		tgKano.PhoneNumber = "+996-552-44-23-99"
		tgKano.TelegramAccount = "доп информация по аккаунту 1"
		kanoTelegramUserID, err := dbM.CreateTelegramUser(tgKano, 1)
		if err != nil {
			panic("CreateCuttedImage")
		} else {
			idString := fmt.Sprintf("Database stored KANO TelegramUserID %d", kanoTelegramUserID)
			dbM.providerLogs.WithFields(log.Fields{"Database ceated": idString}).Info("CreateSampleTelegramUsersIfNotExsist")
		}
	} else {
		kano.TelegramAccount = "доп информация по аккаунту 1"
		_, err := dbM.UpdateTelegramUser(kano, 1)
		if err != nil {
			panic("CreateSampleTelegramUsersIfNotExsist")
		}
	}
	azamat, _ := dbM.TelegramUserForTelegramID(182220854) // aza
	if azamat == nil {
		tgAzamat := new(dbmodels.TelegramUser)
		tgAzamat.UserID = uint64(1200)
		tgAzamat.CompanyID = uint64(1)
		tgAzamat.TelegramID = 182220854
		tgAzamat.FirstName = "Азамат"
		tgAzamat.SecondName = "Буржуев Чолпонбекович"
		tgAzamat.PhoneNumber = "+996 555 330 440"
		tgAzamat.TelegramAccount = "доп информация по аккаунту 2"
		azaTelegramUserID, err := dbM.CreateTelegramUser(tgAzamat, 1)
		if err != nil {
			panic("CreateCuttedImage")
		} else {
			idString := fmt.Sprintf("Database stored AZAMAT TelegramUserID %d", azaTelegramUserID)
			dbM.providerLogs.WithFields(log.Fields{"Database ceated": idString}).Info("CreateSampleTelegramUsersIfNotExsist")
		}
	} else {
		azamat.PhoneNumber = "+996 555 330 440"
		azamat.TelegramAccount = "доп информация по аккаунту 2"
		_, err := dbM.UpdateTelegramUser(azamat, 1)
		if err != nil {
			panic("CreateSampleTelegramUsersIfNotExsist")
		}
	}
	tariel, _ := dbM.TelegramUserForTelegramID(1824029863)
	if tariel == nil {
		tgTarik := new(dbmodels.TelegramUser)
		tgTarik.UserID = uint64(1200)
		tgTarik.CompanyID = uint64(1)
		tgTarik.TelegramID = 1824029863
		tgTarik.FirstName = "Тариэл"
		tgTarik.SecondName = "Алмаматов Мыйманбаевич"
		tgTarik.PhoneNumber = "+996 990066966"
		tgTarik.TelegramAccount = "доп информация по аккаунту 3"
		tarikTelegramUserID, err := dbM.CreateTelegramUser(tgTarik, 1)
		if err != nil {
			panic("CreateSampleTelegramUsersIfNotExsist")
		} else {
			idString := fmt.Sprintf("Database stored TARIEL TelegramUserID %d", tarikTelegramUserID)
			dbM.providerLogs.WithFields(log.Fields{"Database ceated": idString}).Info("CreateSampleTelegramUsersIfNotExsist")
		}
	} else {
		tariel.FirstName = "Тариэл"
		tariel.SecondName = "Алмаматов Мыйманбаевич"
		tariel.PhoneNumber = "+996 990066966"
		tariel.TelegramAccount = "доп информация по аккаунту 3"
		_, err := dbM.UpdateTelegramUser(tariel, 1)
		if err != nil {
			panic("CreateSampleTelegramUsersIfNotExsist")
		}
	}
}

func (dbM *DbManager) CreateUsersIfNotExsist() {
	// 5 sample user for test lets create
	// 1. Пользователь1 // 14c25048-b901-461f-934d-1adefb083912
	// 2. Пользователь2 // 56999956-47b6-442a-941e-6b3d548e6174
	// 3. Пользователь3 // 0b82413e-9ffb-4f10-bf59-a89840e476e1
	// 4. Пользователь4 // ee5a8e61-4b1c-46c3-8bee-2312f8b053e2
	// 5. Пользователь5 // 8551fdc5-e87a-4a68-a536-7938fee5b623
	user1, _ := dbM.UserForUUID("14c25048-b901-461f-934d-1adefb083912")
	if user1 == nil {
		userModel1 := new(dbmodels.UserRequest)
		userModel1.UserUUID = "14c25048-b901-461f-934d-1adefb083912"
		userModel1.UserType = uint32(dbmodels.USER_TYPE_OWNER)
		userModel1.CompanyId = uint64(1)
		userModel1.StockId = uint64(1)
		userModel1.UserImagePath = "UserImagePath1"
		userModel1.FirstName = "Пользователь1"
		userModel1.SecondName = "SecondName1"
		userModel1.Email = "email1@gmail.com"
		userModel1.Password = "password1"
		userModel1.PhoneNumber = "PhoneNumber1"
		userModel1.Address = "Address1"
		storedUserID1, err := dbM.CreateUser(userModel1)
		if err != nil {
			panic("Create user")
		} else {
			idString := fmt.Sprintf("Database stored userID = %d", storedUserID1)
			dbM.providerLogs.WithFields(log.Fields{"Database ceated user": idString}).Info("CreateUsersIfNotExsist")
		}
	}

	user2, _ := dbM.UserForUUID("0b82413e-9ffb-4f10-bf59-a89840e476e1")
	if user2 == nil {
		userModel2 := new(dbmodels.UserRequest)
		userModel2.UserUUID = "0b82413e-9ffb-4f10-bf59-a89840e476e1"
		userModel2.UserType = uint32(dbmodels.USER_TYPE_ADMIN)
		userModel2.CompanyId = uint64(1)
		userModel2.StockId = uint64(1)
		userModel2.UserImagePath = "UserImagePath2"
		userModel2.FirstName = "Пользователь2"
		userModel2.SecondName = "SecondName2"
		userModel2.Email = "email2@gmail.com"
		userModel2.Password = "password2"
		userModel2.PhoneNumber = "PhoneNumber2"
		userModel2.Address = "Address2"
		storedUserID2, err := dbM.CreateUser(userModel2)
		if err != nil {
			panic("Create user")
		} else {
			idString := fmt.Sprintf("Database stored userID = %d", storedUserID2)
			dbM.providerLogs.WithFields(log.Fields{"Database ceated user": idString}).Info("CreateUsersIfNotExsist")
		}
	}

	user3, _ := dbM.UserForUUID("56999956-47b6-442a-941e-6b3d548e6174")
	if user3 == nil {
		userModel3 := new(dbmodels.UserRequest)
		userModel3.UserUUID = "56999956-47b6-442a-941e-6b3d548e6174"
		userModel3.UserType = uint32(dbmodels.USER_TYPE_CAN_ADD_DOC)
		userModel3.CompanyId = uint64(1)
		userModel3.StockId = uint64(1)
		userModel3.UserImagePath = "UserImagePath3"
		userModel3.FirstName = "Пользователь3"
		userModel3.SecondName = "SecondName3"
		userModel3.Email = "email3@gmail.com"
		userModel3.Password = "password3"
		userModel3.PhoneNumber = "PhoneNumber3"
		userModel3.Address = "Address3"
		storedUserID3, err := dbM.CreateUser(userModel3)
		if err != nil {
			panic("Create user")
		} else {
			idString := fmt.Sprintf("Database stored userID = %d", storedUserID3)
			dbM.providerLogs.WithFields(log.Fields{"Database ceated user": idString}).Info("CreateUsersIfNotExsist")
		}
	}
}

func (dbM *DbManager) CreateTelegramUser(telegramUser *dbmodels.TelegramUser, companyId uint64) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		lastInsertId, err := dbmodels.StoreTelegramUser(tx, telegramUser)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(lastInsertId)
	}
	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}

func (dbM *DbManager) UpdateTelegramUser(telegramUser *dbmodels.TelegramUser, companyId uint64) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		rowsAffected, err := dbmodels.UpdateTelegramUser(tx, telegramUser)
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			dbM.sendError(err)
			return
		}

		dbM.sendLastInsertedId(rowsAffected)
	}
	dbM.dbChannel <- f
	return dbM.waitForLastInsertedId()
}

func (dbM *DbManager) AllTelegramUsersForCompany(companyId uint64) ([]*dbmodels.TelegramUser, error) {
	customers, err := dbmodels.AllTelegramUsersForCompany(dbM.DB, companyId)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (dbM *DbManager) TelegramUserForTelegramID(telegramID uint64) (*dbmodels.TelegramUser, error) {
	telegramUser, err := dbmodels.TelegramUserForTelegramID(dbM.DB, telegramID)
	if err != nil {
		return nil, err
	}
	return telegramUser, nil
}

func (dbM *DbManager) TelegramUserForTelegramUserID(telegramUserID uint64) (*dbmodels.TelegramUser, error) {
	telegramUser, err := dbmodels.TelegramUserForTelegramUserID(dbM.DB, telegramUserID)
	if err != nil {
		return nil, err
	}
	return telegramUser, nil
}

// func TelegramUserForTelegramID(db *sqlx.DB, telegramID uint64) (*TelegramUser, error) {

func (dbM *DbManager) AllUpdatedTelegramUsers(suppFilter *dbmodels.TelegramUserFilter, companyId uint64) ([]*dbmodels.TelegramUser, error) {
	customers, err := dbmodels.AllUpdatedTelegramUserForCompany(dbM.DB, suppFilter, companyId)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// ------------------------------------------------------------------------------------------------- //
var schemaDROP = `
DROP TABLE IF EXISTS some_unknown_table;
`

var schema = `
CREATE TABLE IF NOT EXISTS person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE IF NOT EXISTS place (
    country text,
    city text NULL,
    telcode integer
)`

type Persondb struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func (dbM *DbManager) DropAllTables() {
	dbM.DB.MustExec(schemaDROP)
	dbM.DB.MustExec(schema)
}

func (dbM *DbManager) TestSqlxDatabaseConnection() {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	dbM.DB.MustExec(schema)

	tx := dbM.DB.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Persondbdb{}) that you have populated, you can pass it in as &person
	tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Persondb{"Jane", "Citizen", "jane.citzen@example.com"})
	tx.Commit()

	// Query the database, storing results in a []Persondb (wrapped in []interface{})
	people := []Persondb{}
	dbM.DB.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	jason, john := people[0], people[1]

	fmt.Printf("%#v\n%#v", jason, john)
	// Persondb{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}
	// Persondb{FirstName:"John", LastName:"Doe", Email:"johndoeDNE@gmail.net"}

	// You can also get a single result, a la QueryRow
	jason = Persondb{}
	err = dbM.DB.Get(&jason, "SELECT * FROM person WHERE first_name=$1", "Jason")
	fmt.Printf("%#v\n", jason)
	// Persondb{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}

	// if you have null fields and use SELECT *, you must use sql.Null* in your struct
	places := []Place{}
	err = dbM.DB.Select(&places, "SELECT * FROM place ORDER BY telcode ASC")
	if err != nil {
		fmt.Println(err)
		return
	}
	usa, singsing, honkers := places[0], places[1], places[2]

	fmt.Printf("%#v\n%#v\n%#v\n", usa, singsing, honkers)
	// Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	// Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}
	// Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}

	// Loop through rows using only one struct
	place := Place{}
	rows, _ := dbM.DB.Queryx("SELECT * FROM place")
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", place)
	}
	// Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	// Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}
	// Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}

	// Named queries, using `:name` as the bindvar.  Automatic bindvar support
	// which takes into account the dbtype based on the driverName on sqlx.Open/Connect
	_, err = dbM.DB.NamedExec(`INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`,
		map[string]interface{}{
			"first": "Bin",
			"last":  "Smuth",
			"email": "bensmith@allblacks.nz",
		})

	// Selects Mr. Smith from the database
	rows, err = dbM.DB.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "Bin"})
	fmt.Printf("%v\n", rows)

	// Named queries can also use structs.  Their bind names follow the same rules
	// as the name -> db mapping, so struct fields are lowercased and the `db` tag
	// is taken into consideration.
	rows, err = dbM.DB.NamedQuery(`SELECT * FROM person WHERE first_name=:first_name`, jason)
	fmt.Printf("%v\n", rows)

	// batch insert

	// batch insert with structs
	personStructs := []Persondb{
		{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
		{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
		{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	}

	_, err = dbM.DB.NamedExec(`INSERT INTO person (first_name, last_name, email)
        VALUES (:first_name, :last_name, :email)`, personStructs)

	// batch insert with maps
	personMaps := []map[string]interface{}{
		{"first_name": "Ardie", "last_name": "Savea", "email": "asavea@ab.co.nz"},
		{"first_name": "Sonny Bill", "last_name": "Williams", "email": "sbw@ab.co.nz"},
		{"first_name": "Ngani", "last_name": "Laumape", "email": "nlaumape@ab.co.nz"},
	}

	_, err = dbM.DB.NamedExec(`INSERT INTO person (first_name, last_name, email)
        VALUES (:first_name, :last_name, :email)`, personMaps)
}
