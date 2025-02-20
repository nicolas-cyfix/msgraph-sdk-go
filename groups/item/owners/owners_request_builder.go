package owners

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i5d9194bc543b744f49882baf1deba7783f8adf39a671346ee44a908c58264ad5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/application"
    ia466d02a89eb3d9bbe833218240b70a19823a4199a21afd1a1f404a8fdba7453 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/serviceprincipal"
    ia9b32cfcd97ed09c19ec76331c8b226f8dd5bcc3d7fef9a14ef61c2f7f3afd7b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/ref"
    ibc8a35398e3a0b6a21eb36efc279ad0376e9cebe23abf67fb8c6cc48146bfc93 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/group"
    icc6c36c6397f9fb69347da14a089fabf6f2e1565ccd87fb28d2565909b5b806d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/count"
    id7810ec1180f736012f630cb1bea61100276ab3225f9d42378192b69e0caf8ca "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/orgcontact"
    id8fbdcddfb63d171d33d9aaf84732613b8570698cd2b69957b1d36b0af6d9cc1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/device"
    ifc0178cb0ec9e03e9bcb4ec91805fa4dc334f99963059f0458e9d19ac9b0a73e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/owners/user"
)

// OwnersRequestBuilder provides operations to manage the owners property of the microsoft.graph.group entity.
type OwnersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OwnersRequestBuilderGetQueryParameters the owners of the group who can be users or service principals. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
type OwnersRequestBuilderGetQueryParameters struct {
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
// OwnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OwnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OwnersRequestBuilderGetQueryParameters
}
// Application the application property
func (m *OwnersRequestBuilder) Application()(*i5d9194bc543b744f49882baf1deba7783f8adf39a671346ee44a908c58264ad5.ApplicationRequestBuilder) {
    return i5d9194bc543b744f49882baf1deba7783f8adf39a671346ee44a908c58264ad5.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOwnersRequestBuilderInternal instantiates a new OwnersRequestBuilder and sets the default values.
func NewOwnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnersRequestBuilder) {
    m := &OwnersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/owners{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOwnersRequestBuilder instantiates a new OwnersRequestBuilder and sets the default values.
func NewOwnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOwnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *OwnersRequestBuilder) Count()(*icc6c36c6397f9fb69347da14a089fabf6f2e1565ccd87fb28d2565909b5b806d.CountRequestBuilder) {
    return icc6c36c6397f9fb69347da14a089fabf6f2e1565ccd87fb28d2565909b5b806d.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the owners of the group who can be users or service principals. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
func (m *OwnersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the owners of the group who can be users or service principals. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
func (m *OwnersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OwnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OwnersRequestBuilder) Device()(*id8fbdcddfb63d171d33d9aaf84732613b8570698cd2b69957b1d36b0af6d9cc1.DeviceRequestBuilder) {
    return id8fbdcddfb63d171d33d9aaf84732613b8570698cd2b69957b1d36b0af6d9cc1.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the owners of the group who can be users or service principals. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
func (m *OwnersRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the owners of the group who can be users or service principals. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
func (m *OwnersRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *OwnersRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
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
// Group the group property
func (m *OwnersRequestBuilder) Group()(*ibc8a35398e3a0b6a21eb36efc279ad0376e9cebe23abf67fb8c6cc48146bfc93.GroupRequestBuilder) {
    return ibc8a35398e3a0b6a21eb36efc279ad0376e9cebe23abf67fb8c6cc48146bfc93.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *OwnersRequestBuilder) OrgContact()(*id7810ec1180f736012f630cb1bea61100276ab3225f9d42378192b69e0caf8ca.OrgContactRequestBuilder) {
    return id7810ec1180f736012f630cb1bea61100276ab3225f9d42378192b69e0caf8ca.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the ref property
func (m *OwnersRequestBuilder) Ref()(*ia9b32cfcd97ed09c19ec76331c8b226f8dd5bcc3d7fef9a14ef61c2f7f3afd7b.RefRequestBuilder) {
    return ia9b32cfcd97ed09c19ec76331c8b226f8dd5bcc3d7fef9a14ef61c2f7f3afd7b.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *OwnersRequestBuilder) ServicePrincipal()(*ia466d02a89eb3d9bbe833218240b70a19823a4199a21afd1a1f404a8fdba7453.ServicePrincipalRequestBuilder) {
    return ia466d02a89eb3d9bbe833218240b70a19823a4199a21afd1a1f404a8fdba7453.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *OwnersRequestBuilder) User()(*ifc0178cb0ec9e03e9bcb4ec91805fa4dc334f99963059f0458e9d19ac9b0a73e.UserRequestBuilder) {
    return ifc0178cb0ec9e03e9bcb4ec91805fa4dc334f99963059f0458e9d19ac9b0a73e.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
