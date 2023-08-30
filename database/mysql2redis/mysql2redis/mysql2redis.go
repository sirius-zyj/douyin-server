package mysql2redis

import (
	"github.com/duanhf2012/origin/log"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
)

type Mysql2RedisService struct {
	service.Service

	canals map[string]*CommonCanalMeta
}

func init() {
	node.Setup(&Mysql2RedisService{})
}

func (s *Mysql2RedisService) OnInit() error {
	cfg := s.GetServiceCfg().(map[string]interface{})
	path := cfg["ConfigPath"].(string)

	s.canals = make(map[string]*CommonCanalMeta)

	LoadConfigWithFile(path)
	for name, currConfig := range GetConfigure().CanalConfigs {
		go func(cfg *CanalConfig, name string) {
			defer func() {
				err := recover()
				if err != nil {
					log.Error("canal[%s] crash, err:%s", name, err)
				}
			}()
			currCancal := &CommonCanalMeta{}

			currCancal.RunWithConfig(name, cfg)
			s.canals[name] = currCancal

		}(currConfig, name)

	}

	return nil
}

func (s *Mysql2RedisService) OnRelease() {

	for n, c := range s.canals {
		log.Debug("closing mysql2redis inst:%s", n)
		c.Close()
	}

}
