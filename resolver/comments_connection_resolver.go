package resolver

import (
	"pinjihui.com/pinjihui/model"
	"pinjihui.com/pinjihui/service"
)

type commentsConnectionResolver struct {
	users      []*model.User
	totalCount int
	from       *string
	to         *string
}

func (r *commentsConnectionResolver) TotalCount() int32 {
	return int32(r.totalCount)
}

func (r *commentsConnectionResolver) Edges() *[]*usersEdgeResolver {
	l := make([]*usersEdgeResolver, len(r.users))
	for i := range l {
		l[i] = &usersEdgeResolver{
			cursor: service.EncodeCursor(&(r.users[i].ID)),
			model:  r.users[i],
		}
	}
	return &l
}

func (r *commentsConnectionResolver) PageInfo() *pageInfoResolver {
	return &pageInfoResolver{
		startCursor: service.EncodeCursor(r.from),
		endCursor:   service.EncodeCursor(r.to),
		hasNextPage: false,
	}
}
