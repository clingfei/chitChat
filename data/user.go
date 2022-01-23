package data

import "time"

type Session struct {
	Id 			int						//
	Uuid 		string					//为用户随机生成的唯一id， 通过uuid来定位用户
	Email		string					//邮件地址
	UserId		string
	CreatedAt 	time.Time
}
