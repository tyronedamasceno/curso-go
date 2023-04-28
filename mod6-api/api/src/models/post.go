package models

import (
	"errors"
)

type Post struct {
	ID       uint64 `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Content  string `json:"content,omitempty"`
	AuthorID uint64 `json:"authorId,omitempty"`
	Likes    uint64 `json:"likes"`
}

func (post *Post) Prepare() error {
	err := post.validate()
	return err
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("title can't be blank")
	}
	if post.Content == "" {
		return errors.New("content can't be blank")
	}
	return nil
}
