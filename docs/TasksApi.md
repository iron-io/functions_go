# \TasksApi

All URIs are relative to *https://127.0.0.1:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksGet**](TasksApi.md#TasksGet) | **Get** /tasks | Get next task.


# **TasksGet**
> TasksWrapper TasksGet($n)

Get next task.

Gets the next task in the queue, ready for processing. Titan may return <=n tasks. Consumers should start processing tasks in order. Each returned task is set to `status` \"running\" and `started_at` is set to the current time. No other consumer can retrieve this task.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n** | **int32**| Number of tasks to return. | [optional] [default to 1]

### Return type

[**TasksWrapper**](TasksWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

