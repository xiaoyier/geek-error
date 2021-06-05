package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"geek-error/dao"
	"io/ioutil"
	"net/http"
)

type userSignUpReq struct {
	UserName string `json:"user_name"`
	Passwd   string `json:"passwd"`
	RePasswd string `json:"re_passwd"`
	Email    string `json:"email"`
	PhoneNo  string `json:"phone_number"`
}

type userLoginReq struct {
	UserName string `json:"user_name"`
	Passwd   string `json:"passwd"`
}

func UserSignUp(o http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("service.user.UserSignUp: read body error")
		ResponseError(o, ErrParamInvalid)
		return
	}
	defer r.Body.Close()
	var userReq userSignUpReq
	err = json.Unmarshal(body, &userReq)
	if err != nil {
		fmt.Println("service.user.UserSignUp: unmarshal body error")
		ResponseError(o, ErrParamInvalid)
		return
	}

	if userReq.Passwd != userReq.RePasswd {
		ResponseError(o, ErrPasswdNotMatch)
		return
	}

	user, err := dao.QueryUser(userReq.UserName)
	if user != nil {
		ResponseError(o, ErrUserExisted)
		return
	}

	if err != nil {
		fmt.Println(err)
	}

	u := &dao.User{
		UserName:    userReq.UserName,
		Passwd:      md5Encrypt(userReq.Passwd),
		Email:       userReq.Email,
		PhoneNumber: userReq.PhoneNo,
	}

	err = u.Insert()
	if err != nil {
		fmt.Println(err)
		ResponseError(o, ErrServerBusy)
		return
	}

	ResponseSuccess(o)
}

func UserLogin(o http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("service.user.UserLogin: read body error")
		ResponseError(o, ErrParamInvalid)
		return
	}
	defer r.Body.Close()
	var userReq userLoginReq
	err = json.Unmarshal(body, &userReq)
	if err != nil {
		fmt.Println("service.user.UserLogin: unmarshal body error")
		ResponseError(o, ErrParamInvalid)
		return
	}

	user, err := dao.QueryUser(userReq.UserName)
	if err != nil {
		fmt.Println(err)
		ResponseError(o, ErrServerBusy)
		return
	}

	if user == nil {
		ResponseError(o, ErrUserUnExist)
		return
	}

	if md5Encrypt(user.Passwd) != userReq.Passwd {
		ResponseError(o, ErrPasswdInvalid)
		return
	}

	ResponseSuccess(o)
}

func md5Encrypt(src string) string {

	sha := md5.New()
	sha.Write([]byte(src))
	return hex.EncodeToString(sha.Sum(nil))
}
