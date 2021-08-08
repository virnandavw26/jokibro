package request

type CategoryRequest struct {
	ID   uint   `json:"id" query:"id" path:"id"`
	Name string `json:"name" query:"name"`
}
