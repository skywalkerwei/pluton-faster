package config

// SmsConf 短信
type SmsConf struct {
	AliYun struct {
		RegionId        string
		AccessKeyID     string
		AccessKeySecret string
		SignName        string
	}
	Debug         bool
	Length        int
	Life          int
	MaxCheckTimes int
	MagicCode     string
	TestUsers     []string
	Template      struct {
		Reg string
	}
}
