package conf

import "github.com/gookit/ini/v2"

var config *ini.Ini

func LoadConfig() {
	config = ini.NewWithOptions(ini.ParseVar, func(opts *ini.Options) {
		opts.DefSection = "default"
	})
	err := config.LoadFiles("config/app.ini")
	if err != nil {
		panic(err.Error())
	}
}

func Get(env string) string {
	return config.String(env)
}
func GetInt(env string) int {
	return config.Int(env)
}
func GetBool(env string) bool {
	return config.Bool(env)
}
