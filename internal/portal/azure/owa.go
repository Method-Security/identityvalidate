package azure

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	identityvalidate "github.com/Method-Security/identityvalidate/generated/go"
)

type OWALibrary struct{}

func (lib *OWALibrary) ModuleRun(config *identityvalidate.PortalConfig) (*identityvalidate.Trigger, []string) {
	// Initialize structs
	target := "https://login.microsoftonline.com/common/GetCredentialType?mkt=en-US"
	trigger := identityvalidate.Trigger{Target: target}
	errors := []string{}

	// Construct the request data
	headers := map[string]string{
		"Content-Type": "application/json",
		"Connection":   "close",
	}
	if config.AgentHeader != "" {
		headers["User-Agent"] = config.AgentHeader
	}
	reqBody := fmt.Sprintf(`{"username": "%s"}`, config.Username)
	request := identityvalidate.GeneralRequestInfo{Method: identityvalidate.HttpMethodPost, Url: target, Headers: headers, Body: &reqBody}

	// Construct the request
	req, err := http.NewRequest("POST", target, strings.NewReader(reqBody))
	if err != nil {
		return nil, []string{fmt.Sprintf("Error creating request: %s", err.Error())}
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		errorMessage := err.Error()
		errors = append(errors, errorMessage)
		response := identityvalidate.GeneralResponseInfo{Error: &errorMessage}
		GeneralAttemptInfo := identityvalidate.GeneralAttemptInfo{Request: &request, Response: &response}
		trigger.AttemptInfo = identityvalidate.NewAttemptInfoUnionFromGeneralAttempt(&GeneralAttemptInfo)
		return &trigger, errors
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorMessage := err.Error()
		errors = append(errors, errorMessage)
		response := identityvalidate.GeneralResponseInfo{Error: &errorMessage}
		GeneralAttemptInfo := identityvalidate.GeneralAttemptInfo{Request: &request, Response: &response}
		trigger.AttemptInfo = identityvalidate.NewAttemptInfoUnionFromGeneralAttempt(&GeneralAttemptInfo)
		return &trigger, errors
	}
	bodyStr := string(body)

	// Close resp
	err = resp.Body.Close()
	if err != nil {
		errors = append(errors, err.Error())
		return &trigger, errors
	}

	// Marshal data
	response := identityvalidate.GeneralResponseInfo{StatusCode: resp.StatusCode, Body: &bodyStr}
	GeneralAttemptInfo := identityvalidate.GeneralAttemptInfo{Request: &request, Response: &response}
	trigger.AttemptInfo = identityvalidate.NewAttemptInfoUnionFromGeneralAttempt(&GeneralAttemptInfo)
	return &trigger, errors
}
