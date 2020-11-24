package main

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func Test_Init(t *testing.T) {
	var regs = map[string]*Register{"key": &Register{Func: func(l validator.FieldLevel) bool {
		return l.Field().String() == "1@163.com"
	}, Msg: "{0}非法"}}
	Init("zh", regs)
	type User struct {
		Name  string `validate:"required" label:"姓名"`
		Age   uint8  `validate:"gte=0,lte=130" label:"年龄"`
		Email string `validate:"required,email,key" label:"邮箱"`
	}
	var u User
	u.Age = 190
	u.Email = "2@163.com"
	if err := Struct(u); err != nil {
		t.Log(GetMsg(err))
	}

}
