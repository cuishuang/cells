package service

import (
	"errors"
	"fmt"
	"time"

	json "github.com/pydio/cells/v4/common/utils/jsonx"

	"go.uber.org/zap"

	"github.com/pydio/cells/v4/common/log"
	"github.com/pydio/cells/v4/common/registry"
	"github.com/pydio/cells/v4/common/server"
	servicecontext "github.com/pydio/cells/v4/common/service/context"
)

// Service for the pydio app
type service struct {
	opts *ServiceOptions
}

var _ registry.Service = (*service)(nil)

const (
	configSrvKeyFork      = "fork"
	configSrvKeyAutoStart = "autostart"
	configSrvKeyForkDebug = "debugFork"
	configSrvKeyUnique    = "unique"
)

var (
	mandatoryOptions []ServiceOption

	errNoServerAttached = errors.New("no server attached to the service")
)

type Service interface {
	Init(opts ...ServiceOption)
	Options() *ServiceOptions
	Metadata() map[string]string
	ID() string
	Name() string
	Start() error
	Stop() error
	IsGRPC() bool
	IsREST() bool
	IsGeneric() bool
	As(i interface{}) bool
}

type Stopper func() error

func NewService(opts ...ServiceOption) Service {
	s := &service{
		opts: newOptions(append(mandatoryOptions, opts...)...),
	}

	name := s.opts.Name

	s.opts.Context = servicecontext.WithServiceName(s.opts.Context, name)

	if reg := servicecontext.GetRegistry(s.opts.Context); reg != nil {
		if err := reg.Register(s); err != nil {
			log.Logger(s.opts.Context).Warn("could not register", zap.Error(err))
		}
	} else {
		log.Logger(s.opts.Context).Warn("no registry attached")
	}

	return s
}

func (s *service) Init(opts ...ServiceOption) {
	for _, o := range opts {
		if o == nil {
			continue
		}

		o(s.opts)
	}
}

func (s *service) Options() *ServiceOptions {
	return s.opts
}

func (s *service) Metadata() map[string]string {
	return s.opts.Metadata
}

func (s *service) As(i interface{}) bool {
	if v, ok := i.(*Service); ok {
		*v = s
		return true
	}

	if v, ok := i.(*registry.Service); ok {
		*v = s
		return true
	}

	return false
}

func (s *service) Start() (er error) {
	now := time.Now()

	defer func() {
		elapsed := time.Now().Sub(now)
		if elapsed > 5*time.Second {
			fmt.Println(s.Name(), "took ", elapsed, " to start")
		}

		if e := recover(); e != nil {
			log.Logger(s.opts.Context).Error("panic while starting service", zap.Any("p", e))
			er = fmt.Errorf("panic while starting service %v", e)
		}
	}()

	for _, before := range s.opts.BeforeStart {
		if err := before(s.opts.Context); err != nil {
			return err
		}
	}

	if s.opts.serverStart != nil {
		if err := s.opts.serverStart(); err != nil {
			return err
		}
	}

	for _, after := range s.opts.AfterStart {
		if err := after(s.opts.Context); err != nil {
			return err
		}
	}

	s.opts.Server.AfterServe(func() error {
		go UpdateServiceVersion(s.opts)

		return nil
	})

	for _, after := range s.opts.AfterServe {
		s.opts.Server.AfterServe(func() error {
			go after(s.opts.Context)

			return nil
		})
	}

	if reg := servicecontext.GetRegistry(s.opts.Context); reg != nil {
		go func() {
			if err := reg.Register(s); err != nil {
				log.Logger(s.opts.Context).Warn("could not register", zap.Error(err))
			}
		}()
	} else {
		log.Logger(s.opts.Context).Warn("no registry attached")
	}

	// fmt.Println("Going through here ? for ", s.opts.Name)
	log.Logger(s.opts.Context).Info("started")

	return nil
}

func (s *service) Stop() error {
	for _, before := range s.opts.BeforeStop {
		if err := before(s.opts.Context); err != nil {
			return err
		}
	}

	if s.opts.serverStop != nil {
		if err := s.opts.serverStop(); err != nil {
			return err
		}
	}

	for _, after := range s.opts.AfterStop {
		if err := after(s.opts.Context); err != nil {
			return err
		}
	}

	if reg := servicecontext.GetRegistry(s.opts.Context); reg != nil {
		if err := reg.Deregister(s); err != nil {
			log.Logger(s.opts.Context).Warn("could not deregister", zap.Error(err))
		}
	} else {
		log.Logger(s.opts.Context).Warn("no registry attached")
	}

	log.Logger(s.opts.Context).Info("stopped")

	return nil
}

func (s *service) Name() string {
	return s.opts.Name
}
func (s *service) ID() string {
	return s.opts.ID
}
func (s *service) Version() string {
	return s.opts.Version
}
func (s *service) Nodes() []registry.Node {
	if s.opts.Server == nil {
		return []registry.Node{}
	}
	return []registry.Node{s.opts.Server}
}
func (s *service) Tags() []string {
	return s.opts.Tags
}
func (s *service) IsGeneric() bool {
	return s.opts.serverType == server.ServerType_GENERIC
}
func (s *service) IsGRPC() bool {
	return s.opts.serverType == server.ServerType_GRPC
}
func (s *service) IsREST() bool {
	return s.opts.serverType == server.ServerType_HTTP
}

func (s *service) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.opts)
}

func (s *service) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, s.opts)
}
