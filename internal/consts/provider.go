package consts

// Provider 数据存储类型
type Provider string

const (
	MySql      Provider = "mysql"
	Sqlite     Provider = "sqlite"
	Redis      Provider = "redis"
	Memory     Provider = "memory"
	ClickHouse Provider = "clickhouse"
	PgSql      Provider = "pgsql"
)
