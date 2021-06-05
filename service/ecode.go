package service

type ErrCode int32

const (
	ErrSuccess      ErrCode = 0
	ErrParamInvalid ErrCode = 10000 + iota
	ErrPasswdNotMatch
	ErrPasswdInvalid
	ErrUserExisted
	ErrUserUnExist
	ErrServerBusy ErrCode = 20000
)

var definesErrors = map[ErrCode]string{
	ErrSuccess:        "success",
	ErrParamInvalid:   "invalid params",
	ErrPasswdNotMatch: "passwd != re_passwd",
	ErrPasswdInvalid:  "passwd error",
	ErrUserExisted:    "user existed",
	ErrUserUnExist:    "user unexist",
	ErrServerBusy:     "server busy",
}

func (e ErrCode) Message() string {
	return definesErrors[e]
}
