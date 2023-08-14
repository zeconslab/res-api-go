package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBS = "host=localhost user=zeconslab password=zontrox11 dbname=goprueba port=5432"
var DB *gorm.DB

/*
Metodo para verificacion conexion de la base de datos
y retorna una respuesta de conexion exitosa o un error
*/
func DBconection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DBS), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		//Migracion de la base de datos
		Migrates()
		log.Println("DB Connected!")
	}
}
