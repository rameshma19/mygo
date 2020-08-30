package datastore

// Getter Interface
type Getter interface {
	GetAllItems() []Item
}

// Adder Interface
type Adder interface {
	AddItem(item Item)
}

// Item Newsfeed item
type Item struct {
	Title    string `json:"title"`
	NewsPost string `json:"post"`
}

// Repo List of Items
type Repo struct {
	Items []Item
}

// New creates a new repo
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

// AddItem Adds a new item to Repo on Post
func (r *Repo) AddItem(item Item) {
	r.Items = append(r.Items, item)
}

// GetAllItems Return the list of items on Get
func (r *Repo) GetAllItems() []Item {
	return r.Items
}
