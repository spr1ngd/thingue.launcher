package util

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestA(t *testing.T) {
	servers := make(map[string]interface{})
	servers["server"] = []string{"1", "2"}
	data := make(map[string]interface{})
	data["type"] = "config"
	data["peerConnectionOptions"] = []map[string]interface{}{servers}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON编码错误:", err)
		return
	}
	for k, v := range data {
		fmt.Println(k, v)
	}

	fmt.Println(string(jsonBytes))
}

// 定义User模型结构体
type User struct {
	gorm.Model
	Name         string
	Email        string
	PhoneNumbers Work
}

type Work []string

func (t *Work) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t Work) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func TestB(t *testing.T) {
	// 连接到SQLite数据库...
	db, err := gorm.Open(sqlite.Open("your-database-file.db"), &gorm.Config{})
	if err != nil {
		panic("无法连接到SQLite数据库")
	}

	// 自动迁移User模型以确保表存在
	db.AutoMigrate(&User{})

	// 创建一个用户并保存到数据库
	user := User{
		Name:         "John Doe",
		Email:        "john@example.com",
		PhoneNumbers: []string{"123456789", "987654321"},
	}
	db.Create(&user)

	// 查询用户并获取电话号码切片
	var fetchedUser User
	db.First(&fetchedUser, user.ID)
	fmt.Println(fetchedUser)

	// 在fetchedUser.PhoneNumbers中可以访问用户的电话号码切片
}
