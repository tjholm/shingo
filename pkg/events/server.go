package events

import (
	"context"
	"sync"

	v1 "github.com/tjholm/shingo/pkg/api/shingo/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EventsServer - Implementation for Shingo events server
type EventsServer struct {
	v1.UnsafeEventsServiceServer
	subscriptions map[string][]v1.EventsService_SubscribeServer
	subLock       sync.RWMutex
	// maintain a map of topics to and SubscriptionServers to send them on
}

func (e *EventsServer) Emit(ctx context.Context, req *v1.EmitRequest) (*v1.EmitResponse, error) {
	e.subLock.RLock()
	defer e.subLock.RUnlock()
	// If we have active subscriptions then send out the message :)
	if srvs, ok := e.subscriptions[req.Topic]; ok {
		for _, srv := range srvs {
			// TODO: Filter out own subscriber
			srv.Send(&v1.SubscriptionEvent{
				Event: &v1.SubscriptionEvent_OnMessage{
					OnMessage: &v1.OnMessageEvent{
						Event: req.Event,
					},
				},
			})
		}
	}

	// return the response, TODO: We'll actually want to handle the sending out of messages in a separate go-routine
	// and return ASAP rather than blocking...
	return &v1.EmitResponse{}, nil
}

func (e *EventsServer) Subscribe(req *v1.SubscribeRequest, srv v1.EventsService_SubscribeServer) error {
	// Register subscription server
	e.subLock.Lock()
	defer e.subLock.Unlock()

	e.subscriptions[req.Topic] = append(e.subscriptions[req.Topic], srv)

	// TODO: We need to block here, we may want to block on a channel that tracks the subscribers status???
	// Generate an Id for the subscriber and allow them to voluntarily unsubscribe?

	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func New() *EventsServer {
	return &EventsServer{
		subscriptions: make(map[string][]v1.EventsService_SubscribeServer),
		subLock:       sync.RWMutex{},
	}
}
