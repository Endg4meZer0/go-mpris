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

	volume, err := player.GetVolume()
	if err != nil {
		log.Fatal("Could not get current volume")
	}

	log.Printf("The player's volume is %f...", volume)
	err = player.SetVolume(volume - 0.1)
	if err != nil {
		panic(err)
	}
}
