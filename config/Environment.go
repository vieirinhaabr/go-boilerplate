package config

type EnvironmentVariables struct {
	Gin struct {
		Port int
	}
	Gorm struct {
		Host   string
		Port   int
		DBName string
		User   string
		Pass   string
	}
}

func InitEnvVars() *EnvironmentVariables {

	return new(EnvironmentVariables)
}
