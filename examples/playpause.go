package main

import (
	"log"

	"github.com/Endg4meZer0/go-mpris"
	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	names, err := mpris.List(conn)
	if err != nil {
		panic(err)
	}

	if len(names) == 0 {
		log.Fatal("No players found")
	}

	name := names[0]

	player := mpris.New(conn, name)

	status, err := player.GetPlaybackStatus()
	if err != nil {
		log.Fatalf("Could not get current playback status: %s", err)
	}

	log.Printf("The player was %s...", status)

	err = player.PlayPause()
	if err != nil {
		log.Fatalf("Could not play/pause player: %s", err)
	}
}
