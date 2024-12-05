package main

import (
	"fmt"
)

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130"`
	UserName  string `validate:"gte=5,lte=10"`
	Optional  string
}

func main() {
	u1 := &User{
		FirstName: "",
		LastName:  "",
		Age:       135,
		UserName:  "abc",
	}

	if err := UserValidate(u1); err != nil {
		fmt.Printf("User: %+v Error: %s\n", u1, err)
	} else {
		fmt.Printf("User: %+v is valid\n", u1)
	}

	u2 := &User{
		FirstName: "",
		LastName:  "",
		Age:       135,
		UserName:  "mylongusername",
	}

	if err := UserValidate(u2); err != nil {
		fmt.Printf("User: %+v Error: %s\n", u2, err)
	} else {
		fmt.Printf("User: %+v is valid\n", u2)
	}

	u3 := &User{
		FirstName: "First",
		LastName:  "Last",
		Age:       49,
		UserName:  "myusername",
	}

	if err := UserValidate(u3); err != nil {
		fmt.Printf("User: %+v Error: %s\n", u3, err)
	} else {
		fmt.Printf("User: %+v is valid\n", u3)
	}
}
