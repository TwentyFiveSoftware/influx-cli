/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	_context "context"
	_fmt "fmt"
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type DeleteApi interface {

	/*
	 * PostDelete Delete data
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostDeleteRequest
	 */
	PostDelete(ctx _context.Context) ApiPostDeleteRequest

	/*
	 * PostDeleteExecute executes the request
	 */
	PostDeleteExecute(r ApiPostDeleteRequest) error

	/*
	 * PostDeleteExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 */
	PostDeleteExecuteWithHttpInfo(r ApiPostDeleteRequest) (*_nethttp.Response, error)
}

// DeleteApiService DeleteApi service
type DeleteApiService service

type ApiPostDeleteRequest struct {
	ctx                    _context.Context
	ApiService             DeleteApi
	deletePredicateRequest *DeletePredicateRequest
	zapTraceSpan           *string
	org                    *string
	bucket                 *string
	orgID                  *string
	bucketID               *string
}

func (r ApiPostDeleteRequest) DeletePredicateRequest(deletePredicateRequest DeletePredicateRequest) ApiPostDeleteRequest {
	r.deletePredicateRequest = &deletePredicateRequest
	return r
}
func (r ApiPostDeleteRequest) GetDeletePredicateRequest() *DeletePredicateRequest {
	return r.deletePredicateRequest
}

func (r ApiPostDeleteRequest) ZapTraceSpan(zapTraceSpan string) ApiPostDeleteRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostDeleteRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostDeleteRequest) Org(org string) ApiPostDeleteRequest {
	r.org = &org
	return r
}
func (r ApiPostDeleteRequest) GetOrg() *string {
	return r.org
}

func (r ApiPostDeleteRequest) Bucket(bucket string) ApiPostDeleteRequest {
	r.bucket = &bucket
	return r
}
func (r ApiPostDeleteRequest) GetBucket() *string {
	return r.bucket
}

func (r ApiPostDeleteRequest) OrgID(orgID string) ApiPostDeleteRequest {
	r.orgID = &orgID
	return r
}
func (r ApiPostDeleteRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiPostDeleteRequest) BucketID(bucketID string) ApiPostDeleteRequest {
	r.bucketID = &bucketID
	return r
}
func (r ApiPostDeleteRequest) GetBucketID() *string {
	return r.bucketID
}

func (r ApiPostDeleteRequest) Execute() error {
	return r.ApiService.PostDeleteExecute(r)
}

func (r ApiPostDeleteRequest) ExecuteWithHttpInfo() (*_nethttp.Response, error) {
	return r.ApiService.PostDeleteExecuteWithHttpInfo(r)
}

/*
 * PostDelete Delete data
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostDeleteRequest
 */
func (a *DeleteApiService) PostDelete(ctx _context.Context) ApiPostDeleteRequest {
	return ApiPostDeleteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *DeleteApiService) PostDeleteExecute(r ApiPostDeleteRequest) error {
	_, err := a.PostDeleteExecuteWithHttpInfo(r)
	return err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 */
func (a *DeleteApiService) PostDeleteExecuteWithHttpInfo(r ApiPostDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeleteApiService.PostDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.deletePredicateRequest == nil {
		return nil, reportError("deletePredicateRequest is required and must be specified")
	}

	if r.org != nil {
		localVarQueryParams.Add("org", parameterToString(*r.org, ""))
	}
	if r.bucket != nil {
		localVarQueryParams.Add("bucket", parameterToString(*r.bucket, ""))
	}
	if r.orgID != nil {
		localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
	}
	if r.bucketID != nil {
		localVarQueryParams.Add("bucketID", parameterToString(*r.bucketID, ""))
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
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	// body params
	localVarPostBody = r.deletePredicateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	newErr := GenericOpenAPIError{
		buildHeader: localVarHTTPResponse.Header.Get("X-Influxdb-Build"),
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.body = localVarBody
		newErr.error = localVarHTTPResponse.Status
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
				return localVarHTTPResponse, newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
			newErr.model = &v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
				return localVarHTTPResponse, newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
			newErr.model = &v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
				return localVarHTTPResponse, newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
			newErr.model = &v
			return localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarHTTPResponse, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
