package service

import (
	"log"

	"github.com/eqto/api-server"
	"github.com/eqto/config"
)

//Service ..
type Service struct {
	cfg  config.Config
	apis *api.Server

	port int
	log  struct {
		debug func(d ...interface{})
	}
}

func (s *Service) debug(d ...interface{}) {
	if s.log.debug != nil {
		s.log.debug(d...)
	} else {
		log.Println(d...)
	}
}

//Run ...
func (s *Service) Run() error {
	for i := 0; i < 10; i++ {
		e := s.apis.Serve(8800 + i)
		if e != nil {
			s.debug(e)
		}
	}
	return nil

}

//New ...
func New(cfg config.Config) *Service {
	return &Service{cfg: cfg, apis: api.New()}
}
