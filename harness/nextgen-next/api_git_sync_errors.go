/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type GitSyncErrorsApiService service

/*
GitSyncErrorsApiService Get Errors Count for the given scope, Repo and Branch
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GitSyncErrorsApiGetGitSyncErrorsCountOpts - Optional Parameters:
     * @param "AccountIdentifier" (optional.String) -  Account Identifier for the entity
     * @param "OrgIdentifier" (optional.String) -  Organization Identifier for the entity
     * @param "ProjectIdentifier" (optional.String) -  Project Identifier for the entity
     * @param "SearchTerm" (optional.String) -  Search Term
     * @param "Branch" (optional.String) -  Branch Name
     * @param "RepoIdentifier" (optional.String) -  Git Sync Config Identifier
     * @param "GetDefaultFromOtherRepo" (optional.Bool) -  if true, return all the default entities
@return ResponseDtoGitSyncErrorCount
*/

type GitSyncErrorsApiGetGitSyncErrorsCountOpts struct {
	AccountIdentifier       optional.String
	OrgIdentifier           optional.String
	ProjectIdentifier       optional.String
	SearchTerm              optional.String
	Branch                  optional.String
	RepoIdentifier          optional.String
	GetDefaultFromOtherRepo optional.Bool
}

func (a *GitSyncErrorsApiService) GetGitSyncErrorsCount(ctx context.Context, localVarOptionals *GitSyncErrorsApiGetGitSyncErrorsCountOpts) (ResponseDtoGitSyncErrorCount, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResponseDtoGitSyncErrorCount
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/git-sync-errors/count"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.AccountIdentifier.IsSet() {
		localVarQueryParams.Add("accountIdentifier", parameterToString(localVarOptionals.AccountIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrgIdentifier.IsSet() {
		localVarQueryParams.Add("orgIdentifier", parameterToString(localVarOptionals.OrgIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProjectIdentifier.IsSet() {
		localVarQueryParams.Add("projectIdentifier", parameterToString(localVarOptionals.ProjectIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchTerm.IsSet() {
		localVarQueryParams.Add("searchTerm", parameterToString(localVarOptionals.SearchTerm.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoIdentifier.IsSet() {
		localVarQueryParams.Add("repoIdentifier", parameterToString(localVarOptionals.RepoIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GetDefaultFromOtherRepo.IsSet() {
		localVarQueryParams.Add("getDefaultFromOtherRepo", parameterToString(localVarOptionals.GetDefaultFromOtherRepo.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml", "text/yaml", "text/html"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoGitSyncErrorCount
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
GitSyncErrorsApiService Lists Git to Harness Errors by file or connectivity errors for the given scope, Repo and Branch
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GitSyncErrorsApiListGitSyncErrorsOpts - Optional Parameters:
     * @param "PageIndex" (optional.Int32) -
     * @param "PageSize" (optional.Int32) -
     * @param "SortOrders" (optional.Interface of []SortOrder) -
     * @param "AccountIdentifier" (optional.String) -  Account Identifier for the entity
     * @param "OrgIdentifier" (optional.String) -  Organization Identifier for the entity
     * @param "ProjectIdentifier" (optional.String) -  Project Identifier for the entity
     * @param "SearchTerm" (optional.String) -  Search Term
     * @param "Branch" (optional.String) -  Branch Name
     * @param "RepoIdentifier" (optional.String) -  Git Sync Config Identifier
     * @param "GetDefaultFromOtherRepo" (optional.Bool) -  if true, return all the default entities
     * @param "GitToHarness" (optional.Bool) -  This specifies which errors to show - (Git to Harness or Connectivity), Put true to show Git to Harness Errors
@return ResponseDtoPageResponseGitSyncError
*/

type GitSyncErrorsApiListGitSyncErrorsOpts struct {
	PageIndex               optional.Int32
	PageSize                optional.Int32
	SortOrders              optional.Interface
	AccountIdentifier       optional.String
	OrgIdentifier           optional.String
	ProjectIdentifier       optional.String
	SearchTerm              optional.String
	Branch                  optional.String
	RepoIdentifier          optional.String
	GetDefaultFromOtherRepo optional.Bool
	GitToHarness            optional.Bool
}

func (a *GitSyncErrorsApiService) ListGitSyncErrors(ctx context.Context, localVarOptionals *GitSyncErrorsApiListGitSyncErrorsOpts) (ResponseDtoPageResponseGitSyncError, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResponseDtoPageResponseGitSyncError
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/git-sync-errors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.PageIndex.IsSet() {
		localVarQueryParams.Add("pageIndex", parameterToString(localVarOptionals.PageIndex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortOrders.IsSet() {
		localVarQueryParams.Add("sortOrders", parameterToString(localVarOptionals.SortOrders.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.AccountIdentifier.IsSet() {
		localVarQueryParams.Add("accountIdentifier", parameterToString(localVarOptionals.AccountIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrgIdentifier.IsSet() {
		localVarQueryParams.Add("orgIdentifier", parameterToString(localVarOptionals.OrgIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProjectIdentifier.IsSet() {
		localVarQueryParams.Add("projectIdentifier", parameterToString(localVarOptionals.ProjectIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchTerm.IsSet() {
		localVarQueryParams.Add("searchTerm", parameterToString(localVarOptionals.SearchTerm.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoIdentifier.IsSet() {
		localVarQueryParams.Add("repoIdentifier", parameterToString(localVarOptionals.RepoIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GetDefaultFromOtherRepo.IsSet() {
		localVarQueryParams.Add("getDefaultFromOtherRepo", parameterToString(localVarOptionals.GetDefaultFromOtherRepo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GitToHarness.IsSet() {
		localVarQueryParams.Add("gitToHarness", parameterToString(localVarOptionals.GitToHarness.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml", "text/yaml", "text/html"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoPageResponseGitSyncError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
GitSyncErrorsApiService Lists Git to Harness Errors for the given Commit Id
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param commitId Commit Id
 * @param optional nil or *GitSyncErrorsApiListGitToHarnessErrorForCommitOpts - Optional Parameters:
     * @param "PageIndex" (optional.Int32) -
     * @param "PageSize" (optional.Int32) -
     * @param "SortOrders" (optional.Interface of []SortOrder) -
     * @param "AccountIdentifier" (optional.String) -  Account Identifier for the entity
     * @param "OrgIdentifier" (optional.String) -  Organization Identifier for the entity
     * @param "ProjectIdentifier" (optional.String) -  Project Identifier for the entity
     * @param "Branch" (optional.String) -  Branch Name
     * @param "RepoIdentifier" (optional.String) -  Git Sync Config Identifier
     * @param "GetDefaultFromOtherRepo" (optional.Bool) -  if true, return all the default entities
@return ResponseDtoPageResponseGitSyncError
*/

type GitSyncErrorsApiListGitToHarnessErrorForCommitOpts struct {
	PageIndex               optional.Int32
	PageSize                optional.Int32
	SortOrders              optional.Interface
	AccountIdentifier       optional.String
	OrgIdentifier           optional.String
	ProjectIdentifier       optional.String
	Branch                  optional.String
	RepoIdentifier          optional.String
	GetDefaultFromOtherRepo optional.Bool
}

func (a *GitSyncErrorsApiService) ListGitToHarnessErrorForCommit(ctx context.Context, commitId string, localVarOptionals *GitSyncErrorsApiListGitToHarnessErrorForCommitOpts) (ResponseDtoPageResponseGitSyncError, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResponseDtoPageResponseGitSyncError
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/git-sync-errors/commits/{commitId}"
	localVarPath = strings.Replace(localVarPath, "{"+"commitId"+"}", fmt.Sprintf("%v", commitId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.PageIndex.IsSet() {
		localVarQueryParams.Add("pageIndex", parameterToString(localVarOptionals.PageIndex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortOrders.IsSet() {
		localVarQueryParams.Add("sortOrders", parameterToString(localVarOptionals.SortOrders.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.AccountIdentifier.IsSet() {
		localVarQueryParams.Add("accountIdentifier", parameterToString(localVarOptionals.AccountIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrgIdentifier.IsSet() {
		localVarQueryParams.Add("orgIdentifier", parameterToString(localVarOptionals.OrgIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProjectIdentifier.IsSet() {
		localVarQueryParams.Add("projectIdentifier", parameterToString(localVarOptionals.ProjectIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoIdentifier.IsSet() {
		localVarQueryParams.Add("repoIdentifier", parameterToString(localVarOptionals.RepoIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GetDefaultFromOtherRepo.IsSet() {
		localVarQueryParams.Add("getDefaultFromOtherRepo", parameterToString(localVarOptionals.GetDefaultFromOtherRepo.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml", "text/yaml", "text/html"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoPageResponseGitSyncError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
GitSyncErrorsApiService Lists Git to Harness Errors grouped by Commits for the given scope, Repo and Branch
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GitSyncErrorsApiListGitToHarnessErrorsGroupedByCommitsOpts - Optional Parameters:
     * @param "PageIndex" (optional.Int32) -
     * @param "PageSize" (optional.Int32) -
     * @param "SortOrders" (optional.Interface of []SortOrder) -
     * @param "AccountIdentifier" (optional.String) -  Account Identifier for the entity
     * @param "OrgIdentifier" (optional.String) -  Organization Identifier for the entity
     * @param "ProjectIdentifier" (optional.String) -  Project Identifier for the entity
     * @param "SearchTerm" (optional.String) -  Search Term
     * @param "Branch" (optional.String) -  Branch Name
     * @param "RepoIdentifier" (optional.String) -  Git Sync Config Identifier
     * @param "GetDefaultFromOtherRepo" (optional.Bool) -  if true, return all the default entities
     * @param "NumberOfErrorsInSummary" (optional.Int32) -  This is the count of errors to be displayed in the summary
@return ResponseDtoPageResponseGitSyncErrorAggregateByCommit
*/

type GitSyncErrorsApiListGitToHarnessErrorsGroupedByCommitsOpts struct {
	PageIndex               optional.Int32
	PageSize                optional.Int32
	SortOrders              optional.Interface
	AccountIdentifier       optional.String
	OrgIdentifier           optional.String
	ProjectIdentifier       optional.String
	SearchTerm              optional.String
	Branch                  optional.String
	RepoIdentifier          optional.String
	GetDefaultFromOtherRepo optional.Bool
	NumberOfErrorsInSummary optional.Int32
}

func (a *GitSyncErrorsApiService) ListGitToHarnessErrorsGroupedByCommits(ctx context.Context, localVarOptionals *GitSyncErrorsApiListGitToHarnessErrorsGroupedByCommitsOpts) (ResponseDtoPageResponseGitSyncErrorAggregateByCommit, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResponseDtoPageResponseGitSyncErrorAggregateByCommit
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/git-sync-errors/aggregate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.PageIndex.IsSet() {
		localVarQueryParams.Add("pageIndex", parameterToString(localVarOptionals.PageIndex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortOrders.IsSet() {
		localVarQueryParams.Add("sortOrders", parameterToString(localVarOptionals.SortOrders.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.AccountIdentifier.IsSet() {
		localVarQueryParams.Add("accountIdentifier", parameterToString(localVarOptionals.AccountIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrgIdentifier.IsSet() {
		localVarQueryParams.Add("orgIdentifier", parameterToString(localVarOptionals.OrgIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProjectIdentifier.IsSet() {
		localVarQueryParams.Add("projectIdentifier", parameterToString(localVarOptionals.ProjectIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchTerm.IsSet() {
		localVarQueryParams.Add("searchTerm", parameterToString(localVarOptionals.SearchTerm.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoIdentifier.IsSet() {
		localVarQueryParams.Add("repoIdentifier", parameterToString(localVarOptionals.RepoIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GetDefaultFromOtherRepo.IsSet() {
		localVarQueryParams.Add("getDefaultFromOtherRepo", parameterToString(localVarOptionals.GetDefaultFromOtherRepo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NumberOfErrorsInSummary.IsSet() {
		localVarQueryParams.Add("numberOfErrorsInSummary", parameterToString(localVarOptionals.NumberOfErrorsInSummary.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml", "text/yaml", "text/html"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoPageResponseGitSyncErrorAggregateByCommit
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}