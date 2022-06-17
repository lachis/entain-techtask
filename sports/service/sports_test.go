package service

import (
	"database/sql"
	"testing"

	"git.neds.sh/matty/entain/sports/db"
	"git.neds.sh/matty/entain/sports/proto/sports"
)

// ideally there would be more tests but i am keeping it simple as well as
// needing to complete the front end challenge too
func TestBasicSportsEventRequest(t *testing.T) {
	racingDB, err := sql.Open("sqlite3", "../db/racing.db")
	if err != nil {
		t.Errorf("TestVisibleFilter error connecting to db")
	}

	eventsRepo := db.NewEventsRepo(racingDB)
	if err := eventsRepo.Init(); err != nil {
		t.Fatalf("Failed to initialise db")
	}

	s := sportsService{eventsrepo: eventsRepo}
	req := &sports.ListEventRequestFilter{}
	resp, err := s.eventsrepo.List(req)

	if err != nil {
		t.Errorf("Unexpected error in List")
	}
	if resp == nil || len(resp) <= 0 {
		t.Errorf("Response is nil or empty")
	}

}
