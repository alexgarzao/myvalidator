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

	if !(obj.Age >= 0) {
		errs = append(errs, fmt.Errorf("%w: Age must be >= 0", types.ErrValidation))
	}

	if !(obj.Age <= 130) {
		errs = append(errs, fmt.Errorf("%w: Age must be <= 130", types.ErrValidation))
	}

	if !(len(obj.UserName) >= 5) {
		errs = append(errs, fmt.Errorf("%w: length UserName must be >= 5", types.ErrValidation))
	}

	if !(len(obj.UserName) <= 10) {
		errs = append(errs, fmt.Errorf("%w: length UserName must be <= 10", types.ErrValidation))
	}

	return errs
}
