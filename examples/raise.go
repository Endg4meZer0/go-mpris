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
		log.Fatal("No players found.")
	}

	name := names[0]
	log.Println("Found player: ", name)

	player := mpris.New(conn, name)

	identity, err := player.GetIdentity()

	if err != nil {
		panic(err)
	}

	log.Println("Player's identity: ", identity)

	err = player.Raise()
	if err != nil {
		panic(err)
	}
}
