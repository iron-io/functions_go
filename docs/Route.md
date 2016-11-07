# Route

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | App this route belongs to. | [optional] [default to null]
**Path** | **string** | URL path that will be matched to this route | [optional] [default to null]
**Image** | **string** | Name of Docker image to use in this route. You should include the image tag, which should be a version number, to be more accurate. Can be overridden on a per route basis with route.image. | [optional] [default to null]
**Headers** | **string** | Map of http headers that will be sent with the response | [optional] [default to null]
**Memory** | **int64** | Max usable memory for this route (MiB). | [optional] [default to null]
**Type_** | **string** | Route type | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


