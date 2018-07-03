package ov

type Menu struct {
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Path      string `json:"path"`
	Children  []Menu `json:"children"`
	Authority string `json:"authority"`
}
type Page struct {
	List       interface{} `json:"list"`
	Pagination Pagination `json:"pagination"`
}
type Pagination struct {
	Total int
}
type Treemenu struct {
	Title           string `json:"title"`
	Key             string `json:"key"`
	DisableCheckbox bool `json:"disableCheckbox"`
	Disabled        bool `json:"disabled"`
	Icon            string `json:"icon"`
	IsLeaf          bool `json:"isLeaf"`
	Selectable      bool `json:"selectable"`
	Children        []Treemenu `json:"children"`
}
