package models

import (
	"github.com/jinzhu/gorm"
	"time"
	orm "user/database"
)

type User struct {
	ID       int64  `form:"id" json:"id"`       				//
	Username string `form:"username" json:"username"` 			// 英文名
	Realname string `form:"realname" json:"realname"` 			// 姓名
	Email string `form:"email" json:"email"` 					// 邮箱
	Mobile string `form:"mobile" json:"mobile"` 				// 手机
	Status int `form:"status" json:"status"` 					// 状态
	Verify string `form:"verify" json:"verify"` 				// 密钥
	EntryDate int64 `form:"entry_date" json:"entry_date"` 		// 入职时间
	OpRole int `form:"op_role" json:"op_role"` 					// 后台用户
	AdminRole int `form:"admin_role" json:"admin_role"` 		// 管理员
	EmailFlag int `json:"email_flag"` 							// 邮件验证状态
	MobileFlag int `json:"mobile_flag"` 						// 手机验证状态
	CreatedAt time.Time `json:"created_at"` 					// 创建时间
	UpdatedAt time.Time `json:"updated_at"` 					// 更新时间
	DeletedAt *time.Time `json:"deleted_at"` 					// 删除时间
}

// 设置User的表名为`profiles`
func (User) TableName() string {
	return "user"
}

type UserExt struct {
	ID		int64	`json:"id"`       			// 列名为 `id`
	Uid		User	`gorm:"foreignkey:UserId" json:"uid"`       			// 列名为 `uid`	用户id
	Desc	int64	`json:"desc"`       		// 列名为 `desc`	备注
}

// 设置User的表名为`profiles`
func (UserExt) TableName() string {
	return "userExt"
}

var UserField = []string {}
var UserExtField = []string {}

func init() {
	var user User
	UserField = orm.GetFieldName(user)
	var userExt UserExt
	UserExtField = orm.GetFieldName(userExt)
}

//添加
func (user User) Create() (id int64, err error) {
	result := orm.Eloquent.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// 更新
func (user *User) Update(id int64, updateData map[string]interface{}) (err error) {
	err = user.GetOne(id)
	if err != nil {
		return
	}
	if err = orm.Eloquent.Model(&user).Updates(updateData).Error; err != nil {
		return
	}
	return
}

//删除数据
func (user *User) Delete(id int64) (Result User, err error) {
	if err = orm.Eloquent.Select([]string{"id"}).First(&Result, id).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Delete(&Result).Error; err != nil {
		return
	}
	return
}

//获取一个
func (user *User) GetOne(id int64) (err error) {
	err = orm.Eloquent.Select(UserField).First(&user, id).Error
	return
}

//获取查询db
func (user *User) QueryDB(db *gorm.DB, searchKey map[string] string)( returnDB *gorm.DB ){
	// 查询扩展
	// 字段中如果字段 为 realname
	// 则默认支持以下查询
	// 		searchKey["realname"]  				精确查询		"realname = ?", "%" + searchKey["realname_min"] + "%"
	// 		searchKey["realname_not"]  			精确查询		"realname <> ?", "%" + searchKey["realname_min"] + "%"
	//		searchKey["realname_like"]  		模糊查询		"realname like ?", "%" + searchKey["realname_min"] + "%"
	//		searchKey["realname_left_like"]  	模糊左匹配查询		"realname like ?", searchKey["realname_min"] + "%"
	//		searchKey["realname_min"]			查询   		"realname >= ?" , searchKey["realname_min"]
	//		searchKey["realname_max"]			查询   		"realname <= ?" , searchKey["realname_min"]
	db = orm.QueryDbInit( db, UserField, searchKey)

	// 查询条件更多条件请继续增加
	for k,v := range searchKey {
		if v=="" {
			continue
		}
		switch k {
		case "searchkey":
			likeStr := "%" + v + "%"
			db = db.Where("username LIKE ? OR realname LIKE ? OR email LIKE ? OR mobile LIKE ?", likeStr, likeStr, likeStr, likeStr)
			break;
		}
	}

	returnDB = db
	return
}

//列表
func (user *User) Query( searchKey map[string] string, page int64, pageSize int64  ) (users []User, err error) {
	db := orm.Eloquent

	// 设置返回的字段
	db = db.Select(UserField)
	db = user.QueryDB(db, searchKey)

	// 分页
	if page > 0 && pageSize > 0 {
		db = db.Limit(pageSize).Offset((page - 1) * pageSize)
	}

	// 查询数据
	if err = db.Find(&users).Error; err != nil {
		return
	}
	return
}

func (user *User) QueryTotalSize( searchKey map[string] string  ) ( totalSize int64, err error) {
	db := orm.Eloquent

	// 设置返回的字段
	db = db.Select(UserField)
	db = user.QueryDB(db, searchKey)

	err = db.Table( user.TableName() ).Count(&totalSize).Error

	return
}


