package connection

import (
	"api_model_cnn/src/apimodels/utils"
	"api_model_cnn/src/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Connection interface {
	SqlDb() *sql.DB
}

type postgres struct {
	db *sql.DB
}

// SqlDb implements Connection.
func (i *postgres) SqlDb() *sql.DB {
	i.db.SetMaxOpenConns(10)
	i.db.SetMaxIdleConns(2)
	i.db.SetConnMaxIdleTime(30 * time.Millisecond)

	return i.db
}

func CreateConnectionPostgres(config *config.Config) Connection {
	fmt.Println(config)
	conn, errDb := sql.Open("postgres",config.Dns)
	if errDb != nil {
		utils.LogError(errDb,"Error in Open Connection PGSQL")
		return nil
	}
	if err := conn.Ping(); err != nil {
		utils.LogError(err, "Error in PingConnection PGSQL")
		return nil
	}
	return &postgres{db : conn}
}