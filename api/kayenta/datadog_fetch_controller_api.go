/*
 * Kayenta API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kayenta

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
)

// Linger please
var (
	_ context.Context
)

type DatadogFetchControllerApiService service


/* DatadogFetchControllerApiService queryMetrics
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "metricsAccountName" (string) metricsAccountName
     @param "storageAccountName" (string) storageAccountName
     @param "metricSetName" (string) metricSetName
     @param "metricName" (string) metricName
     @param "scope" (string) The scope of the Datadog query. e.g. autoscaling_group:myapp-prod-v002
     @param "start" (string) An ISO format timestamp, e.g.: 2018-03-15T01:23:45Z
     @param "end" (string) An ISO format timestamp, e.g.: 2018-03-15T01:23:45Z
     @param "step" (int64) seconds
 @return interface{}*/
func (a *DatadogFetchControllerApiService) QueryMetricsUsingPOST(ctx context.Context, localVarOptionals map[string]interface{}) (interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/fetch/datadog/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["metricsAccountName"], "string", "metricsAccountName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storageAccountName"], "string", "storageAccountName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["metricSetName"], "string", "metricSetName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["metricName"], "string", "metricName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["scope"], "string", "scope"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["start"], "string", "start"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["end"], "string", "end"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["step"], "int64", "step"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["metricsAccountName"].(string); localVarOk {
		localVarQueryParams.Add("metricsAccountName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storageAccountName"].(string); localVarOk {
		localVarQueryParams.Add("storageAccountName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["metricSetName"].(string); localVarOk {
		localVarQueryParams.Add("metricSetName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["metricName"].(string); localVarOk {
		localVarQueryParams.Add("metricName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["scope"].(string); localVarOk {
		localVarQueryParams.Add("scope", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["start"].(string); localVarOk {
		localVarQueryParams.Add("start", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["end"].(string); localVarOk {
		localVarQueryParams.Add("end", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["step"].(int64); localVarOk {
		localVarQueryParams.Add("step", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"*/*",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}
