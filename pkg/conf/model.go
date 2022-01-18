package conf

type account struct {
	BiliBili bilibili `mapstructure:"bilibili"`
	Huya     huya     `mapstructure:"huya"`
	Douyu    douyu    `mapstructure:"douyu"`
}

type huya struct {
	Enable  bool   `mapstructure:"enable"`
	Cookies string `mapstructure:"cookies"`
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
