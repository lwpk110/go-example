/*
 * Identity
 *
 * Welcome to the Identity HTTP API documentation! You will find documentation for all HTTP APIs here.
 *
 * API version: latest
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package identityclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type SettingsApi interface {

	/*
	 * GetSettingsFlowRequest # Get Settings Flow
	 * When accessing this endpoint through Identity' Public API you must ensure that either the Identity Session Cookie
or the Identity Session Token are set.

Depending on your configuration this endpoint might return a 403 error if the session has a lower Authenticator
Assurance Level (AAL) than is possible for the identity. This can happen if the identity has password + webauthn
credentials (which would result in AAL2) but the session has only AAL1. If this error occurs, ask the user
to sign in with the second factor or change the configuration.

You can access this endpoint without credentials when using Identity' Admin API.

If this endpoint is called via an AJAX request, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`session_inactive`: No Identity Session was found - sign in a user first.
`security_identity_mismatch`: The flow was interrupted with `session_refresh_required` but apparently some other
identity logged in instead.

More information can be found at
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SettingsApiApiGetSettingsFlowRequestRequest
	 */
	GetSettingsFlowRequest(ctx context.Context) SettingsApiApiGetSettingsFlowRequestRequest

	/*
	 * GetSettingsFlowRequestExecute executes the request
	 * @return SettingsFlow
	 */
	GetSettingsFlowRequestExecute(r SettingsApiApiGetSettingsFlowRequestRequest) (*SettingsFlow, *http.Response, error)

	/*
	 * InitBrowserSettingsFlowRequest # Initialize Settings Flow for Browsers
	 * This endpoint initializes a browser-based user settings flow. Once initialized, the browser will be redirected to
`selfservice.flows.settings.ui_url` with the flow ID set as the query parameter `?flow=`. If no valid
Identity Session Cookie is included in the request, a login flow will be initialized.

If this endpoint is opened as a link in the browser, it will be redirected to
`selfservice.flows.settings.ui_url` with the flow ID set as the query parameter `?flow=`. If no valid user session
was set, the browser will be redirected to the login endpoint.

If this endpoint is called via an AJAX request, the response contains the settings flow without any redirects
or a 401 forbidden error if no valid session was set.

Depending on your configuration this endpoint might return a 403 error if the session has a lower Authenticator
Assurance Level (AAL) than is possible for the identity. This can happen if the identity has password + webauthn
credentials (which would result in AAL2) but the session has only AAL1. If this error occurs, ask the user
to sign in with the second factor (happens automatically for server-side browser flows) or change the configuration.

If this endpoint is called via an AJAX request, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`session_inactive`: No Session was found - sign in a user first.
`security_identity_mismatch`: The requested `?return_to` address is not allowed to be used. Adjust this in the configuration!

This endpoint is NOT INTENDED for clients that do not have a browser (Chrome, Firefox, ...) as cookies are needed.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SettingsApiApiInitBrowserSettingsFlowRequestRequest
	 */
	InitBrowserSettingsFlowRequest(ctx context.Context) SettingsApiApiInitBrowserSettingsFlowRequestRequest

	/*
	 * InitBrowserSettingsFlowRequestExecute executes the request
	 * @return SettingsFlow
	 */
	InitBrowserSettingsFlowRequestExecute(r SettingsApiApiInitBrowserSettingsFlowRequestRequest) (*SettingsFlow, *http.Response, error)

	/*
	 * SubmitSettingsFlowRequest # Complete Settings Flow
	 * Use this endpoint to complete a settings flow by sending an identity's updated password. This endpoint
behaves differently for API and browser flows.

API-initiated flows expect `application/json` to be sent in the body and respond with
HTTP 200 and an application/json body with the session token on success;
HTTP 303 redirect to a fresh settings flow if the original flow expired with the appropriate error messages set;
HTTP 400 on form validation errors.
HTTP 401 when the endpoint is called without a valid session token.
HTTP 403 when `selfservice.flows.settings.privileged_session_max_age` was reached or the session's AAL is too low.
Implies that the user needs to re-authenticate.

Browser flows without HTTP Header `Accept` or with `Accept: text/*` respond with
a HTTP 303 redirect to the post/after settings URL or the `return_to` value if it was set and if the flow succeeded;
a HTTP 303 redirect to the Settings UI URL with the flow ID containing the validation errors otherwise.
a HTTP 303 redirect to the login endpoint when `selfservice.flows.settings.privileged_session_max_age` was reached or the session's AAL is too low.

Browser flows with HTTP Header `Accept: application/json` respond with
HTTP 200 and a application/json body with the signed in identity and a `Set-Cookie` header on success;
HTTP 303 redirect to a fresh login flow if the original flow expired with the appropriate error messages set;
HTTP 401 when the endpoint is called without a valid session cookie.
HTTP 403 when the page is accessed without a session cookie or the session's AAL is too low.
HTTP 400 on form validation errors.

Depending on your configuration this endpoint might return a 403 error if the session has a lower Authenticator
Assurance Level (AAL) than is possible for the identity. This can happen if the identity has password + webauthn
credentials (which would result in AAL2) but the session has only AAL1. If this error occurs, ask the user
to sign in with the second factor (happens automatically for server-side browser flows) or change the configuration.

If this endpoint is called with a `Accept: application/json` HTTP header, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`session_refresh_required`: The identity requested to change something that needs a privileged session. Redirect
the identity to the login init endpoint with query parameters `?refresh=true&return_to=<the-current-browser-url>`,
or initiate a refresh login flow otherwise.
`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`session_inactive`: No Identity Session was found - sign in a user first.
`security_identity_mismatch`: The flow was interrupted with `session_refresh_required` but apparently some other
identity logged in instead.
`security_identity_mismatch`: The requested `?return_to` address is not allowed to be used. Adjust this in the configuration!
`browser_location_change_required`: Usually sent when an AJAX request indicates that the browser needs to open a specific URL.
Most likely used in Social Sign In flows.

More information can be found at [Identity User Settings & Profile Management Documentation](../self-service/flows/user-settings).
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return SettingsApiApiSubmitSettingsFlowRequestRequest
	 */
	SubmitSettingsFlowRequest(ctx context.Context) SettingsApiApiSubmitSettingsFlowRequestRequest

	/*
	 * SubmitSettingsFlowRequestExecute executes the request
	 * @return SettingsFlow
	 */
	SubmitSettingsFlowRequestExecute(r SettingsApiApiSubmitSettingsFlowRequestRequest) (*SettingsFlow, *http.Response, error)
}

// SettingsApiService SettingsApi service
type SettingsApiService service

type SettingsApiApiGetSettingsFlowRequestRequest struct {
	ctx context.Context
	ApiService SettingsApi
	id *string
	xSessionToken *string
	cookie *string
}

func (r SettingsApiApiGetSettingsFlowRequestRequest) Id(id string) SettingsApiApiGetSettingsFlowRequestRequest {
	r.id = &id
	return r
}
func (r SettingsApiApiGetSettingsFlowRequestRequest) XSessionToken(xSessionToken string) SettingsApiApiGetSettingsFlowRequestRequest {
	r.xSessionToken = &xSessionToken
	return r
}
func (r SettingsApiApiGetSettingsFlowRequestRequest) Cookie(cookie string) SettingsApiApiGetSettingsFlowRequestRequest {
	r.cookie = &cookie
	return r
}

func (r SettingsApiApiGetSettingsFlowRequestRequest) Execute() (*SettingsFlow, *http.Response, error) {
	return r.ApiService.GetSettingsFlowRequestExecute(r)
}

/*
 * GetSettingsFlowRequest # Get Settings Flow
 * When accessing this endpoint through Identity' Public API you must ensure that either the Identity Session Cookie
or the Identity Session Token are set.

Depending on your configuration this endpoint might return a 403 error if the session has a lower Authenticator
Assurance Level (AAL) than is possible for the identity. This can happen if the identity has password + webauthn
credentials (which would result in AAL2) but the session has only AAL1. If this error occurs, ask the user
to sign in with the second factor or change the configuration.

You can access this endpoint without credentials when using Identity' Admin API.

If this endpoint is called via an AJAX request, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`session_inactive`: No Identity Session was found - sign in a user first.
`security_identity_mismatch`: The flow was interrupted with `session_refresh_required` but apparently some other
identity logged in instead.

More information can be found at
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return SettingsApiApiGetSettingsFlowRequestRequest
 */
func (a *SettingsApiService) GetSettingsFlowRequest(ctx context.Context) SettingsApiApiGetSettingsFlowRequestRequest {
	return SettingsApiApiGetSettingsFlowRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SettingsFlow
 */
func (a *SettingsApiService) GetSettingsFlowRequestExecute(r SettingsApiApiGetSettingsFlowRequestRequest) (*SettingsFlow, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *SettingsFlow
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsApiService.GetSettingsFlowRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/self-service/settings/flows"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id == nil {
		return localVarReturnValue, nil, reportError("id is required and must be specified")
	}

	localVarQueryParams.Add("id", parameterToString(*r.id, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xSessionToken != nil {
		localVarHeaderParams["X-Session-Token"] = parameterToString(*r.xSessionToken, "")
	}
	if r.cookie != nil {
		localVarHeaderParams["Cookie"] = parameterToString(*r.cookie, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 410 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SettingsApiApiInitBrowserSettingsFlowRequestRequest struct {
	ctx context.Context
	ApiService SettingsApi
	returnTo *string
}

func (r SettingsApiApiInitBrowserSettingsFlowRequestRequest) ReturnTo(returnTo string) SettingsApiApiInitBrowserSettingsFlowRequestRequest {
	r.returnTo = &returnTo
	return r
}

func (r SettingsApiApiInitBrowserSettingsFlowRequestRequest) Execute() (*SettingsFlow, *http.Response, error) {
	return r.ApiService.InitBrowserSettingsFlowRequestExecute(r)
}

/*
 * InitBrowserSettingsFlowRequest # Initialize Settings Flow for Browsers
 * This endpoint initializes a browser-based user settings flow. Once initialized, the browser will be redirected to
`selfservice.flows.settings.ui_url` with the flow ID set as the query parameter `?flow=`. If no valid
Identity Session Cookie is included in the request, a login flow will be initialized.

If this endpoint is opened as a link in the browser, it will be redirected to
`selfservice.flows.settings.ui_url` with the flow ID set as the query parameter `?flow=`. If no valid user session
was set, the browser will be redirected to the login endpoint.

If this endpoint is called via an AJAX request, the response contains the settings flow without any redirects
or a 401 forbidden error if no valid session was set.

Depending on your configuration this endpoint might return a 403 error if the session has a lower Authenticator
Assurance Level (AAL) than is possible for the identity. This can happen if the identity has password + webauthn
credentials (which would result in AAL2) but the session has only AAL1. If this error occurs, ask the user
to sign in with the second factor (happens automatically for server-side browser flows) or change the configuration.

If this endpoint is called via an AJAX request, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`session_inactive`: No Session was found - sign in a user first.
`security_identity_mismatch`: The requested `?return_to` address is not allowed to be used. Adjust this in the configuration!

This endpoint is NOT INTENDED for clients that do not have a browser (Chrome, Firefox, ...) as cookies are needed.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return SettingsApiApiInitBrowserSettingsFlowRequestRequest
 */
func (a *SettingsApiService) InitBrowserSettingsFlowRequest(ctx context.Context) SettingsApiApiInitBrowserSettingsFlowRequestRequest {
	return SettingsApiApiInitBrowserSettingsFlowRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SettingsFlow
 */
func (a *SettingsApiService) InitBrowserSettingsFlowRequestExecute(r SettingsApiApiInitBrowserSettingsFlowRequestRequest) (*SettingsFlow, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *SettingsFlow
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsApiService.InitBrowserSettingsFlowRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/self-service/settings/browser"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnTo != nil {
		localVarQueryParams.Add("return_to", parameterToString(*r.returnTo, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SettingsApiApiSubmitSettingsFlowRequestRequest struct {
	ctx context.Context
	ApiService SettingsApi
	flow *string
	submitSettingsFlowBody *SubmitSettingsFlowBody
	xSessionToken *string
	cookie *string
}

func (r SettingsApiApiSubmitSettingsFlowRequestRequest) Flow(flow string) SettingsApiApiSubmitSettingsFlowRequestRequest {
	r.flow = &flow
	return r
}
func (r SettingsApiApiSubmitSettingsFlowRequestRequest) SubmitSettingsFlowBody(submitSettingsFlowBody SubmitSettingsFlowBody) SettingsApiApiSubmitSettingsFlowRequestRequest {
	r.submitSettingsFlowBody = &submitSettingsFlowBody
	return r
}
func (r SettingsApiApiSubmitSettingsFlowRequestRequest) XSessionToken(xSessionToken string) SettingsApiApiSubmitSettingsFlowRequestRequest {
	r.xSessionToken = &xSessionToken
	return r
}
func (r SettingsApiApiSubmitSettingsFlowRequestRequest) Cookie(cookie string) SettingsApiApiSubmitSettingsFlowRequestRequest {
	r.cookie = &cookie
	return r
}

func (r SettingsApiApiSubmitSettingsFlowRequestRequest) Execute() (*SettingsFlow, *http.Response, error) {
	return r.ApiService.SubmitSettingsFlowRequestExecute(r)
}

/*
 * SubmitSettingsFlowRequest # Complete Settings Flow
 * Use this endpoint to complete a settings flow by sending an identity's updated password. This endpoint
behaves differently for API and browser flows.

API-initiated flows expect `application/json` to be sent in the body and respond with
HTTP 200 and an application/json body with the session token on success;
HTTP 303 redirect to a fresh settings flow if the original flow expired with the appropriate error messages set;
HTTP 400 on form validation errors.
HTTP 401 when the endpoint is called without a valid session token.
HTTP 403 when `selfservice.flows.settings.privileged_session_max_age` was reached or the session's AAL is too low.
Implies that the user needs to re-authenticate.

Browser flows without HTTP Header `Accept` or with `Accept: text/*` respond with
a HTTP 303 redirect to the post/after settings URL or the `return_to` value if it was set and if the flow succeeded;
a HTTP 303 redirect to the Settings UI URL with the flow ID containing the validation errors otherwise.
a HTTP 303 redirect to the login endpoint when `selfservice.flows.settings.privileged_session_max_age` was reached or the session's AAL is too low.

Browser flows with HTTP Header `Accept: application/json` respond with
HTTP 200 and a application/json body with the signed in identity and a `Set-Cookie` header on success;
HTTP 303 redirect to a fresh login flow if the original flow expired with the appropriate error messages set;
HTTP 401 when the endpoint is called without a valid session cookie.
HTTP 403 when the page is accessed without a session cookie or the session's AAL is too low.
HTTP 400 on form validation errors.

Depending on your configuration this endpoint might return a 403 error if the session has a lower Authenticator
Assurance Level (AAL) than is possible for the identity. This can happen if the identity has password + webauthn
credentials (which would result in AAL2) but the session has only AAL1. If this error occurs, ask the user
to sign in with the second factor (happens automatically for server-side browser flows) or change the configuration.

If this endpoint is called with a `Accept: application/json` HTTP header, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`session_refresh_required`: The identity requested to change something that needs a privileged session. Redirect
the identity to the login init endpoint with query parameters `?refresh=true&return_to=<the-current-browser-url>`,
or initiate a refresh login flow otherwise.
`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`session_inactive`: No Identity Session was found - sign in a user first.
`security_identity_mismatch`: The flow was interrupted with `session_refresh_required` but apparently some other
identity logged in instead.
`security_identity_mismatch`: The requested `?return_to` address is not allowed to be used. Adjust this in the configuration!
`browser_location_change_required`: Usually sent when an AJAX request indicates that the browser needs to open a specific URL.
Most likely used in Social Sign In flows.

More information can be found at [Identity User Settings & Profile Management Documentation](../self-service/flows/user-settings).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return SettingsApiApiSubmitSettingsFlowRequestRequest
 */
func (a *SettingsApiService) SubmitSettingsFlowRequest(ctx context.Context) SettingsApiApiSubmitSettingsFlowRequestRequest {
	return SettingsApiApiSubmitSettingsFlowRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SettingsFlow
 */
func (a *SettingsApiService) SubmitSettingsFlowRequestExecute(r SettingsApiApiSubmitSettingsFlowRequestRequest) (*SettingsFlow, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *SettingsFlow
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsApiService.SubmitSettingsFlowRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/self-service/settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.flow == nil {
		return localVarReturnValue, nil, reportError("flow is required and must be specified")
	}
	if r.submitSettingsFlowBody == nil {
		return localVarReturnValue, nil, reportError("submitSettingsFlowBody is required and must be specified")
	}

	localVarQueryParams.Add("flow", parameterToString(*r.flow, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xSessionToken != nil {
		localVarHeaderParams["X-Session-Token"] = parameterToString(*r.xSessionToken, "")
	}
	if r.cookie != nil {
		localVarHeaderParams["Cookie"] = parameterToString(*r.cookie, "")
	}
	// body params
	localVarPostBody = r.submitSettingsFlowBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v SettingsFlow
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 410 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v JsonErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}