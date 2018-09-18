package main

import (
	"github.com/astaxie/beego/orm"
	"net/http"
	"golang-crud-api-sederhana/controller/usercontroller"
	"fmt"
	"golang-crud-api-sederhana/model"
)

/** using beego ORM https://beego.me/docs/mvc/model/orm.md
**/
func init()  {
	model.ConnectToDb()
}


func main() {
	orm.Debug = true
	http.HandleFunc("/users", usercontroller.Handler)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
