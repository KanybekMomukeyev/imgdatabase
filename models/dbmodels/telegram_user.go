package dbmodels

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TelegramUserFilter struct {
	TelegramUserKeyword   string `protobuf:"bytes,1,opt,name=customerKeyword,proto3" json:"customerKeyword"`
	TelegramUserUpdatedAt uint64 `protobuf:"varint,2,opt,name=customerUpdatedAt,proto3" json:"customerUpdatedAt"`
}

type TelegramUser struct {
	TelegramUserID  uint64 `db:"telegram_user_id" json:"telegram_user_id"`
	UserID          uint64 `db:"user_id"  json:"user_id"`
	CompanyID       uint64 `db:"company_id" json:"company_id"`
	TelegramID      uint64 `db:"telegram_id" json:"telegram_id"`
	FirstName       string `db:"first_name" json:"first_name"`
	SecondName      string `db:"second_name" json:"second_name"`
	PhoneNumber     string `db:"phone_number" json:"phone_number"`
	TelegramAccount string `db:"tegram_account" json:"tegram_account"`
}

// CREATE TABLE IF NOT EXISTS telegram_users (
//     telegram_user_id BIGSERIAL PRIMARY KEY NOT NULL,
//     user_id BIGINT,
//     company_id BIGINT,
//     telegram_id BIGINT,
//     first_name varchar (400),
//     second_name varchar (400),
//     phone_number varchar (400),
//     tegram_account varchar (400),
//     updated_at BIGINT
// );

// CREATE INDEX IF NOT EXISTS company_id_tegram_users_idx ON telegram_users (company_id);

var selectTelegramUserRow string = "telegram_user_id, " +
	"user_id, " +
	"company_id, " +
	"telegram_id, " +
	"first_name, " +
	"second_name, " +
	"phone_number, " +
	"tegram_account "

func scanTelegramUserRow(rows *sqlx.Rows) ([]*TelegramUser, error) {
	telegramUsers := make([]*TelegramUser, 0)
	for rows.Next() {
		telegramUser := new(TelegramUser)
		err := rows.Scan(
			&telegramUser.TelegramUserID,
			&telegramUser.UserID,
			&telegramUser.CompanyID,
			&telegramUser.TelegramID,
			&telegramUser.FirstName,
			&telegramUser.SecondName,
			&telegramUser.PhoneNumber,
			&telegramUser.TelegramAccount,
		)

		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("")
			return nil, err
		}
		telegramUsers = append(telegramUsers, telegramUser)
	}
	return telegramUsers, nil
}

func StoreTelegramUser(tx *sqlx.Tx, telegramUser *TelegramUser) (uint64, error) {

	var lastInsertId uint64
	err := tx.QueryRow("INSERT INTO telegram_users("+
		"user_id, "+
		"company_id, "+
		"telegram_id, "+
		"first_name, "+
		"second_name, "+
		"phone_number, "+
		"tegram_account, "+
		"updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) returning telegram_user_id;",
		telegramUser.UserID,
		telegramUser.CompanyID,
		telegramUser.TelegramID,
		telegramUser.FirstName,
		telegramUser.SecondName,
		telegramUser.PhoneNumber,
		telegramUser.TelegramAccount,
		UpdatedAt(),
	).Scan(&lastInsertId)

	if err != nil {
		return ErrorFunc(err)
	}

	return lastInsertId, nil
}

func UpdateTelegramUser(tx *sqlx.Tx, telegramUser *TelegramUser) (uint64, error) {

	stmt, err := tx.Prepare("UPDATE telegram_users SET " +
		"user_id=$1, " +
		"company_id=$2, " +
		"telegram_id=$3, " +
		"first_name=$4, " +
		"second_name=$5, " +
		"phone_number=$6, " +
		"tegram_account=$7, " +
		"updated_at=$8 " +
		"WHERE telegram_user_id=$9")

	if err != nil {
		return ErrorFunc(err)
	}

	res, err := stmt.Exec(
		telegramUser.UserID,
		telegramUser.CompanyID,
		telegramUser.TelegramID,
		telegramUser.FirstName,
		telegramUser.SecondName,
		telegramUser.PhoneNumber,
		telegramUser.TelegramAccount,
		UpdatedAt(),
		telegramUser.TelegramUserID,
	)

	if err != nil {
		return ErrorFunc(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return ErrorFunc(err)
	}

	return uint64(affect), nil
}

func AllTelegramUsersForCompany(db *sqlx.DB, companyID uint64) ([]*TelegramUser, error) {

	rows, err := db.Queryx("SELECT "+
		selectTelegramUserRow+
		"FROM telegram_users WHERE company_id=$1 ORDER BY tegram_account ASC", companyID)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	customers, err := scanTelegramUserRow(rows)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func TelegramUserForTelegramID(db *sqlx.DB, telegramID uint64) (*TelegramUser, error) {

	rows, err := db.Queryx("SELECT "+
		selectTelegramUserRow+
		"FROM telegram_users WHERE telegram_id=$1 LIMIT 1", telegramID)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	tgUsers, err := scanTelegramUserRow(rows)
	if err != nil {
		return nil, err
	}
	if len(tgUsers) > 0 {
		return tgUsers[len(tgUsers)-1], nil
	} else {
		return nil, nil
	}
}

func TelegramUserForTelegramUserID(db *sqlx.DB, telegramUserID uint64) (*TelegramUser, error) {

	rows, err := db.Queryx("SELECT "+
		selectTelegramUserRow+
		"FROM telegram_users WHERE telegram_user_id=$1 LIMIT 1", telegramUserID)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	tgUsers, err := scanTelegramUserRow(rows)
	if err != nil {
		return nil, err
	}
	if len(tgUsers) > 0 {
		return tgUsers[len(tgUsers)-1], nil
	} else {
		return nil, nil
	}
}

func AllUpdatedTelegramUserForCompany(db *sqlx.DB, telFilter *TelegramUserFilter, companyID uint64) ([]*TelegramUser, error) {

	rows, err := db.Queryx("SELECT "+
		selectTelegramUserRow+
		"FROM telegram_users WHERE updated_at >= $1 AND cutted_image_id = $2 LIMIT $3", telFilter.TelegramUserUpdatedAt, companyID, 1000)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	customers, err := scanTelegramUserRow(rows)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func DeleteTelegramUsers(tx *sqlx.Tx, TelegramUserID uint64) (uint64, error) {

	stmt, err := tx.Prepare("DELETE FROM telegram_users WHERE telegram_user_id = $1")
	if err != nil {
		return ErrorFunc(err)
	}

	res, err := stmt.Exec(TelegramUserID)
	if err != nil {
		return ErrorFunc(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return ErrorFunc(err)
	}

	return uint64(affect), nil
}
