package mysql2redis

import (
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/duanhf2012/origin/log"
)

type CanalConfig struct {
	Cancalconfigpath string `toml:"cancalconfigpath"`
	Posconfigfile    string `toml:"posconfigfile"`
	Bulksize         int    `toml:"bulksize"`
	// WriteInterval    int          `toml:"writeinterval"`
	Flushbulktime int64        `toml:"flushbulktime"`
	CacheSize     int64        `toml:"cachesize"`
	Redis         *RedisConfig `toml:"redis"`
	LogPos        *Pos
	//	Target           map[string]string    `toml:"target"`
}

type Pos struct {
	Name string `toml:"bin_name"`
	Pos  uint32 `toml:"bin_pos"`
}

type CommonConfig interface {
	GetConfigName() string
}

type RedisConfig struct {
	CommonConfig
	Address  string                       `toml:"address"`
	Password string                       `toml:"password"`
	Tables   map[string]*RedisTableConfig `toml:"tables"`
}

func (this *RedisConfig) GetConfigName() string {
	return "Redis"
}

type RedisTableConfig struct {
	DB int `toml:"db"`

	Tablename string `toml:"tablename"`
	//Struct     string   `toml:"struct"`
	Key string `toml:"key"`
}

type Configure struct {
	CanalConfigs map[string]*CanalConfig `toml:"canal"`
}

var configure *Configure

func LoadConfigWithFile(name string) error {
	data, readFileErr := ioutil.ReadFile(name)
	if readFileErr != nil {
		log.Error(readFileErr.Error())
		return readFileErr
	}

	conf := &Configure{}
	_, decodeTomlErr := toml.Decode(string(data), &conf)
	if decodeTomlErr != nil {
		log.Error(decodeTomlErr.Error())
		return decodeTomlErr
	}

	for _, currCanalConfig := range conf.CanalConfigs {
		if strings.Trim(currCanalConfig.Posconfigfile, " ") != "" {
			currPos, readPosErr := ioutil.ReadFile(currCanalConfig.Posconfigfile)
			if readPosErr != nil {
				log.Error(readPosErr.Error())
				return readPosErr
			}
			pos := &Pos{}
			_, decodePosErr := toml.Decode(string(currPos), &pos)

			if decodePosErr != nil {
				log.Error(decodePosErr.Error())
				return decodePosErr
			}

			if pos.Name != "" {
				currCanalConfig.LogPos = pos
			}
		}

	}

	configure = conf

	return nil
}

func GetConfigure() *Configure {
	return configure
}
