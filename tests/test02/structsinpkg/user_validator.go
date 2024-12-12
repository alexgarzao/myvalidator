package structsinpkg

import (
	"fmt"
)

func UserValidate(obj *User) []error {
	var errs []error

	if !(obj.FirstName != "") {
		errs = append(errs, fmt.Errorf("%w: FirstName required", ErrValidation))
	}

	if !(obj.LastName != "") {
		errs = append(errs, fmt.Errorf("%w: LastName required", ErrValidation))
	}

	if !(obj.Age != 0) {
		errs = append(errs, fmt.Errorf("%w: Age required", ErrValidation))
	}

	return errs
}
