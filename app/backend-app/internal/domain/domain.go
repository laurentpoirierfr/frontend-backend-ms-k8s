package domain

type Info struct {
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MenuItem struct {
	ID        string     `json:"id,omitempty"`
	Title     string     `json:"title"`
	Url       string     `json:"url"`
	Target    string     `json:"target,omitempty"`
	MenuItems []MenuItem `json:"items,omitempty"`
}
