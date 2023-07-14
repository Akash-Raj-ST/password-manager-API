package models

type Data struct{
	Username  string           `cql:"username"`;
	Password  string           `cql:"password"`;
	Items     []Item           `cql:"items"`;
}

type Item struct {
	Key       string `cql:"key"`
	Value     string `cql:"value"`
	IsSecret  bool   `cql:"is_secret"`
}