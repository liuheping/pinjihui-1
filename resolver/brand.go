
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type brandResolver struct {
    //m *model.brand
}

func (r *brandResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *brandResolver) Name() string {
    res := "test string"
    return res
}

func (r *brandResolver) Thumbnaim() *string {
    res := "test string"
    return &res
}

func (r *brandResolver) Description() *string {
    res := "test string"
    return &res
}
