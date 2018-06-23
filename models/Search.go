package models

type SearchText struct {
	Text         string   `bson:"text" json:"text"`
	Title        string   `bson:"title" json:"title"`
	Description  string   `bson:"description" json:"description"`
	Tags         []string `bson:"tags" json:"tags"`
	Category     string   `bson:"category" json:"category"`
	Owner        int      `bson:"owner" json:"owner"`
	Contributors []string `bson:"contributors" json:"contributors"`
	Location     string   `bson:"location" json:"location"`
	Tasks        []Task   `bson:"tasks" json:"tasks"`
}
