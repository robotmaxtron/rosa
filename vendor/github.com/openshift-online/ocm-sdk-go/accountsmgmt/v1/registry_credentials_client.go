/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// RegistryCredentialsClient is the client of the 'registry_credentials' resource.
//
// Manages the collection of registry credentials.
type RegistryCredentialsClient struct {
	transport http.RoundTripper
	path      string
}

// NewRegistryCredentialsClient creates a new client for the 'registry_credentials'
// resource using the given transport to send the requests and receive the
// responses.
func NewRegistryCredentialsClient(transport http.RoundTripper, path string) *RegistryCredentialsClient {
	return &RegistryCredentialsClient{
		transport: transport,
		path:      path,
	}
}

// Add creates a request for the 'add' method.
//
// Creates a new registry credential.
func (c *RegistryCredentialsClient) Add() *RegistryCredentialsAddRequest {
	return &RegistryCredentialsAddRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// List creates a request for the 'list' method.
//
// Retrieves the list of accounts.
func (c *RegistryCredentialsClient) List() *RegistryCredentialsListRequest {
	return &RegistryCredentialsListRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// RegistryCredential returns the target 'registry_credential' resource for the given identifier.
//
// Reference to the service that manages an specific registry credential.
func (c *RegistryCredentialsClient) RegistryCredential(id string) *RegistryCredentialClient {
	return NewRegistryCredentialClient(
		c.transport,
		path.Join(c.path, id),
	)
}

// RegistryCredentialsAddRequest is the request for the 'add' method.
type RegistryCredentialsAddRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *RegistryCredential
}

// Parameter adds a query parameter.
func (r *RegistryCredentialsAddRequest) Parameter(name string, value interface{}) *RegistryCredentialsAddRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *RegistryCredentialsAddRequest) Header(name string, value interface{}) *RegistryCredentialsAddRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Impersonate wraps requests on behalf of another user.
// Note: Services that do not support this feature may silently ignore this call.
func (r *RegistryCredentialsAddRequest) Impersonate(user string) *RegistryCredentialsAddRequest {
	helpers.AddImpersonationHeader(&r.header, user)
	return r
}

// Body sets the value of the 'body' parameter.
//
// Registry credential data.
func (r *RegistryCredentialsAddRequest) Body(value *RegistryCredential) *RegistryCredentialsAddRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *RegistryCredentialsAddRequest) Send() (result *RegistryCredentialsAddResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *RegistryCredentialsAddRequest) SendContext(ctx context.Context) (result *RegistryCredentialsAddResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := &bytes.Buffer{}
	err = writeRegistryCredentialsAddRequest(r, buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "POST",
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &RegistryCredentialsAddResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	reader := bufio.NewReader(response.Body)
	_, err = reader.Peek(1)
	if err == io.EOF {
		err = nil
		return
	}
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readRegistryCredentialsAddResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// RegistryCredentialsAddResponse is the response for the 'add' method.
type RegistryCredentialsAddResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *RegistryCredential
}

// Status returns the response status code.
func (r *RegistryCredentialsAddResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *RegistryCredentialsAddResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *RegistryCredentialsAddResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
// Registry credential data.
func (r *RegistryCredentialsAddResponse) Body() *RegistryCredential {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Registry credential data.
func (r *RegistryCredentialsAddResponse) GetBody() (value *RegistryCredential, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// RegistryCredentialsListRequest is the request for the 'list' method.
type RegistryCredentialsListRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	order     *string
	page      *int
	search    *string
	size      *int
}

// Parameter adds a query parameter.
func (r *RegistryCredentialsListRequest) Parameter(name string, value interface{}) *RegistryCredentialsListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *RegistryCredentialsListRequest) Header(name string, value interface{}) *RegistryCredentialsListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Impersonate wraps requests on behalf of another user.
// Note: Services that do not support this feature may silently ignore this call.
func (r *RegistryCredentialsListRequest) Impersonate(user string) *RegistryCredentialsListRequest {
	helpers.AddImpersonationHeader(&r.header, user)
	return r
}

// Order sets the value of the 'order' parameter.
//
// Order criteria.
//
// The syntax of this parameter is similar to the syntax of the _order by_ clause of
// a SQL statement. For example, in order to sort the
// RegistryCredentials descending by username the value should be:
//
// ```sql
// username desc
// ```
//
// If the parameter isn't provided, or if the value is empty, then the order of the
// results is undefined.
func (r *RegistryCredentialsListRequest) Order(value string) *RegistryCredentialsListRequest {
	r.order = &value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *RegistryCredentialsListRequest) Page(value int) *RegistryCredentialsListRequest {
	r.page = &value
	return r
}

// Search sets the value of the 'search' parameter.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the _where_ clause of a
// SQL statement, but using the names of the attributes of the RegistryCredentials instead
// of the names of the columns of a table. For example, in order to retrieve all the
// RegistryCredentials for a user the value should be:
//
// ```sql
// username = 'abcxyz...'
// ```
//
// If the parameter isn't provided, or if the value is empty, then all the
// RegistryCredentials that the user has permission to see will be returned.
func (r *RegistryCredentialsListRequest) Search(value string) *RegistryCredentialsListRequest {
	r.search = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
func (r *RegistryCredentialsListRequest) Size(value int) *RegistryCredentialsListRequest {
	r.size = &value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *RegistryCredentialsListRequest) Send() (result *RegistryCredentialsListResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *RegistryCredentialsListRequest) SendContext(ctx context.Context) (result *RegistryCredentialsListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	if r.order != nil {
		helpers.AddValue(&query, "order", *r.order)
	}
	if r.page != nil {
		helpers.AddValue(&query, "page", *r.page)
	}
	if r.search != nil {
		helpers.AddValue(&query, "search", *r.search)
	}
	if r.size != nil {
		helpers.AddValue(&query, "size", *r.size)
	}
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "GET",
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &RegistryCredentialsListResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	reader := bufio.NewReader(response.Body)
	_, err = reader.Peek(1)
	if err == io.EOF {
		err = nil
		return
	}
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readRegistryCredentialsListResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// RegistryCredentialsListResponse is the response for the 'list' method.
type RegistryCredentialsListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	items  *RegistryCredentialList
	page   *int
	size   *int
	total  *int
}

// Status returns the response status code.
func (r *RegistryCredentialsListResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *RegistryCredentialsListResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *RegistryCredentialsListResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of registry credentials.
func (r *RegistryCredentialsListResponse) Items() *RegistryCredentialList {
	if r == nil {
		return nil
	}
	return r.items
}

// GetItems returns the value of the 'items' parameter and
// a flag indicating if the parameter has a value.
//
// Retrieved list of registry credentials.
func (r *RegistryCredentialsListResponse) GetItems() (value *RegistryCredentialList, ok bool) {
	ok = r != nil && r.items != nil
	if ok {
		value = r.items
	}
	return
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *RegistryCredentialsListResponse) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
func (r *RegistryCredentialsListResponse) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
func (r *RegistryCredentialsListResponse) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Maximum number of items that will be contained in the returned page.
func (r *RegistryCredentialsListResponse) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *RegistryCredentialsListResponse) Total() int {
	if r != nil && r.total != nil {
		return *r.total
	}
	return 0
}

// GetTotal returns the value of the 'total' parameter and
// a flag indicating if the parameter has a value.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *RegistryCredentialsListResponse) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}
