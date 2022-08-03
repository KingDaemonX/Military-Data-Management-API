package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KingAnointing/go-project/configs"
	"github.com/KingAnointing/go-project/helpers"
	"github.com/KingAnointing/go-project/models"
	"github.com/KingAnointing/go-project/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var collections *mongo.Collection = configs.DbCollection()
var validate = validator.New()

// greeter function
func Greeter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "sucess", Data: map[string]interface{}{"data": "Welcome to my personal skill test on gin, monogoDB & golang CRUD API"}})
	}
}

func hashPassword(c *gin.Context, password string) string {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		log.Fatal(err)
	}
	return string(byte)
}
func checkCount(count int64, err error, c *gin.Context) {
	if count > 0 {
		c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": "Email or Phone Already Exist"}})
		return
	}
}

// create function
func CreateASoldierProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cbg, cancel := context.WithTimeout(context.Background(), time.Second*100)
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

		count, err := collections.CountDocuments(cbg, bson.M{"email": soldier.Soldier.Email})
		checkCount(count, err, ctx)

		count, err = collections.CountDocuments(cbg, bson.M{"phone": soldier.Soldier.Phone})
		checkCount(count, err, ctx)

		soldier.Soldier.Year_of_recruitment, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		soldier.Soldier.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		soldier.Soldier.ID = primitive.NewObjectID()
		soldier.Soldier.User_id = soldier.Soldier.ID.Hex()
		token, refreshToken, err := helpers.GenerateAllToken(*soldier.Soldier.First_name, *soldier.Soldier.Last_name, *soldier.Soldier.Email, soldier.Soldier.User_id, *soldier.Soldier.User_type)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}
		soldier.Soldier.Token = &token
		soldier.Soldier.Refresh_Token = &refreshToken

		hashedPassword := hashPassword(ctx, *soldier.Soldier.Password)
		soldier.Soldier.Password = &hashedPassword

		result, err := collections.InsertOne(cbg, &soldier)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		ctx.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

		// serialize the data into soldier profile
		/*soldierProfile := models.Army{
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
		} */
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

		if err := collections.FindOne(cbg, bson.M{"soldier.id": id}).Decode(&soldier); err != nil {
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

		var soldier models.Army
		// validate updated input
		if err := c.BindJSON(&soldier); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// validate the structure and required field
		if validateErr := validate.Struct(&soldier); validateErr != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validateErr.Error()}})
			return
		}

		update := bson.M{
			"soldier.age":                    soldier.Soldier.Age,
			"soldier.rank":                   soldier.Soldier.Rank,
			"soldier.nextofkin":              soldier.Soldier.NextOfKin,
			"soldier.residentbarracks":       soldier.Soldier.ResidentBarrack,
			"soldier.address":                soldier.Soldier.Address,
			"soldier.placeofservice":         soldier.Soldier.PlaceOfService,
			"soldier.isassignedarm":          soldier.Soldier.IsAssignedArm,
			"soldier.division.divisionname":  soldier.Soldier.Division.DivisionName,
			"soldier.division.commandername": soldier.Soldier.Division.CommanderName,
			"soldier.division.location":      soldier.Soldier.Division.Location,
			"soldier.division.position":      soldier.Soldier.Division.Position,
			"soldier.division.department":    soldier.Soldier.Division.Department,
		}
		result, err := collections.UpdateOne(ctx, bson.M{"soldier.id": id}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedProfile models.Army

		if result.ModifiedCount == 1 {
			err := collections.FindOne(ctx, bson.M{"soldier.id": id}).Decode(&updatedProfile)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedProfile}})
	}
}

func DeleteASoldierProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		id, _ := primitive.ObjectIDFromHex(userId)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		result, err := collections.DeleteOne(ctx, bson.M{"soldier.id": id})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, responses.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User With Specified Army Number not found"}})
			return
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Profile successfully deleted"}})
	}
}

func GetAllSoldierProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		var soldiers []models.Army

		cursor, err := collections.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var singleSoldier models.Army
			if err := cursor.Decode(&singleSoldier); err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			soldiers = append(soldiers, singleSoldier)
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": soldiers}})
	}
}

func DeleteAllSoldierProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		result, _ := collections.DeleteMany(ctx, bson.D{})

		message := fmt.Sprintf("Deleted all user from database with a count of : %v", result.DeletedCount)

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": message}})

	}
}
