## 一、本包功能介绍

在实际项目开发中我们可能会习惯性的手动建立`mysql`数据表，或者在企业中有专门的`DB`来创建数据表，在我们`go`语言开发中不想手动去创建关于数据表的实体类(这里叫`model`，本人习惯性实体类用`Entity`来结尾)，可以使用本包来一件生成实体类



## 二、使用方式

* 1、下载依赖包

  ```properties
  go get -u github.com/kuangshp/generate-model
  go mod tity
  ```

* 2、在项目的根目录下创建一个文件(比如`generate.go`)的文件用来生成实体类

  ```go
  package main
  
  import (
  	"fmt"
  	"github.com/kuangshp/generate-model/converter"
  	"github.com/kuangshp/generate-model/utils"
  	"strings"
  )
  
  func main() {
  	var tableName string
  	fmt.Print("请输入表名:")
  	if _, err := fmt.Scanln(&tableName); err != nil {
  		return
  	}
  	var fileName = utils.Case2Camel(strings.ToUpper(tableName[0:1]) + tableName[1:]) // 转为首字目大写
  	var err = converter.NewTable2Struct().
  		SavePath(fmt.Sprintf("./model/%sEntity.go", fileName)). // 这个地方根据自己的需求来写
  		Dsn("root:123456@tcp(localhost:3306)/test?charset=utf8mb4"). // 配置数据库连接
  		TagKey("gorm").              // orm
  		EnableJsonTag(true).         // json
  		RealNameMethod("TableName"). // 生成表名
  		Table(tableName).            // 表名
  		DateToTime(true).            // 对时间字段进行转换
  		Config(&converter.T2tConfig{
  			JsonTagToHump: true,
  			StructEnd:     "Entity",
  		}). // 配置tag转驼峰
  		Run()
  	if err != nil {
  		fmt.Println("创建数据模型失败")
  		return
  	}
  	fmt.Println("生成数据模型成功")
  }
  
  ```

* 3、运行上面的文件，控制台提示你输入表名，然后回车后自动生成数据模型

  ```go
  package model
  
  import "time"
  
  type AccountEntity struct {
  	Id        int64     `gorm:"id" json:"id"`                // 主键id
  	UserName  string    `gorm:"username" json:"username"`    // vj名称
  	Password  string    `gorm:"password" json:"password"`    // 密码
  	CreatedAt time.Time `gorm:"created_at" json:"createdAt"` // 创建时间
  	UpdatedAt time.Time `gorm:"updated_at" json:"updatedAt"` // 更新时间
  	DeletedAt time.Time `gorm:"deleted_at" json:"deletedAt"` // 软删除时间
  }
  
  func (t *AccountEntity) TableName() string {
  	return "account"
  }
  
  ```

* 4、继续自己的业务开发