package models
type Customer struct {
	CustomerId      string            `json:"customerid" bson:"customerid"`
	Firstname       string            `json:"firstname" bson:"firstname"`
	Lastname        string            `json:"lastname" bson:"lastname"`
	Password        string            `json:"password" bson:"password"`
	Email           string            `json:"email" bson:"email"`
}

type DBResponse struct{
	CustomerId      string            `json:"customerid" bson:"customerid"`
}