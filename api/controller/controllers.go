package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	ProductionKey = "ggVZQsaslSsuVbmfxNUM32J2YPWLAod57_GcDmKS06r7YQfCPChi5A"
	AccountId     = "0b71949e-b245-4360-b55e-af088bada6c8"
	AccountToken  = "LxIkk0zvQXOCKrN8bIxhJPn6RBkhdUdRjCapiLFucqsvdiyJzJsm5w"
)

func SyncAccount() {
	// data := map[string]string{
	// 	"X-Account-Token": APItoken,
	// }
	url := fmt.Sprintf("https://app.merge.dev/linked-accounts/account/sync/%s", AccountId)
	fmt.Println("url", url)

	payload := map[string]interface{}{
		"linked_account_id": AccountId,
		"X-Account-Token":   AccountToken,
		"force_refresh":     true,
	}

	fmt.Println("Payload", payload)

	value, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error occured")
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(value))
	if err != nil {
		fmt.Println("Error occured")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ProductionKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
