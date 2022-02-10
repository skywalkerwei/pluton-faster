package sms

// Conf 短信
type Conf struct {
	AliYun struct {
		RegionId        string
		AccessKeyID     string
		AccessKeySecret string
		SignName        string
	}
	Debug     bool
	Length    int
	Life      int
	MagicCode string
	TestUsers []string
	Template  struct {
		Reg string
	}
}
