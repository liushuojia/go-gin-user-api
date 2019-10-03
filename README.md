# go-gin-user-api

使用golang语言开发
1. gin 框架
2. gorm 处理mysql数据库
3. 使用redis配合存储登录用户的数据及token验证


接口文档说明
/doc/UserAPI说明文档.docx

go环境配置
go安装

    配置环境变量
        sudo vi etc/profile
        或者
        vi  .bashrc（这个是全局变量）
        export GOPATH=$HOME/go
        export PATH=$PATH:$GOPATH/bin
    更新一下文件
        source etc/profile
        或者
        source .bashrc
    查看环境
        go env


环境支持

    设置下代理加快速度
    export GOPROXY=https://goproxy.io

    go get github.com/go-sql-driver/mysql
    go get github.com/jinzhu/gorm
    go get github.com/gin-gonic/gin

    编译时候如果有部分包没有
        mkdir -p $GOPATH/src/golang.org/x
        cd $GOPATH/src/golang.org/x
        git clone https://github.com/golang/sys.git

    读取配置文件
    go get github.com/Unknwon/goconfig

    redigo
    go get github.com/garyburd/redigo/redis

发布支持
    env GOOS=linux GOARCH=amd64 go build main.go


mysql redis 持续化连接


redis启动
　   vim redis.conf
    daemonize => yes    //后台运行

    redis-server /usr/local/etc/redis.conf


userAPI 功能描述
1. 获取授权token
2. 验证授权token
3. 用户管理
4. 获取用户自身资料

