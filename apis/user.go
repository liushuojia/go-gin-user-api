package apis

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	"user/libs"
	model "user/models"
	orm "user/database"
	. "user/conf"
)

// 生成token
func BuildUid( id int64, verify string, timeStamp int64, userAgent string ) ( uid string ) {
	data := map[string]interface{}  {
		"key" : ConfigApi["md5Key"],
		"userAgent" : userAgent,
		"verify" : verify,
		"id" : id,
		"timestamp" : timeStamp,
	}
	uid = libs.BuildToken( data )
	return
}

// 校验token
func CheckUid( token string, userAgent string ) ( user model.User, err error ) {

	chrstr := strings.Split(token,"-");
	switch( len(chrstr) ) {
	case 2:
		uid := BuildUid( 0, "", libs.AnyToDecimal( chrstr[0], libs.IDNUN ), userAgent )
		if uid!=token {
			err = errors.New("token 过期");
		}
		break;
	case 3:
		id := libs.AnyToDecimal( chrstr[0], libs.IDNUN )
		timestamp := libs.AnyToDecimal( chrstr[1], libs.IDNUN )

		if( id<=0 ){
			err = errors.New("token 过期#1");
			break;
		}

		if( (time.Now().Unix() - timestamp)>24*60*60 ) {
			err = errors.New("token 过期#2");
			break;
		}

		// 获取用户信息
		userID := libs.SetValueToType(id,"string").(string)
		keyAdmin := ConfigAdminRedis["redisAdminPrefix"] + userID

		Json,errGet := orm.Redis.Get( keyAdmin )

		if errGet!=nil {
			//临时数据过期
			err = user.GetOne(id);
			if  err != nil {
				return
			}

			//重新获取到数据后, 重新将数据写入redis
			saveTime := libs.SetValueToType(ConfigAdminRedis["redisSaveTime"],"int64").(int64)
			jsonBytes, _ := json.Marshal(user)
			orm.Redis.Set( keyAdmin, string(jsonBytes), saveTime )
		}else{
			err = json.Unmarshal( []byte(Json), &user)
			if err!=nil {
				err = errors.New("缓存数据出错");
				break
			}
		}

		uid := BuildUid( id, user.Verify, timestamp, userAgent )
		if uid!=token {
			err = errors.New("token 过期#3");
			break;
		}

		break;
	}
	return
}

func CheckToken(c *gin.Context) {
	userAgent := c.GetHeader("User-Agent")
	token := c.GetHeader("token")

	user, err := CheckUid(token,userAgent);
	if err!=nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err.Error(),
		})
		return
	}

	if user.ID ==0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   	1,
			"message": 	"临时token 有效",
		})
	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   	1,
			"message": 	"token 有效",
		})
	}

	return
}

//获取token
func GetToken(c *gin.Context) {
	mobile := c.DefaultQuery("mobile", "")
	verify := c.DefaultQuery("verify", "")
	userAgent := c.GetHeader("User-Agent")

	var user model.User;
	if mobile == "" {
		// 临时token
		uid := BuildUid( 0, "", time.Now().Unix(), userAgent )
		c.JSON(http.StatusOK, gin.H{
			"code": 	1,
			"message": 	"获取成功",
			"data": 	uid,
		})
		return
	}

	searchMap := make( map[string]string )
	searchMap["mobile"] = mobile
	searchMap["verify"] = verify
	result, err := user.Query(searchMap, 1, 1)
	if err != nil || len(result)==0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "手机号码错误或密钥错误",
		})
		return
	}
	user = result[0]

	uid := BuildUid( user.ID, user.Verify, time.Now().Unix(), userAgent )

	// 写入缓存
	// 用户json 及 用户的token 对应的用户id
	userID := libs.SetValueToType(user.ID,"string").(string)
	saveTime := libs.SetValueToType(ConfigAdminRedis["redisSaveTime"],"int64").(int64)

	keyAdmin := ConfigAdminRedis["redisAdminPrefix"] + userID
	keyUid := ConfigAdminRedis["redisAdminUid"] + userID

	jsonBytes, _ := json.Marshal(user)

	orm.Redis.Set( keyUid, uid, saveTime )
	orm.Redis.Set( keyAdmin, string(jsonBytes), saveTime )

	c.JSON(http.StatusOK, gin.H{
		"code": 	1,
		"message": 	"获取成功",
		"data": 	uid,
	})
	return
}

//获取一个
func UserGetOne(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   	-1,
			"message": 	"参数传递错误",
		})
	}

	if err := user.GetOne(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   	-1,
			"message": 	err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": user,
	})
	return
}

//添加
func UserCreate(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    	-1,
			"message": 	err.Error(),
		})
		return
	}
	/*
		请在这里做一些私密性的信息限制, 如果创建只是指定部分数据,请清理部分不需要的数据或者建立特定的struct
	 */

	// 用户密钥
	user.Verify = strings.ToUpper( libs.MD5( libs.GetRandomString(10) ) );

	id, err := user.Create()
	if  err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    	-1,
			"message": 	err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  	1,
		"message": 	"添加成功",
		"data":    	id,
	})
	return
}

//删除数据
func UserDelete(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   	-1,
			"message": 	"参数传递错误",
		})
	}

	if result, err := user.Delete(id); err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "删除成功",
	})
	return
}

//修改数据
func UserUpdate(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err!=nil || id<=0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数传递错误",
		})
		return
	}

	updateMap := make( map[string] interface{})
	if err := c.ShouldBindJSON(&updateMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    	-1,
			"message": 	err.Error(),
		})
		return
	}

	delete(updateMap, "id" )

	// 更新用户密钥
	if value,ok := updateMap["verify"]; ok && value!="" {
		updateMap["verify"] = strings.ToUpper( libs.MD5( libs.GetRandomString(10) ) );
	}

		/*
			请在这里过滤一些业务不允许更新的字段
			如下
			delete(updateMap, "mobile" )
		*/
	err = user.Update(id,updateMap)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "修改成功",
	})
	return
}

//列表数据
func UserQuery(c *gin.Context) {
	var user model.User

	//查询条件
	searchMap := make( map[string]string )
	searchMap["status"] = c.DefaultQuery("status", "")
	searchMap["searchkey"] = c.DefaultQuery("searchkey", "")

	//页码 页数
	page,_ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	pageSize,_ := strconv.ParseInt(c.DefaultQuery("pageSize", "30"), 10, 64)
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 30
	}

	/*
		页码
	 */

	totalSize, err := user.QueryTotalSize(searchMap)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "查询信息失败 #pageSize",
		})
		return
	}
	pageMsg := make( map[string] int64 )
	pageMsg["page"] = page
	pageMsg["pageSize"] = pageSize
	pageMsg["totalSize"] = totalSize

	SF1 := strconv.FormatInt(totalSize, 10)
	totalSizeF, _ := strconv.ParseFloat(SF1, 64)
	SF2 := strconv.FormatInt(pageSize, 10)
	pageSizeF, _ := strconv.ParseFloat(SF2, 64)
	totalPageF := math.Ceil( totalSizeF/pageSizeF )
	pageMsg["totalPage"] = int64( totalPageF )

	/*
	查询结果
	 */
	result, err := user.Query(searchMap, page, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "查询信息失败 #dataArray",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
		"page": pageMsg,
	})
}

