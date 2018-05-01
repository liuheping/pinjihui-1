
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type categoryResolver struct {
    //m *model.category
}

func (r *categoryResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *categoryResolver) Name() string {
    res := "test string"
    return res
}

func (r *categoryResolver) ParentId() *graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return &res
}

func (r *categoryResolver) Children() *[]*categoryResolver {
    res := make([]*categoryResolver, 3)
    for i := range res {
        v := categoryResolver{}
        res[i] = &v
    }
    return &res
}
