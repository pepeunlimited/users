package server

import "github.com/pepeunlimited/users/internal/app/app1/ent"

func rolesToString(roles []*ent.Role) []string {
	str := make([]string, 0)
	for _, role := range roles {
		str = append(str, role.Role)
	}
	return str
}