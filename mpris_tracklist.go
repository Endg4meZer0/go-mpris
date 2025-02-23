package mpris

import (
	"errors"

	"github.com/godbus/dbus/v5"
)

/*
    __  _________________  ______  ____  _____
   /  |/  / ____/_  __/ / / / __ \/ __ \/ ___/
  / /|_/ / __/   / / / /_/ / / / / / / /\__ \
 / /  / / /___  / / / __  / /_/ / /_/ /___/ /
/_/  /_/_____/ /_/ /_/ /_/\____/_____//____/
*/

// Returns the specified tracks' metadata.
// See also: https://specifications.freedesktop.org/mpris-spec/latest/Track_List_Interface.html#Method:GetTracksMetadata
func (i *Player) GetTracksMetadata(tracks []string) ([]Metadata, error) {
	variant := dbus.Variant{}
	err := i.obj.Call(getPropertyMethod, 0, TrackListInterface, "GetTracksMetadata", dbus.MakeVariant(tracks)).Store(&variant)

	if err != nil {
		return nil, err
	}
	if variant.Value() == nil {
		return nil, errors.New("variant value is nil")
	}
	value, ok := variant.Value().([]Metadata)
	if !ok {
		return nil, errors.New("variant type is not []Metadata")
	}

	return value, nil
}

// Adds the specified Uri after a specified track and if it should become current track.
// See also: https://specifications.freedesktop.org/mpris-spec/latest/Track_List_Interface.html#Method:AddTrack
func (i *Player) AddTrack(uri string, afterTrack string, setAsCurrent bool) error {
	err := i.obj.Call(getPropertyMethod, 0, TrackListInterface, "AddTrack", dbus.MakeVariant(uri), dbus.MakeVariant(afterTrack), dbus.MakeVariant(setAsCurrent)).Err
	return err
}

// Removes the specified track from the track list.
// See also: https://specifications.freedesktop.org/mpris-spec/latest/Track_List_Interface.html#Method:RemoveTrack
func (i *Player) RemoveTrack(trackId string) error {
	err := i.obj.Call(getPropertyMethod, 0, TrackListInterface, "RemoveTrack", dbus.MakeVariant(trackId)).Err
	return err
}

// Goes to the specified track in the track list.
// See also: https://specifications.freedesktop.org/mpris-spec/latest/Track_List_Interface.html#Method:GoTo
func (i *Player) GoTo(trackId string) error {
	err := i.obj.Call(getPropertyMethod, 0, TrackListInterface, "RemoveTrack", dbus.MakeVariant(trackId)).Err
	return err
}

/*
   _____ ___________   _____    __   _____
  / ___//  _/ ____/ | / /   |  / /  / ___/
  \__ \ / // / __/  |/ / /| | / /   \__ \
 ___/ // // /_/ / /|  / ___ |/ /______/ /
/____/___/\____/_/ |_/_/  |_/_____/____/
*/

// Adds a handler to the TrackListReplaced signal.
// Unfortunately, no ease of conversion function here just yet.
// See also: https://specifications.freedesktop.org/mpris-spec/latest/Track_List_Interface.html#Signal:TrackListReplaced
func (i *Player) OnTrackListReplaced(ch chan<- *dbus.Signal) (err error) {
	err = i.conn.AddMatchSignal(
		dbus.WithMatchInterface(TrackListInterface),
		dbus.WithMatchMember("TrackListReplaced"),
		dbus.WithMatchSender(i.name),
	)
	if err != nil {
		return err
	}

	i.conn.Signal(ch)
	return
}

// Adds a handler to the TrackAdded signal.
// Unfortunately, no ease of conversion function here just yet.
// See also: https://specifications.freedesktop.org/mpris-spec/latest/Track_List_Interface.html#Signal:TrackAdded
func (i *Player) OnTrackAdded(ch chan<- *dbus.Signal) (err error) {
	err = i.conn.AddMatchSignal(
		dbus.WithMatchInterface(TrackListInterface),
		dbus.WithMatchMember("TrackAdded"),
		dbus.WithMatchSender(i.name),
	)
	if err != nil {
		return err
	}

	i.conn.Signal(ch)
	return
}

// Adds a handler to the TrackRemoved signal.
// Unfortunately, no ease of conversion function here just yet.
// See also: https://specifications.freedesktop.org/mpris-spec/latest/Track_List_Interface.html#Signal:TrackRemoved
func (i *Player) OnTrackRemoved(ch chan<- *dbus.Signal) (err error) {
	err = i.conn.AddMatchSignal(
		dbus.WithMatchInterface(TrackListInterface),
		dbus.WithMatchMember("TrackRemoved"),
		dbus.WithMatchSender(i.name),
	)
	if err != nil {
		return err
	}

	i.conn.Signal(ch)
	return
}

// Adds a handler to the TrackMetadataChanged signal.
// Unfortunately, no ease of conversion function here just yet.
// See also: https://specifications.freedesktop.org/mpris-spec/latest/Track_List_Interface.html#Signal:TrackMetadataChanged
func (i *Player) OnTrackMetadataChanged(ch chan<- *dbus.Signal) (err error) {
	err = i.conn.AddMatchSignal(
		dbus.WithMatchInterface(TrackListInterface),
		dbus.WithMatchMember("TrackMetadataChanged"),
		dbus.WithMatchSender(i.name),
	)
	if err != nil {
		return err
	}

	i.conn.Signal(ch)
	return
}

/*
    ____  ____  ____  ____  __________  _____________________
   / __ \/ __ \/ __ \/ __ \/ ____/ __ \/_  __/  _/ ____/ ___/
  / /_/ / /_/ / / / / /_/ / __/ / /_/ / / /  / // __/  \__ \
 / ____/ _, _/ /_/ / ____/ /___/ _, _/ / / _/ // /___ ___/ /
/_/   /_/ |_|\____/_/   /_____/_/ |_| /_/ /___/_____//____/
*/

// Returns the track IDs of the current track list.
// See also: https://specifications.freedesktop.org/mpris-spec/latest/Track_List_Interface.html#Property:Tracks
func (i *Player) GetTracks() ([]string, error) {
	variant, err := getProperty(i.obj, TrackListInterface, "Tracks")
	if err != nil {
		return nil, err
	}
	if variant.Value() == nil {
		return nil, errors.New("variant value is nil")
	}
	value, ok := variant.Value().([]string)
	if !ok {
		return nil, errors.New("variant type is not []string")
	}
	return value, nil
}

// Returns if the track list can be edited by calls.
// See also: https://specifications.freedesktop.org/mpris-spec/latest/Track_List_Interface.html#Property:CanEditTracks
func (i *Player) CanEditTracks() (bool, error) {
	variant, err := getProperty(i.obj, TrackListInterface, "CanEditTracks")
	if err != nil {
		return false, err
	}
	if variant.Value() == nil {
		return false, errors.New("variant value is nil")
	}
	value, ok := variant.Value().(bool)
	if !ok {
		return false, errors.New("variant type is not bool")
	}
	return value, nil
}
