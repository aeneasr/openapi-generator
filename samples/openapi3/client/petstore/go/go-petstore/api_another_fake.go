/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

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

type AnotherFakeApi interface {

	/*
	 * Call123TestSpecialTags To test special tags
	 * To test special tags and operation ID starting with number
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiCall123TestSpecialTagsRequest
	 */
	Call123TestSpecialTags(ctx context.Context) ApiCall123TestSpecialTagsRequest

	/*
	 * Call123TestSpecialTagsExecute executes the request
	 * @return Client
	 */
	Call123TestSpecialTagsExecute(r ApiCall123TestSpecialTagsRequest) (*Client, *http.Response, error)
}

// AnotherFakeApiService AnotherFakeApi service
type AnotherFakeApiService service

type ApiCall123TestSpecialTagsRequest struct {
	ctx context.Context
	ApiService AnotherFakeApi
	client *Client
}

func (r ApiCall123TestSpecialTagsRequest) Client(client Client) ApiCall123TestSpecialTagsRequest {
	r.client = &client
	return r
}

func (r ApiCall123TestSpecialTagsRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.Call123TestSpecialTagsExecute(r)
}

/*
 * Call123TestSpecialTags To test special tags
 * To test special tags and operation ID starting with number
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCall123TestSpecialTagsRequest
 */
func (a *AnotherFakeApiService) Call123TestSpecialTags(ctx context.Context) ApiCall123TestSpecialTagsRequest {
	return ApiCall123TestSpecialTagsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return Client
 */
func (a *AnotherFakeApiService) Call123TestSpecialTagsExecute(r ApiCall123TestSpecialTagsRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnotherFakeApiService.Call123TestSpecialTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/another-fake/dummy"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.client == nil {
		return localVarReturnValue, nil, reportError("client is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.client
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
