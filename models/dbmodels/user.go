package dbmodels

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type UserFilter struct {
	UserKeyword   string `protobuf:"bytes,1,opt,name=userKeyword,proto3" json:"userKeyword"`
	UserUpdatedAt uint64 `protobuf:"varint,2,opt,name=userUpdatedAt,proto3" json:"userUpdatedAt"`
}

type UserRequest struct {
	UserId        uint64 `protobuf:"varint,1,opt,name=userId" json:"userId"`
	UserUUID      string `protobuf:"bytes,2,opt,name=userUUID" json:"userUUID"`
	CompanyId     uint64 `protobuf:"varint,3,opt,name=companyId" json:"companyId"`
	StockId       uint64 `protobuf:"varint,4,opt,name=stockId" json:"stockId"`
	UserType      uint32 `protobuf:"varint,5,opt,name=userType" json:"userType"`
	UserImagePath string `protobuf:"bytes,6,opt,name=userImagePath" json:"userImagePath"`
	FirstName     string `protobuf:"bytes,7,opt,name=firstName" json:"firstName"`
	SecondName    string `protobuf:"bytes,8,opt,name=secondName" json:"secondName"`
	Email         string `protobuf:"bytes,9,opt,name=email" json:"email"`
	Password      string `protobuf:"bytes,10,opt,name=password" json:"password"`
	PhoneNumber   string `protobuf:"bytes,11,opt,name=phoneNumber" json:"phoneNumber"`
	Address       string `protobuf:"bytes,12,opt,name=address" json:"address"`
	ActiveUser    uint32 `protobuf:"varint,13,opt,name=activeUser" json:"activeUser"`
}

const (
	USER_TYPE_ADMIN       AcceptStatus = 11000
	USER_TYPE_SIMPLE      AcceptStatus = 55000
	USER_TYPE_OWNER       AcceptStatus = 88000
	USER_TYPE_CAN_ADD_DOC AcceptStatus = 66000
)

// CREATE TABLE IF NOT EXISTS users (
//     user_id BIGSERIAL PRIMARY KEY NOT NULL,
//     user_uuid VARCHAR (400) UNIQUE,
//     user_type INTEGER,
//     user_image_path VARCHAR (400),
//     first_name VARCHAR (400),
//     second_name VARCHAR (400),
//     email VARCHAR (400) UNIQUE,
//     password VARCHAR (400),
//     phone_number VARCHAR (400),
//     address VARCHAR (400),
//     updated_at BIGINT
// );

func UpdatedAt() uint64 {
	updatedAt := (time.Now().UnixNano() / 1000000)
	return uint64(updatedAt)
}

func ErrorFunc(err error) (uint64, error) {
	log.WithFields(log.Fields{"error": err}).Fatal("SQLX QueryRow breaks")
	panic(err)
	return 0, err
}

func StoreUser(tx *sqlx.Tx, user *UserRequest) (uint64, error) {

	updatedAt := (time.Now().UnixNano() / 1000000)
	var lastInsertId uint64
	err := tx.QueryRow("INSERT INTO users "+
		"(user_uuid, "+
		"user_type, "+
		"user_image_path, "+
		"first_name, "+
		"second_name, "+
		"email, "+
		"password, "+
		"phone_number, "+
		"address, "+
		"company_id, "+
		"stock_id, "+
		"active_user, "+
		"updated_at) "+
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) returning user_id;",
		user.UserUUID,
		user.UserType,
		user.UserImagePath,
		user.FirstName,
		user.SecondName,
		user.Email,
		user.Password,
		user.PhoneNumber,
		user.Address,
		user.CompanyId,
		user.StockId,
		1,
		updatedAt).Scan(&lastInsertId)

	if err != nil {
		return ErrorFunc(err)
	}

	return lastInsertId, nil
}

var selectUserRow string = "user_id, " +
	"company_id, " +
	"stock_id, " +
	"user_uuid, " +
	"user_type, " +
	"user_image_path, " +
	"first_name, " +
	"second_name, " +
	"email, " +
	"password, " +
	"phone_number, " +
	"active_user, " +
	"address "

func scanUserRow(rows *sqlx.Rows) ([]*UserRequest, error) {
	users := make([]*UserRequest, 0)
	for rows.Next() {
		user := new(UserRequest)
		err := rows.Scan(
			&user.UserId,
			&user.CompanyId,
			&user.StockId,
			&user.UserUUID,
			&user.UserType,
			&user.UserImagePath,
			&user.FirstName,
			&user.SecondName,
			&user.Email,
			&user.Password,
			&user.PhoneNumber,
			&user.ActiveUser,
			&user.Address)

		if err != nil {
			log.WithFields(log.Fields{"scanUserRow": err}).Warn("ERROR")
			break
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(tx *sqlx.Tx, user *UserRequest) (uint64, error) {

	stmt, err := tx.Prepare("UPDATE users SET " +
		"user_type=$1, " +
		"user_uuid=$2, " +
		"company_id=$3, " +
		"stock_id=$4, " +
		"user_image_path=$5, " +
		"first_name=$6, " +
		"second_name=$7, " +
		"password=$8, " +
		"phone_number=$9, " +
		"address=$10, " +
		"updated_at=$11 " +
		"WHERE user_id=$12")

	if err != nil {
		return ErrorFunc(err)
	}

	res, err := stmt.Exec(user.UserType,
		user.UserUUID,
		user.CompanyId,
		user.StockId,
		user.UserImagePath,
		user.FirstName,
		user.SecondName,
		user.Password,
		user.PhoneNumber,
		user.Address,
		UpdatedAt(),
		user.UserId)

	if err != nil {
		return ErrorFunc(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return ErrorFunc(err)
	}

	log.WithFields(log.Fields{"update user rows changed": affect}).Info("")
	return uint64(affect), nil
}

func AllUsersForCompany(db *sqlx.DB, company_id uint64) ([]*UserRequest, error) {

	rows, err := db.Queryx("SELECT "+
		selectUserRow+
		"FROM users WHERE company_id = $1", company_id)

	if err != nil {
		log.WithFields(log.Fields{"Queryx": err}).Warn("ERROR")
		return nil, err
	}

	users, err := scanUserRow(rows)

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func AllUsersForUpdate(db *sqlx.DB, filter *UserFilter, company_id uint64) ([]*UserRequest, error) {

	rows, err := db.Queryx("SELECT "+
		selectUserRow+
		"FROM users WHERE updated_at >= $1 AND company_id = $2 LIMIT $3", filter.UserUpdatedAt, company_id, 3000)

	if err != nil {
		log.WithFields(log.Fields{"Queryx": err}).Warn("ERROR")
		return nil, err
	}

	users, err := scanUserRow(rows)

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func UserForEmail(db *sqlx.DB, email string) (*UserRequest, error) {

	rows, err := db.Queryx("SELECT "+
		selectUserRow+
		"FROM users WHERE email=$1", email)

	if err != nil {
		return nil, err
	}

	users, err := scanUserRow(rows)
	if len(users) > 0 {
		return users[0], nil
	}

	return nil, errors.New("No such user for email")
}

func UserForUUID(db *sqlx.DB, uuid string) (*UserRequest, error) {

	rows, err := db.Queryx("SELECT "+
		selectUserRow+
		"FROM users WHERE user_uuid = $1", uuid)

	if err != nil {
		return nil, err
	}

	users, err := scanUserRow(rows)
	if len(users) > 0 {
		return users[0], nil
	}

	return nil, errors.New("No such user for uuid")
}
