package conf

type config struct {
	Server  server  `mapstructure:"server"`
	Socks5  socks5  `mapstructure:"socks5"`
	Account account `mapstructure:"account"`
}
type server struct {
	Port  int    `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
	Path  string `mapstructure:"path"`
}

type account struct {
	BiliBili bilibili `mapstructure:"bilibili"`
	Huya     huya     `mapstructure:"huya"`
	Douyu    douyu    `mapstructure:"douyu"`
}

type huya struct {
	Enable bool `mapstructure:"enable"`
}

type douyu struct {
	Enable bool `mapstructure:"enable"`
}

type bilibili struct {
	Enable          bool   `mapstructure:"enable"`
	DedeUserID      string `mapstructure:"DedeUserID"`      // DedeUserID
	DedeUserIDCkMd5 string `mapstructure:"DedeUserIDCkMd5"` // DedeUserID__ckMd5
	SESSDATA        string `mapstructure:"SESSDATA"`        // SESSDATA
	BiliJCT         string `mapstructure:"BiliJCT"`         // bili_jct
}

type socks5 struct {
	Enable   bool   `mapstructure:"enable"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
