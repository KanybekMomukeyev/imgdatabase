package dbmodels

import (
	"errors"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type ImageFilter struct {
	CustomerKeyword   string `protobuf:"bytes,1,opt,name=customerKeyword,proto3" json:"customerKeyword"`
	CustomerUpdatedAt uint64 `protobuf:"varint,2,opt,name=customerUpdatedAt,proto3" json:"customerUpdatedAt"`
}

// type Cuttedimage struct {
// 	ImageID         int            `db:"image_id"  json:"imageId"`
// 	DocModelID      sql.NullInt64  `db:"docmodel_id"  json:"DocModelID"`
// 	CompanyID       sql.NullInt64  `db:"company_id"  json:"companyId"`
// 	FolderID        sql.NullInt64  `db:"folder_id"  json:"folder_id"`
// 	ParsedImagePath sql.NullString `db:"parsed_image_path"  json:"parsedImagePath"`
// 	FolderName      sql.NullString `db:"folder_name"  json:"folderName"`
// 	SecondName      sql.NullString `db:"second_name"  json:"secondName"`
// 	PhoneNumber     sql.NullString `db:"phone_number"  json:"phoneNumber"`
// 	Address         sql.NullString `db:"address"  json:"address"`
// 	UpdatedAt       sql.NullInt64  `db:"updated_at"  json:"updatedAt"`
// }

type CuttedImage struct {
	CuttedImageID    uint64 `db:"image_id"  json:"image_id"`
	CuttedImageState uint32 `db:"cutted_image_state" json:"cutted_image_state"` // DEFAULT VALUE 1000
	DocModelID       uint64 `db:"docmodel_id"  json:"docmodel_id"`
	CompanyID        uint64 `db:"company_id"  json:"company_id"`
	FolderID         uint64 `db:"folder_id"  json:"folder_id"`
	ParsedImagePath  string `db:"parsed_image_path"  json:"parsed_image_path"`
	FolderName       string `db:"folder_name"  json:"folder_name"`
	SecondName       string `db:"second_name"  json:"second_name"`
	PhoneNumber      string `db:"phone_number"  json:"phone_number"`
	Address          string `db:"address"  json:"address"`
	UpdatedAt        uint64 `db:"updated_at"  json:"updated_at"`
	WaitingCount     int    `json:"waiting_count"` // NOT IN DATABASE
}

type CuttedImageState uint32

const (
	CUTTED_IMAGE_NOT_TRANSLATED CuttedImageState = 1000
	CUTTED_IMAGE_TRANSLATED     CuttedImageState = 3000
	CUTTED_IMAGE_MISSTAKED      CuttedImageState = 5000
	CUTTED_IMAGE_MARKED_INVALID CuttedImageState = 9000
)

// CREATE TABLE IF NOT EXISTS cutted_images (
//     image_id BIGSERIAL PRIMARY KEY NOT NULL,
//     docmodel_id BIGINT,
//     cutted_image_state INTEGER DEFAULT 1000,
//     company_id BIGINT,
//     folder_id BIGINT,
//     parsed_image_path varchar (500),
//     folder_name varchar (500),
//     second_name varchar (500),
//     phone_number varchar (500),
//     address varchar (500),
//     updated_at BIGINT
// );

// CREATE INDEX IF NOT EXISTS company_id_cuttedimages_idx ON cutted_images (company_id);
// CREATE INDEX IF NOT EXISTS folder_id_cuttedimages_idx ON cutted_images (folder_id);
// CREATE INDEX IF NOT EXISTS docmodel_id_cuttedimages_idx ON cutted_images (docmodel_id);

func StoreCuttedImage(tx *sqlx.Tx, imageRequest *CuttedImage, companyId uint64) (uint64, error) {

	var lastInsertId uint64
	err := tx.QueryRow("INSERT INTO cutted_images("+
		"docmodel_id, "+
		"company_id, "+
		"folder_id, "+
		"parsed_image_path, "+
		"folder_name, "+
		"second_name, "+
		"phone_number, "+
		"address, "+
		"updated_at, "+
		"cutted_image_state) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning image_id;",
		imageRequest.DocModelID,
		companyId,
		imageRequest.FolderID,
		imageRequest.ParsedImagePath,
		imageRequest.FolderName,
		imageRequest.SecondName,
		imageRequest.PhoneNumber,
		imageRequest.Address,
		UpdatedAt(),
		imageRequest.CuttedImageState,
	).Scan(&lastInsertId)

	if err != nil {
		return ErrorFunc(err)
	}

	return lastInsertId, nil
}

func UpdateCuttedImage(tx *sqlx.Tx, imageRequest *CuttedImage, companyId uint64) (uint64, error) {

	stmt, err := tx.Prepare("UPDATE cutted_images SET " +
		"docmodel_id=$1, " +
		"company_id=$2, " +
		"folder_id=$3, " +
		"parsed_image_path=$4, " +
		"folder_name=$5, " +
		"second_name=$6, " +
		"phone_number=$7, " +
		"address=$8, " +
		"updated_at=$9, " +
		"cutted_image_state=$10 " +
		"WHERE image_id=$11")

	if err != nil {
		return ErrorFunc(err)
	}

	res, err := stmt.Exec(
		imageRequest.DocModelID,
		companyId,
		imageRequest.FolderID,
		imageRequest.ParsedImagePath,
		imageRequest.FolderName,
		imageRequest.SecondName,
		imageRequest.PhoneNumber,
		imageRequest.Address,
		UpdatedAt(),
		imageRequest.CuttedImageState,
		imageRequest.CuttedImageID,
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

var selectCuttedImageRow string = "image_id, " +
	"docmodel_id, " +
	"folder_id, " +
	"parsed_image_path, " +
	"folder_name, " +
	"second_name, " +
	"phone_number, " +
	"address, " +
	"cutted_image_state "

func AllCuttedImagesForCompany(db *sqlx.DB, companyId uint64) ([]*CuttedImage, error) {

	rows, err := db.Queryx("SELECT "+
		selectCuttedImageRow+
		"FROM cutted_images WHERE company_id=$1 ORDER BY folder_name ASC", companyId)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	customers, err := scanCuttedImageRow(rows)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func CountCuttedImagesForDoc(db *sqlx.DB, docModelID uint64) (int, error) {
	var stateCount int
	err := db.Get(&stateCount, "SELECT count(*) FROM cutted_images WHERE docmodel_id=$1", docModelID)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return 0, err
	}
	return stateCount, nil
}

func StatsCuttedImageForDoc(db *sqlx.DB, docModelID uint64, cuttedImageState CuttedImageState) (int, error) {
	var stateCount int
	err := db.Get(&stateCount, "SELECT count(*) FROM cutted_images WHERE docmodel_id=$1 AND cutted_image_state=$2", docModelID, uint32(cuttedImageState))
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return 0, err
	}
	return stateCount, nil
}

func CuttedImageForID(db *sqlx.DB, ImageID uint64) (*CuttedImage, error) {

	rows, err := db.Queryx("SELECT "+
		selectCuttedImageRow+
		"FROM cutted_images WHERE image_id=$1 LIMIT 1", ImageID)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}
	cuttedImages, err := scanCuttedImageRow(rows)
	if err != nil {
		return nil, err
	}

	if len(cuttedImages) > 0 {
		return cuttedImages[len(cuttedImages)-1], nil
	} else {
		return nil, nil
	}
}

func CuttedImagesForFolder(db *sqlx.DB, folderID uint64) ([]*CuttedImage, error) {

	rows, err := db.Queryx("SELECT "+
		selectCuttedImageRow+
		"FROM cutted_images WHERE folder_id=$1 ORDER BY folder_name ASC", folderID)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	customers, err := scanCuttedImageRow(rows)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func RandomNotTranslatedCuttedImageNOT_IN_ARRAY(db *sqlx.DB, imageIDs []uint64) (*CuttedImage, error) {

	if len(imageIDs) > 0 {
		query, args, err := sqlx.In("SELECT "+
			selectCuttedImageRow+
			"FROM cutted_images WHERE cutted_image_state=1000 AND image_id NOT IN (?) ORDER BY RANDOM() LIMIT 1", imageIDs)

		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("ERROR")
			return nil, err
		}

		query = sqlx.Rebind(sqlx.DOLLAR, query) // only if postgres
		rows, err := db.Queryx(query, args...)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("ERROR")
			return nil, err
		}

		cuttedImages, err := scanCuttedImageRow(rows)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("ERROR")
			return nil, err
		}

		if len(cuttedImages) > 0 {
			return cuttedImages[len(cuttedImages)-1], nil
		} else {
			return nil, errors.New("EMPTY")
		}
	} else {
		rows, err := db.Queryx("SELECT "+
			selectCuttedImageRow+
			"FROM cutted_images WHERE cutted_image_state=$1 ORDER BY RANDOM() LIMIT 1", uint32(CUTTED_IMAGE_NOT_TRANSLATED))

		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("")
			return nil, err
		}
		cuttedImages, err := scanCuttedImageRow(rows)
		if err != nil {
			return nil, err
		}

		if len(cuttedImages) > 0 {
			return cuttedImages[len(cuttedImages)-1], nil
		} else {
			return nil, nil
		}
	}
}

func RandomNotTranslatedCuttedImage__NOT_USED(db *sqlx.DB, companyID uint64, imageIDs []uint64) (*CuttedImage, error) {
	// SELECT column FROM table ORDER BY RANDOM() LIMIT 1

	rows, err := db.Queryx("SELECT "+
		selectCuttedImageRow+
		"FROM cutted_images WHERE cutted_image_state=$1 ORDER BY RANDOM() LIMIT 1", uint32(CUTTED_IMAGE_NOT_TRANSLATED))

	// rows, err := db.Queryx("SELECT " +
	// 	selectCuttedImageRow +
	// 	"FROM cutted_images WHERE cutted_image_state = 1000 ORDER BY RANDOM() LIMIT 1")

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}
	cuttedImages, err := scanCuttedImageRow(rows)
	if err != nil {
		return nil, err
	}

	if len(cuttedImages) > 0 {
		return cuttedImages[len(cuttedImages)-1], nil
	} else {
		return nil, nil
	}
}

func scanCuttedImageRow(rows *sqlx.Rows) ([]*CuttedImage, error) {
	customers := make([]*CuttedImage, 0)
	for rows.Next() {
		customer := new(CuttedImage)
		err := rows.Scan(
			&customer.CuttedImageID,
			&customer.DocModelID,
			&customer.FolderID,
			&customer.ParsedImagePath,
			&customer.FolderName,
			&customer.SecondName,
			&customer.PhoneNumber,
			&customer.Address,
			&customer.CuttedImageState,
		)

		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("")
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func AllUpdatedCuttedImages(db *sqlx.DB, custFilter *ImageFilter, companyId uint64) ([]*CuttedImage, error) {

	rows, err := db.Queryx("SELECT "+
		selectCuttedImageRow+
		"FROM cutted_images WHERE updated_at >= $1 AND company_id = $2 LIMIT $3", custFilter.CustomerUpdatedAt, companyId, 1000)

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("")
		return nil, err
	}

	customers, err := scanCuttedImageRow(rows)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func DeleteCuttedImages(tx *sqlx.Tx, companyId uint64) (uint64, error) {

	stmt, err := tx.Prepare("DELETE FROM cutted_images WHERE company_id = $1")
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
