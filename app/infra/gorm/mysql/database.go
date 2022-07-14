package mysql

import (
	"context"
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Database variavel com implementação do gorm.
var Database *gorm.DB

//Connect conecta no banco de dados.
func Connect() error {
	var err error

	if os.Getenv("DIABETES_MYSQL_HOST") == "" {
		fmt.Println("MySQL vars not set.")
		return errors.New("mySQL vars not set")
	}

	conectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		os.Getenv("DIABETES_MYSQL_LOGIN"),
		os.Getenv("DIABETES_MYSQL_PASSWORD"),
		os.Getenv("DIABETES_MYSQL_HOST"),
		os.Getenv("DIABETES_MYSQL_PORT"),
		os.Getenv("DIABETES_MYSQL_DATABASE"),
	)

	Database, err = gorm.Open(mysql.Open(conectionString), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	return nil
}

//Close fecha conexão com banco de dados.
func Close() error {
	if Database == nil {
		fmt.Println("Database connection not set. No one connection to close.")
		return nil
	}

	sqlDB, err := Database.DB()

	if err != nil {
		fmt.Println(context.TODO(), "Error get connect database", err)
		panic(err)
	}

	err = sqlDB.Close()

	if err != nil {
		fmt.Println(context.TODO(), "Error closing DB connection.", err)
		panic(err)
	}

	return nil
}
