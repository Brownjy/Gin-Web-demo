# Gin-Web-demo
gin网络框架demo

目录结构
```text

- test
|- - api //接口控制器包
|- - - controller.go //控制器文件
|- - common //公共函数包
| - - - functions.go //公共函数文件
|- - config //配置文件包
| - - - config.go //配置加载文件
| - - - config.yml //配置文件
|- - model //模型文件包
| - - - database.go //数据库连接初始化文件
| - - - pw_member.go //表模型文件
|- - resource // 静态文件夹
|- - route //路由文件包
| - - - router.go //路由配置文件
main.go // 入口文件 （图片显示 go-main.go 是我自己定义的文件，改成main.go即可）
go.mod // go 依赖管理文件
```
项目配置
```shell
## gin
go get -u github.com/gin-gonic/gin
## viper
go get github.com/spf13/viper
## mysql & gorm
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

```