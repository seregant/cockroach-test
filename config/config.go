package config

type Config struct {
	Host     string
	Port     string
	User     string
	DbName   string
	HttpPort string
	SrvKey   string
}

func SetConfig() Config {
	var config Config

	//set configuration here
	config.Host = "cockroachdb"
	config.Port = "26257" //cockroachdb's port
	config.User = "root"
	config.DbName = "karyawan"
	config.HttpPort = "1234"
	config.SrvKey = "Aw4s_g4l4k"
	return config
}
