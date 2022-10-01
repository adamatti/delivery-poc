package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var instance *gorm.DB

// FIXME this is a pet project, in a real one must think in a clever/robust way
func runMigrations (entities []interface{}, db *gorm.DB) {
	for _, entity := range entities {
		err:= db.AutoMigrate(entity)
		if err != nil {
			panic("fail during db migration")
		}
	}
}

func StartDatabase(entities []interface{}, config postgres.Config) *gorm.DB {
	connection:= postgres.New(config)

	db, err := gorm.Open(connection, &gorm.Config{})
	if err != nil {
    panic("failed to connect database")
  }

	runMigrations(entities, db)
	
	instance = db
	return db
}

func GetInstance() *gorm.DB {
	return instance
}