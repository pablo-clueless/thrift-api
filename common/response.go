package common

type Response struct {
	Error   bool        `json:"error" bson:"error"`
	Code    int         `json:"code" bson:"code"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}
