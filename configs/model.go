package configs

type Configuration struct {
	Gitlab `mapstructure:"gitlab" json:"gitlab"`
}

type Gitlab struct {
	Url          string `mapstructure:"url" json:"url"`
	PrivateToken string `mapstructure:"private_token" json:"private_token"`
	RootGroupId  int    `mapstructure:"root_group_id" json:"root_group_id"`
}
