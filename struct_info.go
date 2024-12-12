package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

var structValidatorTpl = `package {{.PackageName}}

import (
	"fmt"
)

func {{.Name}}Validate(obj *{{.Name}}) []error {
	var errs []error
{{range .FieldsInfo}}{{condition .Name .Type .Validations}}{{end}}
	return errs
}
`

var packageDefinitionTpl = `package {{.PackageName}}

import (
	"errors"
)

var ErrValidation = errors.New("validation error")
`

type StructInfo struct {
	Name           string
	Path           string
	PackageName    string
	FieldsInfo     []FieldInfo
	HasValidateTag bool
}

// TODO: NewFieldInfo to validate params and build the object.
type FieldInfo struct {
	Name        string
	Type        string
	Tag         string
	Validations []string
}

type FieldTestElements struct {
	loperand     string
	operator     string
	roperand     string
	errorMessage string
}

func (fv *StructInfo) GenerateValidator() (string, error) {
	funcMap := template.FuncMap{
		"condition": condition,
	}

	tmpl, err := template.New("FileValidator").Funcs(funcMap).Parse(structValidatorTpl)
	if err != nil {
		return "", err
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, fv); err != nil {
		return "", err
	}

	return code.String(), nil
}

func condition(fieldName, fieldType string, fieldValidations []string) (string, error) {

	tests := ""
	for _, fieldValidation := range fieldValidations {
		testElements, err := GetFieldTestElements(fieldName, fieldValidation, fieldType)
		if err != nil {
			return "", fmt.Errorf("field %s: %w", fieldName, err)
		}

		tests += fmt.Sprintf(
			`
	if !(%s %s %s) {
		errs = append(errs, fmt.Errorf("%%w: %s", ErrValidation))
	}
`, testElements.loperand, testElements.operator, testElements.roperand, testElements.errorMessage)
	}

	return tests, nil
}

func GetFieldTestElements(fieldName, fieldValidation, fieldType string) (FieldTestElements, error) {
	ifCode := map[string]FieldTestElements{
		"required,string": {"{{.Name}}", "!=", `""`, "{{.Name}} required"},
		"required,uint8":  {"{{.Name}}", "!=", `0`, "{{.Name}} required"},
		"gte,uint8":       {"{{.Name}}", ">=", `{{.Target}}`, "{{.Name}} must be >= {{.Target}}"},
		"lte,uint8":       {"{{.Name}}", "<=", `{{.Target}}`, "{{.Name}} must be <= {{.Target}}"},
		"gte,string":      {"len({{.Name}})", ">=", `{{.Target}}`, "length {{.Name}} must be >= {{.Target}}"},
		"lte,string":      {"len({{.Name}})", "<=", `{{.Target}}`, "length {{.Name}} must be <= {{.Target}}"},
	}

	splitField := strings.Split(fieldValidation, "=")
	validation := splitField[0]
	target := ""
	if len(splitField) > 1 {
		target = splitField[1]
	}

	ifData, ok := ifCode[validation+","+fieldType]
	if !ok {
		return FieldTestElements{}, fmt.Errorf("unsupported validation %s type %s", fieldValidation, fieldType)
	}

	ifData.loperand = strings.Replace(ifData.loperand, "{{.Name}}", "obj."+fieldName, -1)
	ifData.loperand = strings.Replace(ifData.loperand, "{{.Target}}", target, -1)
	ifData.roperand = strings.Replace(ifData.roperand, "{{.Name}}", "obj."+fieldName, -1)
	ifData.roperand = strings.Replace(ifData.roperand, "{{.Target}}", target, -1)
	ifData.errorMessage = strings.Replace(ifData.errorMessage, "{{.Name}}", fieldName, -1)
	ifData.errorMessage = strings.Replace(ifData.errorMessage, "{{.Target}}", target, -1)

	return ifData, nil
}

func (s *StructInfo) GenerateFileValidator() error {
	fmt.Printf("Generating struct %s validations code\n", s.Name)

	code, err := s.GenerateValidator()
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.Path+"/"+strings.ToLower(s.Name)+"_validator.go", []byte(code), 0644); err != nil {
		return err
	}

	return nil
}

func (s *StructInfo) GenerateFilePackageDefinition() error {
	fmt.Println("Generating package definitions code")

	code, err := s.Generate()
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.Path+"/validators.go", []byte(code), 0644); err != nil {
		return err
	}

	return nil
}

func (s *StructInfo) Generate() (string, error) {
	tmpl, err := template.New("PkgDef").Parse(packageDefinitionTpl)
	if err != nil {
		return "", err
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, s); err != nil {
		return "", err
	}

	return code.String(), nil
}

func (s *StructInfo) PrintInfo() {
	fmt.Println("Struct:", s.Name)
	fmt.Println("\tHasValidateTag:", s.HasValidateTag)

	for _, f := range s.FieldsInfo {
		fmt.Println("\tField:", f.Name, f.Type, f.Tag)
	}
}
