package domain

type Info struct {
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Menu struct {
	MenuItems []MenuItem `json:"menuItems,omitempty"`
}

type MenuItem struct {
	ID        string     `json:"id,omitempty"`
	Title     string     `json:"title"`
	Url       string     `json:"url"`
	MenuItems []MenuItem `json:"items,omitempty"`
}
