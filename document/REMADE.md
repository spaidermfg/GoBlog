#使用GoLang搭建一个博客

## 0.奠基仪式
**2022-06-08 14:30**

## 1.初步待实现功能
* 首页
* 新增博客
* 删除博客
* 查看博客
* 查询
* md解析

## 2.技术选型
1. 框架
    - gin
    - 获取命令：go get -u github.com/gin-gonic/gin
    - 项目地址：https://github.com/gin-gonic/gin
    - 官方文档：https://gin-gonic.com/zh-cn/docs/
2. 日志
    - zap
    - 获取命令：go get go.uber.org/zap
    - 项目地址：https://github.com/uber-go/zap
3. 数据库
    - mysql
    - 下载驱动：go get -u github.com/go-sql-driver/mysql
4. 访问数据库
    - database/sql接口
    - sqlx
    - go get github.com/jmoiron/sqlx
5. ORM框架
    - 中文文档： https://gorm.io/zh_CN/
    -
6. 配置管理工具
    - viper
    - 获取命令：go get github.com/spf13/viper
    - 中文文档：https://www.liwenzhou.com/posts/Go/viper_tutorial/
    - 官方文档：https://github.com/spf13/viper/blob/master/README.md
7. 日志切割归档工具
    - Lumberjack
    - 获取命令：go get -u github.com/natefinch/lumberjack
8. 代码管理工具
    - github
    - 初始化仓库：git init
    - 本地仓库关联远程仓库：git remote add origin git@gitee.com:KKKLxxx/study-notes.git
9. JWT跨域认证
    - 获取命令：go get github.com/golang-jwt/jwt/v4
    - 项目地址：https://gorm.io/zh_CN/
10. 热重载
    - 检查代码修改，自动编译重启项目
11. 配置redis
    - 获取命令： go get -u github.com/gomodule/redigo

  
 
