package service

import (
	"database/sql"
	"testing"

	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
)

func TestOrderRacesByAdvertisedStartTime(t *testing.T) {
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

	for i, race := range resp {

		if i+1 != len(resp) {
			next := resp[i+1]

			if next.AdvertisedStartTime.AsTime().Before(race.AdvertisedStartTime.AsTime()) {
				t.Fatalf("Races are not ordered by advertised start time")
			}
		}
	}
}
