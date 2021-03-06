# \CredentialsControllerApi

All URIs are relative to *https://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountUsingGET**](CredentialsControllerApi.md#GetAccountUsingGET) | **Get** /credentials/{account} | Retrieve an account&#39;s details
[**GetAccountsUsingGET**](CredentialsControllerApi.md#GetAccountsUsingGET) | **Get** /credentials | Retrieve a list of accounts


# **GetAccountUsingGET**
> AccountDetails GetAccountUsingGET(ctx, account, optional)
Retrieve an account's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| account | 
 **roles** | [**[]string**](string.md)|  | 
 **allowedAccounts** | [**[]string**](string.md)|  | 
 **email** | **string**|  | 
 **username** | **string**|  | 
 **firstName** | **string**|  | 
 **lastName** | **string**|  | 
 **xRateLimitApp** | **string**| X-RateLimit-App | 

### Return type

[**AccountDetails**](AccountDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountsUsingGET**
> []Account GetAccountsUsingGET(ctx, optional)
Retrieve a list of accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roles** | [**[]string**](string.md)|  | 
 **allowedAccounts** | [**[]string**](string.md)|  | 
 **email** | **string**|  | 
 **username** | **string**|  | 
 **firstName** | **string**|  | 
 **lastName** | **string**|  | 
 **expand** | **bool**| expand | 

### Return type

[**[]Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

