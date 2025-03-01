package mpris

import (
	"testing"
	"time"

	"github.com/godbus/dbus/v5"
)

// All the tests below should be run with an existing player
// (just launch your favorite music player that supports MPRIS)
// and they also require you to change the player name in general_test.go
// to the one you want to use

/*
    __  _________________  ______  ____  _____
   /  |/  / ____/_  __/ / / / __ \/ __ \/ ___/
  / /|_/ / __/   / / / /_/ / / / / / / /\__ \
 / /  / / /___  / / / __  / /_/ / /_/ /___/ /
/_/  /_/_____/ /_/ /_/ /_/\____/_____//____/
*/

func TestGetTracksMetadata(t *testing.T) {
	player := getCurrentPlayer(t)
	ids, err := player.GetTracks()
	if err != nil {
		t.Error(err)
		return
	}
	mds, err := player.GetTracksMetadata(ids)
	if err != nil {
		t.Error(err)
		return
	}
	_ = mds
}

func TestAddTrack(t *testing.T) {
	player := getCurrentPlayer(t)
	ids, err := player.GetTracks()
	if err != nil {
		t.Error(err)
		return
	}
	if len(ids) == 0 {
		t.Error("can't really continue testing with 0 ids")
		return
	}

	if err := player.AddTrack(`file://test.mp3`, ids[0], true); err != nil {
		t.Error(err)
		return
	}
}

func TestRemoveTrack(t *testing.T) {
	player := getCurrentPlayer(t)
	ids, err := player.GetTracks()
	if err != nil {
		t.Error(err)
		return
	}
	if len(ids) == 0 {
		t.Error("can't really continue testing with 0 ids")
		return
	}

	if err := player.RemoveTrack(ids[0]); err != nil {
		t.Error(err)
		return
	}
}

func TestGoTo(t *testing.T) {
	player := getCurrentPlayer(t)
	ids, err := player.GetTracks()
	if err != nil {
		t.Error(err)
		return
	}
	if len(ids) < 2 {
		t.Error("can't really continue testing with less than 2 ids")
		return
	}

	if err := player.GoTo(ids[1]); err != nil {
		t.Error(err)
		return
	}
}

/*
   _____ ___________   _____    __   _____
  / ___//  _/ ____/ | / /   |  / /  / ___/
  \__ \ / // / __/  |/ / /| | / /   \__ \
 ___/ // // /_/ / /|  / ___ |/ /______/ /
/____/___/\____/_/ |_/_/  |_/_____/____/
*/

// This test runs for 10s to collect TrackListReplaced signals.
func TestTrackListReplaced(t *testing.T) {
	player := getCurrentPlayer(t)
	ch := make(chan *dbus.Signal)
	err := player.RegisterSignalReceiver(ch)
	if err != nil {
		t.Error(err)
		return
	}

	timer := time.NewTimer(10 * time.Second)
	sigCol := make([]*dbus.Signal, 0, 8)

	for {
		select {
		case v := <-ch:
			if GetSignalType(v) != SignalTrackListReplaced {
				continue
			}

			sigCol = append(sigCol, v)
		case <-timer.C:
			t.Log(sigCol)
			return
		}
	}
}

// This test runs for 10s to collect TrackAdded signals.
func TestTrackAdded(t *testing.T) {
	player := getCurrentPlayer(t)
	ch := make(chan *dbus.Signal)
	err := player.RegisterSignalReceiver(ch)
	if err != nil {
		t.Error(err)
		return
	}

	timer := time.NewTimer(10 * time.Second)
	sigCol := make([]*dbus.Signal, 0, 8)

	for {
		select {
		case v := <-ch:
			if GetSignalType(v) != SignalTrackAdded {
				continue
			}

			sigCol = append(sigCol, v)
		case <-timer.C:
			t.Log(sigCol)
			return
		}
	}
}

// This test runs for 10s to collect TrackRemoved signals.
func TestTrackRemoved(t *testing.T) {
	player := getCurrentPlayer(t)
	ch := make(chan *dbus.Signal)
	err := player.RegisterSignalReceiver(ch)
	if err != nil {
		t.Error(err)
		return
	}

	timer := time.NewTimer(10 * time.Second)
	sigCol := make([]*dbus.Signal, 0, 8)

	for {
		select {
		case v := <-ch:
			if GetSignalType(v) != SignalTrackRemoved {
				continue
			}

			sigCol = append(sigCol, v)
		case <-timer.C:
			t.Log(sigCol)
			return
		}
	}
}

// This test runs for 10s to collect TrackMetadataChanged signals.
func TestTrackMetadataChanged(t *testing.T) {
	player := getCurrentPlayer(t)
	ch := make(chan *dbus.Signal)
	err := player.RegisterSignalReceiver(ch)
	if err != nil {
		t.Error(err)
		return
	}

	timer := time.NewTimer(10 * time.Second)
	sigCol := make([]*dbus.Signal, 0, 8)

	for {
		select {
		case v := <-ch:
			if GetSignalType(v) != SignalTrackMetadataChanged {
				continue
			}

			sigCol = append(sigCol, v)
		case <-timer.C:
			t.Log(sigCol)
			return
		}
	}
}

/*
    ____  ____  ____  ____  __________  _____________________
   / __ \/ __ \/ __ \/ __ \/ ____/ __ \/_  __/  _/ ____/ ___/
  / /_/ / /_/ / / / / /_/ / __/ / /_/ / / /  / // __/  \__ \
 / ____/ _, _/ /_/ / ____/ /___/ _, _/ / / _/ // /___ ___/ /
/_/   /_/ |_|\____/_/   /_____/_/ |_| /_/ /___/_____//____/
*/

func TestGetTracks(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetTracks()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCanEditTracks(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.CanEditTracks()
	if err != nil {
		t.Error(err)
		return
	}
}
