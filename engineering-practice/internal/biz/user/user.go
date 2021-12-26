package user

import (
	"github.com/sjjwantfish/go-learn/engineering-practice/internal/service"
	"time"
)

type UserInfoDo struct {
	service.UserInfoDto
	CreateTime time.Time
}

func (u UserInfoDo) Create()  {
	// TODO transfer to repo
}