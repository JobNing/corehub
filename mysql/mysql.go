package databases

import (
	"fmt"
	"github.com/JobNing/corehub/config"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	Mysql Mysql `yaml:"mysql"`
}

type Mysql struct {
	Host   string `yaml:"host"`
	Port   int64  `yaml:"port"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

func WithClient(hand func(db *gorm.DB) error) error {
	var conf MysqlConfig
	confContext, err := config.GetConfig()
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(confContext), &conf)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Mysql.User,
		conf.Mysql.Pwd,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.Dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	defer func() {
		dd, _ := db.DB()
		dd.Close()
	}()

	return hand(db)
}
