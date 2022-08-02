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
	ID                primitive.ObjectID `json:"id,omitempty"`
	FirstName         string             `json:"first_name" validate:"required,min=2,max=200"`
	LastName          string             `json:"last_name" validate:"required,min=2,max=200"`
	NickName          string             `json:"nick_name" validate:"required,min=2,max=200"`
	Email             string             `json:"email" validate:"required,email"`
	Password          string             `json:"password" validate:"required,min=8"`
	ArmyNumber        string             `json:"army_id" validate:"required,min=2,max=200"`
	Age               uint               `json:"age" validate:"required"`
	YearOfRecruitment time.Time          `json:"year_recruited" validate:"required min=2,max=200"`
	Updated           time.Time          `json:"updated" validate:"required"`
	Rank              string             `json:"rank" validate:"required"`
	NextOfKin         string             `json:"next_of_kin" validate:"required,min=2,max=200"`
	ResidentBarrack   string             `json:"resident_barracks" validate:"required,min=2,max=200"`
	Address           string             `json:"address" validate:"required,min=2,max=200"`
	PlaceOfService    string             `json:"place_of_service" validate:"required,min=2,max=200"`
	IsAssignedArm     bool               `json:"is_armed"`
	Division          *Division          `json:"division"`
	Token             string             `json:"token"`
	Refresh_Token     string             `json:"refresh_token"`
}

type Division struct {
	DivisionName  string `json:"division_name" validate:"required,min=2,max=200"`
	CommanderName string `json:"commander" validate:"required,min=2,max=200"`
	Location      string `json:"location" validate:"required,min=2,max=200"`
	Position      string `json:"position" validate:"required,min=2,max=200"`
	Department    string `json:"department" validate:"required,min=2,max=200"`
}
