package twirp

import (
	"github.com/pepeunlimited/users/internal/pkg/ent"
	"github.com/pepeunlimited/users/pkg/rpc/user"
)

func rolesToString(roles []*ent.Role) []string {
	str := make([]string, 0)
	for _, role := range roles {
		str = append(str, role.Role)
	}
	return str
}

func ToUser(from *ent.User, roles []*ent.Role) *user.User {
	u := &user.User{
		Id:               int64(from.ID),
		Username:         from.Username,
		Email:            from.Email,
		Roles:            rolesToString(roles),
	}
	if from.ProfilePictureID != nil {
		u.ProfilePictureId = *from.ProfilePictureID
	}
	return u
}