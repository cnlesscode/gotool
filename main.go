package main

import (
	"fmt"

	"github.com/cnlesscode/gotool/db"
)

type Students struct {
	Id      int    `gorm:"column:st_id;primaryKey"`
	ClassId int    `gorm:"column:st_class_id" validate:"gt=0"`
	Name    string `gorm:"column:st_name" validate:"min=3"`
	Age     int    `gorm:"column:st_age" validate:"gt=10"`
	AddTime int    `gorm:"column:st_add_time"`
}

func main() {
	dbObj := db.Init()
	// where 条件组合演示
	whereMap := map[string][]any{
		"st_id":   {"", ">", 265520},
		"st_name": {"and", "like", "%张%"},
	}
	whereSql, whereVal := db.MapToWhere(whereMap)
	fmt.Printf("whereSql: %v\n", whereSql)
	fmt.Printf("whereVal: %v\n", whereVal)
	// 查询演示

	data := make([]Students, 0)
	dbObj.Where(whereSql, whereVal...).Limit(10).Find(&data)
	fmt.Printf("data: %v\n", data)
}
