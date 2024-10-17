package schemas

type EventInput struct {
	Title            string `json:"title" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Location         string `json:"location" binding:"required"`
	Date             string `json:"date" binding:"required"`
	Time             string `json:"time" binding:"required"`
	IsFeatured       bool   `json:"is_featured"`
}

type EventUpdate struct {
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	Location         string `json:"location"`
	Date             string `json:"date"`
	Time             string `json:"time"`
	IsFeatured       bool   `json:"is_featured"`
}

type Event struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	Location         string `json:"location"`
	Date             string `json:"date"`
	Time             string `json:"time"`
	IsFeatured       bool   `json:"is_featured"`
	CreatedBy        int    `json:"created_by"`
}
