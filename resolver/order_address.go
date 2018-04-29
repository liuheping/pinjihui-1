
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type orderAddressResolver struct {
    //m *model.orderAddress
}

func (r *orderAddressResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *orderAddressResolver) Mobie() string {
    res := "test string"
    return res
}

func (r *orderAddressResolver) Consignee() string {
    res := "test string"
    return res
}

func (r *orderAddressResolver) Address() *addressResolver {
    res := addressResolver{}
    return &res
}

func (r *orderAddressResolver) Zipcode() *string {
    res := "test string"
    return &res
}
