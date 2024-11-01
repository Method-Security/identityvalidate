package azure

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	identityvalidate "github.com/Method-Security/identityvalidate/generated/go"
)

type OWALibrary struct{}

func (lib *OWALibrary) ModuleRun(config *identityvalidate.PortalConfig) (*identityvalidate.Trigger, []string) {
	target := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", config.TenantId)
	trigger := identityvalidate.Trigger{Target: target}
	errors := []string{}

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	if config.AgentHeader != "" {
		headers["User-Agent"] = config.AgentHeader
	}

	formData := url.Values{}
	formData.Set("grant_type", "password")
	formData.Set("client_id", config.ClientId)
	formData.Set("scope", "https://graph.microsoft.com/.default")
	formData.Set("username", config.Username)
	formData.Set("password", config.Password)
	requestBody := formData.Encode()

	request := identityvalidate.GeneralRequestInfo{
		Method:  identityvalidate.HttpMethodPost,
		Url:     target,
		Headers: headers,
		Body:    &requestBody,
	}

	req, err := http.NewRequest("POST", target, strings.NewReader(requestBody))
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

	resp, err := client.Do(req)
	if err != nil {
		errorMessage := err.Error()
		errors = append(errors, errorMessage)
		response := identityvalidate.GeneralResponseInfo{Error: &errorMessage}
		GeneralAttemptInfo := identityvalidate.GeneralAttemptInfo{Request: &request, Response: &response}
		trigger.AttemptInfo = identityvalidate.NewAttemptInfoUnionFromGeneralAttempt(&GeneralAttemptInfo)
		return &trigger, errors
	}

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

	err = resp.Body.Close()
	if err != nil {
		errors = append(errors, err.Error())
		return &trigger, errors
	}

	response := identityvalidate.GeneralResponseInfo{
		StatusCode: resp.StatusCode,
		Body:       &bodyStr,
	}
	GeneralAttemptInfo := identityvalidate.GeneralAttemptInfo{
		Request:  &request,
		Response: &response,
	}
	trigger.AttemptInfo = identityvalidate.NewAttemptInfoUnionFromGeneralAttempt(&GeneralAttemptInfo)

	return &trigger, errors
}
