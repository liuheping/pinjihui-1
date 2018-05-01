
package resolver

import (
    //"pinjihui.com/pinjihui/model"
    "github.com/graph-gophers/graphql-go"
)

type productImageResolver struct {
    //m *model.productImage
}

func (r *productImageResolver) ID() graphql.ID {
    res := graphql.ID("xjauwkahsi92h1j")
    return res
}

func (r *productImageResolver) SamllImage() string {
    res := "https://img10.360buyimg.com/mobilecms/s500x500_jfs/t13828/23/991435575/195172/ebf643d9/5a17dc63N9f1003de.jpg"
    return res
}

func (r *productImageResolver) MediumImage() string {
    res := "https://img10.360buyimg.com/mobilecms/s500x500_jfs/t13828/23/991435575/195172/ebf643d9/5a17dc63N9f1003de.jpg"
    return res
}

func (r *productImageResolver) BigImage() string {
    res := "https://img10.360buyimg.com/mobilecms/s500x500_jfs/t13828/23/991435575/195172/ebf643d9/5a17dc63N9f1003de.jpg"
    return res
}
