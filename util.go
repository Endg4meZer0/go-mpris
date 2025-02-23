package mpris

import (
	"errors"

	"github.com/godbus/dbus/v5"
)

func GetPositionFromSignal(signal *dbus.Signal) (int64, error) {
	if len(signal.Body) == 0 {
		return 0, errors.New("signal's body is empty")
	}

	v, ok := signal.Body[0].(int64)
	if !ok {
		return 0, errors.New("body is not an int64")
	}

	return v, nil
}

func GetPropertiesChangedFromSignal(signal *dbus.Signal) (map[string]dbus.Variant, error) {
	if len(signal.Body) == 0 {
		return nil, errors.New("signal's body is empty")
	}

	v, ok := signal.Body[1].(map[string]dbus.Variant)
	if !ok {
		return nil, errors.New("body is not a dict")
	}
	return v, nil
}
