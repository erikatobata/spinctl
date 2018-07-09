# Go API client for kayenta

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./kayenta"
```

## Documentation for API Endpoints

All URIs are relative to *https://spinnaker-kayenta.example.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminControllerApi* | [**SetInstanceEnabledUsingPOST**](docs/AdminControllerApi.md#setinstanceenabledusingpost) | **Post** /admin/orca/instance/enabled | setInstanceEnabled
*CanaryConfigControllerApi* | [**DeleteCanaryConfigUsingDELETE**](docs/CanaryConfigControllerApi.md#deletecanaryconfigusingdelete) | **Delete** /canaryConfig/{canaryConfigId} | Delete a canary config
*CanaryConfigControllerApi* | [**ListAllCanaryConfigsUsingGET**](docs/CanaryConfigControllerApi.md#listallcanaryconfigsusingget) | **Get** /canaryConfig | Retrieve a list of canary config ids and timestamps
*CanaryConfigControllerApi* | [**LoadCanaryConfigUsingGET**](docs/CanaryConfigControllerApi.md#loadcanaryconfigusingget) | **Get** /canaryConfig/{canaryConfigId} | Retrieve a canary config from object storage
*CanaryConfigControllerApi* | [**StoreCanaryConfigUsingPOST**](docs/CanaryConfigControllerApi.md#storecanaryconfigusingpost) | **Post** /canaryConfig | Write a canary config to object storage
*CanaryConfigControllerApi* | [**UpdateCanaryConfigUsingPUT**](docs/CanaryConfigControllerApi.md#updatecanaryconfigusingput) | **Put** /canaryConfig/{canaryConfigId} | Update a canary config
*CanaryControllerApi* | [**GetCanaryResultsByApplicationUsingGET**](docs/CanaryControllerApi.md#getcanaryresultsbyapplicationusingget) | **Get** /canary/executions | Retrieve a list of an application&#39;s canary results
*CanaryControllerApi* | [**GetCanaryResultsUsingGET**](docs/CanaryControllerApi.md#getcanaryresultsusingget) | **Get** /canary/{canaryExecutionId} | Retrieve status and results for a canary run
*CanaryControllerApi* | [**InitiateCanaryUsingPOST**](docs/CanaryControllerApi.md#initiatecanaryusingpost) | **Post** /canary/{canaryConfigId} | Initiate a canary pipeline
*CanaryControllerApi* | [**InitiateCanaryWithConfigUsingPOST**](docs/CanaryControllerApi.md#initiatecanarywithconfigusingpost) | **Post** /canary | Initiate a canary pipeline with CanaryConfig provided
*CanaryJudgeResultControllerApi* | [**DeleteCanaryJudgeResultUsingDELETE**](docs/CanaryJudgeResultControllerApi.md#deletecanaryjudgeresultusingdelete) | **Delete** /canaryJudgeResult/{canaryJudgeResultId} | Delete a canary judge result
*CanaryJudgeResultControllerApi* | [**ListAllCanaryJudgeResultsUsingGET**](docs/CanaryJudgeResultControllerApi.md#listallcanaryjudgeresultsusingget) | **Get** /canaryJudgeResult | Retrieve a list of canary judge result ids and timestamps
*CanaryJudgeResultControllerApi* | [**LoadCanaryJudgeResultUsingGET**](docs/CanaryJudgeResultControllerApi.md#loadcanaryjudgeresultusingget) | **Get** /canaryJudgeResult/{canaryJudgeResultId} | Retrieve a canary judge result from object storage
*CanaryJudgeResultControllerApi* | [**StoreCanaryJudgeResultUsingPOST**](docs/CanaryJudgeResultControllerApi.md#storecanaryjudgeresultusingpost) | **Post** /canaryJudgeResult | Write a canary judge result to object storage
*CanaryJudgesControllerApi* | [**ListUsingGET**](docs/CanaryJudgesControllerApi.md#listusingget) | **Get** /judges | Retrieve a list of all configured canary judges
*CredentialsControllerApi* | [**ListUsingGET1**](docs/CredentialsControllerApi.md#listusingget1) | **Get** /credentials | Retrieve a list of all configured credentials
*DatadogFetchControllerApi* | [**QueryMetricsUsingPOST**](docs/DatadogFetchControllerApi.md#querymetricsusingpost) | **Post** /fetch/datadog/query | queryMetrics
*HealthMvcEndpointApi* | [**InvokeUsingGET**](docs/HealthMvcEndpointApi.md#invokeusingget) | **Get** /health | invoke
*MetricSetListControllerApi* | [**DeleteMetricSetListUsingDELETE**](docs/MetricSetListControllerApi.md#deletemetricsetlistusingdelete) | **Delete** /metricSetList/{metricSetListId} | Delete a metric set list
*MetricSetListControllerApi* | [**ListAllMetricSetListsUsingGET**](docs/MetricSetListControllerApi.md#listallmetricsetlistsusingget) | **Get** /metricSetList | Retrieve a list of metric set list ids and timestamps
*MetricSetListControllerApi* | [**LoadMetricSetListUsingGET**](docs/MetricSetListControllerApi.md#loadmetricsetlistusingget) | **Get** /metricSetList/{metricSetListId} | Retrieve a metric set list from object storage
*MetricSetListControllerApi* | [**StoreMetricSetListUsingPOST**](docs/MetricSetListControllerApi.md#storemetricsetlistusingpost) | **Post** /metricSetList | Write a metric set list to object storage
*MetricSetPairListControllerApi* | [**DeleteMetricSetPairListUsingDELETE**](docs/MetricSetPairListControllerApi.md#deletemetricsetpairlistusingdelete) | **Delete** /metricSetPairList/{metricSetPairListId} | Delete a metric set pair list
*MetricSetPairListControllerApi* | [**ListAllMetricSetPairListsUsingGET**](docs/MetricSetPairListControllerApi.md#listallmetricsetpairlistsusingget) | **Get** /metricSetPairList | Retrieve a list of metric set pair list ids and timestamps
*MetricSetPairListControllerApi* | [**LoadMetricSetPairListUsingGET**](docs/MetricSetPairListControllerApi.md#loadmetricsetpairlistusingget) | **Get** /metricSetPairList/{metricSetPairListId} | Retrieve a metric set pair list from object storage
*MetricSetPairListControllerApi* | [**LoadMetricSetPairUsingGET**](docs/MetricSetPairListControllerApi.md#loadmetricsetpairusingget) | **Get** /metricSetPairList/{metricSetPairListId}/{metricSetPairId} | Retrieve a single metric set pair from a metricSetPairList from object storage
*MetricSetPairListControllerApi* | [**StoreMetricSetPairListUsingPOST**](docs/MetricSetPairListControllerApi.md#storemetricsetpairlistusingpost) | **Post** /metricSetPairList | Write a metric set pair list to object storage
*MetricsServiceMetadataControllerApi* | [**ListMetadataUsingGET**](docs/MetricsServiceMetadataControllerApi.md#listmetadatausingget) | **Get** /metadata/metricsService | Retrieve a list of descriptors for use in populating the canary config ui
*PipelineControllerApi* | [**CancelUsingPUT**](docs/PipelineControllerApi.md#cancelusingput) | **Put** /pipelines/{executionId}/cancel | Cancel a pipeline execution
*PipelineControllerApi* | [**GetPipelineUsingGET**](docs/PipelineControllerApi.md#getpipelineusingget) | **Get** /pipelines/{executionId} | Retrieve a pipeline execution
*PipelineControllerApi* | [**StartUsingPOST**](docs/PipelineControllerApi.md#startusingpost) | **Post** /pipelines/start | Initiate a pipeline execution
*StackdriverFetchControllerApi* | [**QueryMetricsUsingPOST1**](docs/StackdriverFetchControllerApi.md#querymetricsusingpost1) | **Post** /fetch/stackdriver/query | Exercise the Stackdriver Metrics Service directly, without any orchestration or judging


## Documentation For Models

 - [AccountCredentials](docs/AccountCredentials.md)
 - [Artifact](docs/Artifact.md)
 - [AuthenticationDetails](docs/AuthenticationDetails.md)
 - [CanaryAdhocExecutionRequest](docs/CanaryAdhocExecutionRequest.md)
 - [CanaryAnalysisResult](docs/CanaryAnalysisResult.md)
 - [CanaryClassifierConfig](docs/CanaryClassifierConfig.md)
 - [CanaryClassifierThresholdsConfig](docs/CanaryClassifierThresholdsConfig.md)
 - [CanaryConfig](docs/CanaryConfig.md)
 - [CanaryConfigUpdateResponse](docs/CanaryConfigUpdateResponse.md)
 - [CanaryExecutionRequest](docs/CanaryExecutionRequest.md)
 - [CanaryExecutionResponse](docs/CanaryExecutionResponse.md)
 - [CanaryExecutionStatusResponse](docs/CanaryExecutionStatusResponse.md)
 - [CanaryJudge](docs/CanaryJudge.md)
 - [CanaryJudgeConfig](docs/CanaryJudgeConfig.md)
 - [CanaryJudgeGroupScore](docs/CanaryJudgeGroupScore.md)
 - [CanaryJudgeResult](docs/CanaryJudgeResult.md)
 - [CanaryJudgeScore](docs/CanaryJudgeScore.md)
 - [CanaryMetricConfig](docs/CanaryMetricConfig.md)
 - [CanaryMetricSetQueryConfig](docs/CanaryMetricSetQueryConfig.md)
 - [CanaryResult](docs/CanaryResult.md)
 - [CanaryScope](docs/CanaryScope.md)
 - [CanaryScopePair](docs/CanaryScopePair.md)
 - [Duration](docs/Duration.md)
 - [Execution](docs/Execution.md)
 - [ExpectedArtifact](docs/ExpectedArtifact.md)
 - [LastModifiedDetails](docs/LastModifiedDetails.md)
 - [Mapstringobject](docs/Mapstringobject.md)
 - [Mapstringstring](docs/Mapstringstring.md)
 - [Metadata](docs/Metadata.md)
 - [MetricSet](docs/MetricSet.md)
 - [MetricSetPair](docs/MetricSetPair.md)
 - [MetricSetScope](docs/MetricSetScope.md)
 - [ModelMap](docs/ModelMap.md)
 - [PausedDetails](docs/PausedDetails.md)
 - [Serializable](docs/Serializable.md)
 - [Stage](docs/Stage.md)
 - [Task](docs/Task.md)
 - [TemporalUnit](docs/TemporalUnit.md)
 - [Trigger](docs/Trigger.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author


