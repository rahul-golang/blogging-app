package models

type Blog struct {
	Times
	Tittle    string `json:"tittle"`
	RelatedTo string `json:"related_to"`
	Containt  string `json:"containt"`
	UserID    int    `json:"user_id" gorm:"foreignkey:user_id"`
}

type CreateBlogReq struct {
	Blog Blog `json:"blog"`
}

type CreateBlogResp struct {
	Message string `json:"message"`
	Blog    *Blog  `json:"blog"`
}
