package db

import (
	"github.com/eaigner/hood"
)

type Presentation struct {
	Id      hood.Id
	Title   string `sql:"size(255)"`
	Body    string
	Created hood.Created
	Updated hood.Updated
}

func (table *Presentation) Indexes(indexes *hood.Indexes) {
	indexes.Add("title_index", "title")
}
