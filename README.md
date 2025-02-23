# GO-MPRIS

A Go library for MPRIS.

## Install
```
$ go get github.com/Endg4meZer0/go-mpris
```
The dependency `github.com/godbus/dbus/v5` will be installed as well.

## Example
Printing the current playback status and then changing it:
```go
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
		log.Fatal("No player found")
	}

	name := names[0]
	player := mpris.New(conn, name)

	status, err := player.GetPlaybackStatus()
	if err != nil {
		log.Fatal("Could not get current playback status")
	}

	log.Printf("The player was %s...", status)
	err = player.PlayPause()
	if err != nil {
		log.Fatal("Could not play/pause player")
	}
}
```

**For more examples, see the [examples folder](./examples).**

## Go Docs
Read the docs at https://pkg.go.dev/github.com/Endg4meZer0/go-mpris.

## Credits
[Pauloo27](https://github.com/pauloo27/go-mpris) for the original code.

[emersion](https://github.com/emersion/go-mpris) for the original-original code.

[leberKleber](https://github.com/leberKleber/go-mpris) for several additional ideas regarding metadata.
