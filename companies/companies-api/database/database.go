package database

import (
	"gorm.io/gorm"
)

var instance *gorm.DB

type Entities []interface{}

// FIXME this is a pet project, in a real one must think in a clever/robust way
func runMigrations (entities []interface{}, db *gorm.DB) {
	for _, entity := range entities {
		err:= db.AutoMigrate(entity)
		if err != nil {
			panic("fail during db migration")
		}
	}
}

func StartDatabase(entities []interface{}, dialector gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
    panic("failed to connect database")
  }

	runMigrations(entities, db)
	
	instance = db
	return db
}

func GetInstance() *gorm.DB {
	if instance == nil {
		panic("Database not initialized")
	}
	return instance
}