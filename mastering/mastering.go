package mastering

import "errors"

const (
	MINIMUM_MASTER_PIN_SIZE = 2 // to avoid pins jamming sideways
	LEAST_MASTERED_PINS     = 2
	MOST_MASTERED_PINS      = 5
)

func GetKeysAndLocks(rooms []Room, keysToKeep []Key) (keys []Key, locks []Lock, err error) {
	return nil, nil, errors.New("UNIMPLEMENTED")
}

type Room struct {
	Name     string
	Owner    string
	OpenedBy []string
}

type Key struct {
	Owner string
	Cuts  []int
}

// Lock counts pin stacks from outside to the inside.
type Lock struct {
	Room string
	Pins []PinStack
}

func (lock *Lock) OpensWith(k *Key) bool {
	return false
}

// PinStack is a list of pin sizes, starting from the lowest (key) pin.
type PinStack []int
