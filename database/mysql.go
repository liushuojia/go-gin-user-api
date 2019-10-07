package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"log"
	"reflect"
	. "user/conf"
	"user/libs"
)

var Eloquent *gorm.DB

func init() {

	var err error
	Eloquent, err = gorm.Open("mysql", ConfigMysql["host"])

	if err != nil {
		fmt.Printf("mysql connect error %v\n", err)
		panic("")
	}

	if Eloquent.Error != nil {
		fmt.Printf("\ndatabase error %v", Eloquent.Error)
		panic("")
	}

	// 表名禁用复数
	Eloquent.SingularTable(true)

	if ConfigMysql["log"]=="show" {
		Eloquent.LogMode(true)
	}
}

// 获取结构体的字段
func GetFieldName(structName interface{}) (result []string) {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	ObjType,_ := libs.Obj2Map(structName)

	for k, _ := range ObjType {
		result = append(result, k)
	}
	return
}

// 初始化DB的查询条件, 根据map
// 查询扩展
// 字段中如果字段 为 realname
// 则默认支持以下查询
// 		searchKey["realname"]  				精确查询		"realname = ?", "%" + searchKey["realname_min"] + "%"
// 		searchKey["realname_not"]  			精确查询		"realname <> ?", "%" + searchKey["realname_min"] + "%"
//		searchKey["realname_like"]  		模糊查询		"realname like ?", "%" + searchKey["realname_min"] + "%"
//		searchKey["realname_left_like"]  	模糊左匹配查询		"realname like ?", searchKey["realname_min"] + "%"
//		searchKey["realname_min"]			查询   		"realname >= ?" , searchKey["realname_min"]
//		searchKey["realname_max"]			查询   		"realname <= ?" , searchKey["realname_min"]
//
//	如果不想用这个, 请直接在类中写, 这样可以增加效率或执行速度
func QueryDbInit ( db *gorm.DB, FieldMap []string, searchkey map[string]string )  (returnDb *gorm.DB) {
	for _,key := range FieldMap {
		if key == "" {
			continue
		}
		if value, ok := searchkey[key]; ok && value!="" {
			db = db.Where( key + " = ?", value)
		}
		if value, ok := searchkey[key+"_not"]; ok && value!="" {
			db = db.Where( key + " <> ?", value)
		}
		if value, ok := searchkey[key+"_like"]; ok && value!="" {
			db = db.Where( key + " LIKE ?", "%" + value + "%")
		}
		if value, ok := searchkey[key+"_left_like"]; ok && value!="" {
			db = db.Where( key + " LIKE ?", "%" + value )
		}
		if value, ok := searchkey[key+"_min"]; ok && value!="" {
			db = db.Where( key + " >= ?", value )
		}
		if value, ok := searchkey[key+"_max"]; ok && value!="" {
			db = db.Where( key + " <= ?", value )
		}
	}
	returnDb = db
	return
}

