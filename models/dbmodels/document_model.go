package dbmodels

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type DocFilter struct {
	CustomerKeyword   string `protobuf:"bytes,1,opt,name=customerKeyword,proto3" json:"customerKeyword"`
	CustomerUpdatedAt uint64 `protobuf:"varint,2,opt,name=customerUpdatedAt,proto3" json:"customerUpdatedAt"`
}

type DocModel struct {
	DocmodelID       uint64 `db:"docmodel_id"  json:"docmodel_id"`
	CompanyID        uint64 `db:"company_id" json:"company_id"`
	UserID           uint64 `db:"user_id"  json:"user_id"`
	ParsedImageCount uint32 `db:"parsed_image_count"  json:"parsed_image_count"`
	StatCount        uint32 `db:"stat_count"  json:"stat_count"`
	DocmodelName     string `db:"docmodel_name"  json:"docmodel_name"`
	Summary          string `db:"summary"  json:"summary"`
	Comments         string `db:"comments"  json:"comments"`
	Descriptionn     string `db:"descriptionn"  json:"descriptionn"`
	UpdatedAt        uint64 `db:"updated_at"  json:"updated_at"`
}

// CREATE TABLE IF NOT EXISTS docmodels (
//     docmodel_id BIGSERIAL PRIMARY KEY NOT NULL,
//     company_id BIGINT,
//     user_id BIGINT,
//     parsed_image_count INTEGER,
//     stat_count INTEGER,
//     docmodel_name varchar (400),
//     summary varchar (400),
//     comments varchar (400),
//     descriptionn varchar (400),
//     updated_at BIGINT
// );

// CREATE INDEX IF NOT EXISTS company_id_docmodels_idx ON docmodels (company_id);
// CREATE INDEX IF NOT EXISTS user_id_docmodels_idx ON docmodels (user_id);

func StoreDocModel(tx *sqlx.Tx, docRequest *DocModel, companyId uint64) (uint64, error) {

	var lastInsertId uint64
	err := tx.QueryRow("INSERT INTO docmodels("+
		"company_id, "+
		"user_id, "+
		"parsed_image_count, "+
		"stat_count, "+
		"docmodel_name, "+
		"summary, "+
		"comments, "+
		"descriptionn, "+
		"updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) returning docmodel_id;",
		companyId,
		docRequest.UserID,
		docRequest.ParsedImageCount,
		docRequest.StatCount,
		docRequest.DocmodelName,
		docRequest.Summary,
		docRequest.Comments,
		docRequest.Descriptionn,
		UpdatedAt(),
	).Scan(&lastInsertId)

	if err != nil {
		return ErrorFunc(err)
	}

	return lastInsertId, nil
}

func DocModelUpdate(tx *sqlx.Tx, docRequest *DocModel, companyId uint64) (uint64, error) {

	stmt, err := tx.Prepare("UPDATE docmodels SET " +
		"company_id=$1, " +
		"user_id=$2, " +
		"parsed_image_count=$3, " +
		"stat_count=$4, " +
		"docmodel_name=$5, " +
		"summary=$6, " +
		"comments=$7, " +
		"descriptionn=$8, " +
		"updated_at=$9 " +
		"WHERE docmodel_id=$10")

	if err != nil {
		return ErrorFunc(err)
	}

	res, err := stmt.Exec(
		companyId,
		docRequest.UserID,
		docRequest.ParsedImageCount,
		docRequest.StatCount,
		docRequest.DocmodelName,
		docRequest.Summary,
		docRequest.Comments,
		docRequest.Descriptionn,
		UpdatedAt(),
		docRequest.DocmodelID,
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

var selectDocModelRow string = "docmodel_id, " +
	"company_id, " +
	"user_id, " +
	"parsed_image_count, " +
	"stat_count, " +
	"docmodel_name, " +
	"summary, " +
	"comments, " +
	"updated_at, " +
	"descriptionn "

func AllDocModelsForCompany(db *sqlx.DB, companyId uint64) ([]*DocModel, error) {

	rows, err := db.Queryx("SELECT "+
		selectDocModelRow+
		"FROM docmodels WHERE company_id=$1 ORDER BY docmodel_name ASC", companyId)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	customers, err := scanDocModelRow(rows)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func DocForDocModelID(db *sqlx.DB, docModelID uint64) (*DocModel, error) {

	rows, err := db.Queryx("SELECT "+
		selectDocModelRow+
		"FROM docmodels WHERE docmodel_id=$1 LIMIT 1", docModelID)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}
	cuttedImages, err := scanDocModelRow(rows)
	if err != nil {
		return nil, err
	}

	if len(cuttedImages) > 0 {
		return cuttedImages[len(cuttedImages)-1], nil
	} else {
		return nil, nil
	}
}

func DocModelForUserID(db *sqlx.DB, userID uint64) (*DocModel, error) {

	rows, err := db.Queryx("SELECT "+
		selectDocModelRow+
		"FROM docmodels WHERE user_id=$1 LIMIT 1", userID)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}
	cuttedImages, err := scanDocModelRow(rows)
	if err != nil {
		return nil, err
	}

	if len(cuttedImages) > 0 {
		return cuttedImages[len(cuttedImages)-1], nil
	} else {
		return nil, nil
	}
}

func scanDocModelRow(rows *sqlx.Rows) ([]*DocModel, error) {
	customers := make([]*DocModel, 0)
	for rows.Next() {
		docModel := new(DocModel)
		err := rows.Scan(
			&docModel.DocmodelID,
			&docModel.CompanyID,
			&docModel.UserID,
			&docModel.ParsedImageCount,
			&docModel.StatCount,
			&docModel.DocmodelName,
			&docModel.Summary,
			&docModel.Comments,
			&docModel.UpdatedAt,
			&docModel.Descriptionn,
		)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("")
			return nil, err
		}
		customers = append(customers, docModel)
	}
	return customers, nil
}

func DocModelDelete(tx *sqlx.Tx, companyId uint64) (uint64, error) {

	stmt, err := tx.Prepare("DELETE FROM docmodels WHERE company_id = $1")
	if err != nil {
		return ErrorFunc(err)
	}

	res, err := stmt.Exec(companyId)
	if err != nil {
		return ErrorFunc(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return ErrorFunc(err)
	}

	return uint64(affect), nil
}
