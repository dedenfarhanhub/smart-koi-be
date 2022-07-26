package mysql

import (
	"fmt"
	"github.com/dedenfarhanhub/smart-koi-be/drivers/databases/calculate_productions"
	"github.com/dedenfarhanhub/smart-koi-be/drivers/databases/history_productions"
	"github.com/dedenfarhanhub/smart-koi-be/drivers/databases/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ConfigDB struct {
	DbUsername string
	DbPassword string
	DbHost      string
	DbPort     string
	DbDatabase string
}

func (config *ConfigDB) InitialMysqlDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&users.Users{},
		&history_productions.HistoryProductions{},
		&calculate_productions.CalculateProductions{},
	)
	if err != nil {
		return nil
	}

	return db
}