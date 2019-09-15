package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//MySQLClientConn defines mysql connection
type MySQLClientConn struct {
	clientConn *gorm.DB
}

//MySQLClientConnInterface is MySQLClientConn interfcae
type MySQLClientConnInterface interface {
	NewClientConnection() *gorm.DB
}

//NewMySQLClientConn inject dependancies for
func NewMySQLClientConn() MySQLClientConnInterface {
	return &MySQLClientConn{}
}

//NewClientConnection  new mysql client connection
func (mySQLClientConn MySQLClientConn) NewClientConnection() *gorm.DB {

	client, err := gorm.Open("mysql", "rahul:password@tcp(127.0.0.1:3306)/ecommerce?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Error in Create client connection", err)
		panic("Error In Create Client Connection")
	}
	return client

}
