package book

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"github.com/rs/xid"
)

const (
	AppName = "book"
)

var (
	validate = validator.New()
)

func NewCreateBookRequest() *CreateBookRequest {
	return &CreateBookRequest{}
}

func NewBook(req *CreateBookRequest) (*Book, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Book{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMicro(),
		Data:     req,
	}, nil
}

func (req *CreateBookRequest) Validate() error {
	return validate.Struct(req)
}

func NewBookSet() *BookSet {
	return &BookSet{
		Items: []*Book{},
	}
}

func (s *BookSet) Add(item *Book) {
	s.Items = append(s.Items, item)
}

func NewDefaultBook() *Book {
	return &Book{
		Data: &CreateBookRequest{},
	}
}

func (i *Book) Update(req *UpdateBookRequest) {
	i.UpdateAt = time.Now().UnixMicro()
	i.UpdateBy = req.UpdateBy
	i.Data = req.Data
}

func (i *Book) Patch(req *UpdateBookRequest) error {
	i.UpdateAt = time.Now().UnixMicro()
	i.UpdateBy = req.UpdateBy
	return mergo.MergeWithOverwrite(i.Data, req.Data)
}

func NewDescribeBookRequest(id string) *DescribeBookRequest {
	return &DescribeBookRequest{
		Id: id,
	}
}

func NewQueryBookRequestFromHTTP(r *http.Request) *QueryBookRequest {
	qs := r.URL.Query()

	return &QueryBookRequest{
		Page:     &PageRequest{PageSize: 1, PageNumber: 1, Offset: 0},
		Keywords: qs.Get("keywords"),
	}
}

func NewDeleteBookRequestWithID(id string) *DeleteBookRequest {
	return &DeleteBookRequest{
		Id: id,
	}
}
