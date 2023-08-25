package mastering

import "errors"

const (
	MINIMUM_MASTER_PIN_SIZE = 2 // to avoid pins jamming sideways
)

// LockPinSet counts pin stacks from the shoulder of the key to the tip.
type LockPinSet []PinStack

// PinStack is a list of pin sizes, starting from the lowest (key) pin.
type PinStack []int

func foo() (*LockPinSet, error) {
	return nil, errors.New("UNIMPLEMENTED")
}
