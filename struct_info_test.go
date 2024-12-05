package main

import (
	"reflect"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestStructInfoGenerateValidator(t *testing.T) {
	type fields struct {
		StructInfo StructInfo
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Valid struct",
			fields: fields{
				StructInfo: StructInfo{
					Name: "User",
					FieldsInfo: []FieldInfo{
						{
							Name:        "FirstName",
							Type:        "string",
							Tag:         `validate:"required"`,
							Validations: []string{"required"},
						},
						{
							Name:        "MyAge",
							Type:        "uint8",
							Tag:         `validate:"required"`,
							Validations: []string{"required"},
						},
					},
					HasValidateTag: true,
					PackageName:    "main",
				},
			},
			want: `package main

import (
	"fmt"
)

func UserValidate(obj *User) []error {
	var errs []error

	if obj.FirstName == "" {
		errs = append(errs, fmt.Errorf("%w: FirstName required", ErrValidation))
	}

	if obj.MyAge == 0 {
		errs = append(errs, fmt.Errorf("%w: MyAge required", ErrValidation))
	}

	return errs
}
`,
			wantErr: false,
		},
		{
			name: "FirstName must have 5 characters or more",
			fields: fields{
				StructInfo: StructInfo{
					Name: "User",
					FieldsInfo: []FieldInfo{
						{
							Name:        "FirstName",
							Type:        "string",
							Tag:         `validate:"gte=5"`,
							Validations: []string{"gte=5"},
						},
					},
					HasValidateTag: true,
					PackageName:    "main",
				},
			},
			want: `package main

import (
	"fmt"
)

func UserValidate(obj *User) []error {
	var errs []error

	if len(obj.FirstName) < 5 {
		errs = append(errs, fmt.Errorf("%w: length FirstName must be >= 5", ErrValidation))
	}

	return errs
}
`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := tt.fields.StructInfo
			got, err := fv.GenerateValidator()
			if (err != nil) != tt.wantErr {
				t.Errorf("FileValidator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileValidator.Generate() = %v, want %v", got, tt.want)
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(tt.want, got, false)
				if len(diffs) > 1 {
					t.Errorf("FileValidator.Generate() diff = \n%v", dmp.DiffPrettyText(diffs))
				}
			}
		})
	}
}

func TestGetFieldTestElements(t *testing.T) {
	type args struct {
		fieldName       string
		fieldValidation string
		fieldType       string
	}
	tests := []struct {
		name    string
		args    args
		want    FieldTestElements
		wantErr bool
	}{
		{
			name: "Required string",
			args: args{
				fieldName:       "myfield1",
				fieldValidation: "required",
				fieldType:       "string",
			},
			want: FieldTestElements{
				loperand:     "obj.myfield1",
				operator:     "==",
				roperand:     `""`,
				errorMessage: "myfield1 required",
			},
			wantErr: false,
		},
		{
			name: "Required uint8",
			args: args{
				fieldName:       "myfield2",
				fieldValidation: "required",
				fieldType:       "uint8",
			},
			want: FieldTestElements{
				loperand:     "obj.myfield2",
				operator:     "==",
				roperand:     `0`,
				errorMessage: "myfield2 required",
			},
			wantErr: false,
		},
		{
			name: "uint8 >= 0",
			args: args{
				fieldName:       "myfield3",
				fieldValidation: "gte=0",
				fieldType:       "uint8",
			},
			want: FieldTestElements{
				loperand:     "obj.myfield3",
				operator:     "<",
				roperand:     `0`,
				errorMessage: "myfield3 must be >= 0",
			},
			wantErr: false,
		},
		{
			name: "uint8 <= 130",
			args: args{
				fieldName:       "myfield4",
				fieldValidation: "lte=130",
				fieldType:       "uint8",
			},
			want: FieldTestElements{
				loperand:     "obj.myfield4",
				operator:     ">",
				roperand:     `130`,
				errorMessage: "myfield4 must be <= 130",
			},
			wantErr: false,
		},
		{
			name: "String size >= 5",
			args: args{
				fieldName:       "myfield5",
				fieldValidation: "gte=5",
				fieldType:       "string",
			},
			want: FieldTestElements{
				loperand:     "len(obj.myfield5)",
				operator:     "<",
				roperand:     `5`,
				errorMessage: "length myfield5 must be >= 5",
			},
			wantErr: false,
		},
		{
			name: "String size <= 10",
			args: args{
				fieldName:       "myfield6",
				fieldValidation: "lte=10",
				fieldType:       "string",
			},
			want: FieldTestElements{
				loperand:     "len(obj.myfield6)",
				operator:     ">",
				roperand:     `10`,
				errorMessage: "length myfield6 must be <= 10",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFieldTestElements(tt.args.fieldName, tt.args.fieldValidation, tt.args.fieldType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFieldTestElements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFieldTestElements() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
