package entity

type PackDocument struct {
	ID   string `bson:"_id;omitempty"`
	Size int64  `bson:"size"`
}
