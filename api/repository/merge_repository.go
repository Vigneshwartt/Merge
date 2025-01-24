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
	PostClientData(url string, datas []map[string]interface{}) (*models.DealsPayload, error)
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

	return deals
}

func (client *repo) PostClientData(url string, data []map[string]interface{}) (*models.DealsPayload, error) {
	var response models.DealsPayload

	fmt.Println("Processing data", data)
	// transformedData := TransformDataFormat(data)

	fmt.Println("")
	// fmt.Println("Transfomeddata", transformedData)

	value, err := json.Marshal(data)
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
	fmt.Println("Response", resp)
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

// [map[amount:1000 closedate:2025-01-30 12:52:25.81 +0000 UTC created_at:2025-01-23 13:53:43.96251 +0000 UTC description:<nil> id:0a6e70c4-9c98-4314-ae05-89008933dc71 last_activity_at:2025-01-30 12:52:25.81 +0000 UTC modified_at:2025-01-23 13:53:43.962526 +0000 UTC name:Sales Deal remote_created_at:2025-01-23 12:52:45.224 +0000 UTC remote_id:1000 remote_was_deleted:false status:WON] map[amount:520000 closedate:2025-01-31 09:59:35.82 +0000 UTC created_at:2025-01-22 11:50:15.07677 +0000 UTC description:<nil> id:1232d4ce-f6a3-46ac-b868-e5ea8bf013e6 last_activity_at:2025-01-31 09:59:35.82 +0000 UTC modified_at:2025-01-22 11:50:15.076775 +0000 UTC name:lotus remote_created_at:2025-01-22 09:59:52.282 +0000 UTC remote_id:520000 remote_was_deleted:false status:OPEN] map[amount:80000 closedate:2025-01-30 12:30:45 +0000 UTC created_at:2025-01-22 11:50:15.076653 +0000 UTC description:<nil> id:4ae24b35-3a42-4005-a740-8ca621cb6385 last_activity_at:2025-01-30 12:30:45 +0000 UTC modified_at:2025-01-22 11:50:15.07667 +0000 UTC name:muruga remote_created_at:2025-01-21 13:31:24.767 +0000 UTC remote_id:80000 remote_was_deleted:false status:OPEN]]
