package database

import (
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pelletier/go-toml"
	"github.com/theking0912/IrisApiProject/config"
)

var (
	DB = New()
)

/**
*设置数据库连接
*@param diver string
 */
func New() *gorm.DB {

	if isTestEnv() {
		configTree := config.Conf.Get("test").(*toml.Tree)
		DB, err := gorm.Open(configTree.Get("DataBaseDriver").(string), configTree.Get("DataBaseConnect").(string))
		if err != nil {
			panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
		}

		return DB

	} else {

		driver := config.Conf.Get("database.dirver").(string)
		fmt.Println(driver)
		configTree := config.Conf.Get(driver).(*toml.Tree)
		userName := configTree.Get("databaseUserName").(string)
		password := configTree.Get("databasePassword").(string)
		databaseName := configTree.Get("databaseName").(string)
		databaseHost := configTree.Get("databaseHost").(string)
		databasePort := configTree.Get("databasePort").(string)
		connect := userName + ":" + password + "@(" + databaseHost + ":" + databasePort + ")/" + databaseName + "?charset=utf8&parseTime=True&loc=Local"
		fmt.Println(connect)
		DB, err := gorm.Open(driver, connect)

		if err != nil {
			panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
		}

		return DB
	}
}

//获取程序运行环境
// 根据程序运行路径后缀判断
//如果是 test 就是测试环境
func isTestEnv() bool {
	files := os.Args
	for _, v := range files {
		if strings.Contains(v, "test") {
			return true
		}
	}

	return false
}
