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
