package model

type (
	Config struct {
		App     AppConfig     `yaml:"app"`
		Adapter AdapterConfig `yaml:"adapter"`
	}

	AppConfig struct {
		Address string `yaml:"address"`
	}

	AdapterConfig struct {
		Mongo MongoConfig `yaml:"mongo"`
	}

	MongoConfig struct {
		Address    string `yaml:"address"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		AuthSource string `yaml:"auth_source"`
	}
)
