
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type productsEdgeResolver struct {
    //m *model.productsEdge
}

func (r *productsEdgeResolver) Cursor() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *productsEdgeResolver) Node() *productResolver {
    res := productResolver{}
    return &res
}
