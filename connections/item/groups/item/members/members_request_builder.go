package members

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc "github.com/microsoftgraph/msgraph-sdk-go/models/externalconnectors"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    if6edf7b060afe68bdd05a688cee858c25b08120e274e71d7ffdeeefc35ab0ad5 "github.com/microsoftgraph/msgraph-sdk-go/connections/item/groups/item/members/count"
)

// MembersRequestBuilder provides operations to manage the members property of the microsoft.graph.externalConnectors.externalGroup entity.
type MembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MembersRequestBuilderGetQueryParameters a member added to an externalGroup. You can add Azure Active Directory users, Azure Active Directory groups, or other externalGroups as members.
type MembersRequestBuilderGetQueryParameters struct {
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
// MembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MembersRequestBuilderGetQueryParameters
}
// MembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembersRequestBuilder) {
    m := &MembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/connections/{externalConnection%2Did}/groups/{externalGroup%2Did}/members{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMembersRequestBuilder instantiates a new MembersRequestBuilder and sets the default values.
func NewMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *MembersRequestBuilder) Count()(*if6edf7b060afe68bdd05a688cee858c25b08120e274e71d7ffdeeefc35ab0ad5.CountRequestBuilder) {
    return if6edf7b060afe68bdd05a688cee858c25b08120e274e71d7ffdeeefc35ab0ad5.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation a member added to an externalGroup. You can add Azure Active Directory users, Azure Active Directory groups, or other externalGroups as members.
func (m *MembersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration a member added to an externalGroup. You can add Azure Active Directory users, Azure Active Directory groups, or other externalGroups as members.
func (m *MembersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to members for connections
func (m *MembersRequestBuilder) CreatePostRequestInformation(body i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Identityable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to members for connections
func (m *MembersRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Identityable, requestConfiguration *MembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get a member added to an externalGroup. You can add Azure Active Directory users, Azure Active Directory groups, or other externalGroups as members.
func (m *MembersRequestBuilder) Get()(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.IdentityCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler a member added to an externalGroup. You can add Azure Active Directory users, Azure Active Directory groups, or other externalGroups as members.
func (m *MembersRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *MembersRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.IdentityCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.CreateIdentityCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.IdentityCollectionResponseable), nil
}
// Post create new navigation property to members for connections
func (m *MembersRequestBuilder) Post(body i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Identityable)(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Identityable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to members for connections
func (m *MembersRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Identityable, requestConfiguration *MembersRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Identityable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.CreateIdentityFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.Identityable), nil
}
