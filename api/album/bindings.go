package album

type DeleteRequest struct {
	Key string `json:"key" binding:"required"`
}
