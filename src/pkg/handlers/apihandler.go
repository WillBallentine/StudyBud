package handlers

import (
	"context"
	"net/http"
	"time"

	"studybud/src/cmd/utils"
	"studybud/src/pkg/entity"
	"studybud/src/pkg/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiHandler interface {
	Healthcheck(*gin.Context)
	AddUser(*gin.Context)
	GetUserById(*gin.Context)
}

type apiHandler struct {
	client          *mongo.Client
	mongoRepository repository.MongoRepository
	config          utils.Configuration
}

func NewApiHandler(client *mongo.Client, repo repository.MongoRepository, config utils.Configuration) ApiHandler {
	return &apiHandler{client: client, mongoRepository: repo, config: config}
}

func (hand *apiHandler) Healthcheck(c *gin.Context) {
	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(hand.config.App.Timeout)*time.Second)

	if ctxErr != nil {
		logrus.Error("borked", ctxErr)
	}

	if err := hand.client.Ping(ctx, nil); err != nil {
		utils.InternalServerError("mongo unhealthy", err, map[string]interface{}{"Data": "please check client", "time": time.Local})
	}

	c.IndentedJSON(http.StatusOK, utils.Response("healthy", map[string]interface{}{"Data": "mongo healthy", "time": time.Local}))
}

func (hand *apiHandler) AddUser(c *gin.Context) {
	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(hand.config.App.Timeout)*time.Second)

	defer ctxErr()

	var userEntity *entity.User

	userEntity = &entity.User{
		FirstName:         c.Param("first_name"),
		LastName:          c.Param("last_name"),
		Email:             c.Param("email"),
		Password:          c.Param("pass"),
		School:            c.Param("school"),
		SubscriptionLevel: c.Param("sub_level"),
	}

	entity := entity.User(*userEntity)

	oId, err := hand.mongoRepository.AddUser(entity, ctx)

	if err != nil {
		utils.BadRequestError("add_user_handler", err, map[string]interface{}{"Data": entity})
	}

	c.IndentedJSON(http.StatusCreated, utils.Response("add_user_handler", map[string]interface{}{"OId": oId}))
}

func (hand *apiHandler) GetUserById(c *gin.Context) {
	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(hand.config.App.Timeout)*time.Second)

	defer ctxErr()

	id := c.Param("id")

	oId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		utils.BadRequestError("getuserbyid_handler", err, map[string]interface{}{"Data": id})
	}

	logrus.Infof("OId %s", oId)

	result, err := hand.mongoRepository.GetUserById(oId, ctx)
	if err != mongo.ErrNilCursor {
		utils.BadRequestError("getuserbyid_handler", err, map[string]interface{}{"Data": id})
	}

	if result == nil {
		utils.NotFoundRequestError("getuserbyid_handler", err, map[string]interface{}{"Data": id})
	}

	c.IndentedJSON(http.StatusOK, utils.Response("getuserbyid_handler", map[string]interface{}{"Data": result}))

}
