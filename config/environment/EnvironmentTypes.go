package environment

type Variables struct {
	Gin   Gin
	Gorm  Gorm
	Mongo Mongo
}

type Gin struct {
	Port string
}

type Gorm struct {
	Host          string
	Port          string
	DBName        string
	User          string
	Pass          string
	MigrationsDir string
}

type Mongo struct {
	Uri    string
	DBName string
}
