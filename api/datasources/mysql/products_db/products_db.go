// products_db.go

package products_db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CONSTANT VALUES
const (
	DBTYPE = "mysql"
	SCHEMA = "%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local"
)

var (
	Client   *gorm.DB
	username = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	dbName   = os.Getenv("MYSQL_DATABASE")

	datasourceName = fmt.Sprintf(SCHEMA, username, password, dbName)
)

// https://gorm.io/ja_JP/docs/connecting_to_the_database.html#MySQL
func init() {
	var err error
	// user:password@/db_name -> docker.compose.yml - mysql service
	Client, err = gorm.Open(DBTYPE, datasourceName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("database successfully configure")
}
