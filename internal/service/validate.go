package service

import (
	"errors"
	"fmt"
	"regexp"
)

var ErrInvalidParam = errors.New("invalid param value, param must contain only digits and commas")

func validateLimitAndPage(page, limit int32) error {
	if page <= 0 {
		return fmt.Errorf("%s error: %w", "page must be > 0", ErrInvalidArgument)
	}
	if limit < 10 || limit > 100 {
		return fmt.Errorf("%s error: %w", "limit must in range [10;100]", ErrInvalidArgument)
	}

	return nil
}

func checkParam(val string) error {
	exp := regexp.MustCompile("^[!-&!+,0-9]+$")

	if !exp.Match([]byte(val)) {
		return ErrInvalidParam
	}

	return nil
}
