package db

import (
	"database/sql"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"

	"git.neds.sh/matty/entain/sports/proto/sports"
)

//EventsRepo provides repository access to events.
type EventsRepo interface {
	//Init will initialise our events repository.
	Init() error

	//List will return a list of events.
	List(filter *sports.ListEventRequestFilter) ([]*sports.Event, error)
}

type eventsRepo struct {
	db   *sql.DB
	init sync.Once
}

//NewEventsRepo creates a new events repository.
func NewEventsRepo(db *sql.DB) EventsRepo {
	return &eventsRepo{db: db}
}

// Init prepares the event repository dummy data.
func (r *eventsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy events.
		err = r.seed()
	})

	return err
}

func (r *eventsRepo) List(filter *sports.ListEventRequestFilter) ([]*sports.Event, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getEventQueries()[eventsList]

	query, args = r.applyFilter(query, filter)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanEvents(rows)
}

func (r *eventsRepo) applyFilter(query string, filter *sports.ListEventRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.Sport) > 0 {
		clauses = append(clauses, "sport IN ("+strings.Repeat("?,", len(filter.Sport)-1)+"?)")

		for _, sport := range filter.Sport {
			args = append(args, sport)
		}
	}

	if filter.GetRetrieveVisibleOnly() {
		clauses = append(clauses, "visible = 1")
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	query += " ORDER BY advertised_start_time"

	return query, args
}

func (m *eventsRepo) scanEvents(
	rows *sql.Rows,
) ([]*sports.Event, error) {
	var events []*sports.Event

	for rows.Next() {
		var event sports.Event
		var advertisedStart time.Time

		if err := rows.Scan(&event.Id, &event.Name, &event.Sport, &event.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		event.AdvertisedStartTime = ts
		setEventStatus(&event)

		events = append(events, &event)
	}

	return events, nil
}

// set a event Status, abstracted for easier reading - ideally this would be a method
// on event type, dont feel that it belongs in here hence it is not on the eventRepo type
func setEventStatus(event *sports.Event) {
	eventTime := event.GetAdvertisedStartTime().AsTime()
	now := time.Now()
	if eventTime.Before(now) {
		event.Status = "CLOSED"
	} else {
		event.Status = "OPEN"
	}

}
