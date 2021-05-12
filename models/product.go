package models

type Product struct {
	ProductID     int       `json:"_id"`
	ProductImage  []string  `json:"product_image"`
	ProductName   string    `json:"product_name"`
	Rating        int       `json:"rating"`
	TotalReviews  int       `json:"total_reviews"`
	ActualPrice   string    `json:"actual_price"`
	DiscountPrice string    `json:"discount_price"`
	Description   string    `json:"description"`
	Category      string    `json:"category"`
	Comments      []Comment `json:"comment"`
}

type Comment struct {
	ProductId    int    `json:"product_id"`
	UserName     string `json:"user_name"`
	CommentTitle string `json:"comment_title"`
	Comment      string `json:"comment"`
	Rating       string `json:"rating"`
}
