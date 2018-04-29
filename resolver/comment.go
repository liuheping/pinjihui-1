
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "time"
    "github.com/graph-gophers/graphql-go"
)

type commentResolver struct {
    //m *model.comment
}

func (r *commentResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *commentResolver) User() *userResolver {
    res := userResolver{}
    return &res
}

func (r *commentResolver) Product() *productResolver {
    res := productResolver{}
    return &res
}

func (r *commentResolver) Rank() int32 {
    res := int32(3)
    return res
}

func (r *commentResolver) Order() *orderResolver {
    res := orderResolver{}
    return &res
}

func (r *commentResolver) Content() string {
    res := "test string"
    return res
}

func (r *commentResolver) CreatedAt() (graphql.Time, error) {
    res, err := time.Parse(time.RFC3339, "2018-04-01 12:04:56.539453")
    return graphql.Time{Time: res}, err
}

func (r *commentResolver) Reply() *commentResolver {
    res := commentResolver{}
    return &res
}
