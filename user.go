package latch

import (
	"bytes"
	"net/http"
)

const (
	SubscriptionUrl = "/api/" + Version + "/subscription"
	ApplicationUrl  = "/api/" + Version + "/application"
)

type LatchUser struct {
	credentials Credentials
}

func NewLatchUser(credentials Credentials) *LatchUser {
	return &LatchUser{credentials}
}

func (user *LatchUser) GetSubscription() *LatchResponse {

	return sendRequest(http.MethodGet, SubscriptionUrl, nil, nil, user.credentials)

}

func (user *LatchUser) GetApplications() *LatchResponse {

	return sendRequest(http.MethodGet, ApplicationUrl, nil, nil, user.credentials)

}

func (user *LatchUser) CreateApplication(name, twoFactor, lockOnRequest, contactPhone, contactEmail string) *LatchResponse {

	parameters := map[string]string{
		NameParameter:          name,
		TwoFactorParameter:     twoFactor,
		LockOnRequestParameter: lockOnRequest,
		ContactPhone:           contactPhone,
		ContactEmail:           contactEmail,
	}

	return sendRequest(http.MethodPut, ApplicationUrl, nil, parameters, user.credentials)

}

func (user *LatchUser) DeleteApplication(applicationId string) *LatchResponse {

	var urlPath bytes.Buffer

	urlPath.WriteString(ApplicationUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(applicationId)

	return sendRequest(http.MethodDelete, urlPath.String(), nil, nil, user.credentials)

}

func (user *LatchUser) UpdateApplication(applicationId, name, twoFactor, lockOnRequest, contactPhone, contactEmail string) *LatchResponse {

	var urlPath bytes.Buffer

	urlPath.WriteString(ApplicationUrl)
	urlPath.WriteString("/")
	urlPath.WriteString(applicationId)

	parameters := map[string]string{
		NameParameter:          name,
		TwoFactorParameter:     twoFactor,
		LockOnRequestParameter: lockOnRequest,
		ContactPhone:           contactPhone,
		ContactEmail:           contactEmail,
	}

	return sendRequest(http.MethodPost, urlPath.String(), nil, parameters, user.credentials)

}
