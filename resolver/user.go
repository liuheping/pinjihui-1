package resolver

import (
	"pinjihui.com/pinjihui/model"
	"github.com/graph-gophers/graphql-go"
	"time"
)

type userResolver struct {
	u *model.User
}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

func (r *userResolver) Name() *string {
	return &r.u.Name
}

func (r *userResolver) Mobile() string {
	return r.u.Mobile
}

func (r *userResolver) Email() *string {
	return &r.u.Email
}

//func (r *userResolver) Password() *string {
//	maskedPassword := "********"
//	return &maskedPassword
//}

func (r *userResolver) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	return graphql.Time{Time: t}, err
}

func (r *userResolver) LastLoginTime() (*graphql.Time, error) {
	if r.u.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	return &graphql.Time{Time: t}, err
}

func (r *userResolver) LastIp() *string {
    res := "test string"
    return &res
}

func (r *userResolver) Addresses() *[]*shippingAddressResolver {
    res := make([]*shippingAddressResolver, 3)
    return &res
}

func (r *userResolver) Orders() *ordersConnectionResolver {
    res := ordersConnectionResolver{}
    return &res
}

func (r *userResolver) Cart() *[]*productInCartResolver {
    res := make([]*productInCartResolver, 3)
    return &res
}

func (r *userResolver) Favorites() *[]*favoriteResolver {
    res := make([]*favoriteResolver, 3)
    return &res
}
