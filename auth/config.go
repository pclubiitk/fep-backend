package auth

import (
	"os"

	_ "github.com/pclubiitk/fep-backend/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var logf *os.File

func openConnection() {
	host := viper.GetString("DATABASE.HOST")
	port := viper.GetString("DATABASE.PORT")
	password := viper.GetString("DATABASE.PASSWORD")

	dbName := viper.GetString("DBNAME.AUTH")
	user := viper.GetString("DATABASE.USER")

	dsn := "host=" + host + " user=" + user + " password=" + password
	dsn += " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Kolkata"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		logrus.Fatal("Failed to connect to auth database: ", err)
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&User{}, &OTP{})
	if err != nil {
		logrus.Fatal("Failed to migrate auth database: ", err)
		panic(err)
	}

	logrus.Info("Connected to auth database")
}

func init() {
	openConnection()
	go cleanupOTP()
}
