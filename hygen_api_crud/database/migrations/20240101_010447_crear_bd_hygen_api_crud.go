package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearBdHygenApiCrud_20240101_010447 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearBdHygenApiCrud_20240101_010447{}
	m.Created = "20240101_010447"

	migration.Register("CrearBdHygenApiCrud_20240101_010447", m)
}

// Run the migrations
func (m *CrearBdHygenApiCrud_20240101_010447) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CrearBdHygenApiCrud_20240101_010447) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
