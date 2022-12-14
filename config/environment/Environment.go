package environment

import (
	"github.com/joho/godotenv"
	"os"
)

var Vars = new(Variables)

func loadVar(name string, required bool) string {
	env, err := os.LookupEnv(name)
	if err == false && required {
		panic("Env var " + name + " not found")
	}

	return env
}

func SetupVars() {
	_ = godotenv.Load()

	Vars.Gin.Port = loadVar("GIN_PORT", true)

	Vars.Gorm.Host = loadVar("GORM_HOST", true)
	Vars.Gorm.Port = loadVar("GORM_PORT", true)
	Vars.Gorm.DBName = loadVar("GORM_DBNAME", true)
	Vars.Gorm.User = loadVar("GORM_USER", true)
	Vars.Gorm.Pass = loadVar("GORM_PASS", true)
}
