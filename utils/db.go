package utils

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	_ "github.com/lib/pq"
)

// mdprInfo table
type mdprInfo struct {
	ID     int         `json:"id" db:"serial;PRIMARY KEY"`
	URL    string      `json:"url" db:"varchar(100);DEFAULT ''"`
	Source QMediaArray `json:"source" db:"jsonb;DEFAULT '[]'"`
}

type QMediaArray []string

func (m *QMediaArray) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}

func (m QMediaArray) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func DBFiled(reflectType reflect.Type, buffer *bytes.Buffer) {
	if reflectType.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < reflectType.NumField(); i++ {
		jsonTag := reflectType.Field(i).Tag.Get("json")
		dbTag := reflectType.Field(i).Tag.Get("db")

		if jsonTag == "" && dbTag == "" {
			DBFiled(reflectType.Field(i).Type, buffer)
			continue
		}

		dbProfile := strings.Split(dbTag, ";")
		dbFiled := fmt.Sprintf("%s %s", jsonTag, strings.Join(dbProfile, " "))
		buffer.WriteString(dbFiled)
		buffer.WriteString(",")
	}
}

// Init PostgreSQL
type PostgreSQL struct{}

var Psql *PostgreSQL
var pdb *sql.DB

func init() {
	Psql := new(PostgreSQL)
	pdb = Psql.Connect()
}

func (p *PostgreSQL) Connect() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Panic("DB Connect Failed", err)
	}
	return db
}

func (p *PostgreSQL) Exec(sql string, args ...interface{}) {
	stmt, err := pdb.Prepare(sql)
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()
	stmt.Exec(args...)
}

func (p *PostgreSQL) QueryOne(sql string, args ...interface{}) *sql.Row {
	stmt, err := pdb.Prepare(sql)
	if err != nil {
		log.Panic(err)
	}

	defer stmt.Close()
	row := stmt.QueryRow(args...)
	return row
}

func CreateTable() {
	var buffer bytes.Buffer
	rType := reflect.TypeOf(mdprInfo{})
	rName := "mdpr"
	DBFiled(rType, &buffer)
	rFiled := buffer.Bytes()[0 : len(buffer.Bytes())-1]

	sql := fmt.Sprintf("CREATE TABLE %s (%s)", rName, rFiled)
	_, err := pdb.Exec(sql)
	if err != nil {
		log.Println("Create Error:", err)
	}
}
