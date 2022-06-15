package service

import (
	"database/sql"
	"testing"

	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
)

func TestVisibleFilter(t *testing.T) {
	racingDB, err := sql.Open("sqlite3", "../db/racing.db")
	if err != nil {
		t.Errorf("VisibleFilterTest got error")
	}

	racesRepo := db.NewRacesRepo(racingDB)

	s := racingService{racesRepo: racesRepo}
	req := &racing.ListRacesRequestFilter{Visible: true}
	resp, err := s.racesRepo.List(req)

	if err != nil {
		t.Errorf("VisibleFilterTest got error")
	}
	if resp == nil {
		t.Errorf("no resp")
	}
}
