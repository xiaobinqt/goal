package config

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/database"
	"github.com/qbhy/goal/utils"
)

func init() {
	configs["database"] = func(env contracts.Env) interface{} {
		return database.Config{
			Default: utils.StringOr(env.GetString("db.connection"), "mysql"),
			Connections: map[string]contracts.Fields{
				"mysql": {
					"driver":      "mysql",
					"host":        env.GetString("db.host"),
					"port":        env.GetString("db.port"),
					"database":    env.GetString("db.database"),
					"username":    env.GetString("db.username"),
					"password":    env.GetString("db.password"),
					"unix_socket": env.GetString("db.unix_socket"),
					"charset":     utils.StringOr(env.GetString("db.charset"), "utf8mb4"),
					"collation":   utils.StringOr(env.GetString("db.collation"), "utf8mb4_unicode_ci"),
					"prefix":      env.GetString("db.prefix"),
					"strict":      env.GetBool("db.struct"),
				},
			},
		}
	}
}