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

func TestNext(t *testing.T) {
	player := getCurrentPlayer(t)
	if err := player.Next(); err != nil {
		t.Error(err)
	}
}

func TestPrevious(t *testing.T) {
	player := getCurrentPlayer(t)
	if err := player.Previous(); err != nil {
		t.Error(err)
	}
}

func TestPause(t *testing.T) {
	player := getCurrentPlayer(t)
	if err := player.Pause(); err != nil {
		t.Error(err)
	}
}

func TestPlayPause(t *testing.T) {
	player := getCurrentPlayer(t)
	if err := player.PlayPause(); err != nil {
		t.Error(err)
	}
}

func TestStop(t *testing.T) {
	player := getCurrentPlayer(t)
	if err := player.Stop(); err != nil {
		t.Error(err)
	}
}

func TestPlay(t *testing.T) {
	player := getCurrentPlayer(t)
	if err := player.Play(); err != nil {
		t.Error(err)
	}
}

func TestSeekBy(t *testing.T) {
	player := getCurrentPlayer(t)
	if err := player.SeekBy(5 * 1000 * 1000); err != nil {
		t.Error(err)
	}
}

// TestSetTrackPosition is essentially TestSetPosition that is located below.

func TestOpenUri(t *testing.T) {
	player := getCurrentPlayer(t)
	if err := player.OpenUri(`file://test.mp3`); err != nil {
		t.Error(err)
	}
}

/*
   _____ ___________   _____    __   _____
  / ___//  _/ ____/ | / /   |  / /  / ___/
  \__ \ / // / __/  |/ / /| | / /   \__ \
 ___/ // // /_/ / /|  / ___ |/ /______/ /
/____/___/\____/_/ |_/_/  |_/_____/____/
*/

// This test runs for 10s to collect Seeked signals.
func TestOnSeeked(t *testing.T) {
	player := getCurrentPlayer(t)
	ch := make(chan *dbus.Signal)
	err := player.RegisterSignalReceiver(ch)
	if err != nil {
		t.Error(err)
		return
	}

	timer := time.NewTimer(10 * time.Second)
	sigCol := make([]int64, 0, 8)

	for {
		select {
		case v := <-ch:
			if GetSignalType(v) != SignalSeeked {
				continue
			}

			if len(v.Body) == 0 {
				t.Error("signal's body is empty")
			}

			val, ok := v.Body[0].(int64)
			if !ok {
				t.Error("signal's body is not int64")
				return
			}

			sigCol = append(sigCol, val)
		case <-timer.C:
			t.Log(sigCol)
			return
		}
	}
}

// This test runs for 10s to collect PropertiesChanged signals.
func TestOnPropertiesChanged(t *testing.T) {
	player := getCurrentPlayer(t)
	ch := make(chan *dbus.Signal)
	err := player.RegisterSignalReceiver(ch)
	if err != nil {
		t.Error(err)
		return
	}

	timer := time.NewTimer(10 * time.Second)
	sigCol := make([]map[string]dbus.Variant, 0, 8)

	for {
		select {
		case v := <-ch:
			if GetSignalType(v) != SignalPropertiesChanged {
				continue
			}

			if len(v.Body) != 3 {
				t.Error("not a PropertiesChanged signal")
			}

			val, ok := v.Body[1].(map[string]dbus.Variant)
			if !ok {
				t.Error("signal's changed properties is not a dict")
			}

			sigCol = append(sigCol, val)
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

func TestGetPlaybackStatus(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetPlaybackStatus()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetLoopStatus(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetLoopStatus()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSetLoopStatus(t *testing.T) {
	player := getCurrentPlayer(t)
	if err := player.SetLoopStatus(LoopTrack); err != nil {
		t.Error(err)
		return
	}
}

func TestGetRate(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetRate()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetShuffle(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetShuffle()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSetShuffle(t *testing.T) {
	player := getCurrentPlayer(t)
	if err := player.SetShuffle(false); err != nil {
		t.Error(err)
		return
	}
}

func TestGetMetadata(t *testing.T) {
	player := getCurrentPlayer(t)
	md, err := player.GetMetadata()
	if err != nil {
		t.Error(err)
		return
	}

	trackID, err := md.TrackID()
	if err != nil {
		t.Log(trackID)
		t.Error(err)
		return
	}
}

func TestGetVolume(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetVolume()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSetVolume(t *testing.T) {
	player := getCurrentPlayer(t)
	if err := player.SetVolume(0.33); err != nil {
		t.Error(err)
		return
	}
}

func TestGetPosition(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetPosition()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSetPosition(t *testing.T) {
	player := getCurrentPlayer(t)
	metadata, err := player.GetMetadata()
	if err != nil {
		t.Error(err)
		return
	}
	if metadata == nil {
		t.Error("metadata is nil")
		return
	}

	trackId, err := metadata.TrackID()
	if err != nil {
		t.Error(err)
	}

	if err := player.SetTrackPosition(trackId, 11*1000*1000); err != nil {
		t.Error(err)
	}
}

func TestGetMinimumRate(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetMinimumRate()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetMaximumRate(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetMaximumRate()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCanGoNext(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.CanGoNext()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCanGoPrevious(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.CanGoPrevious()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCanPlay(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.CanPlay()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCanPause(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.CanPause()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCanSeek(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.CanSeek()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCanControl(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.CanControl()
	if err != nil {
		t.Error(err)
		return
	}
}
