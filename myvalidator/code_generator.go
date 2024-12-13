package main

func generateCode(structs []StructInfo) error {
	// TODO: validate tags ok?

	for _, structInfo := range structs {
		if !structInfo.HasValidateTag {
			continue
		}

		if err := structInfo.GenerateFileValidator(); err != nil {
			return err
		}
	}

	return nil
}
