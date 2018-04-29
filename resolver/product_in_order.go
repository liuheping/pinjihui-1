
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type productInOrderResolver struct {
    //m *model.productInOrder
}

func (r *productInOrderResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *productInOrderResolver) Order() *orderResolver {
    res := orderResolver{}
    return &res
}

func (r *productInOrderResolver) Product() *productResolver {
    res := productResolver{}
    return &res
}

func (r *productInOrderResolver) Name() string {
    res := "test string"
    return res
}

func (r *productInOrderResolver) ProductCount() int32 {
    res := int32(3)
    return res
}

func (r *productInOrderResolver) Price() float64 {
    res := 0.0
    return res
}

func (r *productInOrderResolver) Image() *string {
    res := "test string"
    return &res
}
