package dbmodelsTest

import (
	"testing"

	dbmodels "github.com/KanybekMomukeyev/imgdatabase/v3/models/dbmodels"
	"github.com/stretchr/testify/assert"
)

// type TelegramUser struct {
// 	TelegramUserID  uint64 `db:"telegram_user_id" json:"telegram_user_id"`
// 	UserID          uint64 `db:"user_id"  json:"user_id"`
// 	CompanyID       uint64 `db:"company_id" json:"company_id"`
// 	TelegramID      uint64 `db:"cutted_image_id" json:"cutted_image_id"`
// 	FirstName       string `db:"first_name" json:"first_name"`
// 	SecondName      string `db:"second_name" json:"second_name"`
// 	PhoneNumber     string `db:"phone_number" json:"phone_number"`
// 	TelegramAccount string `db:"tegram_account" json:"tegram_account"`
// }

func TestCreateTelegramUser(t *testing.T) {

	tgUser := new(dbmodels.TelegramUser)
	tgUser.UserID = uint64(1100)
	tgUser.CompanyID = uint64(40000)
	tgUser.TelegramID = uint64(5234432)
	tgUser.FirstName = "FirstName"
	tgUser.SecondName = "SecondName"
	tgUser.PhoneNumber = "PhoneNumber"
	tgUser.TelegramAccount = "Address"

	telegramUserID, err := DbMng.CreateTelegramUser(tgUser, uint64(40000))

	assert.Equal(t, telegramUserID, uint64(1), "they should be equal")
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, telegramUserID, "userId  should not be equal to zero")
	if assert.NotNil(t, tgUser) {
	}

	tgUsers, err := DbMng.AllTelegramUsersForCompany(uint64(40000))
	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, tgUsers) {
	}
	println(len(tgUsers))
	assert.NotEqual(t, 0, len(tgUsers), "users.count should not be equal to zero")

	for _, tgUser := range tgUsers {
		assert.Equal(t, tgUser.TelegramUserID, uint64(1), "they should be equal")
		assert.Equal(t, tgUser.UserID, uint64(1100), "they should be equal")
		assert.Equal(t, tgUser.CompanyID, uint64(40000), "they should be equal")
		assert.Equal(t, tgUser.TelegramID, uint64(5234432), "they should be equal")
		assert.Equal(t, tgUser.FirstName, "FirstName", "they should be equal")
		assert.Equal(t, tgUser.SecondName, "SecondName", "they should be equal")
		assert.Equal(t, tgUser.PhoneNumber, "PhoneNumber", "they should be equal")
		assert.Equal(t, tgUser.TelegramAccount, "Address", "they should be equal")
	}
}

func TestUpdateTelegramUser(t *testing.T) {

	tgUser := new(dbmodels.TelegramUser)
	tgUser.TelegramUserID = uint64(1)
	tgUser.UserID = uint64(110320)
	tgUser.CompanyID = uint64(40000)
	tgUser.TelegramID = uint64(5234432423)
	tgUser.FirstName = "FirstNameFirstName"
	tgUser.SecondName = "SecondNameFirstName"
	tgUser.PhoneNumber = "PhoneNumberFirstName"
	tgUser.TelegramAccount = "AddressFirstName"

	rowsAffected, err := DbMng.UpdateTelegramUser(tgUser, uint64(40000))
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, rowsAffected, "rowsAffected should not be equal to zero")

	tgUserFor, _ := DbMng.TelegramUserForTelegramID(5234432423)
	if true {
		assert.Equal(t, tgUserFor.TelegramUserID, uint64(1), "they should be equal")
		assert.Equal(t, tgUserFor.UserID, uint64(110320), "they should be equal")
		assert.Equal(t, tgUserFor.CompanyID, uint64(40000), "they should be equal")
		assert.Equal(t, tgUserFor.TelegramID, uint64(5234432423), "they should be equal")
		assert.Equal(t, tgUserFor.FirstName, "FirstNameFirstName", "they should be equal")
		assert.Equal(t, tgUserFor.SecondName, "SecondNameFirstName", "they should be equal")
		assert.Equal(t, tgUserFor.PhoneNumber, "PhoneNumberFirstName", "they should be equal")
		assert.Equal(t, tgUserFor.TelegramAccount, "AddressFirstName", "they should be equal")
	}

	tgUsers, err := DbMng.AllTelegramUsersForCompany(uint64(40000))
	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, tgUsers) {
	}

	println(len(tgUsers))
	assert.NotEqual(t, 0, len(tgUsers), "users.count should not be equal to zero")

	for _, tgUser := range tgUsers {
		assert.Equal(t, tgUser.TelegramUserID, uint64(1), "they should be equal")
		assert.Equal(t, tgUser.UserID, uint64(110320), "they should be equal")
		assert.Equal(t, tgUser.CompanyID, uint64(40000), "they should be equal")
		assert.Equal(t, tgUser.TelegramID, uint64(5234432423), "they should be equal")
		assert.Equal(t, tgUser.FirstName, "FirstNameFirstName", "they should be equal")
		assert.Equal(t, tgUser.SecondName, "SecondNameFirstName", "they should be equal")
		assert.Equal(t, tgUser.PhoneNumber, "PhoneNumberFirstName", "they should be equal")
		assert.Equal(t, tgUser.TelegramAccount, "AddressFirstName", "they should be equal")
	}
}
