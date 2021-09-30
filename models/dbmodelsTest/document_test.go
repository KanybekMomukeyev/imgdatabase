package dbmodelsTest

import (
	"testing"

	dbmodels "github.com/KanybekMomukeyev/imgdatabase/v3/models/dbmodels"
	"github.com/stretchr/testify/assert"
)

func TestCreateDocModel(t *testing.T) {

	docModel := new(dbmodels.DocModel)

	var _companyID uint64 = 589321

	docModel.CompanyID = _companyID
	docModel.UserID = uint64(54321)
	docModel.ParsedImageCount = uint32(54356)
	docModel.StatCount = uint32(54356234)
	docModel.DocmodelName = "DocmodelName"
	docModel.Summary = "Summary"
	docModel.Comments = "Comments"
	docModel.Descriptionn = "Descriptionn"

	docModelId, err := DbMng.CreateDocModel(docModel, _companyID)

	assert.Equal(t, docModelId, uint64(1), "they should be equal")
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, docModelId, "userId  should not be equal to zero")
	if assert.NotNil(t, docModel) {
	}

	docModels, err := DbMng.AllDocModelsForCompany(_companyID)
	assert.NotEqual(t, 0, len(docModels), "users.count should not be equal to zero")

	for _, docModel := range docModels {
		assert.Equal(t, docModel.DocmodelID, uint64(1), "they should be equal")
		assert.Equal(t, docModel.CompanyID, _companyID, "they should be equal")
		assert.Equal(t, docModel.UserID, uint64(54321), "they should be equal")
		assert.Equal(t, docModel.ParsedImageCount, uint32(54356), "they should be equal")
		assert.Equal(t, docModel.StatCount, uint32(54356234), "CategoryName")
		assert.Equal(t, docModel.DocmodelName, "DocmodelName", "CategoryName")
		assert.Equal(t, docModel.Summary, "Summary", "CategoryName")
		assert.Equal(t, docModel.Comments, "Comments", "CategoryName")
		assert.Equal(t, docModel.Descriptionn, "Descriptionn", "CategoryName")
	}

	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, docModels) {
	}
}

func TestUpdateDocModel(t *testing.T) {
	var _companyID uint64 = 589321
	docModel := new(dbmodels.DocModel)
	docModel.DocmodelID = uint64(1)
	docModel.CompanyID = _companyID
	docModel.UserID = uint64(54321456)
	docModel.ParsedImageCount = uint32(354356)
	docModel.StatCount = uint32(513534)
	docModel.DocmodelName = "DocmodelNameDocmodelNameDocmodelName"
	docModel.Summary = "SummarySummarySummary"
	docModel.Comments = "CommentsCommentsComments"
	docModel.Descriptionn = "DescriptionnDescriptionnDescriptionn"

	rowsAffected, err := DbMng.UpdateDocModel(docModel, _companyID)

	assert.Equal(t, rowsAffected, uint64(1), "they should be equal")
	assert.Nil(t, err, "error should be nil")
	assert.NotEqual(t, 0, rowsAffected, "userId  should not be equal to zero")
	if assert.NotNil(t, docModel) {
	}

	docModels, err := DbMng.AllDocModelsForCompany(_companyID)
	assert.NotEqual(t, 0, len(docModels), "users.count should not be equal to zero")

	for _, docModel := range docModels {
		assert.Equal(t, docModel.DocmodelID, uint64(1), "they should be equal")
		assert.Equal(t, docModel.CompanyID, _companyID, "they should be equal")
		assert.Equal(t, docModel.UserID, uint64(54321456), "they should be equal")
		assert.Equal(t, docModel.ParsedImageCount, uint32(354356), "they should be equal")
		assert.Equal(t, docModel.StatCount, uint32(513534), "CategoryName")
		assert.Equal(t, docModel.DocmodelName, "DocmodelNameDocmodelNameDocmodelName", "CategoryName")
		assert.Equal(t, docModel.Summary, "SummarySummarySummary", "CategoryName")
		assert.Equal(t, docModel.Comments, "CommentsCommentsComments", "CategoryName")
		assert.Equal(t, docModel.Descriptionn, "DescriptionnDescriptionnDescriptionn", "CategoryName")
	}

	assert.Nil(t, err, "error should be nil")
	if assert.NotNil(t, docModels) {
	}
}
