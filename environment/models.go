package environment

type PostMessage struct {
  Message string `json:"message" form:"message" binding:"required" query:"message"`
}

type PostMessageResponse struct {
	Digest string `json:"digest"`
}

type GetMessageResponse struct {
	Message string `json:"message"`
}

// Database Json Model

type DbMessageResponse struct {
    MessageValue string `json:"messageValue" bson:"messageValue"`
}