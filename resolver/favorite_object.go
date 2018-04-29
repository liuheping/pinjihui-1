package resolver

type favoriteObjectResolver struct {
    result interface{}
}

func (r *favoriteObjectResolver) ToProduct() (*productResolver, bool) {
    res, ok := r.result.(*productResolver)
    return res, ok
}

func (r *favoriteObjectResolver) ToMerchant() (*merchantResolver, bool) {
    res, ok := r.result.(*merchantResolver)
    return res, ok
}
