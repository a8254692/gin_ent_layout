package data

import (
	"back-end/merchant/internal/data/ent"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var EntClient *ent.Client

// Data .
type Data struct {
	DB *ent.Client
}

// NewData .
func NewData() (*Data, func(), error) {
	var entClient *ent.Client
	if EntClient == nil {
		drv, err := getNewEntDrv()
		if err != nil || drv == nil {
			panic(fmt.Sprintf("failed to get ent client, error:%s", err.Error()))
			return nil, nil, err
		}

		entClient = drv
	} else {
		entClient = EntClient
	}

	d := &Data{
		DB: entClient,
	}

	return d, func() {
		if err := d.DB.Close(); err != nil {
		}
	}, nil
}

func getNewEntDrv() (*ent.Client, error) {
	prefix := fmt.Sprintf("mysql_list.merchant")
	dialect := "mysql"
	user := viper.GetString(prefix + ".user")
	password := viper.GetString(prefix + ".passwd")
	database := viper.GetString(prefix + ".db")
	dbPath := viper.GetString(prefix + ".addr")
	maxIdle := 10
	maxOpen := 100
	//maxLifetime := viper.GetInt(prefix + ".maxLifetime")

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbPath, database)

	db, err := sql.Open(dialect, url)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)
	//db.DB().SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	drv := sql.OpenDB(dialect, db.DB())
	entDrv := ent.NewClient(ent.Driver(drv))
	if entDrv == nil {
		return nil, err
	}

	EntClient = entDrv
	return entDrv, nil
}
