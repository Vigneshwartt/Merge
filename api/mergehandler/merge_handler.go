package mergehandler

import (
	"allcaps/api/service"
	"allcaps/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MergeHandler struct {
	Service service.IServiceMerge
}

func NewMessage() {
	fmt.Println("called after every ten seconds")
}

func (hand *MergeHandler) GetClientData(c *gin.Context) {
	url := "https://api.merge.dev/api/crm/v1/opportunities"

	response, err := hand.Service.GetClientData(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": response,
	})

}

// func (hand *MergeHandler) PostClientData(c *gin.Context) {
// 	var req models.Opportunity
// 	url := "https://api.merge.dev/api/crm/v1/opportunities"
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
// 		return
// 	}
// 	fmt.Println("Bindind time----------", req)
// 	data := map[string]interface{}{
// 		"model": map[string]interface{}{
// 			"id":                 req.Model.ID,
// 			"remote_id":          req.Model.Amount,
// 			"created_at":         req.Model.CreatedAt,
// 			"remote_was_deleted": req.Model.RemoteWasDeleted,
// 			"remote_created_at":  req.Model.RemoteCreatedAt,
// 			"closedate":          req.Model.CloseDate,
// 			"last_activity_at":   req.Model.CloseDate,
// 			"status":             req.Model.Status,
// 			"amount":             req.Model.Amount,
// 			"description":        req.Model.Description,
// 			"name":               req.Model.Name,
// 			"modified_at":        req.Model.ModifiedAt,
// 			// "owner":              req.Model.Owner,
// 		},
// 	}
// 	fmt.Println("")
// 	fmt.Println("")
// 	fmt.Println("Datataa", data)
// 	response, err := hand.Service.PostClientData(url, data)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.Response{
// 			Error: err.Error(),
// 		})
// 	}
// 	fmt.Println("Response", response)
// 	c.JSON(http.StatusOK, gin.H{
// 		"Data": response,
// 	})
// }

func (hand *MergeHandler) PostClientData(c *gin.Context) {
	var dealDetails []map[string]interface{}

	url := "https://api.merge.dev/api/crm/v1/opportunities"

	response, err := hand.Service.GetClientData(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		return
	}

	for _, val := range response.Results {
		dealDetail := map[string]interface{}{
			"id":                val.ID,
			"remote_id":         val.RemoteID,
			"created_at":        val.CreatedAt,
			"remote_created_at": val.RemoteCreatedAt,
			"closedate":         val.CloseDate,
			"last_activity_at":  val.CloseDate,
			"status":            val.Status,
			"amount":            val.Amount,
			"name":              val.Name,
			"modified_at":       val.ModifiedAt,
		}
		dealDetails = append(dealDetails, dealDetail)
	}

	responses, err := hand.Service.PostClientData(url, dealDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		// "Message": "Sucessfully Merged the Deals into my account",
		"Data": responses,
	})
}
