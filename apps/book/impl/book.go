package impl

import (
	"context"
	"database/sql"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"

	"github.com/defeng-hub/mcube-demo/apps/book"
)

func (s *service) CreateBook(ctx context.Context, req *book.CreateBookRequest) (
	*book.Book, error) {
	ins, err := book.NewBook(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create book error, %s", err)
	}

	stmt, err := s.db.Prepare(insertBookSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		ins.Id, ins.CreateAt, ins.Data.CreateBy, ins.UpdateAt, ins.UpdateBy,
		ins.Data.Name, ins.Data.Author,
	)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DescribeBook(ctx context.Context, req *book.DescribeBookRequest) (
	*book.Book, error) {
	query := sqlbuilder.NewQuery(queryBookSQL)
	querySQL, args := query.Where("id = ?", req.Id).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query book error, %s", err.Error())
	}
	defer queryStmt.Close()

	ins := book.NewDefaultBook()
	err = queryStmt.QueryRow(args...).Scan(
		&ins.Id, &ins.CreateAt, &ins.Data.CreateBy, &ins.UpdateAt, &ins.UpdateBy,
		&ins.Data.Name, &ins.Data.Author,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFound("%s not found", req.Id)
		}
		return nil, exception.NewInternalServerError("describe book error, %s", err.Error())
	}

	return ins, nil
}

func (s *service) UpdateBook(ctx context.Context, req *book.UpdateBookRequest) (
	*book.Book, error) {
	ins, err := s.DescribeBook(ctx, book.NewDescribeBookRequest(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case book.UpdateMode_PUT:
		ins.Update(req)
	case book.UpdateMode_PATCH:
		err := ins.Patch(req)
		if err != nil {
			return nil, err
		}
	}

	// 校验更新后数据合法性
	if err := ins.Data.Validate(); err != nil {
		return nil, err
	}

	if err := s.updateBook(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DeleteBook(ctx context.Context, req *book.DeleteBookRequest) (
	*book.Book, error) {
	ins, err := s.DescribeBook(ctx, book.NewDescribeBookRequest(req.Id))
	if err != nil {
		return nil, err
	}

	if err := s.deleteBook(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}
