package resolver

import (
	"pinjihui.com/pinjihui/model"
	"github.com/graph-gophers/graphql-go"
	"time"
    "context"
    "github.com/op/go-logging"
    "pinjihui.com/pinjihui/loader"
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

func (r *userResolver) Mobile() *string {
	return &r.u.Mobile
}

func (r *userResolver) Email() *string {
	return &r.u.Email
}

//func (r *userResolver) Password() *string {
//	maskedPassword := "********"
//	return &maskedPassword
//}

func (r *userResolver) CreatedAt() (*graphql.Time, error) {
	if r.u.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	return &graphql.Time{Time: t}, err
}

func (r *userResolver) LastLoginTime() (*graphql.Time, error) {
	//if r.u.CreatedAt == "" {
		return nil, nil
	//}

	//t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	//return &graphql.Time{Time: t}, err
}

func (r *userResolver) LastIp() (string) {
	return "127.0.0.1"
}

func (r *userResolver) Addresses() (string) {
	return "127.0.0.1"
}

func (r *userResolver) Roles(ctx context.Context) *[]*roleResolver {
	if r.u.Roles == nil {
        //roles, err := ctx.Value("roleRepository").(*repository.roleRepository).FindByUserId(&r.u.ID)
        roles, err := loader.LoadRoles(ctx, r.u.ID)
        if err != nil {
            ctx.Value("log").(*logging.Logger).Errorf("Error in retrieving roles : %v", err)
            return nil
        }
        r.u.Roles = roles
    }
	l := make([]*roleResolver, len(r.u.Roles))
	for i := range l {
		l[i] = &roleResolver{
			role: r.u.Roles[i],
		}
	}
	return &l
}
