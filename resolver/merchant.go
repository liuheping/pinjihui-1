
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type merchantResolver struct {
    //m *model.merchant
}

func (r *merchantResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *merchantResolver) CompanyName() string {
    res := "test string"
    return res
}

func (r *merchantResolver) CompanyAddress() string {
    res := "test string"
    return res
}

func (r *merchantResolver) DeliveryAddress() string {
    res := "test string"
    return res
}

func (r *merchantResolver) CompanyImage() *string {
    res := "test string"
    return &res
}
