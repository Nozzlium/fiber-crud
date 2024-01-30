package lib

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func GetDatabase(config *viper.Viper) (*sql.DB, error) {
	dbConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.GetString("database.username"),
		config.GetString("database.password"),
		config.GetString("database.host"),
		config.GetString("database.port"),
		config.GetString("database.database"),
	)
	return sql.Open("mysql", dbConfig)
}
