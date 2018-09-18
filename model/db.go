package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" //PostgreSQL Driver
)

var ormObject orm.Ormer

// ConnectToDb - Initializes the ORM and Connection to the postgres DB
func ConnectToDb() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=root dbname=gosample_db host=127.0.0.1 port=5432 sslmode=disable")
	orm.RegisterModel(new(User))
	ormObject = orm.NewOrm()
	orm.RunSyncdb("default", false, true)
}

// GetOrmObject - Getter function for the ORM object with which we can query the database
func GetOrmObject() orm.Ormer {
	return ormObject
}
