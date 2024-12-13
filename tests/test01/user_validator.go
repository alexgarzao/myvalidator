package main

import (
	"fmt"

	"github.com/opencodeco/myvalidator/types"
)

func UserValidate(obj *User) []error {
	var errs []error

	if !(obj.FirstName != "") {
		errs = append(errs, fmt.Errorf("%w: FirstName required", types.ErrValidation))
	}

	if !(obj.LastName != "") {
		errs = append(errs, fmt.Errorf("%w: LastName required", types.ErrValidation))
	}

	if !(obj.Age != 0) {
		errs = append(errs, fmt.Errorf("%w: Age required", types.ErrValidation))
	}

	return errs
}
