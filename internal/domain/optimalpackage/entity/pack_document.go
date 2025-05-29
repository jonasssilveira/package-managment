package entity

type PackDocument struct {
	ID   string `bson:"_id"`
	Size int64  `bson:"size"`
}
