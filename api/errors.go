package api

import (
	"errors"
)

var ERRSuccessFieldIsFalse = errors.New("Response header sucess field is false")
var ERRRequestStatusCodeNotOk = errors.New("Response status code != 200")
