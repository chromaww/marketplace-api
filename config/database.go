package config

import (
	"github.com/jinzhu/gorm"
)

type Database struct {
	Driver   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
	SslMode  string
}

func (this *Database) Connect() *gorm.DB {
	var dsn string
	switch driver := this.Driver; { // missing expression means "true"
	case driver == "postgres":
		dsn = "host=" + this.Host +
			" port=" + this.Port +
			" user=" + this.Username +
			" password=" + this.Password +
			" dbname=" + this.Database +
			" sslmode=" + this.SslMode
	default:
		panic("DB driver selected was unsupported")
	}

	db, err := gorm.Open(this.Driver, dsn)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	return db
}
