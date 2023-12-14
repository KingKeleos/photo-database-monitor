package server

type Project struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	State       string `json:"state"`
	Folder      string `json:"folder"`
	Description string `json:"description"`
}

type PeopleList struct {
	People []People `json:"People"`
}

type People struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

type Socials struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Type     string `json:"type"`
	Link     string `json:"link"`
}

type Posts struct {
	Link      string `json:"link"`
	ProjectID string `json:"project_id"`
}
