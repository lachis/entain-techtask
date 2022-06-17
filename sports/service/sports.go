package service

import (
	"git.neds.sh/matty/entain/sports/db"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"golang.org/x/net/context"
)

type Sports interface {
	// ListEvents will return a collection of events.
	ListEvents(ctx context.Context, in *sports.ListEventRequest) (*sports.ListEventResponse, error)
}

// sportsService implements the Sports interface.
type sportsService struct {
	eventsrepo db.EventsRepo
}

// NewEventsService instantiates and returns a new eventsService.
func NewSportsService(eventsRepo db.EventsRepo) Sports {
	return &sportsService{eventsRepo}
}

func (s *sportsService) ListEvents(ctx context.Context, in *sports.ListEventRequest) (*sports.ListEventResponse, error) {
	events, err := s.eventsrepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &sports.ListEventResponse{Events: events}, nil
}
