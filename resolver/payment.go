
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type paymentResolver struct {
    //m *model.payment
}

func (r *paymentResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *paymentResolver) PayName() string {
    res := "test string"
    return res
}

func (r *paymentResolver) PayCode() string {
    res := "test string"
    return res
}

func (r *paymentResolver) PayDesc() *string {
    res := "test string"
    return &res
}

func (r *paymentResolver) IsOnline() bool {
    res := false
    return res
}
