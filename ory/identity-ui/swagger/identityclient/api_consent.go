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

type ConsentApi interface {

	/*
	 * GetConsentFlowRequest # Get Consent Flow
	 * This endpoint returns a consent flow's context with, for example, error details and other information.

Browser flows expect the anti-CSRF cookie to be included in the request's HTTP Cookie Header.
For AJAX requests you must ensure that cookies are included in the request or requests will fail.

If you use the browser-flow for server-side apps, the services need to run on a common top-level-domain
and you need to forward the incoming HTTP Cookie header to this endpoint:

```js
pseudo-code example
router.get('/consent', async function (req, res) {
const flow = await client.getSelfServiceConsentFlow(req.header('cookie'), req.query['flow'])

res.render('consent', flow)
})
```

This request may fail due to several reasons. The `error.id` can be one of:

`session_already_available`: The user is already signed in.
`self_service_flow_expired`: The flow is expired and you should request a new one.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ConsentApiApiGetConsentFlowRequestRequest
	 */
	GetConsentFlowRequest(ctx context.Context) ConsentApiApiGetConsentFlowRequestRequest

	/*
	 * GetConsentFlowRequestExecute executes the request
	 * @return ConsentFlow
	 */
	GetConsentFlowRequestExecute(r ConsentApiApiGetConsentFlowRequestRequest) (*ConsentFlow, *http.Response, error)

	/*
	 * InitBrowserConsentFlowRequest # Initialize Consent Flow for Browsers
	 * This endpoint initializes a browser-based user consent flow. This endpoint will set the appropriate
cookies and anti-CSRF measures required for browser-based flows.

If this endpoint is opened as a link in the browser, it will be redirected to
`selfservice.flows.consent.ui_url` with the flow ID set as the query parameter `?flow=`. If a valid user session
exists already, the browser will be redirected to `urls.default_redirect_url` unless the query parameter
`?refresh=true` was set.

If this endpoint is called via an AJAX request, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`security_identity_mismatch`: The requested `?return_to` address is not allowed to be used. Adjust this in the configuration!

This endpoint is NOT INTENDED for clients that do not have a browser (Chrome, Firefox, ...) as cookies are needed.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ConsentApiApiInitBrowserConsentFlowRequestRequest
	 */
	InitBrowserConsentFlowRequest(ctx context.Context) ConsentApiApiInitBrowserConsentFlowRequestRequest

	/*
	 * InitBrowserConsentFlowRequestExecute executes the request
	 * @return ConsentFlow
	 */
	InitBrowserConsentFlowRequestExecute(r ConsentApiApiInitBrowserConsentFlowRequestRequest) (*ConsentFlow, *http.Response, error)

	/*
	 * SubmitConsentFlowRequest # Submit a Consent Flow
	 * Use this endpoint to complete a consent flow. This endpoint
behaves differently for API and browser flows.

API flows expect `application/json` to be sent in the body and responds with
HTTP 200 and a application/json body with the session token on success;
HTTP 410 if the original flow expired with the appropriate error messages set and optionally a `use_flow_id` parameter in the body;
HTTP 400 on form validation errors.

Browser flows expect a Content-Type of `application/x-www-form-urlencoded` or `application/json` to be sent in the body and respond with
a HTTP 303 redirect to the post/after consent URL or the `return_to` value if it was set and if the consent succeeded;
a HTTP 303 redirect to the consent UI URL with the flow ID containing the validation errors otherwise.

More information can be found at:
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ConsentApiApiSubmitConsentFlowRequestRequest
	 */
	SubmitConsentFlowRequest(ctx context.Context) ConsentApiApiSubmitConsentFlowRequestRequest

	/*
	 * SubmitConsentFlowRequestExecute executes the request
	 * @return ConsentFlow
	 */
	SubmitConsentFlowRequestExecute(r ConsentApiApiSubmitConsentFlowRequestRequest) (*ConsentFlow, *http.Response, error)
}

// ConsentApiService ConsentApi service
type ConsentApiService service

type ConsentApiApiGetConsentFlowRequestRequest struct {
	ctx context.Context
	ApiService ConsentApi
	id *string
	cookie *string
}

func (r ConsentApiApiGetConsentFlowRequestRequest) Id(id string) ConsentApiApiGetConsentFlowRequestRequest {
	r.id = &id
	return r
}
func (r ConsentApiApiGetConsentFlowRequestRequest) Cookie(cookie string) ConsentApiApiGetConsentFlowRequestRequest {
	r.cookie = &cookie
	return r
}

func (r ConsentApiApiGetConsentFlowRequestRequest) Execute() (*ConsentFlow, *http.Response, error) {
	return r.ApiService.GetConsentFlowRequestExecute(r)
}

/*
 * GetConsentFlowRequest # Get Consent Flow
 * This endpoint returns a consent flow's context with, for example, error details and other information.

Browser flows expect the anti-CSRF cookie to be included in the request's HTTP Cookie Header.
For AJAX requests you must ensure that cookies are included in the request or requests will fail.

If you use the browser-flow for server-side apps, the services need to run on a common top-level-domain
and you need to forward the incoming HTTP Cookie header to this endpoint:

```js
pseudo-code example
router.get('/consent', async function (req, res) {
const flow = await client.getSelfServiceConsentFlow(req.header('cookie'), req.query['flow'])

res.render('consent', flow)
})
```

This request may fail due to several reasons. The `error.id` can be one of:

`session_already_available`: The user is already signed in.
`self_service_flow_expired`: The flow is expired and you should request a new one.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ConsentApiApiGetConsentFlowRequestRequest
 */
func (a *ConsentApiService) GetConsentFlowRequest(ctx context.Context) ConsentApiApiGetConsentFlowRequestRequest {
	return ConsentApiApiGetConsentFlowRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ConsentFlow
 */
func (a *ConsentApiService) GetConsentFlowRequestExecute(r ConsentApiApiGetConsentFlowRequestRequest) (*ConsentFlow, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *ConsentFlow
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConsentApiService.GetConsentFlowRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/self-service/consent/flows"

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

type ConsentApiApiInitBrowserConsentFlowRequestRequest struct {
	ctx context.Context
	ApiService ConsentApi
	consentChallenge *string
	returnTo *string
	cookie *string
}

func (r ConsentApiApiInitBrowserConsentFlowRequestRequest) ConsentChallenge(consentChallenge string) ConsentApiApiInitBrowserConsentFlowRequestRequest {
	r.consentChallenge = &consentChallenge
	return r
}
func (r ConsentApiApiInitBrowserConsentFlowRequestRequest) ReturnTo(returnTo string) ConsentApiApiInitBrowserConsentFlowRequestRequest {
	r.returnTo = &returnTo
	return r
}
func (r ConsentApiApiInitBrowserConsentFlowRequestRequest) Cookie(cookie string) ConsentApiApiInitBrowserConsentFlowRequestRequest {
	r.cookie = &cookie
	return r
}

func (r ConsentApiApiInitBrowserConsentFlowRequestRequest) Execute() (*ConsentFlow, *http.Response, error) {
	return r.ApiService.InitBrowserConsentFlowRequestExecute(r)
}

/*
 * InitBrowserConsentFlowRequest # Initialize Consent Flow for Browsers
 * This endpoint initializes a browser-based user consent flow. This endpoint will set the appropriate
cookies and anti-CSRF measures required for browser-based flows.

If this endpoint is opened as a link in the browser, it will be redirected to
`selfservice.flows.consent.ui_url` with the flow ID set as the query parameter `?flow=`. If a valid user session
exists already, the browser will be redirected to `urls.default_redirect_url` unless the query parameter
`?refresh=true` was set.

If this endpoint is called via an AJAX request, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`security_identity_mismatch`: The requested `?return_to` address is not allowed to be used. Adjust this in the configuration!

This endpoint is NOT INTENDED for clients that do not have a browser (Chrome, Firefox, ...) as cookies are needed.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ConsentApiApiInitBrowserConsentFlowRequestRequest
 */
func (a *ConsentApiService) InitBrowserConsentFlowRequest(ctx context.Context) ConsentApiApiInitBrowserConsentFlowRequestRequest {
	return ConsentApiApiInitBrowserConsentFlowRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ConsentFlow
 */
func (a *ConsentApiService) InitBrowserConsentFlowRequestExecute(r ConsentApiApiInitBrowserConsentFlowRequestRequest) (*ConsentFlow, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *ConsentFlow
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConsentApiService.InitBrowserConsentFlowRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/self-service/consent/browser"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.consentChallenge == nil {
		return localVarReturnValue, nil, reportError("consentChallenge is required and must be specified")
	}

	if r.returnTo != nil {
		localVarQueryParams.Add("return_to", parameterToString(*r.returnTo, ""))
	}
	localVarQueryParams.Add("consent_challenge", parameterToString(*r.consentChallenge, ""))
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

type ConsentApiApiSubmitConsentFlowRequestRequest struct {
	ctx context.Context
	ApiService ConsentApi
	flow *string
	submitConsentFlowBody *SubmitConsentFlowBody
	xSessionToken *string
	cookie *string
}

func (r ConsentApiApiSubmitConsentFlowRequestRequest) Flow(flow string) ConsentApiApiSubmitConsentFlowRequestRequest {
	r.flow = &flow
	return r
}
func (r ConsentApiApiSubmitConsentFlowRequestRequest) SubmitConsentFlowBody(submitConsentFlowBody SubmitConsentFlowBody) ConsentApiApiSubmitConsentFlowRequestRequest {
	r.submitConsentFlowBody = &submitConsentFlowBody
	return r
}
func (r ConsentApiApiSubmitConsentFlowRequestRequest) XSessionToken(xSessionToken string) ConsentApiApiSubmitConsentFlowRequestRequest {
	r.xSessionToken = &xSessionToken
	return r
}
func (r ConsentApiApiSubmitConsentFlowRequestRequest) Cookie(cookie string) ConsentApiApiSubmitConsentFlowRequestRequest {
	r.cookie = &cookie
	return r
}

func (r ConsentApiApiSubmitConsentFlowRequestRequest) Execute() (*ConsentFlow, *http.Response, error) {
	return r.ApiService.SubmitConsentFlowRequestExecute(r)
}

/*
 * SubmitConsentFlowRequest # Submit a Consent Flow
 * Use this endpoint to complete a consent flow. This endpoint
behaves differently for API and browser flows.

API flows expect `application/json` to be sent in the body and responds with
HTTP 200 and a application/json body with the session token on success;
HTTP 410 if the original flow expired with the appropriate error messages set and optionally a `use_flow_id` parameter in the body;
HTTP 400 on form validation errors.

Browser flows expect a Content-Type of `application/x-www-form-urlencoded` or `application/json` to be sent in the body and respond with
a HTTP 303 redirect to the post/after consent URL or the `return_to` value if it was set and if the consent succeeded;
a HTTP 303 redirect to the consent UI URL with the flow ID containing the validation errors otherwise.

More information can be found at:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ConsentApiApiSubmitConsentFlowRequestRequest
 */
func (a *ConsentApiService) SubmitConsentFlowRequest(ctx context.Context) ConsentApiApiSubmitConsentFlowRequestRequest {
	return ConsentApiApiSubmitConsentFlowRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ConsentFlow
 */
func (a *ConsentApiService) SubmitConsentFlowRequestExecute(r ConsentApiApiSubmitConsentFlowRequestRequest) (*ConsentFlow, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *ConsentFlow
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConsentApiService.SubmitConsentFlowRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/self-service/consent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.flow == nil {
		return localVarReturnValue, nil, reportError("flow is required and must be specified")
	}
	if r.submitConsentFlowBody == nil {
		return localVarReturnValue, nil, reportError("submitConsentFlowBody is required and must be specified")
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
	localVarPostBody = r.submitConsentFlowBody
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
			var v ConsentFlow
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