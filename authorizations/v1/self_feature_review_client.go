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

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// SelfFeatureReviewClient is the client of the 'self_feature_review' resource.
//
// Manages feature review
type SelfFeatureReviewClient struct {
	transport http.RoundTripper
	path      string
}

// NewSelfFeatureReviewClient creates a new client for the 'self_feature_review'
// resource using the given transport to send the requests and receive the
// responses.
func NewSelfFeatureReviewClient(transport http.RoundTripper, path string) *SelfFeatureReviewClient {
	return &SelfFeatureReviewClient{
		transport: transport,
		path:      path,
	}
}

// Post creates a request for the 'post' method.
//
// Reviews ability to toggle a feature
func (c *SelfFeatureReviewClient) Post() *SelfFeatureReviewPostRequest {
	return &SelfFeatureReviewPostRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// SelfFeatureReviewPostRequest is the request for the 'post' method.
type SelfFeatureReviewPostRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	request   *SelfFeatureReviewRequest
}

// Parameter adds a query parameter.
func (r *SelfFeatureReviewPostRequest) Parameter(name string, value interface{}) *SelfFeatureReviewPostRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *SelfFeatureReviewPostRequest) Header(name string, value interface{}) *SelfFeatureReviewPostRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Request sets the value of the 'request' parameter.
//
//
func (r *SelfFeatureReviewPostRequest) Request(value *SelfFeatureReviewRequest) *SelfFeatureReviewPostRequest {
	r.request = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *SelfFeatureReviewPostRequest) Send() (result *SelfFeatureReviewPostResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *SelfFeatureReviewPostRequest) SendContext(ctx context.Context) (result *SelfFeatureReviewPostResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := &bytes.Buffer{}
	err = writeSelfFeatureReviewPostRequest(r, buffer)
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
	result = &SelfFeatureReviewPostResponse{}
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
	err = readSelfFeatureReviewPostResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// SelfFeatureReviewPostResponse is the response for the 'post' method.
type SelfFeatureReviewPostResponse struct {
	status  int
	header  http.Header
	err     *errors.Error
	request *SelfFeatureReviewResponse
}

// Status returns the response status code.
func (r *SelfFeatureReviewPostResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *SelfFeatureReviewPostResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *SelfFeatureReviewPostResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Request returns the value of the 'request' parameter.
//
//
func (r *SelfFeatureReviewPostResponse) Request() *SelfFeatureReviewResponse {
	if r == nil {
		return nil
	}
	return r.request
}

// GetRequest returns the value of the 'request' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *SelfFeatureReviewPostResponse) GetRequest() (value *SelfFeatureReviewResponse, ok bool) {
	ok = r != nil && r.request != nil
	if ok {
		value = r.request
	}
	return
}
