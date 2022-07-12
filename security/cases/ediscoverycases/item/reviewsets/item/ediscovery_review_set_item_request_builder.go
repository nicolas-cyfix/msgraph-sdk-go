package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
    i71a807439fe726577969ea482cb6fabb1e281b489f397bd37ccd6296ea3f7ab1 "github.com/microsoftgraph/msgraph-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/addtoreviewset"
    ie1a6e6e14b2692558465ec1193db07dbecbf71856c138a59e1967fcc3e0891b4 "github.com/microsoftgraph/msgraph-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/queries"
    i98df3683835786b49bff9f6c83edbd31f24eccfb5acde40aa98bd405ca92c0b4 "github.com/microsoftgraph/msgraph-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/queries/item"
)

// EdiscoveryReviewSetItemRequestBuilder provides operations to manage the reviewSets property of the microsoft.graph.security.ediscoveryCase entity.
type EdiscoveryReviewSetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryReviewSetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryReviewSetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryReviewSetItemRequestBuilderGetQueryParameters returns a list of eDiscoveryReviewSet objects in the case.
type EdiscoveryReviewSetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryReviewSetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryReviewSetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryReviewSetItemRequestBuilderGetQueryParameters
}
// EdiscoveryReviewSetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryReviewSetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddToReviewSet the addToReviewSet property
func (m *EdiscoveryReviewSetItemRequestBuilder) AddToReviewSet()(*i71a807439fe726577969ea482cb6fabb1e281b489f397bd37ccd6296ea3f7ab1.AddToReviewSetRequestBuilder) {
    return i71a807439fe726577969ea482cb6fabb1e281b489f397bd37ccd6296ea3f7ab1.NewAddToReviewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEdiscoveryReviewSetItemRequestBuilderInternal instantiates a new EdiscoveryReviewSetItemRequestBuilder and sets the default values.
func NewEdiscoveryReviewSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryReviewSetItemRequestBuilder) {
    m := &EdiscoveryReviewSetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdiscoveryReviewSetItemRequestBuilder instantiates a new EdiscoveryReviewSetItemRequestBuilder and sets the default values.
func NewEdiscoveryReviewSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryReviewSetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryReviewSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reviewSets for security
func (m *EdiscoveryReviewSetItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property reviewSets for security
func (m *EdiscoveryReviewSetItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EdiscoveryReviewSetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation returns a list of eDiscoveryReviewSet objects in the case.
func (m *EdiscoveryReviewSetItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration returns a list of eDiscoveryReviewSet objects in the case.
func (m *EdiscoveryReviewSetItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EdiscoveryReviewSetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property reviewSets in security
func (m *EdiscoveryReviewSetItemRequestBuilder) CreatePatchRequestInformation(body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property reviewSets in security
func (m *EdiscoveryReviewSetItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetable, requestConfiguration *EdiscoveryReviewSetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property reviewSets for security
func (m *EdiscoveryReviewSetItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property reviewSets for security
func (m *EdiscoveryReviewSetItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *EdiscoveryReviewSetItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get returns a list of eDiscoveryReviewSet objects in the case.
func (m *EdiscoveryReviewSetItemRequestBuilder) Get()(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler returns a list of eDiscoveryReviewSet objects in the case.
func (m *EdiscoveryReviewSetItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *EdiscoveryReviewSetItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryReviewSetFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetable), nil
}
// Patch update the navigation property reviewSets in security
func (m *EdiscoveryReviewSetItemRequestBuilder) Patch(body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property reviewSets in security
func (m *EdiscoveryReviewSetItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetable, requestConfiguration *EdiscoveryReviewSetItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Queries the queries property
func (m *EdiscoveryReviewSetItemRequestBuilder) Queries()(*ie1a6e6e14b2692558465ec1193db07dbecbf71856c138a59e1967fcc3e0891b4.QueriesRequestBuilder) {
    return ie1a6e6e14b2692558465ec1193db07dbecbf71856c138a59e1967fcc3e0891b4.NewQueriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// QueriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.security.cases.ediscoveryCases.item.reviewSets.item.queries.item collection
func (m *EdiscoveryReviewSetItemRequestBuilder) QueriesById(id string)(*i98df3683835786b49bff9f6c83edbd31f24eccfb5acde40aa98bd405ca92c0b4.EdiscoveryReviewSetQueryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryReviewSetQuery%2Did"] = id
    }
    return i98df3683835786b49bff9f6c83edbd31f24eccfb5acde40aa98bd405ca92c0b4.NewEdiscoveryReviewSetQueryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
