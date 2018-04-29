
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "time"
    "github.com/graph-gophers/graphql-go"
)

type orderResolver struct {
    //m *model.order
}

func (r *orderResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *orderResolver) User() *userResolver {
    res := userResolver{}
    return &res
}

func (r *orderResolver) OrderStatus() string {
    res := "invalid"
    return res
}

func (r *orderResolver) ShippingStatus() string {
    res := "unshipped"
    return res
}

func (r *orderResolver) PayStatus() string {
    res := "unpaid"
    return res
}

func (r *orderResolver) ChildrenOrders() *[]*orderResolver {
    res := make([]*orderResolver, 3)
    return &res
}

func (r *orderResolver) Address() *orderAddressResolver {
    res := orderAddressResolver{}
    return &res
}

func (r *orderResolver) Merchant() *merchantResolver {
    res := merchantResolver{}
    return &res
}

func (r *orderResolver) Products() *[]*productInOrderResolver {
    res := make([]*productInOrderResolver, 3)
    return &res
}

func (r *orderResolver) ShippingName() string {
    res := "test string"
    return res
}

func (r *orderResolver) PayName() string {
    res := "test string"
    return res
}

func (r *orderResolver) HowOos() *string {
    res := "cancel"
    return &res
}

func (r *orderResolver) Amount() float64 {
    res := 0.0
    return res
}

func (r *orderResolver) ShippingFee() float64 {
    res := 0.0
    return res
}

func (r *orderResolver) MoneyPaid() float64 {
    res := 0.0
    return res
}

func (r *orderResolver) OrderAmount() float64 {
    res := 0.0
    return res
}

func (r *orderResolver) CreatedAt() (graphql.Time, error) {
    res, err := time.Parse(time.RFC3339, "2018-04-01 12:04:56.539453")
    return graphql.Time{Time: res}, err
}

func (r *orderResolver) ConfirmTime() (*graphql.Time, error) {
    res, err := time.Parse(time.RFC3339, "2018-04-01 12:04:56.539453")
    return &graphql.Time{Time: res}, err
}

func (r *orderResolver) PayTime() (*graphql.Time, error) {
    res, err := time.Parse(time.RFC3339, "2018-04-01 12:04:56.539453")
    return &graphql.Time{Time: res}, err
}

func (r *orderResolver) ShippingTime() (*graphql.Time, error) {
    res, err := time.Parse(time.RFC3339, "2018-04-01 12:04:56.539453")
    return &graphql.Time{Time: res}, err
}
