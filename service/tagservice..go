package service

import (
	"github.com/thanhvdt/vcs-week2/data/request"
	"github.com/thanhvdt/vcs-week2/data/response"
)

type TagService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
