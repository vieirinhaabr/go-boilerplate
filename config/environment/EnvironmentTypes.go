package environment

type Variables struct {
	Gin struct {
		Port string
	}
	Gorm struct {
		Host   string
		Port   string
		DBName string
		User   string
		Pass   string
	}
}
