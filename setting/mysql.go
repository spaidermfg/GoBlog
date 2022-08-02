package setting


import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

//mysql全局配置

var db *sqlx.DB

type regions struct {
	REGION_ID int
	REGION_NAME string
}

//初始化连接
func InitDB(cfg *MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		cfg.Name,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect to database fail")
		return
	}

	db.SetMaxOpenConns(10)//设置连接池中最大连接数
	db.SetMaxIdleConns(10)//设置连接池中最大闲置连接数
	zap.L().Debug("init mysql success")


	//建立数据库查询
	sqlStr := fmt.Sprintf("select * from %s where region_id=?", "regions")
	fmt.Println(sqlStr)
	var r regions
	err = db.Get(&r, sqlStr, 1)

	if err != nil {
		fmt.Println("select data fail",err)
	}
	fmt.Printf("order_id: %v, order_name: %v\n", r.REGION_ID, r.REGION_NAME)

	return nil
}

//关闭连接
func CloseMdb() {
	_ = db.Close()
}