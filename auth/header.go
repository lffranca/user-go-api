package auth

// Header Header
type Header struct {
	Authorization string `header:"Authorization" json:"authorization" binding:"required"`
}
