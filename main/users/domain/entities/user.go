package entities

type User struct {
	Id       string `mapstructure:"id"`
	Name     string `mapstructure:"name"`
	Code     string `mapstructure:"code"`
	Password string `mapstructure:"password"`
	Root     bool   `json:"root"`
}
