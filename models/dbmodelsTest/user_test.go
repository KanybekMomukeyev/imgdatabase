package dbmodelsTest

import (
	"fmt"
	"testing"
	"time"

	models "github.com/KanybekMomukeyev/imgdatabase/v3/models"
	dbmodels "github.com/KanybekMomukeyev/imgdatabase/v3/models/dbmodels"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	user := new(dbmodels.UserRequest)
	uuid1 := uuid.NewV4()
	var uuid = uuid1.String()

	user.UserUUID = uuid
	user.UserType = uint32(1100)
	user.CompanyId = uint64(40000)
	user.StockId = uint64(5234432)
	user.UserImagePath = "UserImagePath"
	user.FirstName = "FirstName"
	user.SecondName = "SecondName"
	user.Email = "Email"
	user.Password = "Password"
	user.PhoneNumber = "PhoneNumber"
	user.Address = "Address"

	userId, err := DbMng.CreateUser(user)

	assert.Equal(t, userId, uint64(1), "they should be equal")
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, userId, "userId  should not be equal to zero")
	if assert.NotNil(t, user) {
	}

	users, err := DbMng.AllUsers(uint64(40000))
	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, users) {
	}
	println(len(users))
	assert.NotEqual(t, 0, len(users), "users.count should not be equal to zero")

	for _, user := range users {
		assert.Equal(t, user.UserId, uint64(1), "they should be equal")
		assert.Equal(t, user.StockId, uint64(5234432), "they should be equal")
		assert.Equal(t, user.UserType, uint32(1100), "they should be equal")
		assert.Equal(t, user.CompanyId, uint64(40000), "they should be equal")
		assert.Equal(t, user.UserUUID, uuid, "they should be equal")
		assert.Equal(t, user.UserImagePath, "UserImagePath", "they should be equal")
		assert.Equal(t, user.FirstName, "FirstName", "they should be equal")
		assert.Equal(t, user.SecondName, "SecondName", "they should be equal")
		assert.Equal(t, user.Email, "Email", "they should be equal")
		assert.Equal(t, user.Password, "Password", "they should be equal")
		assert.Equal(t, user.PhoneNumber, "PhoneNumber", "they should be equal")
		assert.Equal(t, user.Address, "Address", "they should be equal")
	}
}

func TestUpdateUser(t *testing.T) {

	uuid1 := uuid.NewV4()
	var uuid = uuid1.String()

	user := new(dbmodels.UserRequest)
	user.UserId = 1
	user.UserUUID = uuid
	user.UserType = uint32(1200)
	user.UserImagePath = "UserImagePathUserImagePath"
	user.FirstName = "FirstNameFirstName"
	user.SecondName = "SecondNameSecondName"
	user.Email = "EmailEmail"
	user.Password = "PasswordPassword"
	user.PhoneNumber = "PhoneNumberPhoneNumber"
	user.Address = "AddressAddressAddress"
	user.CompanyId = uint64(40000)
	user.StockId = uint64(1234432)

	rowsAffected, err := DbMng.UpdateUser(user)
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, rowsAffected, "rowsAffected should not be equal to zero")

	users, err := DbMng.AllUsers(uint64(40000))
	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, users) {
	}

	println(len(users))
	assert.NotEqual(t, 0, len(users), "users.count should not be equal to zero")

	for _, user := range users {
		assert.Equal(t, user.UserId, uint64(1), "they should be equal")
		assert.Equal(t, user.UserType, uint32(1200), "they should be equal")
		assert.Equal(t, user.StockId, uint64(1234432), "they should be equal")
		assert.Equal(t, user.CompanyId, uint64(40000), "they should be equal")
		assert.Equal(t, user.UserUUID, uuid, "they should be equal")
		assert.Equal(t, user.UserImagePath, "UserImagePathUserImagePath", "they should be equal")
		assert.Equal(t, user.FirstName, "FirstNameFirstName", "they should be equal")
		assert.Equal(t, user.SecondName, "SecondNameSecondName", "they should be equal")
		assert.Equal(t, user.Email, "Email", "they should be equal")
		assert.Equal(t, user.Password, "PasswordPassword", "they should be equal")
		assert.Equal(t, user.PhoneNumber, "PhoneNumberPhoneNumber", "they should be equal")
		assert.Equal(t, user.Address, "AddressAddressAddress", "they should be equal")
		assert.Equal(t, user.ActiveUser, uint32(1), "they should be equal")
	}
}

func TestAllUsers(t *testing.T) {

	users, err := DbMng.AllUsers(uint64(40000))
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, len(users), "users.count should not be equal to zero")

	for _, user := range users {
		assert.Equal(t, user.UserId, uint64(1), "they should be equal")
		assert.Equal(t, user.UserType, uint32(1200), "they should be equal")
		assert.Equal(t, user.StockId, uint64(1234432), "they should be equal")
		assert.Equal(t, user.CompanyId, uint64(40000), "they should be equal")
		assert.Equal(t, user.UserImagePath, "UserImagePathUserImagePath", "they should be equal")
		assert.Equal(t, user.FirstName, "FirstNameFirstName", "they should be equal")
		assert.Equal(t, user.SecondName, "SecondNameSecondName", "they should be equal")
		assert.Equal(t, user.Email, "Email", "they should be equal")
		assert.Equal(t, user.Password, "PasswordPassword", "they should be equal")
		assert.Equal(t, user.PhoneNumber, "PhoneNumberPhoneNumber", "they should be equal")
		assert.Equal(t, user.Address, "AddressAddressAddress", "they should be equal")
		assert.Equal(t, user.ActiveUser, uint32(1), "they should be equal")
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < 1000; i++ {
		go func(k int, dBMan *models.DbManager) {

			user := new(dbmodels.UserRequest)

			user.UserType = uint32(1200)
			user.UserImagePath = "UserImagePath"
			user.FirstName = fmt.Sprintf("FirstName is %d", k)
			user.SecondName = fmt.Sprintf("SecondName is %d", k)
			user.Email = fmt.Sprintf("Email is %d", k)
			user.Password = fmt.Sprintf("Password is %d", k)
			user.PhoneNumber = "PhoneNumber"
			user.Address = "Address"

			//var companyId uint64
			//companyId = 20000 + uint64(k)

			userId, err := dBMan.CreateUser(user)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(userId)
			}
		}(i, DbMng)
	}
	time.Sleep(2 * 1e9)
}
