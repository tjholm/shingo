package events

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	v1 "github.com/tjholm/shingo/pkg/api/shingo/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Subscriber struct {
	stream v1.EventsService_SubscribeServer
}

// EventsServer - Implementation for Shingo events server
type EventsServer struct {
	v1.UnimplementedEventsServiceServer
	subscriptions map[string]map[string]*Subscriber
	subLock       sync.RWMutex
}

func (e *EventsServer) Emit(ctx context.Context, req *v1.EmitRequest) (*v1.EmitResponse, error) {
	// send messages in a goroutine
	// so don't block sending back an ack to the event emitter
	go func() {
		e.subLock.RLock()
		defer e.subLock.RUnlock()
		// If we have active subscriptions then send out the message :)
		if subs, ok := e.subscriptions[req.Topic]; ok {
			for id, sub := range subs {
				// skip the actual subscriber (this can be provided by the subscriber to filter out their own requests)
				if req.SubscriberId == id {
					continue
				}

				sub.stream.Send(&v1.SubscriptionEvent{
					Event: &v1.SubscriptionEvent_OnMessage{
						OnMessage: &v1.OnMessageEvent{
							Event: req.Event,
						},
					},
				})
			}
		}
	}()

	// return the response, TODO: We'll actually want to handle the sending out of messages in a separate go-routine
	// and return ASAP rather than blocking...
	return &v1.EmitResponse{}, nil
}

func (e *EventsServer) Subscribe(req *v1.SubscribeRequest, srv v1.EventsService_SubscribeServer) error {
	// Register subscription server
	e.subLock.Lock()
	subid := uuid.NewString()

	if e.subscriptions[req.Topic] == nil {
		e.subscriptions[req.Topic] = make(map[string]*Subscriber)
	}

	e.subscriptions[req.Topic][subid] = &Subscriber{
		stream: srv,
	}
	e.subLock.Unlock()

	if err := srv.Send(&v1.SubscriptionEvent{
		Event: &v1.SubscriptionEvent_OnSubscribe{
			OnSubscribe: &v1.OnSubscribeEvent{
				SubscriberId: subid,
			},
		},
	}); err != nil {
		return status.Errorf(codes.Unavailable, "unable to send subscription confirmation")
	}

	<-srv.Context().Done()

	fmt.Printf("subscriber %s has ended their stream\n", subid)
	// remove them from subscriptions
	e.subLock.Lock()
	delete(e.subscriptions[req.Topic], subid)
	e.subLock.Unlock()

	// return when subscription is cancelled
	return nil
}

func New() *EventsServer {
	return &EventsServer{
		subscriptions: make(map[string]map[string]*Subscriber),
		subLock:       sync.RWMutex{},
	}
}
