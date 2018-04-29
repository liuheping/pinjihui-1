package resolver

import "golang.org/x/net/context"

func (r *Resolver) Products(ctx context.Context, args struct {
	Key *string
}) *productsConnectionResolver {
    return &productsConnectionResolver{}
}

func (r *Resolver) Merchant(ctx context.Context, args struct{Id string}) *merchantResolver {
    return &merchantResolver{}
}
