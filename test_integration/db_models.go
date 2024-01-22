package testintegration

import (
	"log"

	"github.com/elliotchance/pie/v2"
	"gorm.io/gorm"

	"github.com/FrancoLiberali/stori_challenge/app/models"
)

var ListOfTables = []any{
	models.User{},
	models.Transaction{},
}

func CleanDB(db *gorm.DB) {
	CleanDBTables(db, pie.Reverse(ListOfTables))
}

func CleanDBTables(db *gorm.DB, listOfTables []any) {
	// clean database to ensure independency between tests
	for _, table := range listOfTables {
		err := db.Unscoped().Where("1 = 1").Delete(table).Error
		if err != nil {
			log.Fatalln("could not clean database: ", err)
		}
	}
}
