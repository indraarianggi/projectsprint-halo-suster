package lib

import "errors"

var (
	ErrorNotFound            = errors.New("data not found")
	ErrorProductNotAvailable = errors.New("product is not available")
	ErrorNoRowsAffected      = errors.New("no rows affected")
	ErrorNoRowsResult        = errors.New("no rows in result set")
	ErrConstraintKey         = errors.New(`duplicate key value violates unique constraint`)
	ErrInsufficientStock     = errors.New("insufficient stock product")
	ErrInsufficientPayment   = errors.New("insufficient payment")
	ErrWrongChange           = errors.New("wrong change")
)
