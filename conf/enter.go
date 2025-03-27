package conf

type Config struct {
	System System `yaml:"system"`
	Log    log    `yaml:"log"`
	DB     []DB   `yaml:"db"` //数据库链接列表
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
	Site   Site   `yaml:"site"`
	Email  Email  `yaml:"email"`
	QQ     QQ     `yaml:"qq"`
	QiNiu  QiNiu  `yaml:"qiNiu"`
	Ai     Ai     `yaml:"ai"`
	Upload Upload `yaml:"upload"`
	ES     ES     `yaml:"es"`
	River  River  `yaml:"river"`
}
