package config

type Config struct {
	Host   string
	Port   string
	User   string
	DbName string
}

func SetConfig() Config {
	var config Config

	//set configuration here
	config.Host = "localhost"
	config.Port = "26257" //cockroachdb's port
	config.User = "maxcroach"
	config.DbName = "karyawan"

	return config
}
