// This file was auto-generated by Fern from our API Definition.

package identityvalidate

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/Method-Security/identityvalidate/generated/go/core"
	time "time"
)

type HttpMethod string

const (
	HttpMethodGet     HttpMethod = "GET"
	HttpMethodPost    HttpMethod = "POST"
	HttpMethodPut     HttpMethod = "PUT"
	HttpMethodDelete  HttpMethod = "DELETE"
	HttpMethodPatch   HttpMethod = "PATCH"
	HttpMethodOptions HttpMethod = "OPTIONS"
	HttpMethodHead    HttpMethod = "HEAD"
	HttpMethodConnect HttpMethod = "CONNECT"
	HttpMethodTrace   HttpMethod = "TRACE"
)

func NewHttpMethodFromString(s string) (HttpMethod, error) {
	switch s {
	case "GET":
		return HttpMethodGet, nil
	case "POST":
		return HttpMethodPost, nil
	case "PUT":
		return HttpMethodPut, nil
	case "DELETE":
		return HttpMethodDelete, nil
	case "PATCH":
		return HttpMethodPatch, nil
	case "OPTIONS":
		return HttpMethodOptions, nil
	case "HEAD":
		return HttpMethodHead, nil
	case "CONNECT":
		return HttpMethodConnect, nil
	case "TRACE":
		return HttpMethodTrace, nil
	}
	var t HttpMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (h HttpMethod) Ptr() *HttpMethod {
	return &h
}

type AttemptInfoUnion struct {
	Type           string
	GeneralAttempt *GeneralAttemptInfo
}

func NewAttemptInfoUnionFromGeneralAttempt(value *GeneralAttemptInfo) *AttemptInfoUnion {
	return &AttemptInfoUnion{Type: "GeneralAttempt", GeneralAttempt: value}
}

func (a *AttemptInfoUnion) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", a)
	}
	switch unmarshaler.Type {
	case "GeneralAttempt":
		value := new(GeneralAttemptInfo)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.GeneralAttempt = value
	}
	return nil
}

func (a AttemptInfoUnion) MarshalJSON() ([]byte, error) {
	switch a.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", a.Type, a)
	case "GeneralAttempt":
		return core.MarshalJSONWithExtraProperty(a.GeneralAttempt, "type", "GeneralAttempt")
	}
}

type AttemptInfoUnionVisitor interface {
	VisitGeneralAttempt(*GeneralAttemptInfo) error
}

func (a *AttemptInfoUnion) Accept(visitor AttemptInfoUnionVisitor) error {
	switch a.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", a.Type, a)
	case "GeneralAttempt":
		return visitor.VisitGeneralAttempt(a.GeneralAttempt)
	}
}

type GeneralAttemptInfo struct {
	Timestamp time.Time            `json:"timestamp" url:"timestamp"`
	Request   *GeneralRequestInfo  `json:"request,omitempty" url:"request,omitempty"`
	Response  *GeneralResponseInfo `json:"response,omitempty" url:"response,omitempty"`

	extraProperties map[string]interface{}
}

func (g *GeneralAttemptInfo) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GeneralAttemptInfo) UnmarshalJSON(data []byte) error {
	type embed GeneralAttemptInfo
	var unmarshaler = struct {
		embed
		Timestamp *core.DateTime `json:"timestamp"`
	}{
		embed: embed(*g),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*g = GeneralAttemptInfo(unmarshaler.embed)
	g.Timestamp = unmarshaler.Timestamp.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	return nil
}

func (g *GeneralAttemptInfo) MarshalJSON() ([]byte, error) {
	type embed GeneralAttemptInfo
	var marshaler = struct {
		embed
		Timestamp *core.DateTime `json:"timestamp"`
	}{
		embed:     embed(*g),
		Timestamp: core.NewDateTime(g.Timestamp),
	}
	return json.Marshal(marshaler)
}

func (g *GeneralAttemptInfo) String() string {
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type GeneralRequestInfo struct {
	Method  HttpMethod        `json:"method" url:"method"`
	Url     string            `json:"url" url:"url"`
	Params  map[string]string `json:"params,omitempty" url:"params,omitempty"`
	Headers map[string]string `json:"headers,omitempty" url:"headers,omitempty"`
	Body    *string           `json:"body,omitempty" url:"body,omitempty"`

	extraProperties map[string]interface{}
}

func (g *GeneralRequestInfo) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GeneralRequestInfo) UnmarshalJSON(data []byte) error {
	type unmarshaler GeneralRequestInfo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GeneralRequestInfo(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	return nil
}

func (g *GeneralRequestInfo) String() string {
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type GeneralResponseInfo struct {
	StatusCode int               `json:"statusCode" url:"statusCode"`
	Body       *string           `json:"body,omitempty" url:"body,omitempty"`
	Headers    map[string]string `json:"headers,omitempty" url:"headers,omitempty"`
	Error      *string           `json:"error,omitempty" url:"error,omitempty"`

	extraProperties map[string]interface{}
}

func (g *GeneralResponseInfo) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GeneralResponseInfo) UnmarshalJSON(data []byte) error {
	type unmarshaler GeneralResponseInfo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GeneralResponseInfo(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	return nil
}

func (g *GeneralResponseInfo) String() string {
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type ModuleName string

const (
	ModuleNameOwaLogin ModuleName = "OWA_LOGIN"
)

func NewModuleNameFromString(s string) (ModuleName, error) {
	switch s {
	case "OWA_LOGIN":
		return ModuleNameOwaLogin, nil
	}
	var t ModuleName
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (m ModuleName) Ptr() *ModuleName {
	return &m
}

type PortalConfig struct {
	PortalType  PortalType `json:"portalType" url:"portalType"`
	ModuleName  ModuleName `json:"moduleName" url:"moduleName"`
	Attempts    int        `json:"attempts" url:"attempts"`
	Username    string     `json:"username" url:"username"`
	Password    string     `json:"password" url:"password"`
	AgentHeader string     `json:"agentHeader" url:"agentHeader"`
	Interval    int        `json:"interval" url:"interval"`
	Timeout     int        `json:"timeout" url:"timeout"`

	extraProperties map[string]interface{}
}

func (p *PortalConfig) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PortalConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler PortalConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PortalConfig(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	return nil
}

func (p *PortalConfig) String() string {
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PortalReport struct {
	PortalType PortalType `json:"portalType" url:"portalType"`
	ModuleName ModuleName `json:"moduleName" url:"moduleName"`
	Triggers   []*Trigger `json:"triggers,omitempty" url:"triggers,omitempty"`
	Errors     []string   `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
}

func (p *PortalReport) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PortalReport) UnmarshalJSON(data []byte) error {
	type unmarshaler PortalReport
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PortalReport(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	return nil
}

func (p *PortalReport) String() string {
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PortalType string

const (
	PortalTypeAzure PortalType = "AZURE"
)

func NewPortalTypeFromString(s string) (PortalType, error) {
	switch s {
	case "AZURE":
		return PortalTypeAzure, nil
	}
	var t PortalType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PortalType) Ptr() *PortalType {
	return &p
}

type ResponseUnion struct {
	Type            string
	GeneralResponse *GeneralResponseInfo
}

func NewResponseUnionFromGeneralResponse(value *GeneralResponseInfo) *ResponseUnion {
	return &ResponseUnion{Type: "GeneralResponse", GeneralResponse: value}
}

func (r *ResponseUnion) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	r.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", r)
	}
	switch unmarshaler.Type {
	case "GeneralResponse":
		value := new(GeneralResponseInfo)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.GeneralResponse = value
	}
	return nil
}

func (r ResponseUnion) MarshalJSON() ([]byte, error) {
	switch r.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", r.Type, r)
	case "GeneralResponse":
		return core.MarshalJSONWithExtraProperty(r.GeneralResponse, "type", "GeneralResponse")
	}
}

type ResponseUnionVisitor interface {
	VisitGeneralResponse(*GeneralResponseInfo) error
}

func (r *ResponseUnion) Accept(visitor ResponseUnionVisitor) error {
	switch r.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", r.Type, r)
	case "GeneralResponse":
		return visitor.VisitGeneralResponse(r.GeneralResponse)
	}
}

type Trigger struct {
	Target      string            `json:"target" url:"target"`
	AttemptInfo *AttemptInfoUnion `json:"AttemptInfo,omitempty" url:"AttemptInfo,omitempty"`

	extraProperties map[string]interface{}
}

func (t *Trigger) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *Trigger) UnmarshalJSON(data []byte) error {
	type unmarshaler Trigger
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = Trigger(value)

	extraProperties, err := core.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties

	return nil
}

func (t *Trigger) String() string {
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}
