package initializers



import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )


  var DB *grom.DB

func ConnectToDB(){

var err error
dsn := "host=drona.db.elephantsql.com user=ztkfslgf password=z2xF43ftoskRTOTJZ9_iFsjReoRzwqpN dbname=ztkfslgf port=5432 sslmode=disable"
DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

if err != nil{
	log.Fatal("Failed to connec to db")
}
}