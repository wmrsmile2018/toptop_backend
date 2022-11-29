package awesomeProject

type Place struct {
	Id               string `json:"id"`
	IdCategory       string `json:"id_category"`
	IdUser           string `json:"id_user"`
	IdTags           string `json:"id_tags"`
	Label            string `json:"label"`
	Address          string `json:"address"`
	CreationDate     string `json:"creation_date"` // Todo change to date
	FullDescription  string `json:"full_description"`
	ShortDescription string `json:"short_description"`
}

type Category struct {
	Name string `json:"name"`
}
