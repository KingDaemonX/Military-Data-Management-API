package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/KingAnointing/go-project/configs"
	"github.com/KingAnointing/go-project/models"
	"github.com/KingAnointing/go-project/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collections *mongo.Collection = configs.DbCollection()
var validate = validator.New()

// greeter function
func Greeter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "sucess", Data: map[string]interface{}{"data": "Welcome to my personal skill test on gin, monogoDB & golang CRUD API"}})
	}
}

// create function
func CreateASoldierProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cbg, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var soldier models.Army

		// validate json response
		if err := ctx.BindJSON(&soldier); err != nil {
			ctx.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// validate for required field and validation rule giving
		if validateErr := validate.Struct(&soldier); validateErr != nil {
			ctx.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validateErr.Error()}})
			return
		}

		// serialize the data into soldier profile
		soldierProfile := models.Army{
			&models.Soldier{
				ID:              primitive.NewObjectID(),
				FirstName:       soldier.Soldier.FirstName,
				LastName:        soldier.Soldier.LastName,
				NickName:        soldier.Soldier.NickName,
				ArmyNumber:      soldier.Soldier.ArmyNumber,
				Age:             soldier.Soldier.Age,
				Rank:            soldier.Soldier.Rank,
				NextOfKin:       soldier.Soldier.NextOfKin,
				ResidentBarrack: soldier.Soldier.ResidentBarrack,
				Address:         soldier.Soldier.Address,
				PlaceOfService:  soldier.Soldier.PlaceOfService,
				IsAssignedArm:   soldier.Soldier.IsAssignedArm,
				Division: &models.Division{
					DivisionName:  soldier.Soldier.Division.DivisionName,
					CommanderName: soldier.Soldier.Division.CommanderName,
					Location:      soldier.Soldier.Division.Location,
					Position:      soldier.Soldier.Division.Position,
					Department:    soldier.Soldier.Division.Department,
				},
			},
		}
		// &models.Division{
		// 	DivisionName:  soldier.Division.DivisionName,
		// 	CommanderName: soldier.Division.CommanderName,
		// 	Location:      soldier.Division.Location,
		// 	Position:      soldier.Division.Position,
		// 	Department:    soldier.Division.Department,
		// },
		// },
		result, err := collections.InsertOne(cbg, soldierProfile)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		ctx.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "created", Data: map[string]interface{}{"data": result}})
	}
}

func GetOneSoldierProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("userId")
		id, _ := primitive.ObjectIDFromHex(userId)

		cbg, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var soldier models.Army

		if err := collections.FindOne(cbg, bson.M{"_id": id}).Decode(&soldier); err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": soldier}})
	}
}

func UpdateSoldierProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		id, _ := primitive.ObjectIDFromHex(userId)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		
	}
}
