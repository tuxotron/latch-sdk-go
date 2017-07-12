package latch

import (
	"bytes"
	"net/http"
	"strconv"
	"time"
)

// URLs
const (
	checkStatusUrl = "/api/" + Version + "/status"
	pairTokenUrl   = "/api/" + Version + "/pair"
	pairIdUrl      = "/api/" + Version + "/pairWithId"
	unpairUrl      = "/api/" + Version + "/unpair"
	lockUrl        = "/api/" + Version + "/lock"
	unlockUrl      = "/api/" + Version + "/unlock"
	historyUrl     = "/api/" + Version + "/history"
	operationUrl   = "/api/" + Version + "/operation"
	instanceUrl    = "/api/" + Version + "/instance"
)

type LatchApplication struct {
	credentials Credentials
}

func NewLatchApplication(credentials Credentials) *LatchApplication {
	return &LatchApplication{credentials}
}

func (application *LatchApplication) PairWithId(accountId string) *LatchResponse {

	var urlPath bytes.Buffer

	urlPath.WriteString(pairIdUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) PairWithToken(token string) *LatchResponse {

	var urlPath bytes.Buffer

	urlPath.WriteString(pairTokenUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(token)

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) Unpair(accountId string) *LatchResponse {

	var urlPath bytes.Buffer

	urlPath.WriteString(unpairUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)
	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) Status(accountId string, nootp, silent bool) *LatchResponse {

	var urlPath bytes.Buffer

	urlPath.WriteString(checkStatusUrl)
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

func (application *LatchApplication) OperationStatus(accountId, operationId string, nootp, silent bool) *LatchResponse {

	var urlPath bytes.Buffer

	urlPath.WriteString(checkStatusUrl)
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

func (application *LatchApplication) Lock(accountId string) *LatchResponse {

	var urlPath bytes.Buffer

	urlPath.WriteString(lockUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)
	return sendRequest(http.MethodPost, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) Unlock(accountId string) *LatchResponse {

	var urlPath bytes.Buffer

	urlPath.WriteString(unlockUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)

	return sendRequest(http.MethodPost, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) History(accountId string, from, to time.Time) *LatchResponse {

	var urlPath bytes.Buffer
	fromMilis := from.UnixNano() / int64(time.Millisecond)
	toMilis := to.UnixNano() / int64(time.Millisecond)

	urlPath.WriteString(historyUrl)
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

func (application *LatchApplication) CreateOperation(parentId, name, twoFactor, lockOnRequest string) *LatchResponse {

	parameters := map[string]string{
		ParentIdParameter:      parentId,
		NameParameter:          name,
		TwoFactorParameter:     twoFactor,
		LockOnRequestParameter: lockOnRequest,
	}

	return sendRequest(http.MethodPut, operationUrl, nil, parameters, application.credentials)

}

func (application *LatchApplication) UpdateOperation(operationId, name, twoFactor, lockOnRequest string) *LatchResponse {

	var urlPath bytes.Buffer
	urlPath.WriteString(operationUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(operationId)

	parameters := map[string]string{
		NameParameter:          name,
		TwoFactorParameter:     twoFactor,
		LockOnRequestParameter: lockOnRequest,
	}

	return sendRequest(http.MethodPost, urlPath.String(), nil, parameters, application.credentials)

}

func (application *LatchApplication) DeleteOperation(operationId string) *LatchResponse {

	var urlPath bytes.Buffer
	urlPath.WriteString(operationUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(operationId)

	return sendRequest(http.MethodDelete, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) GetOperations() *LatchResponse {

	var urlPath bytes.Buffer
	urlPath.WriteString(operationUrl)

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) GetOperation(operationId string) *LatchResponse {

	var urlPath bytes.Buffer
	urlPath.WriteString(operationUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(operationId)

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

/*
 * Instances
 */

func (application *LatchApplication) GetInstances(accountId string) *LatchResponse {

	var urlPath bytes.Buffer
	urlPath.WriteString(instanceUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)

	return sendRequest(http.MethodGet, urlPath.String(), nil, nil, application.credentials)

}

func (application *LatchApplication) CreateInstance(accountId, operationId, name string) *LatchResponse {

	var urlPath bytes.Buffer
	urlPath.WriteString(instanceUrl)
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

func (application *LatchApplication) UpdateInstance(instanceId, accountId, operationId, name, two_factor, lock_on_request string) *LatchResponse {

	var urlPath bytes.Buffer
	urlPath.WriteString(instanceUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(accountId)
	if len(operationId) > 0 {
		urlPath.WriteString("/op/")
		urlPath.WriteString(operationId)
	}
	urlPath.WriteString("/i/")
	urlPath.WriteString(instanceId)

	parameters := map[string]string{
		NameParameter:          name,
		TwoFactorParameter:     two_factor,
		LockOnRequestParameter: lock_on_request,
	}

	return sendRequest(http.MethodPost, urlPath.String(), nil, parameters, application.credentials)

}

func (application *LatchApplication) DeleteInstance(instanceId, accountId, operationId string) *LatchResponse {

	var urlPath bytes.Buffer
	urlPath.WriteString(instanceUrl)
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
