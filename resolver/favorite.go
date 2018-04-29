
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type favoriteResolver struct {
    //m *model.favorite
}

func (r *favoriteResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *favoriteResolver) User() *userResolver {
    res := userResolver{}
    return &res
}

func (r *favoriteResolver) Obj() *favoriteObjectResolver {
    res := favoriteObjectResolver{}
    return &res
}
