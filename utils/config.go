package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var Config = &viperConfigReader{}

type viperConfigReader struct {
	viper *viper.Viper
}

type Reader interface {
	GetAllKeys() []string
	Get(key string) interface{}
	GetBool(key string) bool
	GetString(key string) string
}

type Conf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
}

func (c *Conf) GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Errorf("cannot unmarshal data: %v\n", err.Error())
	}
	return c
}

func (configReader *viperConfigReader) LoadConfig() {
	configReader.viper = configReader.LoadConfigFromYaml("app", "yaml", ".")
}

func (configReader *viperConfigReader) LoadConfigFromYaml(configName, configType, configPath string) (v *viper.Viper) {
	v = viper.New()
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	v.AddConfigPath(configPath)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//go func() {
	//	for {
	//		time.Sleep(time.Second * 5)
	//		v.WatchConfig()
	//		v.OnConfigChange(func(in fsnotify.Event) {
	//			log.Println("config file changed", in.Name)
	//		})
	//	}
	//}()

	return
}

func (configReader *viperConfigReader) GetAllKeys() []string {
	return configReader.viper.AllKeys()
}

func (configReader *viperConfigReader) Get(key string) interface{} {
	return configReader.viper.Get(key)
}

func (configReader *viperConfigReader) GetBool(key string) bool {
	return configReader.viper.GetBool(key)
}

func (configReader *viperConfigReader) GetString(key string) string {
	return configReader.viper.GetString(key)
}
