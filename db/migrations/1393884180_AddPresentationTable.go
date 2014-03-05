package main

import (
	"github.com/eaigner/hood"
  _ "github.com/ziutek/mymysql/godrv"
)

type Presentation struct {
  Id               hood.Id
  Title            string `sql:"size(255)"`
  Body             string

  Created          hood.Created
  Updated          hood.Updated
}

func (m *M) AddPresentationTable_1393884180_Up(hd *hood.Hood) {
  hd.CreateTable(&Presentation{})

  hd.CreateIndex(&Presentation{}, "title_index", false, "title")
}

func (m *M) AddPresentationTable_1393884180_Down(hd *hood.Hood) {
  hd.DropTable(&Presentation{})
  hd.DropIndex(&Presentation{}, "title_index")
}