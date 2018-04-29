
package resolver

import (
    //"pinjihui.com/pinjihui/model"
)

type addressResolver struct {
    //m *model.address
}

func (r *addressResolver) ProvinceId() int32 {
    res := int32(3)
    return res
}

func (r *addressResolver) CityId() int32 {
    res := int32(3)
    return res
}

func (r *addressResolver) AreaId() int32 {
    res := int32(3)
    return res
}

func (r *addressResolver) RegionName() *string {
    res := "test string"
    return &res
}

func (r *addressResolver) Address() string {
    res := "test string"
    return res
}
