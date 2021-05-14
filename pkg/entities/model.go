package entities

type EntityUser struct {
	ID       string `bson:"id"`
	Email    string `bson:"email"`
	Username string `bson:"username"`
}
