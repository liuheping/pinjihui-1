
package resolver

import (
    //"pinjihui.com/pinjihui/model"
)

type commentsConnectionResolver struct {
    //m *model.commentsConnection
}

func (r *commentsConnectionResolver) TotalCount() int32 {
    res := int32(3)
    return res
}

func (r *commentsConnectionResolver) Edges() *[]*commentsEdgeResolver {
    res := make([]*commentsEdgeResolver, 3)
    return &res
}

func (r *commentsConnectionResolver) PageInfo() *pageInfoResolver {
    res := pageInfoResolver{}
    return &res
}
