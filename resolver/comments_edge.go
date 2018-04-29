
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type commentsEdgeResolver struct {
    //m *model.commentsEdge
}

func (r *commentsEdgeResolver) Cursor() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *commentsEdgeResolver) Node() *commentResolver {
    res := commentResolver{}
    return &res
}
