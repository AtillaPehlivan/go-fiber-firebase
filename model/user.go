package model

type User struct {
	Uid           string `firestore:"uid" json:"uid,omitempty"`
	Email         string `firestore:"email" json:"email,omitempty"`
	EmailVerified bool   `firestore:"emailVerified" json:"emailVerified,omitempty"`
	DisplayName   string `firestore:"displayName" json:"displayName,omitempty"`
	PhoneNumber   string `firestore:"phoneNumber" json:"phoneNumber,omitempty"`
	PhotoURL      string `firestore:"photoURL" json:"photoURL,omitempty"`
	ProviderId    string `firestore:"ProviderId" json:"ProviderId,omitempty"`
}
