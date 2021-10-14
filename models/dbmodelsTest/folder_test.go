package dbmodelsTest

import (
	"testing"

	dbmodels "github.com/KanybekMomukeyev/imgdatabase/v3/models/dbmodels"
	"github.com/stretchr/testify/assert"
)

// type FolderModel struct {
// 	FolderID         uint64 `db:"folder_id"  json:"folder_id"`
// 	CompanyID        uint64 `db:"company_id"  json:"company_id"`
// 	DocmodelID       uint64 `db:"docmodel_id"  json:"docmodel_id"`
// 	WordCount        uint64 `db:"word_count"  json:"word_count"`
// 	ParsedImageCount uint64 `db:"parsed_image_count"  json:"parsed_image_count"`
// 	FolderImagePath  string `db:"folder_image_path"  json:"folder_image_path"`
// 	FolderName       string `db:"folder_name"  json:"folder_name"`
// 	ContactFname     string `db:"contact_fname"  json:"contact_fname"`
// 	PhoneNumber      string `db:"phone_number"  json:"phone_number"`
// 	Address          string `db:"address"  json:"address"`
// 	UpdatedAt        uint64 `db:"updated_at"  json:"updated_at"`
// }

func TestCreateFolder(t *testing.T) {

	folderModel := new(dbmodels.FolderModel)

	folderModel.CompanyID = uint64(99589321)
	folderModel.DocmodelID = uint64(87589321)
	folderModel.WordCount = uint64(54321)
	folderModel.ParsedImageCount = uint64(54356)
	folderModel.FolderImagePath = "SupplierImagePath/SupplierImagePath/SupplierImagePath"
	folderModel.FolderName = "CompanyNameCompanyName"
	folderModel.ContactFname = "ContactFnameContactFname"
	folderModel.PhoneNumber = "PhoneNumberPhoneNumber"
	folderModel.Address = "AddressAddressAddress"

	folderModelId, err := DbMng.CreateFolderModel(folderModel, folderModel.CompanyID)

	assert.Equal(t, folderModelId, uint64(1), "they should be equal")
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, folderModelId, "userId  should not be equal to zero")
	if assert.NotNil(t, folderModel) {
	}

	folderModels, err := DbMng.AllFolderModelsForInitial(folderModel.CompanyID)
	assert.NotEqual(t, 0, len(folderModels), "users.count should not be equal to zero")

	for _, folderModel := range folderModels {
		assert.Equal(t, folderModel.FolderID, uint64(1), "they should be equal")
		assert.Equal(t, folderModel.CompanyID, uint64(99589321), "they should be equal")
		assert.Equal(t, folderModel.DocmodelID, uint64(87589321), "they should be equal")
		assert.Equal(t, folderModel.WordCount, uint64(54321), "they should be equal")
		assert.Equal(t, folderModel.ParsedImageCount, uint64(54356), "they should be equal")
		assert.Equal(t, folderModel.FolderImagePath, "SupplierImagePath/SupplierImagePath/SupplierImagePath", "CategoryName")
		assert.Equal(t, folderModel.FolderName, "CompanyNameCompanyName", "CategoryName")
		assert.Equal(t, folderModel.ContactFname, "ContactFnameContactFname", "CategoryName")
		assert.Equal(t, folderModel.PhoneNumber, "PhoneNumberPhoneNumber", "CategoryName")
		assert.Equal(t, folderModel.Address, "AddressAddressAddress", "CategoryName")
	}

	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, folderModels) {
	}
}

func TestUpdateFolder(t *testing.T) {

	folderModel := new(dbmodels.FolderModel)
	folderModel.FolderID = uint64(1)
	folderModel.CompanyID = uint64(9958932100)
	folderModel.DocmodelID = uint64(58932143)
	folderModel.WordCount = uint64(124321)
	folderModel.ParsedImageCount = uint64(124522)
	folderModel.FolderImagePath = "SupplierImagePathSupplierImagePath"
	folderModel.FolderName = "CompanyNameCompanyName"
	folderModel.ContactFname = "ContactFnameContactFname"
	folderModel.PhoneNumber = "PhoneNumberPhoneNumber"
	folderModel.Address = "AddressAddress"

	rowsAffected, err := DbMng.UpdateFolderModel(folderModel, folderModel.CompanyID)

	assert.Equal(t, rowsAffected, uint64(1), "they should be equal")
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, rowsAffected, "userId  should not be equal to zero")
	if assert.NotNil(t, folderModel) {
	}

	folderModels, err := DbMng.AllFolderModelsForInitial(folderModel.CompanyID)
	assert.NotEqual(t, 0, len(folderModels), "folderModels should not be equal to zero")

	for _, folderModel := range folderModels {
		assert.Equal(t, folderModel.FolderID, uint64(1), "they should be equal")
		assert.Equal(t, folderModel.CompanyID, uint64(9958932100), "they should be equal")
		assert.Equal(t, folderModel.DocmodelID, uint64(58932143), "they should be equal")
		assert.Equal(t, folderModel.WordCount, uint64(124321), "they should be equal")
		assert.Equal(t, folderModel.ParsedImageCount, uint64(124522), "they should be equal")
		assert.Equal(t, folderModel.FolderImagePath, "SupplierImagePathSupplierImagePath", "CategoryName")
		assert.Equal(t, folderModel.FolderName, "CompanyNameCompanyName", "CategoryName")
		assert.Equal(t, folderModel.ContactFname, "ContactFnameContactFname", "CategoryName")
		assert.Equal(t, folderModel.PhoneNumber, "PhoneNumberPhoneNumber", "CategoryName")
		assert.Equal(t, folderModel.Address, "AddressAddress", "CategoryName")
	}

	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, folderModels) {
	}
}
