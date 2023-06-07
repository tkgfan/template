// Package mongos
// author gmfan
// date 2023/2/27
package mongos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Model interface {
	SetID(id primitive.ObjectID)
}

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UpdatedBy int64              `bson:"updated_by"`
	UpdatedAt int64              `bson:"updated_at"`
	CreatedBy int64              `bson:"created_by"`
	CreatedAt int64              `bson:"created_at"`
	IsDeleted int8               `bson:"is_deleted"`
}

func (b *BaseModel) CreatedByUID(uid int64) {
	b.CreatedAt = time.Now().Unix()
	b.CreatedBy = uid
}

func (b *BaseModel) UpdatedByUID(uid int64) {
	b.UpdatedAt = time.Now().Unix()
	b.UpdatedBy = uid
}

func (b *BaseModel) SetID(id primitive.ObjectID) {
	b.ID = id
}
