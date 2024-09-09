package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Conversation struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	SenderId   primitive.ObjectID `bson:"sender_id,omitempty"`
	ReceiverId primitive.ObjectID `bson:"receiver_id,omitempty"`
	CreatedAt  time.Time          `bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty"`
}

type Message struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ConversationId primitive.ObjectID `bson:"conversation_id,omitempty"`
	SenderId       primitive.ObjectID `bson:"sender_id,omitempty"`
	Message        string             `json:"message"`
	CreatedAt      time.Time          `bson:"created_at,omitempty"`
	UpdatedAt      time.Time          `bson:"updated_at,omitempty"`
}
