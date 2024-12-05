package structsinpkg

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"required"`
}

type NoValidateInfo struct {
	Name    string
	Address string
}
