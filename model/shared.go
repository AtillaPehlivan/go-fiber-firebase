package model

type UidName struct {
	Uid  string `firestore:"uid" json:"uid"`
	Name string `firestore:"name" json:"name"`
}
