
package resolver

import (
    //"pinjihui.com/pinjihui/model"
)

type ordersConnectionResolver struct {
    //m *model.ordersConnection
}

func (r *ordersConnectionResolver) TotalCount() int32 {
    res := int32(3)
    return res
}

func (r *ordersConnectionResolver) Edges() *[]*ordersEdgeResolver {
    res := make([]*ordersEdgeResolver, 3)
    for i := range res {
        v := ordersEdgeResolver{}
        res[i] = &v
    }
    return &res
}

func (r *ordersConnectionResolver) PageInfo() *pageInfoResolver {
    res := pageInfoResolver{}
    return &res
}
