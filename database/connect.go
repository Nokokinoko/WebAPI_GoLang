package database

import (
	"database/sql"
	"webapi/handler"

	_ "github.com/go-sql-driver/mysql"
)

// ------ Database ------
type singletonDb struct {
	db *sql.DB
}

var instance_main *singletonDb
var instance_sub *singletonDb

func Instance(dbName string, isMain bool) *singletonDb {
	// TODO: stop replication, connect only main
	isMain = true

	if (isMain && instance_main == nil) || (!isMain && instance_sub == nil) {
		user := "user" // TODO: mysql user
		port := "3306"

		var password string
		var host string
		if handler.IsDocker() {
			password = "password" // TODO: mysql password
			host = "mysql-container"
		} else if isMain {
			// main
			// TODO: aws main host, parameter store
			password = ""
			host = ""
		} else {
			// sub
			// TODO: aws sub host, parameter store
			password = ""
			host = ""
		}

		dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4"
		db, err := sql.Open("mysql", dbconf)
		if err != nil {
			panic(err)
		}

		if isMain {
			instance_main = &singletonDb{db: db}
		} else {
			instance_sub = &singletonDb{db: db}
		}
		defer db.Close()
	}

	if isMain {
		return instance_main
	} else {
		return instance_sub
	}
}

// ------ Record ------
type IRecord interface {
	columns() []string
	assignRow(row *sql.Row) error
	assignRows(rows *sql.Rows) error
}

type abstractRecord struct {
}

func (record abstractRecord) columns() []string {
	return []string{}
}

func (record abstractRecord) assignRow(row *sql.Row) error {
	return nil
}

func (record abstractRecord) assignRows(rows *sql.Rows) error {
	return nil
}

// ------ Table ------
type ITable interface {
	specifyKeys(usePrimaryKey bool) []string
	emptyRecord() abstractRecord

	Update(record abstractRecord) error
	Insert(record abstractRecord) (abstractRecord, error)
	Delete(mapKeyVal map[string]string) error

	Rows(mapKeyVal map[string]string) ([]abstractRecord, error)
	Row(usePrimaryKey bool, mapKeyVal map[string]string) (abstractRecord, error)
	Count(mapKeyVal map[string]string) (int, error)
}

type abstractTable struct {
	databaseName string
	tableName    string
	primaryKey   string
	uniqueKeys   []string
}

func (table abstractTable) specifyKeys(usePrimaryKey bool) []string {
	var specifyKeys []string
	if usePrimaryKey {
		// use primary key
		if 0 < len(table.primaryKey) {
			specifyKeys = append(specifyKeys, table.primaryKey)
		}
	} else {
		// use unique key
		if 0 < len(table.uniqueKeys) {
			specifyKeys = append(specifyKeys, table.uniqueKeys...)
		}
	}
	return specifyKeys
}

func (table abstractTable) emptyRecord() abstractRecord {
	return abstractRecord{}
}
