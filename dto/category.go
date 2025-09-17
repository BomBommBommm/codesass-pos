package dto

type CategoryRequest struct {
	// {"type":"Something"}				<<เพิ่มแท็คให้ type เข้ากับ string
	// Name string `json:"type"`
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}