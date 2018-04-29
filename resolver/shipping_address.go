
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type shippingAddressResolver struct {
    //m *model.shippingAddress
}

func (r *shippingAddressResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *shippingAddressResolver) Mobile() string {
    res := "test string"
    return res
}

func (r *shippingAddressResolver) Consignee() string {
    res := "test string"
    return res
}

func (r *shippingAddressResolver) Address() *addressResolver {
    res := addressResolver{}
    return &res
}

func (r *shippingAddressResolver) Zipcode() *string {
    res := "test string"
    return &res
}

func (r *shippingAddressResolver) IsDefault() bool {
    res := false
    return res
}
