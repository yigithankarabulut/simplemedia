package config

//
//type Config struct {
//	Database Database `mapstructure:"database"`
//	Port     string   `mapstructure:"port"`
//}
//
//type Database struct {
//	Name    string `mapstructure:"name"`
//	Host    string `mapstructure:"host"`
//	Pass    string `mapstructure:"pass"`
//	User    string `mapstructure:"user"`
//	Port    string `mapstructure:"port"`
//	Migrate bool   `mapstructure:"migrate"`
//}
//
//func Load(configName string) (*Config, error) {
//	var BaseConfig Config
//	pwd, err := os.Getwd()
//	if err != nil {
//		return nil, errors.New("System Configuration Error: " + err.Error())
//	}
//
//	viper.SetConfigName(configName)
//	viper.SetConfigType("yml")
//	viper.SetTypeByDefaultValue(true)
//	viper.AddConfigPath(pwd)
//
//	err = viper.ReadInConfig()
//	if err != nil {
//		return nil, errors.New("System Configuration Error: " + err.Error())
//	}
//
//	err = viper.Unmarshal(&BaseConfig)
//	if err != nil {
//		return nil, errors.New("System Configuration Error: " + err.Error())
//	}
//
//	return &BaseConfig, nil
//}
//
//func (c *Config) Validate() error {
//	if c.Database.Host == "" {
//		return errors.New("Database host cannot found")
//	}
//	if c.Database.Name == "" {
//		return errors.New("Database name cannot found")
//	}
//	if c.Database.Pass == "" {
//		return errors.New("Database pass cannot found")
//	}
//	if c.Database.User == "" {
//		return errors.New("Database user cannot found")
//	}
//	if c.Database.Port == "" {
//		return errors.New("Database port cannot found")
//	}
//	return nil
//}
