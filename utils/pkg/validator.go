package pkg

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}

func SetupValidator() *validator.Validate {
	v := validator.New()

	v.RegisterValidation("nip", validateNIP)

	return v
}

func BindValidate(c echo.Context, req interface{}) (err error) {
	if err = c.Bind(req); err != nil {
		err = fmt.Errorf("[Utils][Pkg][Validator] failed to bind request, err: %s", err.Error())
		return
	}

	if err = c.Validate(req); err != nil {
		err = fmt.Errorf("[Utils][Pkg][Validator] failed to validate request, err: %s", err.Error())
		return
	}

	return
}

func validateNIP(fl validator.FieldLevel) bool {
	value := fl.Field().Int()
	valueString := strconv.FormatInt(value, 10)
	role := fl.Param()

	// check if the input matches the general pattern
	pattern := `^(615|303)[12]\d{6}\d{3}$`
	match, _ := regexp.MatchString(pattern, valueString)
	if !match {
		return false
	}

	// extract the role, year, mont, and random digits
	roleCode, err := strconv.Atoi(valueString[:3])
	if err != nil {
		return false
	}

	year, err := strconv.Atoi(valueString[4:8])
	if err != nil {
		return false
	}

	month, err := strconv.Atoi(valueString[8:10])
	if err != nil {
		return false
	}

	randomDigits, err := strconv.Atoi(valueString[10:])
	if err != nil {
		return false
	}

	// validate role
	if (role == "it" && roleCode == 303) || (role == "nurse" && roleCode == 615) {
		return false
	}

	// validate year
	currentYear := time.Now().Year()
	if year < 2000 || year > currentYear {
		return false
	}

	// validate month
	if month < 1 || month > 12 {
		return false
	}

	// validate the random digits (although regex already covers this)
	if randomDigits < 0 || randomDigits > 999 {
		return false
	}

	// if all checks pass, return true
	return true
}
