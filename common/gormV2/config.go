package gormV2

type Conf struct {
	SlowThreshold int
	IsOpenReadDb  bool
	Write         struct {
		Host               string
		Port               int
		User               string
		Pass               string
		Charset            string
		DataBase           string
		Prefix             string
		SetMaxIdleConns    int
		SetMaxOpenConns    int
		SetConnMaxLifetime int
	}
	Read struct {
		Host               string
		Port               int
		User               string
		Pass               string
		Charset            string
		DataBase           string
		Prefix             string
		SetMaxIdleConns    int
		SetMaxOpenConns    int
		SetConnMaxLifetime int
	}
}
