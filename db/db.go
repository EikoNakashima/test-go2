package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

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
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

//DB追加
func Insert(text string, status string) {
	db := gormConnect()
	// if err != nil {
	// 	panic("データベース開けず！（dbInsert)")
	// }
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func Update(id int, text string, status string) {
	db := gormConnect()
	// if err != nil {
	// 	panic("データベース開けず！（dbUpdate)")
	// }
	var todo Todo
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
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB全取得
func GetAll() []Todo {
	db := gormConnect()
	// if err != nil {
	// 	panic("データベース開けず！(dbGetAll())")
	// }
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DB一つ取得
func GetOne(id int) Todo {
	db := gormConnect()
	// if err != nil {
	// 	panic("データベース開けず！(dbGetOne())")
	// }
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
