package repository

import (
	"allcaps/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	ProductionKey  = "ggVZQsaslSsuVbmfxNUM32J2YPWLAod57_GcDmKS06r7YQfCPChi5A"
	AccountId      = "0b71949e-b245-4360-b55e-af088bada6c8"
	ClientToken    = "Iy65i25DgmwFxeT_a1Ho8ZgAIcnshnU4E9_fKsnKPOsQIl8QaJWGQg"
	MyAccountToken = "f3bCr4SAWwRw6bI9yHhgmgOfTlGLfPdDNnDpfMWIGFqA3ZIbcbj9Bg"
)

type IRepoInterface interface {
	GetClientData(url string) (*models.ApiResponse, error)
	PostClientData(url string, datas []map[string]interface{}) (*models.HubSpotDeal, error)
}

type repo struct {
	http.Client
}

func InitRepoClient(connection *http.Client) IRepoInterface {
	return &repo{*connection}
}

func (client *repo) GetClientData(url string) (*models.ApiResponse, error) {
	var opportunityDetails models.ApiResponse

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ProductionKey)
	req.Header.Set("X-Account-Token", ClientToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println("Response", string(body))

	if err := json.Unmarshal(body, &opportunityDetails); err != nil {
		return nil, err
	}
	// fmt.Println("All Data", opportunityDetails)

	return &opportunityDetails, nil

}

func TransformDataFormat(input []map[string]interface{}) []models.HubSpotDeal {
	var deals []models.HubSpotDeal

	fmt.Println("-----", deals)
	fmt.Println("")
	for _, item := range input {
		deal := models.HubSpotDeal{
			Properties: map[string]interface{}{
				"amount":             item["amount"],
				"dealname":           item["name"],
				"dealstage":          item["status"],
				"closedate":          item["closedate"].(time.Time).Format(time.RFC3339),
				"createdate":         item["created_at"].(time.Time).Format(time.RFC3339),
				"last_activity_date": item["last_activity_at"].(time.Time).Format(time.RFC3339),
				"hubspot_owner_id":   item["id"],
				"pipeline":           "default",
				"dealtype":           "newbusiness",
				"deal_status":        "pending",
			},
		}
		deals = append(deals, deal)
	}
	// for _, item := range input {
	// 	deal := models.HubSpotDeal{
	// 		Propert: map[string]interface{}{
	// 			"amount":             item["amount"],
	// 			"dealname":           item["name"],
	// 			"dealstage":          item["status"],
	// 			"closedate":          item["closedate"].(time.Time).Format(time.RFC3339),
	// 			"createdate":         item["created_at"].(time.Time).Format(time.RFC3339),
	// 			"last_activity_date": item["last_activity_at"].(time.Time).Format(time.RFC3339),
	// 			"hubspot_owner_id":   item["id"],
	// 			"pipeline":           "default",
	// 			"dealtype":           "newbusiness",
	// 			"deal_status":        "pending",
	// 		},
	// 	}
	// 	deals = append(deals, deal)
	// }

	return deals
}

func (client *repo) PostClientData(url string, data []map[string]interface{}) (*models.HubSpotDeal, error) {
	var response models.HubSpotDeal

	fmt.Println("")

	fmt.Println("Processing data", data)
	transformedData := TransformDataFormat(data)

	fmt.Println("")
	fmt.Println("Transfomeddata", transformedData)

	value, err := json.Marshal(transformedData)
	if err != nil {
		fmt.Println("Error occurred while marshalling data:", err)
		return nil, err
	}

	fmt.Println("")
	fmt.Println("Value", string(value))
	fmt.Println("")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(value))
	if err != nil {
		fmt.Println("Error occurred while creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ProductionKey)
	req.Header.Set("X-Account-Token", MyAccountToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error occurred while sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("")
	// fmt.Println("Response", resp)
	fmt.Println("")

	// if http.StatusBadRequest == resp.StatusCode {
	// 	return nil, fmt.Errorf("faced a issue in request,check it")
	// }

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error occurred while reading response:", err)
		return nil, err
	}

	fmt.Println("Response Body:", string(body))

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
