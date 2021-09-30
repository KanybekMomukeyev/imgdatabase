package dbmodelsTest

import (
	"testing"

	dbmodels "github.com/KanybekMomukeyev/imgdatabase/models/dbmodels"
	"github.com/stretchr/testify/assert"
)

func TestCreateFolder(t *testing.T) {

	folderModel := new(dbmodels.FolderModel)

	folderModel.DocmodelID = uint64(589321)
	folderModel.WordCount = uint64(54321)
	folderModel.ParsedImageCount = uint64(54356)
	folderModel.FolderImagePath = "SupplierImagePath"
	folderModel.FolderName = "CompanyName"
	folderModel.ContactFname = "ContactFname"
	folderModel.PhoneNumber = "PhoneNumber"
	folderModel.Address = "Address"

	folderModelId, err := DbMng.CreateFolderModel(folderModel, 3000)

	assert.Equal(t, folderModelId, uint64(1), "they should be equal")
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, folderModelId, "userId  should not be equal to zero")
	if assert.NotNil(t, folderModel) {
	}

	folderModels, err := DbMng.AllFolderModelsForInitial(3000)
	assert.NotEqual(t, 0, len(folderModels), "users.count should not be equal to zero")

	for _, folderModel := range folderModels {
		assert.Equal(t, folderModel.FolderID, uint64(1), "they should be equal")
		assert.Equal(t, folderModel.DocmodelID, uint64(589321), "they should be equal")
		assert.Equal(t, folderModel.WordCount, uint64(54321), "they should be equal")
		assert.Equal(t, folderModel.ParsedImageCount, uint64(54356), "they should be equal")
		assert.Equal(t, folderModel.FolderImagePath, "SupplierImagePath", "CategoryName")
		assert.Equal(t, folderModel.FolderName, "CompanyName", "CategoryName")
		assert.Equal(t, folderModel.ContactFname, "ContactFname", "CategoryName")
		assert.Equal(t, folderModel.PhoneNumber, "PhoneNumber", "CategoryName")
		assert.Equal(t, folderModel.Address, "Address", "CategoryName")
	}

	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, folderModels) {
	}
}

func TestUpdateFolder(t *testing.T) {

	folderModel := new(dbmodels.FolderModel)
	folderModel.FolderID = uint64(1)
	folderModel.DocmodelID = uint64(58932143)
	folderModel.WordCount = uint64(124321)
	folderModel.ParsedImageCount = uint64(124522)
	folderModel.FolderImagePath = "SupplierImagePathSupplierImagePath"
	folderModel.FolderName = "CompanyNameCompanyName"
	folderModel.ContactFname = "ContactFnameContactFname"
	folderModel.PhoneNumber = "PhoneNumberPhoneNumber"
	folderModel.Address = "AddressAddress"

	rowsAffected, err := DbMng.UpdateFolderModel(folderModel, 3000)

	assert.Equal(t, rowsAffected, uint64(1), "they should be equal")
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, rowsAffected, "userId  should not be equal to zero")
	if assert.NotNil(t, folderModel) {
	}

	folderModels, err := DbMng.AllFolderModelsForInitial(3000)
	assert.NotEqual(t, 0, len(folderModels), "users.count should not be equal to zero")

	for _, folderModel := range folderModels {
		assert.Equal(t, folderModel.FolderID, uint64(1), "they should be equal")
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
