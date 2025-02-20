package owneddevices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i355f16bd42aea358e8e430601708832725117055ecb49d88907fc3b331872ed9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/owneddevices/endpoint"
    i55394c7779df4b8da3a51e8d1f83fca12c066072829318f80daad28c0cc0954d "github.com/microsoftgraph/msgraph-sdk-go/users/item/owneddevices/count"
    i7690090169a34cb9e7600b834812a49bef81e78171b44454312003371a939455 "github.com/microsoftgraph/msgraph-sdk-go/users/item/owneddevices/device"
    i89097c36c495d7597e64b82c8566f859a6197db2d13442b71c501bdd7dab17af "github.com/microsoftgraph/msgraph-sdk-go/users/item/owneddevices/approleassignment"
)

// OwnedDevicesRequestBuilder provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
type OwnedDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OwnedDevicesRequestBuilderGetQueryParameters devices that are owned by the user. Read-only. Nullable. Supports $expand.
type OwnedDevicesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// OwnedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OwnedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OwnedDevicesRequestBuilderGetQueryParameters
}
// AppRoleAssignment the appRoleAssignment property
func (m *OwnedDevicesRequestBuilder) AppRoleAssignment()(*i89097c36c495d7597e64b82c8566f859a6197db2d13442b71c501bdd7dab17af.AppRoleAssignmentRequestBuilder) {
    return i89097c36c495d7597e64b82c8566f859a6197db2d13442b71c501bdd7dab17af.NewAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOwnedDevicesRequestBuilderInternal instantiates a new OwnedDevicesRequestBuilder and sets the default values.
func NewOwnedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnedDevicesRequestBuilder) {
    m := &OwnedDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/ownedDevices{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOwnedDevicesRequestBuilder instantiates a new OwnedDevicesRequestBuilder and sets the default values.
func NewOwnedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOwnedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *OwnedDevicesRequestBuilder) Count()(*i55394c7779df4b8da3a51e8d1f83fca12c066072829318f80daad28c0cc0954d.CountRequestBuilder) {
    return i55394c7779df4b8da3a51e8d1f83fca12c066072829318f80daad28c0cc0954d.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation devices that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *OwnedDevicesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration devices that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *OwnedDevicesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OwnedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device the device property
func (m *OwnedDevicesRequestBuilder) Device()(*i7690090169a34cb9e7600b834812a49bef81e78171b44454312003371a939455.DeviceRequestBuilder) {
    return i7690090169a34cb9e7600b834812a49bef81e78171b44454312003371a939455.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Endpoint the endpoint property
func (m *OwnedDevicesRequestBuilder) Endpoint()(*i355f16bd42aea358e8e430601708832725117055ecb49d88907fc3b331872ed9.EndpointRequestBuilder) {
    return i355f16bd42aea358e8e430601708832725117055ecb49d88907fc3b331872ed9.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get devices that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *OwnedDevicesRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler devices that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *OwnedDevicesRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *OwnedDevicesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
