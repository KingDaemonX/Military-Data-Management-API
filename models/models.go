package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Army struct {
	Soldier *Soldier `json:"soldier"`
	// Division *Division `json:"division"`
}

type Soldier struct {
	ID                  primitive.ObjectID `bson:"_id"`
	User_id             string             `json:"user_id"`
	First_name          string             `json:"first_name" validate:"required,min=2,max=200"`
	Last_name           string             `json:"last_name" validate:"required,min=2,max=200"`
	Nick_name           string             `json:"nick_name" validate:"required,min=2,max=200"`
	Email               string             `json:"email" validate:"required,email"`
	Password            string             `json:"password" validate:"required,min=8"`
	User_type           string             `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Army_number         string             `json:"army_id" validate:"required,min=2,max=200"`
	Age                 uint               `json:"age" validate:"required"`
	Year_of_recruitment time.Time          `json:"year_recruited" validate:"required min=2,max=200"`
	Updated_at          time.Time          `json:"updated" validate:"required"`
	Rank                string             `json:"rank" validate:"required"`
	Next_of_kin         string             `json:"next_of_kin" validate:"required,min=2,max=200"`
	Resident_barrack    string             `json:"resident_barracks" validate:"required,min=2,max=200"`
	Address             string             `json:"address" validate:"required,min=2,max=200"`
	Place_of_service    string             `json:"place_of_service" validate:"required,min=2,max=200"`
	Is_assigned_arm     bool               `json:"is_armed"`
	Division            *Division          `json:"division"`
	Token               string             `json:"token"`
	Refresh_Token       string             `json:"refresh_token"`
}

type Division struct {
	Division_name  string `json:"division_name" validate:"required,min=2,max=200"`
	Commander_name string `json:"commander" validate:"required,min=2,max=200"`
	Location       string `json:"location" validate:"required,min=2,max=200"`
	Position       string `json:"position" validate:"required,min=2,max=200"`
	Department     string `json:"department" validate:"required,min=2,max=200"`
}
