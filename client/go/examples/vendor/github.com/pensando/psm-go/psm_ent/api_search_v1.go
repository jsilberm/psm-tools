/*
 * Search API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// SearchV1ApiService SearchV1Api service
type SearchV1ApiService service

type ApiSearchGetPolicyQuery1Request struct {
	ctx _context.Context
	ApiService *SearchV1ApiService
	tenant *string
	namespace *string
	app *string
	protocol *string
	port *string
	fromIpAddress *string
	toIpAddress *string
	fromSecurityGroup *string
	toSecurityGroup *string
	kinds *[]string
	name *string
	action *string
}

func (r ApiSearchGetPolicyQuery1Request) Tenant(tenant string) ApiSearchGetPolicyQuery1Request {
	r.tenant = &tenant
	return r
}
func (r ApiSearchGetPolicyQuery1Request) Namespace(namespace string) ApiSearchGetPolicyQuery1Request {
	r.namespace = &namespace
	return r
}
func (r ApiSearchGetPolicyQuery1Request) App(app string) ApiSearchGetPolicyQuery1Request {
	r.app = &app
	return r
}
func (r ApiSearchGetPolicyQuery1Request) Protocol(protocol string) ApiSearchGetPolicyQuery1Request {
	r.protocol = &protocol
	return r
}
func (r ApiSearchGetPolicyQuery1Request) Port(port string) ApiSearchGetPolicyQuery1Request {
	r.port = &port
	return r
}
func (r ApiSearchGetPolicyQuery1Request) FromIpAddress(fromIpAddress string) ApiSearchGetPolicyQuery1Request {
	r.fromIpAddress = &fromIpAddress
	return r
}
func (r ApiSearchGetPolicyQuery1Request) ToIpAddress(toIpAddress string) ApiSearchGetPolicyQuery1Request {
	r.toIpAddress = &toIpAddress
	return r
}
func (r ApiSearchGetPolicyQuery1Request) FromSecurityGroup(fromSecurityGroup string) ApiSearchGetPolicyQuery1Request {
	r.fromSecurityGroup = &fromSecurityGroup
	return r
}
func (r ApiSearchGetPolicyQuery1Request) ToSecurityGroup(toSecurityGroup string) ApiSearchGetPolicyQuery1Request {
	r.toSecurityGroup = &toSecurityGroup
	return r
}
func (r ApiSearchGetPolicyQuery1Request) Kinds(kinds []string) ApiSearchGetPolicyQuery1Request {
	r.kinds = &kinds
	return r
}
func (r ApiSearchGetPolicyQuery1Request) Name(name string) ApiSearchGetPolicyQuery1Request {
	r.name = &name
	return r
}
func (r ApiSearchGetPolicyQuery1Request) Action(action string) ApiSearchGetPolicyQuery1Request {
	r.action = &action
	return r
}

func (r ApiSearchGetPolicyQuery1Request) Execute() (SearchPolicySearchResponse, *_nethttp.Response, error) {
	return r.ApiService.GetPolicyQuery1Execute(r)
}

/*
 * GetPolicyQuery1 Security Policy Query
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSearchGetPolicyQuery1Request
 */
func (a *SearchV1ApiService) GetPolicyQuery1(ctx _context.Context) ApiSearchGetPolicyQuery1Request {
	return ApiSearchGetPolicyQuery1Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SearchPolicySearchResponse
 */
func (a *SearchV1ApiService) GetPolicyQuery1Execute(r ApiSearchGetPolicyQuery1Request) (SearchPolicySearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SearchPolicySearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchV1ApiService.GetPolicyQuery1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search/v1/policy-query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.tenant != nil {
		localVarQueryParams.Add("tenant", parameterToString(*r.tenant, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.app != nil {
		localVarQueryParams.Add("app", parameterToString(*r.app, ""))
	}
	if r.protocol != nil {
		localVarQueryParams.Add("protocol", parameterToString(*r.protocol, ""))
	}
	if r.port != nil {
		localVarQueryParams.Add("port", parameterToString(*r.port, ""))
	}
	if r.fromIpAddress != nil {
		localVarQueryParams.Add("from-ip-address", parameterToString(*r.fromIpAddress, ""))
	}
	if r.toIpAddress != nil {
		localVarQueryParams.Add("to-ip-address", parameterToString(*r.toIpAddress, ""))
	}
	if r.fromSecurityGroup != nil {
		localVarQueryParams.Add("from-security-group", parameterToString(*r.fromSecurityGroup, ""))
	}
	if r.toSecurityGroup != nil {
		localVarQueryParams.Add("to-security-group", parameterToString(*r.toSecurityGroup, ""))
	}
	if r.kinds != nil {
		localVarQueryParams.Add("kinds", parameterToString(*r.kinds, "csv"))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.action != nil {
		localVarQueryParams.Add("action", parameterToString(*r.action, ""))
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode == 401 {
		a.client.cfg.PSMCfg.Login()
		a.client.cfg.PSMCfg.SaveConfig()
		return a.GetPolicyQuery1Execute(r)
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchGetQuery1Request struct {
	ctx _context.Context
	ApiService *SearchV1ApiService
	queryString *string
	from *int32
	maxResults *int32
	sortBy *string
	sortOrder *string
	mode *string
	queryCategories *[]string
	queryKinds *[]string
	tenants *[]string
	aggregate *bool
}

func (r ApiSearchGetQuery1Request) QueryString(queryString string) ApiSearchGetQuery1Request {
	r.queryString = &queryString
	return r
}
func (r ApiSearchGetQuery1Request) From(from int32) ApiSearchGetQuery1Request {
	r.from = &from
	return r
}
func (r ApiSearchGetQuery1Request) MaxResults(maxResults int32) ApiSearchGetQuery1Request {
	r.maxResults = &maxResults
	return r
}
func (r ApiSearchGetQuery1Request) SortBy(sortBy string) ApiSearchGetQuery1Request {
	r.sortBy = &sortBy
	return r
}
func (r ApiSearchGetQuery1Request) SortOrder(sortOrder string) ApiSearchGetQuery1Request {
	r.sortOrder = &sortOrder
	return r
}
func (r ApiSearchGetQuery1Request) Mode(mode string) ApiSearchGetQuery1Request {
	r.mode = &mode
	return r
}
func (r ApiSearchGetQuery1Request) QueryCategories(queryCategories []string) ApiSearchGetQuery1Request {
	r.queryCategories = &queryCategories
	return r
}
func (r ApiSearchGetQuery1Request) QueryKinds(queryKinds []string) ApiSearchGetQuery1Request {
	r.queryKinds = &queryKinds
	return r
}
func (r ApiSearchGetQuery1Request) Tenants(tenants []string) ApiSearchGetQuery1Request {
	r.tenants = &tenants
	return r
}
func (r ApiSearchGetQuery1Request) Aggregate(aggregate bool) ApiSearchGetQuery1Request {
	r.aggregate = &aggregate
	return r
}

func (r ApiSearchGetQuery1Request) Execute() (SearchSearchResponse, *_nethttp.Response, error) {
	return r.ApiService.GetQuery1Execute(r)
}

/*
 * GetQuery1 Structured or free-form search
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSearchGetQuery1Request
 */
func (a *SearchV1ApiService) GetQuery1(ctx _context.Context) ApiSearchGetQuery1Request {
	return ApiSearchGetQuery1Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SearchSearchResponse
 */
func (a *SearchV1ApiService) GetQuery1Execute(r ApiSearchGetQuery1Request) (SearchSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SearchSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchV1ApiService.GetQuery1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search/v1/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.queryString != nil {
		localVarQueryParams.Add("query-string", parameterToString(*r.queryString, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.maxResults != nil {
		localVarQueryParams.Add("max-results", parameterToString(*r.maxResults, ""))
	}
	if r.sortBy != nil {
		localVarQueryParams.Add("sort-by", parameterToString(*r.sortBy, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sort-order", parameterToString(*r.sortOrder, ""))
	}
	if r.mode != nil {
		localVarQueryParams.Add("mode", parameterToString(*r.mode, ""))
	}
	if r.queryCategories != nil {
		localVarQueryParams.Add("query.categories", parameterToString(*r.queryCategories, "csv"))
	}
	if r.queryKinds != nil {
		localVarQueryParams.Add("query.kinds", parameterToString(*r.queryKinds, "csv"))
	}
	if r.tenants != nil {
		localVarQueryParams.Add("tenants", parameterToString(*r.tenants, "csv"))
	}
	if r.aggregate != nil {
		localVarQueryParams.Add("aggregate", parameterToString(*r.aggregate, ""))
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode == 401 {
		a.client.cfg.PSMCfg.Login()
		a.client.cfg.PSMCfg.SaveConfig()
		return a.GetQuery1Execute(r)
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchPostPolicyQueryRequest struct {
	ctx _context.Context
	ApiService *SearchV1ApiService
	body *SearchPolicySearchRequest
}

func (r ApiSearchPostPolicyQueryRequest) Body(body SearchPolicySearchRequest) ApiSearchPostPolicyQueryRequest {
	r.body = &body
	return r
}

func (r ApiSearchPostPolicyQueryRequest) Execute() (SearchPolicySearchResponse, *_nethttp.Response, error) {
	return r.ApiService.PostPolicyQueryExecute(r)
}

/*
 * PostPolicyQuery Security Policy Query
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSearchPostPolicyQueryRequest
 */
func (a *SearchV1ApiService) PostPolicyQuery(ctx _context.Context) ApiSearchPostPolicyQueryRequest {
	return ApiSearchPostPolicyQueryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SearchPolicySearchResponse
 */
func (a *SearchV1ApiService) PostPolicyQueryExecute(r ApiSearchPostPolicyQueryRequest) (SearchPolicySearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SearchPolicySearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchV1ApiService.PostPolicyQuery")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search/v1/policy-query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode == 401 {
		a.client.cfg.PSMCfg.Login()
		a.client.cfg.PSMCfg.SaveConfig()
		return a.PostPolicyQueryExecute(r)
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchPostQueryRequest struct {
	ctx _context.Context
	ApiService *SearchV1ApiService
	body *SearchSearchRequest
}

func (r ApiSearchPostQueryRequest) Body(body SearchSearchRequest) ApiSearchPostQueryRequest {
	r.body = &body
	return r
}

func (r ApiSearchPostQueryRequest) Execute() (SearchSearchResponse, *_nethttp.Response, error) {
	return r.ApiService.PostQueryExecute(r)
}

/*
 * PostQuery Structured or free-form search
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSearchPostQueryRequest
 */
func (a *SearchV1ApiService) PostQuery(ctx _context.Context) ApiSearchPostQueryRequest {
	return ApiSearchPostQueryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SearchSearchResponse
 */
func (a *SearchV1ApiService) PostQueryExecute(r ApiSearchPostQueryRequest) (SearchSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SearchSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchV1ApiService.PostQuery")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search/v1/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode == 401 {
		a.client.cfg.PSMCfg.Login()
		a.client.cfg.PSMCfg.SaveConfig()
		return a.PostQueryExecute(r)
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
