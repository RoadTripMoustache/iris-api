package api

type Response struct {
	Links Links       `json:"links"`
	Data  interface{} `json:"data"`
}
