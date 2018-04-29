
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type productInCartResolver struct {
    //m *model.productInCart
}

func (r *productInCartResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *productInCartResolver) Product() *productResolver {
    res := productResolver{}
    return &res
}

func (r *productInCartResolver) User() *userResolver {
    res := userResolver{}
    return &res
}

func (r *productInCartResolver) ProductCount() int32 {
    res := int32(3)
    return res
}

func (r *productInCartResolver) Merchant() *merchantResolver {
    res := merchantResolver{}
    return &res
}
