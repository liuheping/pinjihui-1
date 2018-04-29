
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type stockResolver struct {
    //m *model.stock
}

func (r *stockResolver) ProductId() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *stockResolver) MerchantId() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *stockResolver) Stock() int32 {
    res := int32(3)
    return res
}
