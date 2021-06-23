package models

// UserFollowUser is the model that saves the relation between user1(account) and user2(to follow)
type UserFollowUser struct {
	User1ID string `bson:"user1ID,omitempty" json:"user1ID,omitempty"`
	User2ID string `bson:"user2ID,omitempty" json:"user2ID,omitempty"`
}
