package router

import (
	"allcaps/api/mergehandler"
	"allcaps/api/repository"
	"allcaps/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handlerouter(r *gin.Engine, c *http.Client) {

	repos := repository.InitRepoClient(c)
	service := service.InitServiceClient(repos)
	handler := mergehandler.MergeHandler{Service: service}

	r.GET("get-values", handler.GetClientData)
	r.GET("post-values", handler.PostClientData)
}
