package response

type Response struct {
	Links Links       `json:"links"`
	Data  interface{} `json:"data"`
}
