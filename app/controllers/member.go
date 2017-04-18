package controllers

import (
	"strconv"

	"tordowngo/app/model"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type MemberController struct {
	GormController
}

func (c MemberController) Login() revel.Result {
	return c.Render()
}

func (c MemberController) Sign() revel.Result {
	return c.Render()
}

func (c MemberController) Create(member model.Member, userPwConfirm string) revel.Result {
	c.Validation.Required(userPwConfirm)
	c.Validation.Required(userPwConfirm == member.UserPws).
		Message("페스워드와 페스워드 확인이 일치하지 않습니다.")
	member.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(MemberController.Sign)
	}
	member.UserPw, _ = bcrypt.GenerateFromPassword([]byte(member.UserPws), bcrypt.DefaultCost)
	if err := c.Txn.Create(&member).Error; err != nil {
		revel.ERROR.Println(err)
		c.Flash.Error("데이터베이스 작성 중 문제가 발생하였습니다.")
		return c.Redirect(MemberController.Sign)
	}
	c.Flash.Success("성공적으로 가입되었습니다.")
	return c.Redirect(MemberController.Login)
}

func (c MemberController) Get(member model.Member) revel.Result {
	member.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(MemberController.Login)
	}
	result := model.Member{}
	if err := c.Txn.Where(member).First(&result).Error; err != nil && err != gorm.ErrRecordNotFound {
		revel.ERROR.Println(err)
		c.Flash.Error("데이터베이스 작성 중 문제가 발생하였습니다.")
		return c.Redirect(MemberController.Login)
	}
	if result.UserCheck != true {
		c.Flash.Error("가입승인이 되지 않았습니다.")
		return c.Redirect(MemberController.Login)
	}
	if result.UserNum == 0 {
		c.Flash.Error("아이디나 비밀번호가 틀립니다.")
		return c.Redirect(MemberController.Login)
	}
	if err := bcrypt.CompareHashAndPassword(result.UserPw, []byte(member.UserPws)); err != nil {
		c.Flash.Error("아이디나 비밀번호가 틀립니다.")
		return c.Redirect(MemberController.Login)
	}
	c.Session["userNum"] = strconv.FormatInt(int64(result.UserNum), 32)
	return c.Redirect(TorrentController.List)
}
