package latch

import (
	"bytes"
	"net/http"
	"strconv"
	"time"
)

const (
	CheckStatusUrl = "/api/" + Version + "/status"
	PairTokenUrl   = "/api/" + Version + "/pair"
	PairIdUrl      = "/api/" + Version + "/pairWithId"
	UnpairUrl      = "/api/" + Version + "/unpair"
	LockUrl        = "/api/" + Version + "/lock"
	UnlockUrl      = "/api/" + Version + "/unlock"
	HistoryUrl     = "/api/" + Version + "/history"
	OperationUrl   = "/api/" + Version + "/operation"
	InstanceUrl    = "/api/" + Version + "/instance"
)

type LatchApplication struct {
	credentials Credentials
}

func NewLatchApplication(credentials Credentials) *LatchApplication {
	return &LatchApplication{credentials}
}

func (application *LatchApplication) PairWithId(accountId string) *Response {

	var urlPath bytes.Buffer

	urlPath.WriteString(PairIdUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) PairWithToken(token string) *Response {

	var urlPath bytes.Buffer

	urlPath.WriteString(PairTokenUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(token)

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) Unpair(accountId string) *Response {

	var urlPath bytes.Buffer

	urlPath.WriteString(UnpairUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)
	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) Status(accountId string, nootp, silent bool) *Response {

	var urlPath bytes.Buffer

	urlPath.WriteString(CheckStatusUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)

	if nootp {
		urlPath.WriteString("/nootp")
	}

	if silent {
		urlPath.WriteString("/silent")
	}

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) OperationStatus(accountId, operationId string, nootp, silent bool) *Response {

	var urlPath bytes.Buffer

	urlPath.WriteString(CheckStatusUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)
	urlPath.WriteString("/op/")
	urlPath.WriteString(operationId)

	if nootp {
		urlPath.WriteString("/nootp")
	}

	if silent {
		urlPath.WriteString("/silent")
	}

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) Lock(accountId string) *Response {

	var urlPath bytes.Buffer

	urlPath.WriteString(LockUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)
	return sendRequest(http.MethodPost, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) Unlock(accountId string) *Response {

	var urlPath bytes.Buffer

	urlPath.WriteString(UnlockUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)

	return sendRequest(http.MethodPost, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) History(accountId string, from, to time.Time) *Response {

	var urlPath bytes.Buffer
	fromMilis := from.UnixNano() / int64(time.Millisecond)
	toMilis := to.UnixNano() / int64(time.Millisecond)

	urlPath.WriteString(HistoryUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)
	urlPath.WriteString("/")
	urlPath.WriteString(strconv.FormatInt(fromMilis, 10))
	urlPath.WriteString("/")
	urlPath.WriteString(strconv.FormatInt(toMilis, 10))

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

/*
 * Operations
 */

func (application *LatchApplication) CreateOperation(parentId, name, twoFactor, lockOnRequest string) *Response {

	parameters := map[string]string{
		"parentId":        parentId,
		"name":            name,
		"two_factor":      twoFactor,
		"lock_on_request": lockOnRequest,
	}

	return sendRequest(http.MethodPut, OperationUrl, nil, parameters, application.credentials)

}

func (application *LatchApplication) UpdateOperation(operationId, name, twoFactor, lockOnRequest string) *Response {

	var urlPath bytes.Buffer
	urlPath.WriteString(OperationUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(operationId)

	parameters := map[string]string{
		"name":            name,
		"two_factor":      twoFactor,
		"lock_on_request": lockOnRequest,
	}

	return sendRequest(http.MethodPost, urlPath.String(), nil, parameters, application.credentials)

}

func (application *LatchApplication) DeleteOperation(operationId string) *Response {

	var urlPath bytes.Buffer
	urlPath.WriteString(OperationUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(operationId)

	return sendRequest(http.MethodDelete, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) GetOperations() *Response {

	var urlPath bytes.Buffer
	urlPath.WriteString(OperationUrl)

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) GetOperation(operationId string) *Response {

	var urlPath bytes.Buffer
	urlPath.WriteString(OperationUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(operationId)

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

/*
 * Instances
 */

func (application *LatchApplication) GetInstances(accountId string) *Response {

	var urlPath bytes.Buffer
	urlPath.WriteString(InstanceUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) CreateInstance(accountId, operationId, name string) *Response {

	var urlPath bytes.Buffer
	urlPath.WriteString(InstanceUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)

	if len(operationId) > 0 {
		urlPath.WriteString("/op/")
		urlPath.WriteString(operationId)
	}

	parameters := map[string]string{
		"instances": name,
	}

	return sendRequest(http.MethodPut, urlPath.String(), nil, parameters, application.credentials)

}

func (application *LatchApplication) UpdateInstance(instanceId, accountId, operationId, name, two_factor, lock_on_request string) *Response {

	var urlPath bytes.Buffer
	urlPath.WriteString(InstanceUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)
	if len(operationId) > 0 {
		urlPath.WriteString("/op/")
		urlPath.WriteString(operationId)
	}
	urlPath.WriteString("/i/")
	urlPath.WriteString(instanceId)

	parameters := map[string]string{
		"name":            name,
		"two_factor":      two_factor,
		"lock_on_request": lock_on_request,
	}

	return sendRequest(http.MethodPost, urlPath.String(), nil, parameters, application.credentials)

}

func (application *LatchApplication) DeleteInstance(instanceId, accountId, operationId string) *Response {

	var urlPath bytes.Buffer
	urlPath.WriteString(InstanceUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)
	if len(operationId) > 0 {
		urlPath.WriteString("/op/")
		urlPath.WriteString(operationId)
	}
	urlPath.WriteString("/i/")
	urlPath.WriteString(instanceId)

	return sendRequest(http.MethodDelete, urlPath.String(), nil, nil, application.credentials)

}
