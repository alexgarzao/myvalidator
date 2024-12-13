# MyValidator (terrible name, fix it!)

[Validator](https://github.com/go-playground/validator) is an amazing project. But in applications with high frequency use, the way that validator works (i.e. reflection) can cause performance penalty.

MyValidator born to solve that gap. Instead of use reflection, MyValidator uses the code generating approach.

This project aims to be compatible with validator tag syntax.

At this time it is an unstable project and should not be used in production environments.

# How to build MyValidator

The following requirements are needed to build the project:
- Git
- Go >= 1.22
- Make

The steps to build are:
```
# Clone the project repository
git clone git@github.com:opencodeco/myvalidator.git

# Run the tests (optional)
make test

# Build the binary
make build
```

After that the executable will be in bin/myvalidator.

# Steps to run the tests

## Steps to run test01

Test01 aims to be a case where all the files are in the same package (in this case, the main package).

```
# Runs myvalidator to generate structs validator and common definitions.
./bin/myvalidator ./tests/test01
```

After that two files will be generated:
- validators.go: contains common definitions
- user_validator.go: contains UserValidate function that is responsible to check if User object has a valid content

## Steps to run test02

Test02 aims to be an example where the structs to be validated are in another package (structsinpkg in this test).

```
# Runs myvalidator to generate structs validator and common definitions.
./bin/myvalidator ./tests/test02
```

After that two files will be generated:
- validators.go: contains common definitions
- user_validator.go: contains UserValidate function that is responsible to check if User object has a valid content


# License

MyValidator uses [MIT License](LICENSE). 
