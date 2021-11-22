package model

type Mock struct {
	Uid         string   `firestore:"uid" json:"uid,omitempty"`
	UserUid     string   `firestore:"user_uid" json:"user_uid,omitempty"`
	Position    string   `firestore:"position" json:"position,omitempty"`
	Level       string   `firestore:"level" json:"level,omitempty"`
	Stacks      []string `firestore:"stacks" json:"stacks,omitempty"`
	Language    string   `firestore:"language" json:"language,omitempty"`
	CoverLetter string   `firestore:"cover_letter" json:"cover_letter,omitempty"`
	Status      string   `firestore:"status" json:"status,omitempty"`
}
