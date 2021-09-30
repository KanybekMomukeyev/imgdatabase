package dbmodelsTest

import (
	"testing"

	dbmodels "github.com/KanybekMomukeyev/imgdatabase/v3/models/dbmodels"
	"github.com/stretchr/testify/assert"
)

// type CuttedImageTranslate struct {
// 	CuttedImageTranslateID uint64 `db:"cutted_image_translate_id"  json:"cutted_image_translate_id"`
// 	CuttedImageID          uint64 `db:"cutted_image_id" json:"cutted_image_id"`
// 	TelegramUserID         uint64 `db:"telegram_user_id" json:"telegram_user_id"`
// 	TranslatedWord         string `db:"translated_word" json:"translated_word"`
// 	Comments               string `db:"comments" json:"comments"`
// 	Summary                string `db:"summary" json:"summary"`
// 	AcceptStatus           uint32 `db:"accept_status" json:"accept_status"`
// }

func TestCreateCuttedImage(t *testing.T) {

	cuttedimage := new(dbmodels.CuttedImage)

	cuttedimage.FolderID = uint64(1243)
	cuttedimage.ParsedImagePath = "CustomerImagePath"
	cuttedimage.FolderName = "FirstName"
	cuttedimage.SecondName = "SecondName"
	cuttedimage.PhoneNumber = "PhoneNumber"
	cuttedimage.Address = "Address"
	cuttedimage.DocModelID = uint64(1200)
	cuttedimage.CuttedImageState = uint32(54200)

	cuttedimageId, err := DbMng.CreateCuttedImage(cuttedimage, 3000)

	assert.Equal(t, cuttedimageId, uint64(1), "they should be equal")
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, cuttedimageId, "userId  should not be equal to zero")
	if assert.NotNil(t, cuttedimage) {
	}

	cuttedimages, err := DbMng.AllCuttedImagesForCompany(3000)
	assert.NotEqual(t, 0, len(cuttedimages), "users.count should not be equal to zero")

	for _, cuttedimage := range cuttedimages {
		assert.Equal(t, cuttedimage.CuttedImageID, uint64(1), "they should be equal")
		assert.Equal(t, cuttedimage.FolderID, uint64(1243), "they should be equal")
		assert.Equal(t, cuttedimage.ParsedImagePath, "CustomerImagePath", "CategoryName")
		assert.Equal(t, cuttedimage.FolderName, "FirstName", "CategoryName")
		assert.Equal(t, cuttedimage.SecondName, "SecondName", "CategoryName")
		assert.Equal(t, cuttedimage.PhoneNumber, "PhoneNumber", "CategoryName")
		assert.Equal(t, cuttedimage.Address, "Address", "CategoryName")
		assert.Equal(t, cuttedimage.DocModelID, uint64(1200), "CategoryName")
		assert.Equal(t, cuttedimage.CuttedImageState, uint32(54200), "CategoryName")
	}

	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, cuttedimages) {
	}

	cuttedImageSaved, _ := DbMng.CuttedImageForID(cuttedimageId)
	if true {
		assert.Equal(t, cuttedImageSaved.CuttedImageID, uint64(1), "they should be equal")
		assert.Equal(t, cuttedImageSaved.FolderID, uint64(1243), "they should be equal")
		assert.Equal(t, cuttedImageSaved.ParsedImagePath, "CustomerImagePath", "CategoryName")
		assert.Equal(t, cuttedImageSaved.FolderName, "FirstName", "CategoryName")
		assert.Equal(t, cuttedImageSaved.SecondName, "SecondName", "CategoryName")
		assert.Equal(t, cuttedImageSaved.PhoneNumber, "PhoneNumber", "CategoryName")
		assert.Equal(t, cuttedImageSaved.Address, "Address", "CategoryName")
		assert.Equal(t, cuttedImageSaved.DocModelID, uint64(1200), "CategoryName")
		assert.Equal(t, cuttedImageSaved.CuttedImageState, uint32(54200), "CategoryName")
	}
}

func TestUpdateCuttedImage(t *testing.T) {

	cuttedimage := new(dbmodels.CuttedImage)
	cuttedimage.CuttedImageID = uint64(1)
	cuttedimage.FolderID = uint64(3211)
	cuttedimage.ParsedImagePath = "CustomerImagePathCustomerImagePath"
	cuttedimage.FolderName = "FirstNameFirstName"
	cuttedimage.SecondName = "SecondNameSecondName"
	cuttedimage.PhoneNumber = "PhoneNumberPhoneNumber"
	cuttedimage.Address = "AddressAddress"
	cuttedimage.DocModelID = uint64(120012)
	cuttedimage.CuttedImageState = uint32(4312)

	rowsAffected, err := DbMng.UpdateCuttedImage(cuttedimage, 3000)
	assert.NotEqual(t, 0, rowsAffected, "userId  should not be equal to zero")
	assert.Nil(t, err, "error should be nil")

	cuttedimages, err := DbMng.AllCuttedImagesForCompany(3000)
	assert.NotEqual(t, 0, len(cuttedimages), "users.count should not be equal to zero")

	for _, cuttedimage := range cuttedimages {
		assert.Equal(t, cuttedimage.CuttedImageID, uint64(1), "they should be equal")
		assert.Equal(t, cuttedimage.FolderID, uint64(3211), "they should be equal")
		assert.Equal(t, cuttedimage.ParsedImagePath, "CustomerImagePathCustomerImagePath", "CategoryName")
		assert.Equal(t, cuttedimage.FolderName, "FirstNameFirstName", "CategoryName")
		assert.Equal(t, cuttedimage.SecondName, "SecondNameSecondName", "CategoryName")
		assert.Equal(t, cuttedimage.PhoneNumber, "PhoneNumberPhoneNumber", "CategoryName")
		assert.Equal(t, cuttedimage.Address, "AddressAddress", "CategoryName")
		assert.Equal(t, cuttedimage.DocModelID, uint64(120012), "CategoryName")
		assert.Equal(t, cuttedimage.CuttedImageState, uint32(4312), "CategoryName")
	}

	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, cuttedimages) {
	}
}
