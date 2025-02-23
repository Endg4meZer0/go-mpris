package mpris

import (
	"strings"
	"testing"

	"github.com/godbus/dbus/v5"
)

const currentPlayerName = "cmus"

func getCurrentPlayer(t *testing.T) *Player {
	conn, err := dbus.SessionBus()
	if err != nil {
		t.Error(err)
		return nil
	}

	names, err := List(conn)
	if err != nil {
		t.Error(err)
		return nil
	}

	if len(names) == 0 {
		t.Error("No players found")
		return nil
	}

	var name string
	for _, v := range names {
		if strings.Contains(v, currentPlayerName) {
			name = v
			break
		}
	}
	t.Logf("Found player %s", name)

	return New(conn, name)
}

func TestList(t *testing.T) {
	conn, err := dbus.SessionBus()
	if err != nil {
		t.Error(err)
		return
	}

	names, err := List(conn)
	if err != nil {
		t.Error(err)
		return
	}

	for _, v := range names {
		if !strings.HasPrefix(v, BaseInterface) {
			t.Errorf("Expected %v to start with %v", v, BaseInterface)
		}
	}
}

// All the tests below should be run with an existing player
// (just launch your favorite music player that supports MPRIS)
// and they also require you to change the player name in `currentPlayerName` above
// to the one you want to use

func TestGetName(t *testing.T) {
	player := getCurrentPlayer(t)
	playerName := player.GetName()

	if !strings.HasPrefix(playerName, BaseInterface+".") || !strings.Contains(playerName, currentPlayerName) {
		t.Errorf("Expected something like %v, got %v", BaseInterface+"."+currentPlayerName, playerName)
	}
}

func TestGetShortName(t *testing.T) {
	player := getCurrentPlayer(t)
	playerName := player.GetShortName()

	if !strings.Contains(playerName, currentPlayerName) {
		t.Errorf("Expected a string that contains %v, got %v", currentPlayerName, playerName)
	}
}
