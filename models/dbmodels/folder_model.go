package dbmodels

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type FolderFilter struct {
	CompanyID         sql.NullInt64 `db:"company_id"`
	SupplierKeyword   string        `protobuf:"bytes,1,opt,name=supplierKeyword,proto3" json:"supplierKeyword"`
	SupplierUpdatedAt uint64        `protobuf:"varint,2,opt,name=supplierUpdatedAt,proto3" json:"supplierUpdatedAt"`
}

type FolderModel struct {
	FolderID         uint64 `db:"folder_id"  json:"folder_id"`
	CompanyID        uint64 `db:"company_id"  json:"company_id"`
	DocmodelID       uint64 `db:"docmodel_id"  json:"docmodel_id"`
	WordCount        uint64 `db:"word_count"  json:"word_count"`
	ParsedImageCount uint64 `db:"parsed_image_count"  json:"parsed_image_count"`
	FolderImagePath  string `db:"folder_image_path"  json:"folder_image_path"`
	FolderName       string `db:"folder_name"  json:"folder_name"`
	ContactFname     string `db:"contact_fname"  json:"contact_fname"`
	PhoneNumber      string `db:"phone_number"  json:"phone_number"`
	Address          string `db:"address"  json:"address"`
	UpdatedAt        uint64 `db:"updated_at"  json:"updated_at"`
}

// type FolderModel struct {
// 	FolderID        int            `db:"folder_id"`
// 	CompanyID       sql.NullInt64  `db:"company_id"`
// 	StockID         sql.NullInt64  `db:"stock_id"`
// 	FolderImagePath sql.NullString `db:"folder_image_path"`
// 	CompanyName     sql.NullString `db:"company_name"`
// 	ContactFname    sql.NullString `db:"contact_fname"`
// 	PhoneNumber     sql.NullString `db:"phone_number"`
// 	Address         sql.NullString `db:"address"`
// 	UpdatedAt       sql.NullInt64  `db:"updated_at"`
// }

// CREATE TABLE IF NOT EXISTS folders (
//     folder_id BIGSERIAL PRIMARY KEY NOT NULL,
//     company_id BIGINT,
//     docmodel_id BIGINT,
//     word_count BIGINT,
//     parsed_image_count BIGINT,
//     folder_image_path varchar (400),
//     folder_name varchar (400),
//     contact_fname varchar (400),
//     phone_number varchar (400),
//     address varchar (400),
//     updated_at BIGINT
// );

// CREATE INDEX IF NOT EXISTS company_id_folders_idx ON folders (company_id);
// CREATE INDEX IF NOT EXISTS word_count_folders_idx ON folders (word_count);
// CREATE INDEX IF NOT EXISTS docmodel_id_folders_idx ON folders (docmodel_id);

func StoreFolderModel(tx *sqlx.Tx, folderRequest *FolderModel, companyId uint64) (uint64, error) {

	var lastInsertId uint64
	err := tx.QueryRow("INSERT INTO folders("+
		"company_id, "+
		"docmodel_id, "+
		"word_count, "+
		"parsed_image_count, "+
		"folder_image_path, "+
		"folder_name, "+
		"contact_fname, "+
		"phone_number, "+
		"address, "+
		"updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning folder_id;",
		companyId,
		folderRequest.DocmodelID,
		folderRequest.WordCount,
		folderRequest.ParsedImageCount,
		folderRequest.FolderImagePath,
		folderRequest.FolderName,
		folderRequest.ContactFname,
		folderRequest.PhoneNumber,
		folderRequest.Address,
		UpdatedAt(),
	).Scan(&lastInsertId)

	if err != nil {
		return ErrorFunc(err)
	}

	return lastInsertId, nil
}

func UpdateFolderModel(tx *sqlx.Tx, folderRequest *FolderModel, companyId uint64) (uint64, error) {

	stmt, err := tx.Prepare("UPDATE folders SET " +
		"company_id=$1, " +
		"docmodel_id=$2, " +
		"word_count=$3, " +
		"parsed_image_count=$4, " +
		"folder_image_path=$5, " +
		"folder_name=$6, " +
		"contact_fname=$7, " +
		"phone_number=$8, " +
		"address=$9, " +
		"updated_at=$10 WHERE folder_id=$11")

	if err != nil {
		return ErrorFunc(err)
	}

	res, err := stmt.Exec(
		companyId,
		folderRequest.DocmodelID,
		folderRequest.WordCount,
		folderRequest.ParsedImageCount,
		folderRequest.FolderImagePath,
		folderRequest.FolderName,
		folderRequest.ContactFname,
		folderRequest.PhoneNumber,
		folderRequest.Address,
		UpdatedAt(),
		folderRequest.FolderID,
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

var selectFolderModelRow string = "folder_id, " +
	"docmodel_id, " +
	"word_count, " +
	"parsed_image_count, " +
	"folder_image_path, " +
	"folder_name, " +
	"contact_fname, " +
	"phone_number, " +
	"address "

func scanFolderModelRows(rows *sqlx.Rows) ([]*FolderModel, error) {
	suppliers := make([]*FolderModel, 0)
	for rows.Next() {
		supplier := new(FolderModel)
		err := rows.Scan(
			&supplier.FolderID,
			&supplier.DocmodelID,
			&supplier.WordCount,
			&supplier.ParsedImageCount,
			&supplier.FolderImagePath,
			&supplier.FolderName,
			&supplier.ContactFname,
			&supplier.PhoneNumber,
			&supplier.Address,
		)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return suppliers, nil
}

func AllFoldersForCompany(db *sqlx.DB, companyId uint64) ([]*FolderModel, error) {

	rows, err := db.Queryx("SELECT "+
		selectFolderModelRow+
		"FROM folders WHERE company_id=$1 ORDER BY folder_id ASC", companyId)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("")
		return nil, err
	}

	suppliers, err := scanFolderModelRows(rows)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("")
		return nil, err
	}

	return suppliers, nil
}

func AllFoldersForDoc(db *sqlx.DB, docModelID uint64) ([]*FolderModel, error) {

	rows, err := db.Queryx("SELECT "+
		selectFolderModelRow+
		"FROM folders WHERE docmodel_id=$1 ORDER BY folder_id ASC", docModelID)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("")
		return nil, err
	}

	suppliers, err := scanFolderModelRows(rows)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("")
		return nil, err
	}

	return suppliers, nil
}

func FolderForFolderId(db *sqlx.DB, folderModelID uint64) (*FolderModel, error) {
	rows, err := db.Queryx("SELECT "+
		selectFolderModelRow+
		"FROM folders WHERE folder_id=$1 ORDER BY folder_id ASC", folderModelID)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("")
		return nil, err
	}

	folders, err := scanFolderModelRows(rows)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("")
		return nil, err
	}

	if len(folders) > 0 {
		return folders[len(folders)-1], nil
	} else {
		return nil, nil
	}
}

func AllFoldersForUpdate(db *sqlx.DB, filter *FolderFilter, companyId uint64) ([]*FolderModel, error) {

	rows, err := db.Queryx("SELECT "+
		selectFolderModelRow+
		"FROM folders WHERE updated_at >= $1 AND company_id=$2 LIMIT $3", filter.SupplierUpdatedAt, companyId, 1000)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("")
		print("error")
	}

	suppliers, err := scanFolderModelRows(rows)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("")
		return nil, err
	}

	return suppliers, nil
}

func DeleteFolder(tx *sqlx.Tx, companyId uint64) (uint64, error) {

	stmt, err := tx.Prepare("DELETE FROM folders WHERE company_id = $1")
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
