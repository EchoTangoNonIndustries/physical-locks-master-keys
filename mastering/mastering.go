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
	Room      string
	PinStacks []PinStack
}

func (lock *Lock) OpensWith(key *Key) bool {
	if lock == nil || len(lock.PinStacks) == 0 || len(key.Cuts) == 0 {
		return false
	}

	for kc, keyCut := range key.Cuts {
	NextCut:
		if kc >= len(lock.PinStacks) {
			return false
		}
		pinStack := lock.PinStacks[kc]
		shearLine := 0
		for _, pin := range pinStack {
			shearLine += pin
			if keyCut == shearLine {
				continue NextCut
			}
		}
		return false
	}
	return true
}

// PinStack is a list of pin sizes, starting from the lowest (key) pin.
type PinStack []int
