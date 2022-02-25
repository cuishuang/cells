package broker

import (
	"fmt"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"

	"github.com/pydio/cells/v4/common"
	"github.com/pydio/cells/v4/common/broker"
	pb "github.com/pydio/cells/v4/common/proto/broker"
)

type Handler struct {
	pb.UnimplementedBrokerServer

	broker      broker.Broker
	subscribers map[string][]string
}

func NewHandler(b broker.Broker) *Handler {
	return &Handler{
		broker: b,
	}
}

func (h *Handler) Name() string {
	return common.ServiceGrpcNamespace_ + common.ServiceBroker
}

func (h *Handler) Publish(stream pb.Broker_PublishServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if !strings.Contains(err.Error(), "context canceled") {
				fmt.Println("[discovery/grpc/broker] Publish: Error is ", err)
			}
			return err
		}

		for _, message := range req.Messages {
			if err := h.broker.PublishRaw(stream.Context(), req.Topic, message.Body, message.Header); err != nil {
				return err
			}
		}
	}
}

func (h *Handler) Subscribe(stream pb.Broker_SubscribeServer) error {
	var unSubscribers []broker.UnSubscriber
	defer func() {
		for _, unSubscriber := range unSubscribers {
			_ = unSubscriber()
			//fmt.Println("Unsubscribing on stream closed", e)
		}
	}()
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		topic := req.GetTopic()
		// queue := req.GetQueue()
		subPID := ""
		if md, ok := metadata.FromIncomingContext(stream.Context()); ok {
			subPID = strings.Join(md["subscriberid"], "")
		}

		/*
			// TODO V4 : we could batch ? but it makes the events slower
			responses := make(chan broker.Message)
			var buffer []broker.Message
			timer := time.NewTimer(100 * time.Millisecond)
			maxSize := 10
			go func() {
				flush := func() {
					if len(buffer) == 0 {
						return
					}
					mm := &pb.SubscribeResponse{Id: req.Id}
					for _, msg := range buffer {
						var target = &pb.Message{}
						target.Header, target.Body = msg.RawData()
						mm.Messages = append(mm.Messages, target)
					}
					done := make(chan bool)
					defer close(done)
					go func() {
						select {
						case <-done:
							return
						case <-time.After(10 * time.Second):
							fmt.Println("GRPC broker stream.Send stuck after 10s - "+topic, "subscriber was "+subPID)
							return
						}
					}()
					fmt.Println("GRPC broker sending messages", len(mm.Messages))
					if err := stream.Send(mm); err != nil {
						fmt.Println("GRPC broker: stream.Send error", err)
					}
					buffer = []broker.Message{}
				}
				for {
					select {
					case <-timer.C:
						flush()
						timer = time.NewTimer(100 * time.Millisecond)
					case m := <-responses:
						buffer = append(buffer, m)
						if len(buffer) >= maxSize {
							flush()
						}
					}
				}
			}()

		*/

		unSub, e := h.broker.Subscribe(stream.Context(), topic, func(msg broker.Message) error {
			// For batching instead
			// responses <- msg
			mm := &pb.SubscribeResponse{}
			var target = &pb.Message{}
			target.Header, target.Body = msg.RawData()

			// fmt.Println("And the message we've received is ? ", string(target.Body))
			// mm.Queue = queue
			mm.Id = req.Id
			mm.Messages = append(mm.Messages, target)
			done := make(chan bool)
			defer close(done)
			go func() {
				select {
				case <-done:
					return
				case <-time.After(10 * time.Second):
					fmt.Println("GRPC broker stream.Send stuck after 10s - "+topic, "subscriber was "+subPID)
					return
				}
			}()
			if err := stream.Send(mm); err != nil {
				return err
			}
			return nil
		})
		if e != nil {
			return e
		}
		unSubscribers = append(unSubscribers, unSub)
	}
}
