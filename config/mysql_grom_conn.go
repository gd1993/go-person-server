package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var sqlDB *gorm.DB

func MysqlGormConn() (db2 *gorm.DB) {
	username := "root"      //账号
	password := "Hzh122427" //密码
	host := "127.0.0.1"     //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "go_person"   //数据库名
	timeout := "10s"        //连接超时，10秒
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	sqlDB, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("mysql open failed", err)
		return nil
	}
	return sqlDB
}
