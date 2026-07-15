package models

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/pkg/errors"
	"time"

	_ "github.com/lib/pq"
)

type BaseModel struct {
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);index" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime);index" json:"updated_at"`
	IsRemoved int8      `orm:"default(0);index"` // 0: 正常，1: 删除
}

func init() {
	postgresConfig, err := beego.AppConfig.GetSection("postgres")
	if err != nil {
		panic(errors.Wrap(err, "get postgres config failed"))
	}

	fmt.Println("init database start")

	dbAlias := postgresConfig["db_alias"]
	dbType := postgresConfig["db_type"]
	dbUser := postgresConfig["db_user"]
	dbPass := postgresConfig["db_pass"]
	dbHost := postgresConfig["db_host"]
	dbPort := postgresConfig["db_port"]
	dbName := postgresConfig["db_name"]
	dbSSLMode := postgresConfig["db_sslmode"]

	if dbAlias == "" {
		dbAlias = "default"
	}

	if dbType == "" {
		dbType = "postgres"
	}

	if dbPort == "" {
		dbPort = "5432"
	}

	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}

	dbURL := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		dbSSLMode,
	)

	if err := orm.RegisterDataBase(dbAlias, dbType, dbURL); err != nil {
		panic(errors.Wrap(err, "register database failed"))
	}

	orm.RegisterModel(
		new(User),
		new(UserInfo),
		new(UserWallet),
		new(UserIntegral),
		new(UserCoupon),
		new(CrfrUserTree),
	)
	key, err := beego.AppConfig.String("runmode")
	if err != nil {
		logs.Error(err)
		return
	}
	if key == "dev" {
		orm.Debug = true
	}
	err = orm.RunSyncdb(dbAlias, false, true)
	if err != nil {
		logs.Error(err.Error())
	}
	fmt.Println("init database success")
}
