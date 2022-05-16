package database

type App struct {
	Id        uint `gorm:"PRIMAR_KEY"`
	Name      string
	Hostname  string
	Language  string
	CodeRepo  string
	ImageRepo string
}
