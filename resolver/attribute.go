
package resolver

import (
    //"pinjihui.com/pinjihui/model"
)

type attributeResolver struct {
    //m *model.attribute
}

func (r *attributeResolver) Name() string {
    res := "test string"
    return res
}

func (r *attributeResolver) Value() *string {
    res := "test string"
    return &res
}
