package account

import "errors"

var ErrAccountNotFound = errors.New("account not found")
var ErrAccountInactive = errors.New("account inactive")
var ErrInvalidUserID = errors.New("invalid user id")
