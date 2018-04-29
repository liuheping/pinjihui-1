
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "time"
    "github.com/graph-gophers/graphql-go"
)

type productResolver struct {
    //m *model.product
}

func (r *productResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *productResolver) Name() string {
    res := "test string"
    return res
}

func (r *productResolver) Price() float64 {
    res := 0.0
    return res
}

func (r *productResolver) Category() *categoryResolver {
    res := categoryResolver{}
    return &res
}

func (r *productResolver) RelatedProducts() *[]*productResolver {
    res := make([]*productResolver, 3)
    return &res
}

func (r *productResolver) Content() *string {
    res := "test string"
    return &res
}

func (r *productResolver) Brand() *brandResolver {
    res := brandResolver{}
    return &res
}

func (r *productResolver) CreatedAt() (graphql.Time, error) {
    res, err := time.Parse(time.RFC3339, "2018-04-01 12:04:56.539453")
    return graphql.Time{Time: res}, err
}

func (r *productResolver) UpdatedAt() (graphql.Time, error) {
    res, err := time.Parse(time.RFC3339, "2018-04-01 12:04:56.539453")
    return graphql.Time{Time: res}, err
}

func (r *productResolver) Tags() *[]string {
    res := make([]string, 3)
    return &res
}

func (r *productResolver) ProductImages() *[]*productImageResolver {
    res := make([]*productImageResolver, 3)
    return &res
}

func (r *productResolver) Merchants() *[]*merchantResolver {
    res := make([]*merchantResolver, 3)
    return &res
}

func (r *productResolver) Stock() *[]*stockResolver {
    res := make([]*stockResolver, 3)
    return &res
}

func (r *productResolver) Attrs() *[]*attributeResolver {
    res := make([]*attributeResolver, 3)
    return &res
}

func (r *productResolver) Comments() *commentsConnectionResolver {
    res := commentsConnectionResolver{}
    return &res
}
