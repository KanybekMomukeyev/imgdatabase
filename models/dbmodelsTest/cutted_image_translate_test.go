package dbmodelsTest

import (
	"fmt"
	"testing"

	dbmodels "github.com/KanybekMomukeyev/imgdatabase/models/dbmodels"
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

func TestCreateCuttedImageTranslate(t *testing.T) {

	cuttedimagetranslate := new(dbmodels.CuttedImageTranslate)

	cuttedimagetranslate.CuttedImageID = uint64(12432)
	cuttedimagetranslate.TelegramUserID = uint64(1243234)
	cuttedimagetranslate.TranslatedWord = "Translated Word KANO MOmuke"
	cuttedimagetranslate.Comments = "CommentsCommentsComments"
	cuttedimagetranslate.Summary = "SummarySummarySummary"
	cuttedimagetranslate.AcceptStatus = uint32(1200)

	cuttedimageId, err := DbMng.CreateCuttedImageTranslate(cuttedimagetranslate)

	assert.Equal(t, cuttedimageId, uint64(1), "they should be equal")
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, cuttedimageId, "userId  should not be equal to zero")
	if assert.NotNil(t, cuttedimagetranslate) {
	}

	cuttedImageTranslates, err := DbMng.CuttedImageTranslatesFor(uint64(12432))
	assert.NotEqual(t, 0, len(cuttedImageTranslates), "users.count should not be equal to zero")

	for _, cuttedImageTranslate := range cuttedImageTranslates {
		assert.Equal(t, cuttedImageTranslate.CuttedImageTranslateID, uint64(1), "they should be equal")
		assert.Equal(t, cuttedImageTranslate.CuttedImageID, uint64(12432), "they should be equal")
		assert.Equal(t, cuttedImageTranslate.TelegramUserID, uint64(1243234), "CategoryName")
		assert.Equal(t, cuttedImageTranslate.TranslatedWord, "Translated Word KANO MOmuke", "TranslatedWordKANOOO")
		assert.Equal(t, cuttedImageTranslate.Comments, "CommentsCommentsComments", "Comments")
		assert.Equal(t, cuttedImageTranslate.Summary, "SummarySummarySummary", "CategoryName")
		assert.Equal(t, cuttedImageTranslate.AcceptStatus, uint32(1200), "CategoryName")
	}

	respAffected, _ := DbMng.AutoMigrateTranslates()
	assert.Equal(t, uint64(1000), respAffected, "respAffected")

	translates, err := DbMng.TranslatesForSearchKeyword("Wor")
	fmt.Printf("\n\n len(translates): %d \n\n", len(translates))
	assert.NotEqual(t, 0, len(translates), "translates.count should not be equal to zero")
	//

	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, cuttedImageTranslates) {
	}
}

func TestUpdateCuttedImageTranslate(t *testing.T) {

	cuttedimagetranslate := new(dbmodels.CuttedImageTranslate)
	cuttedimagetranslate.CuttedImageTranslateID = uint64(1)
	cuttedimagetranslate.CuttedImageID = uint64(12432)
	cuttedimagetranslate.TelegramUserID = uint64(1243234)
	cuttedimagetranslate.TranslatedWord = "Translated Word KANO MOmuke 22222222"
	cuttedimagetranslate.Comments = "CommentsCommentsComments33333"
	cuttedimagetranslate.Summary = "SummarySummarySummary44444"
	cuttedimagetranslate.AcceptStatus = uint32(4200)

	rowsAffected, err := DbMng.UpdateCuttedImageTranslate(cuttedimagetranslate)
	assert.NotEqual(t, 0, rowsAffected, "userId  should not be equal to zero")
	assert.Nil(t, err, "error should be nil")

	cuttedImageTranslates, err := DbMng.CuttedImageTranslatesFor(uint64(12432))
	assert.NotEqual(t, 0, len(cuttedImageTranslates), "users.count should not be equal to zero")

	for _, cuttedImageTranslate := range cuttedImageTranslates {
		assert.Equal(t, cuttedImageTranslate.CuttedImageTranslateID, uint64(1), "they should be equal")
		assert.Equal(t, cuttedImageTranslate.CuttedImageID, uint64(12432), "they should be equal")
		assert.Equal(t, cuttedImageTranslate.TelegramUserID, uint64(1243234), "CategoryName")
		assert.Equal(t, cuttedImageTranslate.TranslatedWord, "Translated Word KANO MOmuke 22222222", "CategoryName")
		assert.Equal(t, cuttedImageTranslate.Comments, "CommentsCommentsComments33333", "CategoryName")
		assert.Equal(t, cuttedImageTranslate.Summary, "SummarySummarySummary44444", "CategoryName")
		assert.Equal(t, cuttedImageTranslate.AcceptStatus, uint32(4200), "CategoryName")
	}

	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, cuttedImageTranslates) {
	}
}
