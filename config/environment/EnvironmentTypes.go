package environment

type Variables struct {
	Gin      Gin
	Grpc     Grpc
	Gorm     Gorm
	Mongo    Mongo
	Redis    Redis
	Services Services
}

type Gin struct {
	Port string
}

type Grpc struct {
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

type Redis struct {
	Host string
	Pass string
}

type Services struct {
	BoilerplateHost string
}
