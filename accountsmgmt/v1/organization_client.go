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
	"time"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// OrganizationClient is the client of the 'organization' resource.
//
// Manages a specific organization.
type OrganizationClient struct {
	transport http.RoundTripper
	path      string
}

// NewOrganizationClient creates a new client for the 'organization'
// resource using the given transport to send the requests and receive the
// responses.
func NewOrganizationClient(transport http.RoundTripper, path string) *OrganizationClient {
	return &OrganizationClient{
		transport: transport,
		path:      path,
	}
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the organization.
func (c *OrganizationClient) Get() *OrganizationGetRequest {
	return &OrganizationGetRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// Update creates a request for the 'update' method.
//
// Updates the organization.
func (c *OrganizationClient) Update() *OrganizationUpdateRequest {
	return &OrganizationUpdateRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// Labels returns the target 'generic_labels' resource.
//
// Reference to the list of labels of a specific organization.
func (c *OrganizationClient) Labels() *GenericLabelsClient {
	return NewGenericLabelsClient(
		c.transport,
		path.Join(c.path, "labels"),
	)
}

// QuotaCost returns the target 'quota_cost' resource.
//
// Reference to the service that returns a summary of quota cost for this organization
func (c *OrganizationClient) QuotaCost() *QuotaCostClient {
	return NewQuotaCostClient(
		c.transport,
		path.Join(c.path, "quota_cost"),
	)
}

// ResourceQuota returns the target 'resource_quotas' resource.
//
// Reference to the service that manages the resource quotas for this
// organization.
func (c *OrganizationClient) ResourceQuota() *ResourceQuotasClient {
	return NewResourceQuotasClient(
		c.transport,
		path.Join(c.path, "resource_quota"),
	)
}

// SummaryDashboard returns the target 'summary_dashboard' resource.
//
// Reference to the service that manages the resource quotas for this
// organization.
func (c *OrganizationClient) SummaryDashboard() *SummaryDashboardClient {
	return NewSummaryDashboardClient(
		c.transport,
		path.Join(c.path, "summary_dashboard"),
	)
}

// OrganizationPollRequest is the request for the Poll method.
type OrganizationPollRequest struct {
	request    *OrganizationGetRequest
	interval   time.Duration
	statuses   []int
	predicates []func(interface{}) bool
}

// Parameter adds a query parameter to all the requests that will be used to retrieve the object.
func (r *OrganizationPollRequest) Parameter(name string, value interface{}) *OrganizationPollRequest {
	r.request.Parameter(name, value)
	return r
}

// Header adds a request header to all the requests that will be used to retrieve the object.
func (r *OrganizationPollRequest) Header(name string, value interface{}) *OrganizationPollRequest {
	r.request.Header(name, value)
	return r
}

// Interval sets the polling interval. This parameter is mandatory and must be greater than zero.
func (r *OrganizationPollRequest) Interval(value time.Duration) *OrganizationPollRequest {
	r.interval = value
	return r
}

// Status set the expected status of the response. Multiple values can be set calling this method
// multiple times. The response will be considered successful if the status is any of those values.
func (r *OrganizationPollRequest) Status(value int) *OrganizationPollRequest {
	r.statuses = append(r.statuses, value)
	return r
}

// Predicate adds a predicate that the response should satisfy be considered successful. Multiple
// predicates can be set calling this method multiple times. The response will be considered successful
// if all the predicates are satisfied.
func (r *OrganizationPollRequest) Predicate(value func(*OrganizationGetResponse) bool) *OrganizationPollRequest {
	r.predicates = append(r.predicates, func(response interface{}) bool {
		return value(response.(*OrganizationGetResponse))
	})
	return r
}

// StartContext starts the polling loop. Responses will be considered successful if the status is one of
// the values specified with the Status method and if all the predicates specified with the Predicate
// method return nil.
//
// The context must have a timeout or deadline, otherwise this method will immediately return an error.
func (r *OrganizationPollRequest) StartContext(ctx context.Context) (response *OrganizationPollResponse, err error) {
	result, err := helpers.PollContext(ctx, r.interval, r.statuses, r.predicates, r.task)
	if result != nil {
		response = &OrganizationPollResponse{
			response: result.(*OrganizationGetResponse),
		}
	}
	return
}

// task adapts the types of the request/response types so that they can be used with the generic
// polling function from the helpers package.
func (r *OrganizationPollRequest) task(ctx context.Context) (status int, result interface{}, err error) {
	response, err := r.request.SendContext(ctx)
	if response != nil {
		status = response.Status()
		result = response
	}
	return
}

// OrganizationPollResponse is the response for the Poll method.
type OrganizationPollResponse struct {
	response *OrganizationGetResponse
}

// Status returns the response status code.
func (r *OrganizationPollResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.response.Status()
}

// Header returns header of the response.
func (r *OrganizationPollResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.response.Header()
}

// Error returns the response error.
func (r *OrganizationPollResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.response.Error()
}

// Body returns the value of the 'body' parameter.
//
//
func (r *OrganizationPollResponse) Body() *Organization {
	return r.response.Body()
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *OrganizationPollResponse) GetBody() (value *Organization, ok bool) {
	return r.response.GetBody()
}

// Poll creates a request to repeatedly retrieve the object till the response has one of a given set
// of states and satisfies a set of predicates.
func (c *OrganizationClient) Poll() *OrganizationPollRequest {
	return &OrganizationPollRequest{
		request: c.Get(),
	}
}

// OrganizationGetRequest is the request for the 'get' method.
type OrganizationGetRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *OrganizationGetRequest) Parameter(name string, value interface{}) *OrganizationGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *OrganizationGetRequest) Header(name string, value interface{}) *OrganizationGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *OrganizationGetRequest) Send() (result *OrganizationGetResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *OrganizationGetRequest) SendContext(ctx context.Context) (result *OrganizationGetResponse, err error) {
	query := helpers.CopyQuery(r.query)
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
	result = &OrganizationGetResponse{}
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
	err = readOrganizationGetResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// OrganizationGetResponse is the response for the 'get' method.
type OrganizationGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Organization
}

// Status returns the response status code.
func (r *OrganizationGetResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *OrganizationGetResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *OrganizationGetResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *OrganizationGetResponse) Body() *Organization {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *OrganizationGetResponse) GetBody() (value *Organization, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// OrganizationUpdateRequest is the request for the 'update' method.
type OrganizationUpdateRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *Organization
}

// Parameter adds a query parameter.
func (r *OrganizationUpdateRequest) Parameter(name string, value interface{}) *OrganizationUpdateRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *OrganizationUpdateRequest) Header(name string, value interface{}) *OrganizationUpdateRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
//
func (r *OrganizationUpdateRequest) Body(value *Organization) *OrganizationUpdateRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *OrganizationUpdateRequest) Send() (result *OrganizationUpdateResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *OrganizationUpdateRequest) SendContext(ctx context.Context) (result *OrganizationUpdateResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := &bytes.Buffer{}
	err = writeOrganizationUpdateRequest(r, buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "PATCH",
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
	result = &OrganizationUpdateResponse{}
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
	err = readOrganizationUpdateResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// OrganizationUpdateResponse is the response for the 'update' method.
type OrganizationUpdateResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Organization
}

// Status returns the response status code.
func (r *OrganizationUpdateResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *OrganizationUpdateResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *OrganizationUpdateResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *OrganizationUpdateResponse) Body() *Organization {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *OrganizationUpdateResponse) GetBody() (value *Organization, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}
