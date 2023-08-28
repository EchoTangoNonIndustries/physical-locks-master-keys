package mastering

const (
	MINIMUM_MASTER_PIN_SIZE = 2 // to avoid pins jamming sideways
	LEAST_MASTERED_PINS     = 2
	MOST_MASTERED_PINS      = 5
)

func GetKeysAndLocks(rooms []Room, alreadyCutKeys []Key) ([]Key, []Lock, error) {
	ownedKeys := map[string]*Key{}
	keysToTheSameRoom := map[string][]*Key{}
	for _, room := range rooms {
		for _, person := range room.OpenedBy {
			ownedKeys[person] = &Key{Owner: person}
		}
	}
	for _, key := range alreadyCutKeys {
		ownedKeys[key.Owner] = &key
	}
	for _, room := range rooms {
		for _, person := range room.OpenedBy {
			key, exists := ownedKeys[person]
			if !exists {
				key = &Key{Owner: person}
			}
			keysToTheSameRoom[room.Name] = append(keysToTheSameRoom[room.Name], key)
		}
	}

	for room, roomKeys := range keysToTheSameRoom {
		newKeys, err := getNewKeysForTheSameRoom(room, roomKeys)
		if err != nil {
			return SOMETHING()
		}
		for _, key := range newKeys {
			ownedKeys[key.Owner] = key
		}
	}
}

func getNewKeysForTheSameRoom(room string, keys []*Key) ([]*Key, error) {
	var keysToKeep []*Key
	var keysToCut []*Key
	for _, key := range keys {
		if key == nil {
			keysToCut = append(keysToCut, key)
		} else {
			keysToKeep = append(keysToKeep, key)
		}
	}
}

type Room struct {
	Name     string
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
		for _, pin := range pinStack.KeyPins {
			shearLine += pin
			if keyCut == shearLine {
				continue NextCut
			}
		}
		return false
	}
	return true
}

type PinStack struct {
	KeyPins []int // counts from the lowest to the highest master pin
	TopPin  int
}
