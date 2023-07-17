package models

import (
	"github.com/gocql/gocql"
);
	

type Data struct{
	DataID    gocql.UUID   `cql:"dataid"   json:"dataid"`;
	Username  string       `cql:"username" json:"username"`;
	Password  string       `cql:"password" json:"password"`;
	Email     string       `cql:"email"    json:"email"`;
	Items     []Item       `cql:"items"    json:"items"`;
}

type Item struct {
	Key       string     `cql:"key"       json:"key"`
	Value     string     `cql:"value"     json:"value"`
	IsSecret  bool       `cql:"is_secret" json:"is_secret"`
}