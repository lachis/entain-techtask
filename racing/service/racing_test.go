package service

import (
	"database/sql"
	"testing"

	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
)

func TestVisibleFilterTrueOnlyReturnsVisibleRaces(t *testing.T) {
	racingDB, err := sql.Open("sqlite3", "../db/racing.db")
	if err != nil {
		t.Errorf("TestVisibleFilter error connecting to db")
	}

	racesRepo := db.NewRacesRepo(racingDB)

	s := racingService{racesRepo: racesRepo}
	req := &racing.ListRacesRequestFilter{RetrieveVisibleOnly: true}
	resp, err := s.racesRepo.List(req)

	if err != nil {
		t.Errorf("Unexpected error in List")
	}
	if resp == nil || len(resp) <= 0 {
		t.Errorf("Response is nil or empty")
	}

	for _, race := range resp {
		if !race.GetVisible() {
			t.Fatalf(`Race %q has visible false, %v`, race.Id, race.Visible)
		}
	}
}

func TestVisibleFilterFalseOnlyReturnsNonVisibleRaces(t *testing.T) {
	racingDB, err := sql.Open("sqlite3", "../db/racing.db")
	if err != nil {
		t.Errorf("TestVisibleFilter error connecting to db")
	}

	racesRepo := db.NewRacesRepo(racingDB)

	s := racingService{racesRepo: racesRepo}
	req := &racing.ListRacesRequestFilter{RetrieveVisibleOnly: false}
	resp, err := s.racesRepo.List(req)

	if err != nil {
		t.Errorf("Unexpected error in List")
	}
	if resp == nil || len(resp) <= 0 {
		t.Errorf("Response is nil or empty")
	}

	containsVisibleFalse := false
	containsVisibleTrue := false
	for _, race := range resp {
		// array contains visible item and flag has not been set
		if race.GetVisible() && !containsVisibleTrue {
			containsVisibleTrue = true
			// array contins visible false item and flag has not been set
		} else if !race.GetVisible() && !containsVisibleFalse {
			containsVisibleFalse = true
		}
	}

	if !containsVisibleFalse || !containsVisibleTrue {
		t.Fatalf("Array should contain both visible and non-visible races")
	}
}

func TestVisibleFilterNotPresentOnRequest(t *testing.T) {
	racingDB, err := sql.Open("sqlite3", "../db/racing.db")
	if err != nil {
		t.Errorf("TestVisibleFilter error connecting to db")
	}

	racesRepo := db.NewRacesRepo(racingDB)

	s := racingService{racesRepo: racesRepo}
	req := &racing.ListRacesRequestFilter{}
	resp, err := s.racesRepo.List(req)

	if err != nil {
		t.Errorf("Unexpected error in List")
	}
	if resp == nil || len(resp) <= 0 {
		t.Errorf("Response is nil or empty")
	}

	containsVisibleFalse := false
	containsVisibleTrue := false
	for _, race := range resp {
		// array contains visible item and flag has not been set
		if race.GetVisible() && !containsVisibleTrue {
			containsVisibleTrue = true
			// array contins visible false item and flag has not been set
		} else if !race.GetVisible() && !containsVisibleFalse {
			containsVisibleFalse = true
		}
	}

	if !containsVisibleFalse || !containsVisibleTrue {
		t.Fatalf("Array should contain both visible and non-visible races")
	}
}
