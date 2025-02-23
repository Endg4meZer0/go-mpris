package mpris

import "testing"

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

func TestRaise(t *testing.T) {
	player := getCurrentPlayer(t)

	err := player.Raise()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestQuit(t *testing.T) {
	player := getCurrentPlayer(t)

	err := player.Quit()
	if err != nil {
		t.Error(err)
		return
	}
}

/*
    ____  ____  ____  ____  __________  _____________________
   / __ \/ __ \/ __ \/ __ \/ ____/ __ \/_  __/  _/ ____/ ___/
  / /_/ / /_/ / / / / /_/ / __/ / /_/ / / /  / // __/  \__ \
 / ____/ _, _/ /_/ / ____/ /___/ _, _/ / / _/ // /___ ___/ /
/_/   /_/ |_|\____/_/   /_____/_/ |_| /_/ /___/_____//____/
*/

func TestCanQuit(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.CanQuit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetFullscreen(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetFullscreen()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSetFullscreen(t *testing.T) {
	player := getCurrentPlayer(t)
	err := player.SetFullscreen(false)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCanSetFullscreen(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.CanSetFullscreen()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCanRaise(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.CanRaise()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestHasTrackList(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.HasTrackList()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetIdentity(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetIdentity()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetDesktopEntry(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetDesktopEntry()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetSupportedUriSchemes(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetSupportedUriSchemes()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetSupportedMimeTypes(t *testing.T) {
	player := getCurrentPlayer(t)
	_, err := player.GetSupportedMimeTypes()
	if err != nil {
		t.Error(err)
		return
	}
}
