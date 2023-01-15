package postgresql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
	Timezone string
}

func Init(conData PostgresConnection) (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		conData.Host, conData.User, conData.Password, conData.DBname, conData.Port, conData.Timezone)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
