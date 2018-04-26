package resolver

import (
	"github.com/op/go-logging"
	"golang.org/x/net/context"
	"pinjihui.com/pinjihui/model"
)

func (r *Resolver) Me(ctx context.Context) (*userResolver, error) {
	//Without using dataloader:
	//user, err := ctx.Value("userRepository").(*repository.userRepository).FindByEmail(args.Email)
	userId := ctx.Value("user_id").(*string)
	user := &model.User{}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved user by user_id[%s] : %v", *userId, *user)
	return &userResolver{user}, nil
}
