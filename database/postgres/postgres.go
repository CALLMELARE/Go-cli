package postgres

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var db *sql.DB

func Init() (err error) {
	connect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", viper.GetString("postgres.host"), viper.GetInt("postgres.port"), viper.GetString("postgres.user"), viper.GetString("postgres.password"), viper.GetString("postgres.dbname"))
	db, err = sql.Open("postgres", connect)
	if err != nil {
		zap.L().Error("Fatal error:Cannot connect to database")
		return
	}
	db.SetMaxOpenConns(viper.GetInt("postgres.max_open_connects"))
	db.SetMaxIdleConns(viper.GetInt("postgres.max_idle_connects"))
	return
}

func Close() {
	_ = db.Close()
}
