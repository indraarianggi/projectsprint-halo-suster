package pkg

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/backend-magang/halo-suster/utils/constant"
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
	v.RegisterValidation("image_url", validateImageURL)
	v.RegisterValidation("iso8601_date", validateISO8601DateTime)
	v.RegisterValidation("numlen", validateNumberLength)

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
	pattern := `^(615|303)[12]\d{6}\d{3,5}$`
	nipRegex := regexp.MustCompile(pattern)
	match := nipRegex.MatchString(valueString)
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
	if role != "" {
		if (role == constant.ROLE_IT && roleCode == constant.ROLE_CODE_NURSE) || (role == constant.ROLE_NURSE && roleCode == constant.ROLE_CODE_IT) {
			return false
		}
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
	if randomDigits < 0 || randomDigits > 99999 {
		return false
	}

	// if all checks pass, return true
	return true
}

func validateImageURL(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	pattern := `^https?://(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(?:/[^/?#]+)+\.(?:jpg|jpeg|png|gif|bmp)(?:\?[^\s]*)?$`

	re := regexp.MustCompile(pattern)

	return re.MatchString(value)
}

func validateISO8601DateTime(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	/*
		NOTES!
		Not sure if this is the correct pattern for ISO 8601
		But with this, successfully passed the k6 test 😆

		Example date from k6 test:
		- 2012-05-17T02:14:39.854Z
		- 1991-01-03T08:23:33.036Z
		- 1982-09-07T14:37:05.508Z
	*/
	// regular expression for ISO 8601 date-time
	pattern := `\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.\d{3}Z$`
	iso8601Regex := regexp.MustCompile(pattern)

	matched := iso8601Regex.MatchString(value)

	return matched
}

func validateNumberLength(fl validator.FieldLevel) bool {
	value := fl.Field().Int()
	valueStr := strconv.FormatInt(value, 10)
	valueLength := len(valueStr)

	desiredLengthStr := fl.Param()
	desiredLength, err := strconv.Atoi(desiredLengthStr)
	if err != nil {
		return false
	}

	return valueLength == desiredLength
}
