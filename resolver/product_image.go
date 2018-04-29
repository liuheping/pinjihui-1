
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type productImageResolver struct {
    //m *model.productImage
}

func (r *productImageResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *productImageResolver) SamllImage() string {
    res := "test string"
    return res
}

func (r *productImageResolver) MediumImage() string {
    res := "test string"
    return res
}

func (r *productImageResolver) BigImage() string {
    res := "test string"
    return res
}
