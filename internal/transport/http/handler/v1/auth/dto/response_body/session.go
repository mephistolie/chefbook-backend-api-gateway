package response_body

import (
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"time"
)

type Sessions []Session

type Session struct {
	Id          int64     `json:"id"`
	Current     bool      `json:"current"`
	Ip          string    `json:"ip"`
	AccessPoint string    `json:"accessPoint"`
	Mobile      bool      `json:"mobile"`
	AccessTime  time.Time `json:"accessTime"`
	Location    string    `json:"location"`
}

func BySessions(sessions []*api.Session, ip string) Sessions {
	dtos := Sessions{}
	for _, session := range sessions {
		dto := Session{
			Id:          session.Id,
			Current:     session.Ip == ip,
			Ip:          session.Ip,
			AccessPoint: session.AccessPoint,
			Mobile:      session.Mobile,
			AccessTime:  session.AccessTime.AsTime(),
			Location:    session.Location,
		}
		dtos = append(dtos, dto)
	}
	return dtos
}
