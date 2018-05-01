package resolver

import (
    "golang.org/x/net/context"
    "github.com/graph-gophers/graphql-go"
)

type productSearchInput struct {
    Key *string
    Recommended *bool
    Brand *graphql.ID
    Category *graphql.ID
}

func (r *Resolver) Products(ctx context.Context, args struct {
    First *int32
	After *string
    Search *productSearchInput
}) *productsConnectionResolver {
    return &productsConnectionResolver{}
}

func (r *Resolver) Merchant(ctx context.Context, args struct{Id string}) *merchantResolver {
    return &merchantResolver{}
}

func (r *Resolver) Brands(args struct{First *int32}) *[]*brandResolver {
    res := make([]*brandResolver, 10)
    for i := range res {
        v := brandResolver{}
        res[i] = &v
    }
    return &res
}

func (r *Resolver) Categories() *[]*categoryResolver {
    res := make([]*categoryResolver, 10)
    for i := range res {
        v := categoryResolver{}
        res[i] = &v
    }
    return &res
}
