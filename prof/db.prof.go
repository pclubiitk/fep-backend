package prof

import (
	"github.com/gin-gonic/gin"
)

func getAllProf(ctx *gin.Context, companies *[]Prof) error {
	tx := db.WithContext(ctx).Find(companies)
	return tx.Error
}

func getProf(ctx *gin.Context, company *Prof, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).First(company)
	return tx.Error
}

func getLimitedProf(ctx *gin.Context, companies *[]Prof, lastFetchedId uint, pageSize int) error {
	tx := db.WithContext(ctx).Order("id asc").Where("id >= ?", lastFetchedId).Limit(pageSize).Find(companies)
	return tx.Error
}

func updateProf(ctx *gin.Context, company *Prof) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ?", company.ID).Updates(company)
	return tx.RowsAffected > 0, tx.Error
}

func createProf(ctx *gin.Context, company *Prof) error {
	tx := db.WithContext(ctx).Create(company)
	return tx.Error
}

// func createCompanies(ctx *gin.Context, company *[]Company) error {
// 	tx := db.WithContext(ctx).Create(company)
// 	return tx.Error
// }

func deleteProf(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).Delete(&Prof{})
	return tx.Error
}

func GetProfName(ctx *gin.Context, id uint) (string, error) {
	var c Prof
	err := getProf(ctx, &c, id)
	if err != nil {
		return "", err
	}
	return c.ProfessorName, nil
}

func FetchProfIDByEmail(ctx *gin.Context, email string) (uint, error) {
	var prof Prof
	tx := db.WithContext(ctx).Where("professor_email_id = ?", email).First(&prof)
	return prof.ID, tx.Error
}
