<<<<<<< HEAD
package model

import (
	"errors"
)

//自定义登录或注册是的错误
var (
	ErrUserNotExist  = errors.New("user not exist")
	ErrInvalidPasswd = errors.New("Passwd or username not right")
	ErrInvalidParams = errors.New("Invalid params")
	ErrUserExist     = errors.New("user exist")
)
=======
package model

import (
	"errors"
)

//自定义登录或注册是的错误
var (
	ErrUserNotExist  = errors.New("user not exist")
	ErrInvalidPasswd = errors.New("Passwd or username not right")
	ErrInvalidParams = errors.New("Invalid params")
	ErrUserExist     = errors.New("user exist")
)
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
