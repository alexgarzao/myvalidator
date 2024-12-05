package main

import (
	"fmt"

	"github.com/alexgarzao/myvalidator/tests/test02/structsinpkg"
)

func main() {
	u1 := &structsinpkg.User{}
	if err := structsinpkg.UserValidate(u1); err != nil {
		fmt.Printf("User: %+v Error: %s\n", u1, err)
	} else {
		fmt.Printf("User: %+v is valid\n", u1)
	}

	u2 := &structsinpkg.User{
		FirstName: "First",
		LastName:  "Last",
		Age:       1,
	}

	if err := structsinpkg.UserValidate(u2); err != nil {
		fmt.Printf("User: %+v Error: %s\n", u2, err)
	} else {
		fmt.Printf("User: %+v is valid\n", u2)
	}
}
