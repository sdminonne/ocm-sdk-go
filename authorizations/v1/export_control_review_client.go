/*
Copyright (c) 2019 Red Hat, Inc.

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
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// ExportControlReviewClient is the client of the 'export_control_review' resource.
//
// Manages export control review.
type ExportControlReviewClient struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewExportControlReviewClient creates a new client for the 'export_control_review'
// resource using the given transport to sned the requests and receive the
// responses.
func NewExportControlReviewClient(transport http.RoundTripper, path string, metric string) *ExportControlReviewClient {
	client := new(ExportControlReviewClient)
	client.transport = transport
	client.path = path
	client.metric = metric
	return client
}

// Post creates a request for the 'post' method.
//
// Screens a user by account user name.
func (c *ExportControlReviewClient) Post() *ExportControlReviewPostRequest {
	request := new(ExportControlReviewPostRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// ExportControlReviewPostRequest is the request for the 'post' method.
type ExportControlReviewPostRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
	request   *ExportControlReviewRequest
}

// Parameter adds a query parameter.
func (r *ExportControlReviewPostRequest) Parameter(name string, value interface{}) *ExportControlReviewPostRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *ExportControlReviewPostRequest) Header(name string, value interface{}) *ExportControlReviewPostRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Request sets the value of the 'request' parameter.
//
//
func (r *ExportControlReviewPostRequest) Request(value *ExportControlReviewRequest) *ExportControlReviewPostRequest {
	r.request = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *ExportControlReviewPostRequest) Send() (result *ExportControlReviewPostResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *ExportControlReviewPostRequest) SendContext(ctx context.Context) (result *ExportControlReviewPostResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	buffer := new(bytes.Buffer)
	err = r.marshal(buffer)
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
	result = new(ExportControlReviewPostResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// marshall is the method used internally to marshal requests for the
// 'post' method.
func (r *ExportControlReviewPostRequest) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.request.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ExportControlReviewPostResponse is the response for the 'post' method.
type ExportControlReviewPostResponse struct {
	status   int
	header   http.Header
	err      *errors.Error
	response *ExportControlReviewResponse
}

// Status returns the response status code.
func (r *ExportControlReviewPostResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *ExportControlReviewPostResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *ExportControlReviewPostResponse) Error() *errors.Error {
	return r.err
}

// Response returns the value of the 'response' parameter.
//
//
func (r *ExportControlReviewPostResponse) Response() *ExportControlReviewResponse {
	if r == nil {
		return nil
	}
	return r.response
}

// GetResponse returns the value of the 'response' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *ExportControlReviewPostResponse) GetResponse() (value *ExportControlReviewResponse, ok bool) {
	ok = r != nil && r.response != nil
	if ok {
		value = r.response
	}
	return
}

// unmarshal is the method used internally to unmarshal responses to the
// 'post' method.
func (r *ExportControlReviewPostResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(exportControlReviewResponseData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.response, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}
