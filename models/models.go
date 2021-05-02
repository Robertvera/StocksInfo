package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Value struct {
	Date  string `bson:"date,omitempty"`
	Price string `bson:"price,omitempty"`
}

type ETF struct {
	ISIN     primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Type     string             `bson:"type,omitempty"`
	Currency string             `bson:"currency,omitempty"`
	Values   []Value            `bson:"values,omitempty"`
}
