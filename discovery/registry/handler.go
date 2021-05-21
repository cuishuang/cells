package grpc

import (
	"context"
	"time"

	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/registry"
	"github.com/pydio/cells/common/micro/registry/service"
	pb "github.com/pydio/cells/common/proto/registry"
)

var (
	NotImplemented = errors.New("notimplemented", "service not implemented", 501)
)

type Handler struct{}

func (h *Handler) GetService(ctx context.Context, req *pb.GetRequest, resp *pb.GetResponse) error {
	ss, err := registry.GetService(req.GetService())
	if err != nil {
		return err
	}

	for _, s := range ss {
		resp.Services = append(resp.Services, service.ToProto(s))
	}

	return nil
}
func (h *Handler) Register(ctx context.Context, s *pb.Service, resp *pb.EmptyResponse) error {
	return registry.Register(service.ToService(s), registry.RegisterTTL(time.Duration(s.GetOptions().GetTtl())*time.Second))
}

func (h *Handler) Deregister(ctx context.Context, s *pb.Service, resp *pb.EmptyResponse) error {
	return registry.Deregister(service.ToService(s))
}

func (h *Handler) ListServices(ctx context.Context, req *pb.ListRequest, resp *pb.ListResponse) error {
	ss, err := registry.ListServices()
	if err != nil {
		return err
	}

	for _, s := range ss {
		resp.Services = append(resp.Services, service.ToProto(s))
	}

	return nil
}

func (h *Handler) Watch(ctx context.Context, req *pb.WatchRequest, stream pb.Registry_WatchStream) error {
	// defer stream.Close()

	var opts []registry.WatchOption
	if s := req.GetService(); s != "" {
		opts = append(opts, registry.WatchService(s))
	}

	w, err := registry.Watch(opts...)
	if err != nil {
		return err
	}

	for {
		res, err := w.Next()
		if err != nil {
			return err
		}

		stream.Send(&pb.Result{
			Action:  res.Action,
			Service: service.ToProto(res.Service),
		})
	}

	return nil
}
