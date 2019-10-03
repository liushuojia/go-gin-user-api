package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "user/apis"
)

// 全局中间件
func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		/*
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		//fmt.Printf("%s::%s \t %s \t %s ", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		/*
		context.JSON(http.StatusBadRequest, gin.H{
			"code":   	-1,
			"message": 	"参数传递错误",
		})
		context.Abort()
		 */
		context.Next()

		//执行完后提示
		//fmt.Println(context.Writer.Status())
	}
}

// 局部中间件, 要求token必须已经登录
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		userAgent := c.GetHeader("User-Agent")
		token := c.GetHeader("token")

		user,err := CheckUid(token,userAgent)
		if err != nil || user.ID ==0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":   	-1,
				"message": 	"token 失效",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 404
	router.NoRoute(NoResponse)

	// 全局中间件
	//router.Use(Logger(), gin.Recovery())


	// user
	user := router.Group("user", Auth() )//通过Group第二个参数，使用中间件
	{
		user.GET("",  UserQuery)
		user.POST("", UserCreate)

		user.GET(":id", UserGetOne)
		user.PUT(":id", UserUpdate)
		user.DELETE(":id", UserDelete)
	}

	// 用户操作自己的信息
	myself := router.Group("my", Auth() )//通过Group第二个参数，使用中间件
	{
		myself.GET("", UserGetOne)
		myself.PUT("", UserGetOne)
	}

	// token 操作
	router.GET("/token/get", GetToken)
	router.GET("/token/check", CheckToken)

	/*
		router.GET("/user",  UserQuery)
		router.POST("/user", UserCreate)
		router.GET("/user/:id", Auth(), UserGetOne)
		router.PUT("/user/:id", UserUpdate)
		router.DELETE("/user/:id", UserDelete)
	*/

	return router
}

//NoResponse 请求的url不存在，返回404
func NoResponse(c *gin.Context) {
	//返回404状态码
	c.JSON(http.StatusNotFound, gin.H{
		"code":   	-1,
		"message": 	"404, page not exists!",
	})
	return
}