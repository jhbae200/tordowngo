package model

import "github.com/revel/revel"

type Member struct {
	UserNum   int    `gorm:"primary_key;AUTO_INCREMENT"`
	UserId    string `gorm:"not null;unique"`
	UserPw    []byte
	UserPws   string `gorm:"-"`
	UserCheck bool   `gorm:"default:'0'"`
}

func (member *Member) Validate(v *revel.Validation) {
	v.Check(member.UserId,
		revel.Required{},
	)

	v.Check(member.UserPws,
		revel.Required{},
	)
}

