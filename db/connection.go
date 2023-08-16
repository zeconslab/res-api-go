package db

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

/*
Metodo para verificacion conexion de la base de datos
y retorna una respuesta de conexion exitosa o un error
*/
func DBconection(dbhost string, dbport string, dbuser string, dbpass string, dbname string) {
	var DBS = "sqlserver://" + dbuser + ":" + dbpass + "@" + dbhost + ":" + dbport + "?database=" + dbname
	log.Println("Connecting to DB...")
	var error error
	DB, error = gorm.Open(sqlserver.Open(DBS), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Connecting to DB... OK")
		//Migracion de la base de datos
		Migrates()
	}
}
