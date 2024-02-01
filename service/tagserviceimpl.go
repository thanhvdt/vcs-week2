package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/thanhvdt/vcs-week2/data/request"
	"github.com/thanhvdt/vcs-week2/data/response"
	"github.com/thanhvdt/vcs-week2/model"
	"github.com/thanhvdt/vcs-week2/repository"
)

type TagServiceImpl struct {
	TagRepository repository.TagsRepository
	Validate      *validator.Validate
}

func (t *TagServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	if err != nil {
		panic(err)
	}
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagRepository.Save(tagModel)
}

func (t *TagServiceImpl) Update(tags request.UpdateTagsRequest) {
	err := t.Validate.Struct(tags)
	if err != nil {
		panic(err)
	}
	tagModel := model.Tags{
		Id:   tags.Id,
		Name: tags.Name,
	}
	t.TagRepository.Update(tagModel)
}

func (t TagServiceImpl) Delete(tagsId int) {
	//TODO implement me
	panic("implement me")
}

func (t TagServiceImpl) FindById(tagsId int) response.TagsResponse {
	//TODO implement me
	panic("implement me")
}

func (t TagServiceImpl) FindAll() []response.TagsResponse {
	//TODO implement me
	panic("implement me")
}

func NewTagServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagService {
	return &TagServiceImpl{TagRepository: tagRepository, Validate: validate}
}
