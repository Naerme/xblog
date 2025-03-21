package conf

type Config struct {
	System System `yaml:"system"`
	Log    log    `yaml:"log"`
	DB     DB     `yaml:"db"`
	DB1    DB     `yaml:"db1"`
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
	Site   Site   `yaml:"site"`
	Email  Email  `yaml:"email"`
	QQ     QQ     `yaml:"qq"`
	QiNiu  QiNiu  `yaml:"qiNiu"`
	Ai     Ai     `yaml:"ai"`
	Upload Upload `yaml:"upload"`
}
