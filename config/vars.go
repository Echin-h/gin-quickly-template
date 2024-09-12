package config

type GlobalConfig struct {
	MODE    string `yaml:"Mode"`
	AppName string `yaml:"AppName"`
	AUTHOR  string `yaml:"Author"`
	VERSION string `yaml:"Version"`
	Host    string `yaml:"Host"`
	Port    string `yaml:"Port"`
	Log     struct {
		LogPath string `yaml:"LogPath"`
		CLS     struct {
			Endpoint    string `yaml:"Endpoint"`
			AccessKey   string `yaml:"AccessKey"`
			AccessToken string `yaml:"AccessToken"`
			TopicID     string `yaml:"TopicID"`
		} `yaml:"CLS"`
	} `yaml:"Log"`
	Database struct {
		MODE  string //  logger mode
		Mysql struct {
			Host      string `yaml:"Host"`
			Port      string `yaml:"Port"`
			Username  string `yaml:"Username"`
			Password  string `yaml:"Password"`
			DBName    string `yaml:"DBName"`
			Charset   string `yaml:"Charset"`
			ParseTime string `yaml:"ParseTime"`
			Loc       string `yaml:"Loc"`
		} `yaml:"Mysql"`
		Postgres struct {
			Host     string `yaml:"Host"`
			Port     string `yaml:"Port"`
			Username string `yaml:"Username"`
			Password string `yaml:"Password"`
			DBName   string `yaml:"DBName"`
			SSLMode  string `yaml:"SSLMode"`
			TimeZone string `yaml:"TimeZone"`
		} `yaml:"Postgres"`
		Redis struct {
			Addr     string `yaml:"Addr"`
			Password string `yaml:"Password"`
			DB       int    `yaml:"DB"`
		} `yaml:"Redis"`
	} `yaml:"Database"`
	Auth struct {
		Secret string `yaml:"Secret"`
		Issuer string `yaml:"Issuer"`
	} `yaml:"Auth"`
	//Databases []Datasource `yaml:"Databases"`
	//Caches    []Cache      `yaml:"Caches"`
	//OSS       Oss          `yaml:"Oss"`
	//Mail      Mail         `yaml:"Mail"`
	//CMS       Cms          `yaml:"Cms"`
}
