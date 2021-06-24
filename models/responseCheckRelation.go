package models

/* ResponseCheckRelation handles the json response to the client to check if a relation exists
   between to users*/
type ResponseCheckRelation struct {
	Status bool `json:"status"`
}
