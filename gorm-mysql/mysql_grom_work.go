package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-person-server/config"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	sqlDb := config.MysqlGormConn()
	if sqlDb != nil {
		fmt.Println("gorm conn mysql success")
	}
	defer sqlDb.Close()
	//自动检查product是否有变化
	sqlDb.AutoMigrate(&Product{})
	//增
	sqlDb.Create(&Product{
		Code:  "L12",
		Price: 120,
	})

	//查
	var product Product
	sqlDb.First(&product, 1)
	sqlDb.First(&product, "code = ?", "L12")
	//修改价格
	sqlDb.Model(&product).Update("Price", 150)
	//删除
	sqlDb.Delete(&product)
}
