package data

//This is a video library struct
type VideoLibrary struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Langague string `json:"lang"`
	Link     string `json:"link"`
}
