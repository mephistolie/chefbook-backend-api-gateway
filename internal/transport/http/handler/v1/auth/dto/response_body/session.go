package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/contract"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
)

type Session = contract.Session
type Sessions = contract.SessionsResponse

func BySessions(sessions []*api.Session) Sessions {
	dtos := make([]contract.Session, 0, len(sessions))
	for _, s := range sessions {
		client := contract.SessionClient{Platform: "unknown", Type: "unknown"}
		if s.Client != nil {
			if s.Client.Name != "" {
				name := s.Client.Name
				client.Name = &name
			}
			if s.Client.Platform != "" {
				client.Platform = contract.SessionClientPlatform(s.Client.Platform)
			}
			if s.Client.Type != "" {
				client.Type = contract.SessionClientType(s.Client.Type)
			}
		}
		var location *string
		if s.Location != "" {
			l := s.Location
			location = &l
		}
		dtos = append(dtos, contract.Session{Id: s.Id, Ip: s.Ip, LastRefreshTimestamp: s.AccessTime.AsTime(), Location: location, Client: client})
	}
	return Sessions{Sessions: dtos}
}
