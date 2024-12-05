package main

import (
	"fmt"
)

func UserValidate(obj *User) []error {
	var errs []error

	if obj.FirstName == "" {
		errs = append(errs, fmt.Errorf("%w: FirstName required", ErrValidation))
	}

	if obj.LastName == "" {
		errs = append(errs, fmt.Errorf("%w: LastName required", ErrValidation))
	}

	if obj.Age < 0 {
		errs = append(errs, fmt.Errorf("%w: Age must be >= 0", ErrValidation))
	}

	if obj.Age > 130 {
		errs = append(errs, fmt.Errorf("%w: Age must be <= 130", ErrValidation))
	}

	if len(obj.UserName) < 5 {
		errs = append(errs, fmt.Errorf("%w: length UserName must be >= 5", ErrValidation))
	}

	if len(obj.UserName) > 10 {
		errs = append(errs, fmt.Errorf("%w: length UserName must be <= 10", ErrValidation))
	}

	return errs
}
