
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type ordersEdgeResolver struct {
    //m *model.ordersEdge
}

func (r *ordersEdgeResolver) Cursor() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *ordersEdgeResolver) Node() *orderResolver {
    res := orderResolver{}
    return &res
}
