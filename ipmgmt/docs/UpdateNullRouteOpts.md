# UpdateNullRouteOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomatedUnnullingAt** | Pointer to **time.Time** | The date and time when the null route is to be deactivated. The date and time should be specified using the format yyyy-mm-ddThh:mm:ss±hh:mm (with time zone designator) or yyyy-mm-ddThh:mm:ssZ (UTC). If this field is not present then the null route will not be automatically removed. | [optional] 
**Comment** | Pointer to **string** | A comment to be stored with the null route (e.g. null route reason) | [optional] 
**TicketId** | Pointer to **string** | A reference to be stored with the null route | [optional] 

## Methods

### NewUpdateNullRouteOpts

`func NewUpdateNullRouteOpts() *UpdateNullRouteOpts`

NewUpdateNullRouteOpts instantiates a new UpdateNullRouteOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNullRouteOptsWithDefaults

`func NewUpdateNullRouteOptsWithDefaults() *UpdateNullRouteOpts`

NewUpdateNullRouteOptsWithDefaults instantiates a new UpdateNullRouteOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomatedUnnullingAt

`func (o *UpdateNullRouteOpts) GetAutomatedUnnullingAt() time.Time`

GetAutomatedUnnullingAt returns the AutomatedUnnullingAt field if non-nil, zero value otherwise.

### GetAutomatedUnnullingAtOk

`func (o *UpdateNullRouteOpts) GetAutomatedUnnullingAtOk() (*time.Time, bool)`

GetAutomatedUnnullingAtOk returns a tuple with the AutomatedUnnullingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedUnnullingAt

`func (o *UpdateNullRouteOpts) SetAutomatedUnnullingAt(v time.Time)`

SetAutomatedUnnullingAt sets AutomatedUnnullingAt field to given value.

### HasAutomatedUnnullingAt

`func (o *UpdateNullRouteOpts) HasAutomatedUnnullingAt() bool`

HasAutomatedUnnullingAt returns a boolean if a field has been set.

### GetComment

`func (o *UpdateNullRouteOpts) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateNullRouteOpts) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateNullRouteOpts) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateNullRouteOpts) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetTicketId

`func (o *UpdateNullRouteOpts) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *UpdateNullRouteOpts) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *UpdateNullRouteOpts) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *UpdateNullRouteOpts) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


