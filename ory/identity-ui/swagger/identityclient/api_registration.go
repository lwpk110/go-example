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

type RegistrationApi interface {

	/*
	 * GetRegistrationFlowRequest # Get Registration Flow
	 * This endpoint returns a registration flow's context with, for example, error details and other information.

Browser flows expect the anti-CSRF cookie to be included in the request's HTTP Cookie Header.
For AJAX requests you must ensure that cookies are included in the request or requests will fail.

If you use the browser-flow for server-side apps, the services need to run on a common top-level-domain
and you need to forward the incoming HTTP Cookie header to this endpoint:

```js
pseudo-code example
router.get('/registration', async function (req, res) {
const flow = await client.getRegistrationFlowRequest(req.header('cookie'), req.query['flow'])

res.render('registration', flow)
})
```

This request may fail due to several reasons. The `error.id` can be one of:

`session_already_available`: The user is already signed in.
`self_service_flow_expired`: The flow is expired and you should request a new one.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return RegistrationApiApiGetRegistrationFlowRequestRequest
	 */
	GetRegistrationFlowRequest(ctx context.Context) RegistrationApiApiGetRegistrationFlowRequestRequest

	/*
	 * GetRegistrationFlowRequestExecute executes the request
	 * @return RegistrationFlow
	 */
	GetRegistrationFlowRequestExecute(r RegistrationApiApiGetRegistrationFlowRequestRequest) (*RegistrationFlow, *http.Response, error)

	/*
	 * InitBrowserRegistrationFlowRequest # Initialize Registration Flow for Browsers
	 * This endpoint initializes a browser-based user registration flow. This endpoint will set the appropriate
cookies and anti-CSRF measures required for browser-based flows.

:::info

This endpoint is EXPERIMENTAL and subject to potential breaking changes in the future.

:::

If this endpoint is opened as a link in the browser, it will be redirected to
`selfservice.flows.registration.ui_url` with the flow ID set as the query parameter `?flow=`. If a valid user session
exists already, the browser will be redirected to `urls.default_redirect_url`.

If this endpoint is called via an AJAX request, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`session_already_available`: The user is already signed in.
`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`security_identity_mismatch`: The requested `?return_to` address is not allowed to be used. Adjust this in the configuration!

If this endpoint is called via an AJAX request, the response contains the registration flow without a redirect.

This endpoint is NOT INTENDED for clients that do not have a browser (Chrome, Firefox, ...) as cookies are needed.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return RegistrationApiApiInitBrowserRegistrationFlowRequestRequest
	 */
	InitBrowserRegistrationFlowRequest(ctx context.Context) RegistrationApiApiInitBrowserRegistrationFlowRequestRequest

	/*
	 * InitBrowserRegistrationFlowRequestExecute executes the request
	 * @return RegistrationFlow
	 */
	InitBrowserRegistrationFlowRequestExecute(r RegistrationApiApiInitBrowserRegistrationFlowRequestRequest) (*RegistrationFlow, *http.Response, error)

	/*
	 * SubmitRegistrationFlowRequest # Submit a Registration Flow
	 * Use this endpoint to complete a registration flow by sending an identity's traits and password. This endpoint
behaves differently for API and browser flows.

API flows expect `application/json` to be sent in the body and respond with
HTTP 200 and a application/json body with the created identity success - if the session hook is configured the
`session` and `session_token` will also be included;
HTTP 410 if the original flow expired with the appropriate error messages set and optionally a `use_flow_id` parameter in the body;
HTTP 400 on form validation errors.

Browser flows expect a Content-Type of `application/x-www-form-urlencoded` or `application/json` to be sent in the body and respond with
a HTTP 303 redirect to the post/after registration URL or the `return_to` value if it was set and if the registration succeeded;
a HTTP 303 redirect to the registration UI URL with the flow ID containing the validation errors otherwise.

Browser flows with an accept header of `application/json` will not redirect but instead respond with
HTTP 200 and a application/json body with the signed in identity and a `Set-Cookie` header on success;
HTTP 303 redirect to a fresh login flow if the original flow expired with the appropriate error messages set;
HTTP 400 on form validation errors.

If this endpoint is called with `Accept: application/json` in the header, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`session_already_available`: The user is already signed in.
`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`security_identity_mismatch`: The requested `?return_to` address is not allowed to be used. Adjust this in the configuration!
`browser_location_change_required`: Usually sent when an AJAX request indicates that the browser needs to open a specific URL.
Most likely used in Social Sign In flows.

More information can be found at:
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return RegistrationApiApiSubmitRegistrationFlowRequestRequest
	 */
	SubmitRegistrationFlowRequest(ctx context.Context) RegistrationApiApiSubmitRegistrationFlowRequestRequest

	/*
	 * SubmitRegistrationFlowRequestExecute executes the request
	 * @return SubmitRegistrationFlowResponse
	 */
	SubmitRegistrationFlowRequestExecute(r RegistrationApiApiSubmitRegistrationFlowRequestRequest) (*SubmitRegistrationFlowResponse, *http.Response, error)
}

// RegistrationApiService RegistrationApi service
type RegistrationApiService service

type RegistrationApiApiGetRegistrationFlowRequestRequest struct {
	ctx context.Context
	ApiService RegistrationApi
	id *string
	cookie *string
}

func (r RegistrationApiApiGetRegistrationFlowRequestRequest) Id(id string) RegistrationApiApiGetRegistrationFlowRequestRequest {
	r.id = &id
	return r
}
func (r RegistrationApiApiGetRegistrationFlowRequestRequest) Cookie(cookie string) RegistrationApiApiGetRegistrationFlowRequestRequest {
	r.cookie = &cookie
	return r
}

func (r RegistrationApiApiGetRegistrationFlowRequestRequest) Execute() (*RegistrationFlow, *http.Response, error) {
	return r.ApiService.GetRegistrationFlowRequestExecute(r)
}

/*
 * GetRegistrationFlowRequest # Get Registration Flow
 * This endpoint returns a registration flow's context with, for example, error details and other information.

Browser flows expect the anti-CSRF cookie to be included in the request's HTTP Cookie Header.
For AJAX requests you must ensure that cookies are included in the request or requests will fail.

If you use the browser-flow for server-side apps, the services need to run on a common top-level-domain
and you need to forward the incoming HTTP Cookie header to this endpoint:

```js
pseudo-code example
router.get('/registration', async function (req, res) {
const flow = await client.getRegistrationFlowRequest(req.header('cookie'), req.query['flow'])

res.render('registration', flow)
})
```

This request may fail due to several reasons. The `error.id` can be one of:

`session_already_available`: The user is already signed in.
`self_service_flow_expired`: The flow is expired and you should request a new one.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return RegistrationApiApiGetRegistrationFlowRequestRequest
 */
func (a *RegistrationApiService) GetRegistrationFlowRequest(ctx context.Context) RegistrationApiApiGetRegistrationFlowRequestRequest {
	return RegistrationApiApiGetRegistrationFlowRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RegistrationFlow
 */
func (a *RegistrationApiService) GetRegistrationFlowRequestExecute(r RegistrationApiApiGetRegistrationFlowRequestRequest) (*RegistrationFlow, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *RegistrationFlow
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrationApiService.GetRegistrationFlowRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/self-service/registration/flows"

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

type RegistrationApiApiInitBrowserRegistrationFlowRequestRequest struct {
	ctx context.Context
	ApiService RegistrationApi
	returnTo *string
	loginChallenge *string
}

func (r RegistrationApiApiInitBrowserRegistrationFlowRequestRequest) ReturnTo(returnTo string) RegistrationApiApiInitBrowserRegistrationFlowRequestRequest {
	r.returnTo = &returnTo
	return r
}
func (r RegistrationApiApiInitBrowserRegistrationFlowRequestRequest) LoginChallenge(loginChallenge string) RegistrationApiApiInitBrowserRegistrationFlowRequestRequest {
	r.loginChallenge = &loginChallenge
	return r
}

func (r RegistrationApiApiInitBrowserRegistrationFlowRequestRequest) Execute() (*RegistrationFlow, *http.Response, error) {
	return r.ApiService.InitBrowserRegistrationFlowRequestExecute(r)
}

/*
 * InitBrowserRegistrationFlowRequest # Initialize Registration Flow for Browsers
 * This endpoint initializes a browser-based user registration flow. This endpoint will set the appropriate
cookies and anti-CSRF measures required for browser-based flows.

:::info

This endpoint is EXPERIMENTAL and subject to potential breaking changes in the future.

:::

If this endpoint is opened as a link in the browser, it will be redirected to
`selfservice.flows.registration.ui_url` with the flow ID set as the query parameter `?flow=`. If a valid user session
exists already, the browser will be redirected to `urls.default_redirect_url`.

If this endpoint is called via an AJAX request, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`session_already_available`: The user is already signed in.
`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`security_identity_mismatch`: The requested `?return_to` address is not allowed to be used. Adjust this in the configuration!

If this endpoint is called via an AJAX request, the response contains the registration flow without a redirect.

This endpoint is NOT INTENDED for clients that do not have a browser (Chrome, Firefox, ...) as cookies are needed.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return RegistrationApiApiInitBrowserRegistrationFlowRequestRequest
 */
func (a *RegistrationApiService) InitBrowserRegistrationFlowRequest(ctx context.Context) RegistrationApiApiInitBrowserRegistrationFlowRequestRequest {
	return RegistrationApiApiInitBrowserRegistrationFlowRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RegistrationFlow
 */
func (a *RegistrationApiService) InitBrowserRegistrationFlowRequestExecute(r RegistrationApiApiInitBrowserRegistrationFlowRequestRequest) (*RegistrationFlow, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *RegistrationFlow
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrationApiService.InitBrowserRegistrationFlowRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/self-service/registration/browser"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnTo != nil {
		localVarQueryParams.Add("return_to", parameterToString(*r.returnTo, ""))
	}
	if r.loginChallenge != nil {
		localVarQueryParams.Add("login_challenge", parameterToString(*r.loginChallenge, ""))
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

type RegistrationApiApiSubmitRegistrationFlowRequestRequest struct {
	ctx context.Context
	ApiService RegistrationApi
	flow *string
	submitRegistrationFlowBody *SubmitRegistrationFlowBody
	cookie *string
}

func (r RegistrationApiApiSubmitRegistrationFlowRequestRequest) Flow(flow string) RegistrationApiApiSubmitRegistrationFlowRequestRequest {
	r.flow = &flow
	return r
}
func (r RegistrationApiApiSubmitRegistrationFlowRequestRequest) SubmitRegistrationFlowBody(submitRegistrationFlowBody SubmitRegistrationFlowBody) RegistrationApiApiSubmitRegistrationFlowRequestRequest {
	r.submitRegistrationFlowBody = &submitRegistrationFlowBody
	return r
}
func (r RegistrationApiApiSubmitRegistrationFlowRequestRequest) Cookie(cookie string) RegistrationApiApiSubmitRegistrationFlowRequestRequest {
	r.cookie = &cookie
	return r
}

func (r RegistrationApiApiSubmitRegistrationFlowRequestRequest) Execute() (*SubmitRegistrationFlowResponse, *http.Response, error) {
	return r.ApiService.SubmitRegistrationFlowRequestExecute(r)
}

/*
 * SubmitRegistrationFlowRequest # Submit a Registration Flow
 * Use this endpoint to complete a registration flow by sending an identity's traits and password. This endpoint
behaves differently for API and browser flows.

API flows expect `application/json` to be sent in the body and respond with
HTTP 200 and a application/json body with the created identity success - if the session hook is configured the
`session` and `session_token` will also be included;
HTTP 410 if the original flow expired with the appropriate error messages set and optionally a `use_flow_id` parameter in the body;
HTTP 400 on form validation errors.

Browser flows expect a Content-Type of `application/x-www-form-urlencoded` or `application/json` to be sent in the body and respond with
a HTTP 303 redirect to the post/after registration URL or the `return_to` value if it was set and if the registration succeeded;
a HTTP 303 redirect to the registration UI URL with the flow ID containing the validation errors otherwise.

Browser flows with an accept header of `application/json` will not redirect but instead respond with
HTTP 200 and a application/json body with the signed in identity and a `Set-Cookie` header on success;
HTTP 303 redirect to a fresh login flow if the original flow expired with the appropriate error messages set;
HTTP 400 on form validation errors.

If this endpoint is called with `Accept: application/json` in the header, the response contains the flow without a redirect. In the
case of an error, the `error.id` of the JSON response body can be one of:

`session_already_available`: The user is already signed in.
`security_csrf_violation`: Unable to fetch the flow because a CSRF violation occurred.
`security_identity_mismatch`: The requested `?return_to` address is not allowed to be used. Adjust this in the configuration!
`browser_location_change_required`: Usually sent when an AJAX request indicates that the browser needs to open a specific URL.
Most likely used in Social Sign In flows.

More information can be found at:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return RegistrationApiApiSubmitRegistrationFlowRequestRequest
 */
func (a *RegistrationApiService) SubmitRegistrationFlowRequest(ctx context.Context) RegistrationApiApiSubmitRegistrationFlowRequestRequest {
	return RegistrationApiApiSubmitRegistrationFlowRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SubmitRegistrationFlowResponse
 */
func (a *RegistrationApiService) SubmitRegistrationFlowRequestExecute(r RegistrationApiApiSubmitRegistrationFlowRequestRequest) (*SubmitRegistrationFlowResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *SubmitRegistrationFlowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrationApiService.SubmitRegistrationFlowRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/self-service/registration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.flow == nil {
		return localVarReturnValue, nil, reportError("flow is required and must be specified")
	}
	if r.submitRegistrationFlowBody == nil {
		return localVarReturnValue, nil, reportError("submitRegistrationFlowBody is required and must be specified")
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
	if r.cookie != nil {
		localVarHeaderParams["Cookie"] = parameterToString(*r.cookie, "")
	}
	// body params
	localVarPostBody = r.submitRegistrationFlowBody
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
			var v RegistrationFlow
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
