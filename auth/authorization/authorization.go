package authorization

type ResourceScope struct {
	Resource map[string]string
	Scope    map[string]string
	ECAID    string
}

type Resource struct {
	ResourceID int    `json:"id"`
	Resource   string `json:"resource"`
	Policy     Policy `json:"policy"`
}

type Policy struct {
	CanAdd    bool `json:"can_add"`
	CanDelete bool `json:"can_delete"`
	CanEdit   bool `json:"can_edit"`
	CanView   bool `json:"can_view"`
}

type Role struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Resources []Resource `json:"resources"`
}