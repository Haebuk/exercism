package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

var ErrNegativeFodder = errors.New("negative fodder")

// TODO: Define the SillyNephewError type here.
type sn struct {
	numCows int
}

func (e *sn) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.numCows)
}

func SillyNephewError(cows int) error {
	return &sn{numCows: cows}
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	fodder, err := weightFodder.FodderAmount()

	if err == ErrScaleMalfunction {
		if fodder > 0 {
			fodder *= 2
		} else if fodder < 0 {
			return 0, ErrNegativeFodder
		}

	} else if err != nil {
		return 0, err
	}

	if fodder < 0 {
		return 0, ErrNegativeFodder
	}

	if cows < 0 {
		return 0, SillyNephewError(cows)
	}

	return fodder / float64(cows), nil
}
