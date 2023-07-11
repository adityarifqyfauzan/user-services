package validator

import "fmt"

func ValidateRole(name string) error {
	// do something

	if len(name) == 0 {
		return fmt.Errorf("name is required")
	}
	return nil
}
