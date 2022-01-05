package dbmodels

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type CuttedImageTranslateFilter struct {
	CuttedImageTranslatedKeyword   string `protobuf:"bytes,1,opt,name=customerKeyword,proto3" json:"customerKeyword"`
	CuttedImageTranslatedUpdatedAt uint64 `protobuf:"varint,2,opt,name=customerUpdatedAt,proto3" json:"customerUpdatedAt"`
}

type CuttedImageTranslate struct {
	CuttedImageTranslateID uint64 `db:"cutted_image_translate_id"  json:"cutted_image_translate_id"`
	CuttedImageID          uint64 `db:"cutted_image_id" json:"cutted_image_id"`
	TelegramUserID         uint64 `db:"telegram_user_id" json:"telegram_user_id"`
	TranslatedWord         string `db:"translated_word" json:"translated_word"`
	Comments               string `db:"comments" json:"comments"`
	Summary                string `db:"summary" json:"summary"`
	AcceptStatus           uint32 `db:"accept_status" json:"accept_status"`
	UserID                 uint64 `db:"user_id" json:"user_id"`
	// TSV                 tsvector `db:"tsv" json:"tsv"`
	UpdatedAt uint64 `db:"updated_at"  json:"updated_at"`
}

func (translate CuttedImageTranslate) LowercasedTranslatedWord() string {
	return strings.ToLower(translate.TranslatedWord)
}

type DocIDsSelectResponse struct {
	DocModelID uint64 `db:"docmodel_id" json:"docmodel_id"`
}

type TelegramUserIDsSelectResponse struct {
	TelegramUserID uint64 `db:"docmodel_id" json:"docmodel_id"`
}

type AcceptStatus uint32

const (
	TRANSLATE_ACCEPTED          AcceptStatus = 1000
	TRANSLATE_WAITING           AcceptStatus = 5000
	TRANSLATE_REJECTED          AcceptStatus = 8000
	TRANSLATE_MARKED_AS_INVALID AcceptStatus = 6000
)

// AcceptStatus For TelegramUser:  // see to rentau conversions
// WAITING
// ACCEPTED
// REJECTED

// CREATE TABLE IF NOT EXISTS cutted_image_translates (
//     cutted_image_translate_id BIGSERIAL PRIMARY KEY NOT NULL,
//     cutted_image_id BIGINT,
//     telegram_user_id BIGINT,
//     translated_word varchar (400),
//     comments varchar (400),
//     summary varchar (400),
//     accept_status INTEGER,
//     updated_at BIGINT,
//     user_id BIGINT,
// );

// CREATE INDEX IF NOT EXISTS cutted_image_id_cutted_image_translates_idx ON cutted_image_translates (cutted_image_id);
// CREATE INDEX IF NOT EXISTS telegram_user_id_cutted_image_translates_idx ON cutted_image_translates (telegram_user_id);

// ALTER TABLE cutted_image_translates ADD COLUMN IF NOT EXISTS tsv tsvector;
// UPDATE cutted_image_translates SET tsv = setweight(to_tsvector(translated_word), 'A');
// CREATE INDEX cutted_image_translates_tsv ON cutted_image_translates USING GIN(tsv);

var selectCuttedImageTranslateRow string = "cutted_image_translate_id, " +
	"cutted_image_id, " +
	"telegram_user_id, " +
	"translated_word, " +
	"comments, " +
	"summary, " +
	"updated_at, " +
	"user_id, " +
	"accept_status "

func scanCuttedImageTranslateRow(rows *sqlx.Rows) ([]*CuttedImageTranslate, error) {
	cuttedImageTranslates := make([]*CuttedImageTranslate, 0)
	for rows.Next() {
		cuttedImageTranslate := new(CuttedImageTranslate)
		err := rows.Scan(
			&cuttedImageTranslate.CuttedImageTranslateID,
			&cuttedImageTranslate.CuttedImageID,
			&cuttedImageTranslate.TelegramUserID,
			&cuttedImageTranslate.TranslatedWord,
			&cuttedImageTranslate.Comments,
			&cuttedImageTranslate.Summary,
			&cuttedImageTranslate.UpdatedAt,
			&cuttedImageTranslate.UserID,
			&cuttedImageTranslate.AcceptStatus)

		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("")
			return nil, err
		}
		cuttedImageTranslates = append(cuttedImageTranslates, cuttedImageTranslate)
	}
	return cuttedImageTranslates, nil
}

func StoreCuttedImageTranslate(tx *sqlx.Tx, cuttedImageTranslate *CuttedImageTranslate) (uint64, error) {

	var lastInsertId uint64
	err := tx.QueryRow("INSERT INTO cutted_image_translates("+
		"cutted_image_id, "+
		"telegram_user_id, "+
		"translated_word, "+
		"comments, "+
		"summary, "+
		"accept_status, "+
		"tsv, "+
		"user_id, "+
		"updated_at) VALUES($1, $2, $3, $4, $5, $6, to_tsvector($7), $8, $9) returning cutted_image_translate_id;",
		cuttedImageTranslate.CuttedImageID,
		cuttedImageTranslate.TelegramUserID,
		cuttedImageTranslate.TranslatedWord,
		cuttedImageTranslate.Comments,
		cuttedImageTranslate.Summary,
		cuttedImageTranslate.AcceptStatus,
		strings.ToLower(cuttedImageTranslate.TranslatedWord),
		cuttedImageTranslate.UserID,
		UpdatedAt(),
	).Scan(&lastInsertId)

	if err != nil {
		return ErrorFunc(err)
	}

	return lastInsertId, nil
}

func UpdateTranslateFTS(tx *sqlx.Tx, cuttedImageTranslate *CuttedImageTranslate) (uint64, error) {
	stmt, err := tx.Prepare("UPDATE cutted_image_translates SET " +
		"tsv=to_tsvector($1) WHERE cutted_image_translate_id=$2")

	if err != nil {
		return ErrorFunc(err)
	}
	res, err := stmt.Exec(
		strings.ToLower(cuttedImageTranslate.TranslatedWord),
		cuttedImageTranslate.CuttedImageTranslateID)

	if err != nil {
		return ErrorFunc(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return ErrorFunc(err)
	}
	return uint64(affect), nil
}

func UpdateCuttedImageTranslate(tx *sqlx.Tx, cuttedImageTranslate *CuttedImageTranslate) (uint64, error) {

	stmt, err := tx.Prepare("UPDATE cutted_image_translates SET " +
		"cutted_image_id=$1, " +
		"telegram_user_id=$2, " +
		"translated_word=$3, " +
		"comments=$4, " +
		"summary=$5, " +
		"accept_status=$6, " +
		"user_id=$7, " +
		"updated_at=$8 " +
		"WHERE cutted_image_translate_id=$9")

	if err != nil {
		return ErrorFunc(err)
	}

	res, err := stmt.Exec(
		cuttedImageTranslate.CuttedImageID,
		cuttedImageTranslate.TelegramUserID,
		cuttedImageTranslate.TranslatedWord,
		cuttedImageTranslate.Comments,
		cuttedImageTranslate.Summary,
		cuttedImageTranslate.AcceptStatus,
		cuttedImageTranslate.UserID,
		UpdatedAt(),
		cuttedImageTranslate.CuttedImageTranslateID,
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

func Translated_STATUSES_forTgUser(db *sqlx.DB, telegramUserID uint64, translateStatus AcceptStatus) ([]*CuttedImageTranslate, error) {
	rows, err := db.Queryx("SELECT "+
		selectCuttedImageTranslateRow+
		"FROM cutted_image_translates WHERE telegram_user_id=$1 AND accept_status=$2 ORDER BY updated_at DESC LIMIT $3", telegramUserID, uint32(translateStatus), 300)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}
	translates, err := scanCuttedImageTranslateRow(rows)
	if err != nil {
		return nil, err
	}
	return translates, nil
}

func LastActivityForTgUser(db *sqlx.DB, telegramUserID uint64) (*CuttedImageTranslate, error) {
	rows, err := db.Queryx("SELECT "+
		selectCuttedImageTranslateRow+
		"FROM cutted_image_translates WHERE telegram_user_id=$1 ORDER BY updated_at DESC LIMIT 1", telegramUserID)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}
	translates, err := scanCuttedImageTranslateRow(rows)
	if err != nil {
		return nil, err
	}

	if len(translates) > 0 {
		return translates[len(translates)-1], nil
	} else {
		return nil, nil
	}
}

// const (
// 	TRANSLATE_ACCEPTED AcceptStatus = 1000
// 	TRANSLATE_WAITING  AcceptStatus = 5000
// 	TRANSLATE_REJECTED AcceptStatus = 8000
// )
func Count_WAITING_TranslatesForCuttedImage(db *sqlx.DB, cuttedImageID uint64, translateStatus AcceptStatus) (int, error) {
	var stateCount int
	err := db.Get(&stateCount, "SELECT count(*) FROM cutted_image_translates WHERE cutted_image_id=$1 AND accept_status=$2", cuttedImageID, uint32(translateStatus))
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return 0, err
	}
	return stateCount, nil
}

func CountFor_Translated_STATUSES_FotTgUser(db *sqlx.DB, telegramUserID uint64, translateStatus AcceptStatus) (int, error) {

	log.WithFields(log.Fields{"telegramUserID %d": telegramUserID,
		"translateStatus %d":    translateStatus,
		"translateStatus+++ %d": translateStatus}).Info("CountFor_Translated_STATUSES_FotTgUser")

	var stateCount int
	err := db.Get(&stateCount, "SELECT count(*) FROM cutted_image_translates WHERE telegram_user_id=$1 AND accept_status=$2", telegramUserID, uint32(translateStatus))
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return 0, err
	}
	return stateCount, nil
}

func CuttedImageIDsForTelegramUser(db *sqlx.DB, telegramUserID uint64) ([]*CuttedImageTranslate, error) {

	rows, err := db.Queryx("SELECT cutted_image_id FROM cutted_image_translates WHERE telegram_user_id=$1", telegramUserID)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	cuttedImageTranslates := make([]*CuttedImageTranslate, 0)
	for rows.Next() {
		cuttedImageTranslate := new(CuttedImageTranslate)
		err := rows.Scan(&cuttedImageTranslate.CuttedImageID)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("")
			return nil, err
		}
		cuttedImageTranslates = append(cuttedImageTranslates, cuttedImageTranslate)
	}
	return cuttedImageTranslates, nil
}

// -------------------------------------------------------
type CustomTranslate struct {
	CuttedImageTranslateID uint64 `json:"cutted_image_translate_id"`
	CuttedImageID          uint64 `json:"cutted_image_id"`
	TelegramUserID         uint64 `json:"telegram_user_id"`
	TranslatedWord         string `json:"translated_word"`
	AcceptStatus           uint32 `json:"accept_status"`
	ParsedImagePath        string `json:"parsed_image_path"`
	UpdatedAt              uint64 `json:"updated_at"`
	UserID                 uint64 `json:"user_id"`
}

func TranslatesForDocModelAndAcceptedStatus(db *sqlx.DB, docModelID uint64, translateStatus AcceptStatus) ([]*CustomTranslate, error) {

	uniqTelegramUserIDs := `select d.cutted_image_translate_id, d.cutted_image_id, d.telegram_user_id, d.translated_word, d.accept_status, c.parsed_image_path, d.updated_at, d.user_id 
	from cutted_images as c
	join cutted_image_translates as d on
	c.image_id = d.cutted_image_id 
	WHERE c.docmodel_id = $1 and d.accept_status = $2 order by d.telegram_user_id`

	rows, err := db.Queryx(uniqTelegramUserIDs, docModelID, uint32(translateStatus))

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	customTranslates := make([]*CustomTranslate, 0)
	for rows.Next() {
		customTranslate := new(CustomTranslate)
		err := rows.Scan(
			&customTranslate.CuttedImageTranslateID,
			&customTranslate.CuttedImageID,
			&customTranslate.TelegramUserID,
			&customTranslate.TranslatedWord,
			&customTranslate.AcceptStatus,
			&customTranslate.ParsedImagePath,
			&customTranslate.UpdatedAt,
			&customTranslate.UserID)

		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("")
			return nil, err
		}
		customTranslates = append(customTranslates, customTranslate)
	}
	return customTranslates, nil
}

func TranslatesForDocModelAndTelegramUser(db *sqlx.DB, docModelID uint64, telegramUserID uint64) ([]*CustomTranslate, error) {

	uniqTelegramUserIDs := `select d.cutted_image_translate_id, d.cutted_image_id, d.telegram_user_id, d.translated_word, d.accept_status, c.parsed_image_path, d.updated_at, d.user_id 
	from cutted_images as c
	join cutted_image_translates as d on
	c.image_id = d.cutted_image_id 
	WHERE c.docmodel_id = $1 and d.telegram_user_id = $2 order by d.accept_status`

	rows, err := db.Queryx(uniqTelegramUserIDs, docModelID, telegramUserID)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	customTranslates := make([]*CustomTranslate, 0)
	for rows.Next() {
		customTranslate := new(CustomTranslate)
		err := rows.Scan(
			&customTranslate.CuttedImageTranslateID,
			&customTranslate.CuttedImageID,
			&customTranslate.TelegramUserID,
			&customTranslate.TranslatedWord,
			&customTranslate.AcceptStatus,
			&customTranslate.ParsedImagePath,
			&customTranslate.UpdatedAt,
			&customTranslate.UserID)

		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("")
			return nil, err
		}
		customTranslates = append(customTranslates, customTranslate)
	}
	return customTranslates, nil
}

func TelegramUserIdsDocModel(db *sqlx.DB, docModelID uint64) ([]*TelegramUserIDsSelectResponse, error) {

	uniqTelegramUserIDs := `select DISTINCT d.telegram_user_id   
	from cutted_images as c
	join cutted_image_translates as d on
	c.image_id = d.cutted_image_id 
	WHERE c.docmodel_id = $1`

	rows, err := db.Queryx(uniqTelegramUserIDs, docModelID)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	telegramIDsSelectResponses := make([]*TelegramUserIDsSelectResponse, 0)
	for rows.Next() {
		telegramIDsSelectResponse := new(TelegramUserIDsSelectResponse)
		err := rows.Scan(&telegramIDsSelectResponse.TelegramUserID)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("")
			return nil, err
		}
		telegramIDsSelectResponses = append(telegramIDsSelectResponses, telegramIDsSelectResponse)
	}
	return telegramIDsSelectResponses, nil
}

func DocIdsForTelegramUser(db *sqlx.DB, telegramUserID uint64) ([]*DocIDsSelectResponse, error) {

	uniqDocIDs := `select DISTINCT c.docmodel_id   
	from cutted_images as c
	join cutted_image_translates as d on
	c.image_id = d.cutted_image_id 
	WHERE d.telegram_user_id = $1`

	rows, err := db.Queryx(uniqDocIDs, telegramUserID)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	docIDsSelectResponses := make([]*DocIDsSelectResponse, 0)
	for rows.Next() {
		docIDsSelectResponse := new(DocIDsSelectResponse)
		err := rows.Scan(&docIDsSelectResponse.DocModelID)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("")
			return nil, err
		}
		docIDsSelectResponses = append(docIDsSelectResponses, docIDsSelectResponse)
	}
	return docIDsSelectResponses, nil
}

func FolderIdsForCuttedImageTranslate(db *sqlx.DB, cuttedImageTranslateIDs []uint64) ([]uint64, error) {

	folderIDs := make([]uint64, 0)
	for _, cuttedImageTranslateID := range cuttedImageTranslateIDs {
		uniqDocIDs := `select DISTINCT c.folder_id   
	from cutted_images as c
	join cutted_image_translates as d on
	c.image_id = d.cutted_image_id 
	WHERE d.cutted_image_translate_id = $1`

		rows, err := db.Queryx(uniqDocIDs, cuttedImageTranslateID)

		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("")
			return nil, err
		}

		for rows.Next() {
			var folderModelID uint64 = 0
			err := rows.Scan(&folderModelID)
			if err != nil {
				log.WithFields(log.Fields{"error": err}).Warn("")
				return nil, err
			}
			folderIDs = append(folderIDs, folderModelID)
		}
	}

	// ---------------------------------- //
	// HERE CLEAN ALL UNNEEDED DATA FROM IDS
	return make_unique_folder_ids(folderIDs), nil
}

func make_unique_folder_ids(intSlice []uint64) []uint64 {
	keys := make(map[uint64]bool)
	list := make([]uint64, 0, len(intSlice))

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// func TranslatesForCuttedImageIDs(db *sqlx.DB, cuttedImageIDs []uint64) ([]*CuttedImageTranslate, error) {
// 	folderIDs := make([]uint64, 0)
// 	for _, cuttedImageID := range cuttedImageIDs {
// 	}
// 	return nil, nil
// }

func CuttedImageTranslatesFor(db *sqlx.DB, cuttedImageID uint64) ([]*CuttedImageTranslate, error) {

	rows, err := db.Queryx("SELECT "+
		selectCuttedImageTranslateRow+
		"FROM cutted_image_translates WHERE cutted_image_id=$1 ORDER BY translated_word ASC", cuttedImageID)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	customers, err := scanCuttedImageTranslateRow(rows)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func AllUpdatedCuttedImageTranslatesForImage(db *sqlx.DB, custFilter *CuttedImageTranslateFilter, cuttedImageID uint64) ([]*CuttedImageTranslate, error) {

	rows, err := db.Queryx("SELECT "+
		selectCuttedImageTranslateRow+
		"FROM cutted_image_translates WHERE updated_at >= $1 AND cutted_image_id = $2 LIMIT $3", custFilter.CuttedImageTranslatedUpdatedAt, cuttedImageID, 1000)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	customers, err := scanCuttedImageTranslateRow(rows)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func AutoTranslatesForSearchKeyword(db *sqlx.DB, searchKey string) ([]*CuttedImageTranslate, error) {
	// # SELECT fieldNames FROM tableName WHERE to_tsvector(fieldName) @@ to_tsquery(conditions)

	lowerCased := strings.ToLower(searchKey)
	likePatternAfter := fmt.Sprintf("%s:*", lowerCased)
	// likePatternBefore := fmt.Sprintf(":*%s", lowerCased)
	// finalPattern := fmt.Sprintf("%s | %s", likePatternBefore, likePatternAfter)

	rows, err := db.Queryx("SELECT "+
		selectCuttedImageTranslateRow+
		"FROM cutted_image_translates WHERE tsv @@ to_tsquery($1) LIMIT 100", likePatternAfter)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	translates, err := scanCuttedImageTranslateRow(rows)
	if err != nil {
		return nil, err
	}
	return translates, nil
}

func FixedTranslatesForFirstSearchKeyword(db *sqlx.DB, firstSearchWord string) ([]*CuttedImageTranslate, error) {

	lowerCased := strings.ToLower(firstSearchWord)
	likePatternAfter := fmt.Sprintf("%s:*", lowerCased)

	rows, err := db.Queryx("SELECT "+
		selectCuttedImageTranslateRow+
		"FROM cutted_image_translates WHERE tsv @@ to_tsquery($1)", likePatternAfter)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	translates, err := scanCuttedImageTranslateRow(rows)
	if err != nil {
		return nil, err
	}
	return translates, nil
}

func AllTranslates(db *sqlx.DB) ([]*CuttedImageTranslate, error) {
	rows, err := db.Queryx("SELECT " + selectCuttedImageTranslateRow + "FROM cutted_image_translates LIMIT 1000000")
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}
	translates, err := scanCuttedImageTranslateRow(rows)
	if err != nil {
		return nil, err
	}
	return translates, nil
}

func DeleteCuttedImageTranslates(tx *sqlx.Tx, CuttedImageTranslateID uint64) (uint64, error) {

	stmt, err := tx.Prepare("DELETE FROM cutted_image_translates WHERE cutted_image_translate_id = $1")
	if err != nil {
		return ErrorFunc(err)
	}

	res, err := stmt.Exec(CuttedImageTranslateID)
	if err != nil {
		return ErrorFunc(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return ErrorFunc(err)
	}

	return uint64(affect), nil
}
