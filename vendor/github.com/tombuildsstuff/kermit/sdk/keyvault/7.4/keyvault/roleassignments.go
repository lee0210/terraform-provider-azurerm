package keyvault

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// RoleAssignmentsClient is the the key vault client performs cryptographic key operations and vault operations against
// the Key Vault service.
type RoleAssignmentsClient struct {
	BaseClient
}

// NewRoleAssignmentsClient creates an instance of the RoleAssignmentsClient client.
func NewRoleAssignmentsClient() RoleAssignmentsClient {
	return RoleAssignmentsClient{New()}
}

// Create creates a role assignment.
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
// scope - the scope of the role assignment to create.
// roleAssignmentName - the name of the role assignment to create. It can be any valid GUID.
// parameters - parameters for the role assignment.
func (client RoleAssignmentsClient) Create(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters) (result RoleAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.Properties.RoleDefinitionID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.Properties.PrincipalID", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("keyvault.RoleAssignmentsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, vaultBaseURL, scope, roleAssignmentName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client RoleAssignmentsClient) CreatePreparer(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"roleAssignmentName": autorest.Encode("path", roleAssignmentName),
		"scope":              scope,
	}

	const APIVersion = "7.4"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) CreateResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a role assignment.
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
// scope - the scope of the role assignment to delete.
// roleAssignmentName - the name of the role assignment to delete.
func (client RoleAssignmentsClient) Delete(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string) (result RoleAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, vaultBaseURL, scope, roleAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RoleAssignmentsClient) DeletePreparer(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"roleAssignmentName": autorest.Encode("path", roleAssignmentName),
		"scope":              scope,
	}

	const APIVersion = "7.4"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) DeleteResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get the specified role assignment.
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
// scope - the scope of the role assignment.
// roleAssignmentName - the name of the role assignment to get.
func (client RoleAssignmentsClient) Get(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string) (result RoleAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultBaseURL, scope, roleAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RoleAssignmentsClient) GetPreparer(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"roleAssignmentName": autorest.Encode("path", roleAssignmentName),
		"scope":              scope,
	}

	const APIVersion = "7.4"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) GetResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForScope gets role assignments for a scope.
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
// scope - the scope of the role assignments.
// filter - the filter to apply on the operation. Use $filter=atScope() to return all role assignments at or
// above the scope. Use $filter=principalId eq {id} to return all role assignments at, above or below the scope
// for the specified principal.
func (client RoleAssignmentsClient) ListForScope(ctx context.Context, vaultBaseURL string, scope string, filter string) (result RoleAssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.ListForScope")
		defer func() {
			sc := -1
			if result.ralr.Response.Response != nil {
				sc = result.ralr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listForScopeNextResults
	req, err := client.ListForScopePreparer(ctx, vaultBaseURL, scope, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "ListForScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.ralr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "ListForScope", resp, "Failure sending request")
		return
	}

	result.ralr, err = client.ListForScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "ListForScope", resp, "Failure responding to request")
		return
	}
	if result.ralr.hasNextLink() && result.ralr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListForScopePreparer prepares the ListForScope request.
func (client RoleAssignmentsClient) ListForScopePreparer(ctx context.Context, vaultBaseURL string, scope string, filter string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "7.4"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForScopeSender sends the ListForScope request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) ListForScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListForScopeResponder handles the response to the ListForScope request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) ListForScopeResponder(resp *http.Response) (result RoleAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listForScopeNextResults retrieves the next set of results, if any.
func (client RoleAssignmentsClient) listForScopeNextResults(ctx context.Context, lastResults RoleAssignmentListResult) (result RoleAssignmentListResult, err error) {
	req, err := lastResults.roleAssignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "listForScopeNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "listForScopeNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListForScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleAssignmentsClient", "listForScopeNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListForScopeComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleAssignmentsClient) ListForScopeComplete(ctx context.Context, vaultBaseURL string, scope string, filter string) (result RoleAssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.ListForScope")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListForScope(ctx, vaultBaseURL, scope, filter)
	return
}