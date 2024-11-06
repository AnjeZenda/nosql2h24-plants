package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Photo       string               `bson:"photo" json:"photo"`
	Surname     string               `bson:"surname" json:"surname"`
	Name        string               `bson:"name" json:"name"`
	FatherName  string               `bson:"father_name" json:"father_name"`
	PhoneNumber string               `bson:"phone_number" json:"phone_number"`
	Email       string               `bson:"email" json:"email"`
	Password    string               `bson:"password" json:"-"`
	CreatedAt   time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time            `bson:"updated_at" json:"updated_at"`
	Plants      []Plant              `bson:"plants" json:"plants"`
	Trades      []primitive.ObjectID `bson:"trades" json:"trades"`
}

type Plant struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Image             string             `bson:"image" json:"image"`
	UserID            primitive.ObjectID `bson:"user_id" json:"user_id"`
	Size              string             `bson:"size" json:"size"`
	Price             float64            `bson:"price" json:"price"`
	LightCondition    string             `bson:"light_condition" json:"light_condition"`
	WateringFrequency string             `bson:"watering_frequency" json:"watering_frequency"`
	TemperatureRegime string             `bson:"temperature_regime" json:"temperature_regime"`
	CareComplexity    string             `bson:"care_complexity" json:"care_complexity"`
	Description       string             `bson:"description" json:"description"`
	Type              string             `bson:"type" json:"type"`
	Species           string             `bson:"species" json:"species"`
	CareRules         primitive.ObjectID `bson:"care_rules" json:"care_rules"`
	CreatedAt         time.Time          `bson:"created_at" json:"created_at"`
}

type TradeUser struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Plant TradePlant         `bson:"plant" json:"plant"`
}

type Trade struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Offerer   TradeUser          `bson:"offerer" json:"offerer"`
	Accepter  TradeUser          `bson:"accepter" json:"accepter"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Status    int                `bson:"status" json:"status"`
}

type TradePlant struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
}

type CareRules struct {
	ID          primitive.ObjectID    `bson:"_id,omitempty" json:"id"`
	Species     string                `bson:"species" json:"species"`
	Description []CareDescriptionPart `bson:"description" json:"description"`
	CreatedAt   time.Time             `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time             `bson:"updated_at" json:"updated_at"`
}

type CareDescriptionPart struct {
	UserID              primitive.ObjectID `bson:"user_id" json:"user_id"`
	DescriptionAddition string             `bson:"description_addition" json:"description_addition"`
	CreatedAt           time.Time          `bson:"created_at" json:"created_at"`
}
