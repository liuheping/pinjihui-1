
package resolver

import (
    //"pinjihui.com/pinjihui/model"
)

type productsConnectionResolver struct {
    //m *model.productsConnection
}

func (r *productsConnectionResolver) TotalCount() int32 {
    res := int32(3)
    return res
}

func (r *productsConnectionResolver) Edges() *[]*productsEdgeResolver {
    res := make([]*productsEdgeResolver, 3)
    for i := range res {
        v := productsEdgeResolver{}
        res[i] = &v
    }
    return &res
}

func (r *productsConnectionResolver) PageInfo() *pageInfoResolver {
    res := pageInfoResolver{}
    return &res
}
