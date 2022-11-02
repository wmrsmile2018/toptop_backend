package awesomeProject

type Post struct {
	Id               int    `json:"id"`
	IdCategory       int    `json:"id_category"`
	IdUser           int    `json:"id_user"`
	IdTags           int    `json:"id_tags"`
	Label            string `json:"label"`
	Address          string `json:"address"`
	CreationDate     string `json:"creation_date"` // Todo change to date
	FullDescription  string `json:"full_description"`
	ShortDescription string `json:"short_description"`
}

type Category struct {
	Name string `json:"name"`
}
