package dbmodelsTest

import (
	"testing"

	dbmodels "github.com/KanybekMomukeyev/imgdatabase/v3/models/dbmodels"
	"github.com/stretchr/testify/assert"
)

// type CuttedImage struct {
// 	CuttedImageID    uint64 `db:"image_id"  json:"image_id"`
// 	CuttedImageState uint32 `db:"cutted_image_state" json:"cutted_image_state"` // DEFAULT VALUE 1000
// 	CuttedImageType  uint32 `db:"cutted_image_type" json:"cutted_image_type"`   // DEFAULT VALUE 101010
// 	DocModelID       uint64 `db:"docmodel_id"  json:"docmodel_id"`
// 	CompanyID        uint64 `db:"company_id"  json:"company_id"`
// 	FolderID         uint64 `db:"folder_id"  json:"folder_id"`
// 	ParsedImagePath  string `db:"parsed_image_path"  json:"parsed_image_path"`
// 	FolderName       string `db:"folder_name"  json:"folder_name"`
// 	SecondName       string `db:"second_name"  json:"second_name"`
// 	PhoneNumber      string `db:"phone_number"  json:"phone_number"`
// 	Address          string `db:"address"  json:"address"`
// 	UpdatedAt        uint64 `db:"updated_at"  json:"updated_at"`         12 TOTAL FIELDS, ID, UPDATEAT REMOVED 10 FIELDS SHOUDL BE TEST
// }

func TestCreateCuttedImage(t *testing.T) {

	cuttedimage := new(dbmodels.CuttedImage)

	cuttedimage.CuttedImageState = uint32(dbmodels.CUTTED_IMAGE_MARKED_INVALID)
	cuttedimage.CuttedImageType = uint32(dbmodels.TYPE_CUTTED_IMAGE_HANDWRITTEN)
	cuttedimage.DocModelID = uint64(124383456)
	cuttedimage.CompanyID = uint64(9876383)
	cuttedimage.FolderID = uint64(1847243)
	cuttedimage.ParsedImagePath = "CustomerImagePath/SomePath/DIDI"
	cuttedimage.FolderName = "FolderNameFolderNameFolderName"
	cuttedimage.SecondName = "SecondNameSecondNameSecondName"
	cuttedimage.PhoneNumber = "PhoneNumberPhoneNumberPhoneNumber"
	cuttedimage.Address = "AddressAddressAddressAddress"

	cuttedimageId, err := DbMng.CreateCuttedImage(cuttedimage, cuttedimage.CompanyID)

	assert.Equal(t, cuttedimageId, uint64(1), "they should be equal")
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, cuttedimageId, "userId  should not be equal to zero")
	if assert.NotNil(t, cuttedimage) {
	}

	cuttedimages, err := DbMng.AllCuttedImagesForCompany(cuttedimage.CompanyID)
	assert.NotEqual(t, 0, len(cuttedimages), "users.count should not be equal to zero")

	for _, cuttedimage := range cuttedimages {
		assert.Equal(t, cuttedimage.CuttedImageID, uint64(1), "they should be equal")

		assert.Equal(t, cuttedimage.CuttedImageState, uint32(dbmodels.CUTTED_IMAGE_MARKED_INVALID), "CategoryName")
		assert.Equal(t, cuttedimage.CuttedImageType, uint32(dbmodels.TYPE_CUTTED_IMAGE_HANDWRITTEN), "CategoryName")
		assert.Equal(t, cuttedimage.DocModelID, uint64(124383456), "they should be equal")
		assert.Equal(t, cuttedimage.CompanyID, uint64(9876383), "they should be equal")
		assert.Equal(t, cuttedimage.FolderID, uint64(1847243), "they should be equal")
		assert.Equal(t, cuttedimage.ParsedImagePath, "CustomerImagePath/SomePath/DIDI", "CategoryName")
		assert.Equal(t, cuttedimage.FolderName, "FolderNameFolderNameFolderName", "CategoryName")
		assert.Equal(t, cuttedimage.SecondName, "SecondNameSecondNameSecondName", "CategoryName")
		assert.Equal(t, cuttedimage.PhoneNumber, "PhoneNumberPhoneNumberPhoneNumber", "CategoryName")
		assert.Equal(t, cuttedimage.Address, "AddressAddressAddressAddress", "CategoryName")
	}

	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, cuttedimages)

	cuttedImageSaved, _ := DbMng.CuttedImageForID(cuttedimageId)
	if true {
		assert.Equal(t, cuttedImageSaved.CuttedImageState, uint32(dbmodels.CUTTED_IMAGE_MARKED_INVALID), "CategoryName")
		assert.Equal(t, cuttedImageSaved.CuttedImageType, uint32(dbmodels.TYPE_CUTTED_IMAGE_HANDWRITTEN), "CategoryName")
		assert.Equal(t, cuttedImageSaved.DocModelID, uint64(124383456), "they should be equal")
		assert.Equal(t, cuttedImageSaved.CompanyID, uint64(9876383), "they should be equal")
		assert.Equal(t, cuttedImageSaved.FolderID, uint64(1847243), "they should be equal")
		assert.Equal(t, cuttedImageSaved.ParsedImagePath, "CustomerImagePath/SomePath/DIDI", "CategoryName")
		assert.Equal(t, cuttedImageSaved.FolderName, "FolderNameFolderNameFolderName", "CategoryName")
		assert.Equal(t, cuttedImageSaved.SecondName, "SecondNameSecondNameSecondName", "CategoryName")
		assert.Equal(t, cuttedImageSaved.PhoneNumber, "PhoneNumberPhoneNumberPhoneNumber", "CategoryName")
		assert.Equal(t, cuttedImageSaved.Address, "AddressAddressAddressAddress", "CategoryName")
	}
}

func TestUpdateCuttedImage(t *testing.T) {

	cuttedimage := new(dbmodels.CuttedImage)
	cuttedimage.CuttedImageID = uint64(1)

	cuttedimage.CuttedImageState = uint32(dbmodels.CUTTED_IMAGE_TRANSLATED)
	cuttedimage.CuttedImageType = uint32(dbmodels.TYPE_CUTTED_IMAGE_UNKNOWN)
	cuttedimage.DocModelID = uint64(12438345600)
	cuttedimage.CompanyID = uint64(987638300)
	cuttedimage.FolderID = uint64(184724300)
	cuttedimage.ParsedImagePath = "CustomerImagePath/SomePath/DIDI00"
	cuttedimage.FolderName = "FolderNameFolderNameFolderName00"
	cuttedimage.SecondName = "SecondNameSecondNameSecondName00"
	cuttedimage.PhoneNumber = "PhoneNumberPhoneNumberPhoneNumber00"
	cuttedimage.Address = "AddressAddressAddressAddress00"

	rowsAffected, err := DbMng.UpdateCuttedImage(cuttedimage, cuttedimage.CompanyID)
	assert.NotEqual(t, 0, rowsAffected, "userId  should not be equal to zero")
	assert.Nil(t, err, "error should be nil")

	unknownRandom, _ := DbMng.RandomUnknownType()
	assert.NotNil(t, unknownRandom)

	cuttedimages, err := DbMng.AllCuttedImagesForCompany(cuttedimage.CompanyID)
	assert.NotEqual(t, 0, len(cuttedimages), "users.count should not be equal to zero")

	for _, cuttedimage := range cuttedimages {
		assert.Equal(t, cuttedimage.CuttedImageID, uint64(1), "they should be equal")
		assert.Equal(t, cuttedimage.CuttedImageState, uint32(dbmodels.CUTTED_IMAGE_TRANSLATED), "CategoryName")
		assert.Equal(t, cuttedimage.CuttedImageType, uint32(dbmodels.TYPE_CUTTED_IMAGE_UNKNOWN), "CategoryName")
		assert.Equal(t, cuttedimage.DocModelID, uint64(12438345600), "they should be equal")
		assert.Equal(t, cuttedimage.CompanyID, uint64(987638300), "they should be equal")
		assert.Equal(t, cuttedimage.FolderID, uint64(184724300), "they should be equal")
		assert.Equal(t, cuttedimage.ParsedImagePath, "CustomerImagePath/SomePath/DIDI00", "CategoryName")
		assert.Equal(t, cuttedimage.FolderName, "FolderNameFolderNameFolderName00", "CategoryName")
		assert.Equal(t, cuttedimage.SecondName, "SecondNameSecondNameSecondName00", "CategoryName")
		assert.Equal(t, cuttedimage.PhoneNumber, "PhoneNumberPhoneNumberPhoneNumber00", "CategoryName")
		assert.Equal(t, cuttedimage.Address, "AddressAddressAddressAddress00", "CategoryName")
	}

	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, cuttedimages) {
	}
}
