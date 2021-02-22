package db

import (
	"github.com/EikoNakashima/test-go3/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "test"
	PASS := "12345678"
	DBNAME := "test"
	// MySQLだと文字コードの問題で"?parseTime=true"を末尾につける必要がある
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

//DB初期化
func Init() {
	db := gormConnect()
	// if err != nil {
	// 	panic("データベース開けず！（dbInit）")
	// }
	db.AutoMigrate(&model.Todo{})
	defer db.Close()
}

//DB追加
func Insert(text string, status string) {
	db := gormConnect()
	// if err != nil {
	// 	panic("データベース開けず！（dbInsert)")
	// }
	db.Create(&model.Todo{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func Update(id int, text string, status string) {
	db := gormConnect()
	// if err != nil {
	// 	panic("データベース開けず！（dbUpdate)")
	// }
	var todo model.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB削除
func Delete(id int) {
	db := gormConnect()
	// if err != nil {
	// 	panic("データベース開けず！（dbDelete)")
	// }
	var todo model.Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB全取得
func GetAll() []model.Todo {
	db := gormConnect()
	// if err != nil {
	// 	panic("データベース開けず！(dbGetAll())")
	// }
	var todos []model.Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DB一つ取得
func GetOne(id int) model.Todo {
	db := gormConnect()
	// if err != nil {
	// 	panic("データベース開けず！(dbGetOne())")
	// }
	var todo model.Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
