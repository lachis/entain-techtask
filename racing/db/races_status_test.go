package db

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"

	"git.neds.sh/matty/entain/racing/proto/racing"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSetRaceStatus_OpenStatus(t *testing.T) {
	ts, err := ptypes.TimestampProto(timestamppb.Now().AsTime().Add(-time.Hour))
	if err != nil {
		t.Fatalf("TestRaceStatusAreSet fatal error converting timestamp")
	}
	race := racing.Race{Id: 1, MeetingId: 1, Name: "Test", Number: 1, Visible: false, AdvertisedStartTime: ts}
	setRaceStatus(&race)

	if race.GetStatus() == "OPEN" {
		t.Errorf("Race should be CLOSED but is OPEN")
	}
}

func TestSetRaceStatus_ClosedStatus(t *testing.T) {
	ts, err := ptypes.TimestampProto(timestamppb.Now().AsTime().Add(time.Hour))
	if err != nil {
		t.Fatalf("TestRaceStatusAreSet fatal error converting timestamp")
	}
	race := racing.Race{Id: 1, MeetingId: 1, Name: "Test", Number: 1, Visible: false, AdvertisedStartTime: ts}
	setRaceStatus(&race)

	if race.GetStatus() == "CLOSED" {
		t.Errorf("Race should be OPEN but is CLOSED")
	}
}
