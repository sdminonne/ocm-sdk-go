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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// CloudRegionsServer represents the interface the manages the 'cloud_regions' resource.
type CloudRegionsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of regions of the cloud provider.
	//
	// IMPORTANT: This collection doesn't currently support paging or searching, so the returned
	// `page` will always be 1 and `size` and `total` will always be the total number of regions
	// of the provider.
	List(ctx context.Context, request *CloudRegionsListServerRequest, response *CloudRegionsListServerResponse) error

	// Region returns the target 'cloud_region' server for the given identifier.
	//
	// Reference to the service that manages an specific region.
	Region(id string) CloudRegionServer
}

// CloudRegionsListServerRequest is the request for the 'list' method.
type CloudRegionsListServerRequest struct {
}

// CloudRegionsListServerResponse is the response for the 'list' method.
type CloudRegionsListServerResponse struct {
	status int
	err    *errors.Error
	items  *CloudRegionList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of cloud providers.
func (r *CloudRegionsListServerResponse) Items(value *CloudRegionList) *CloudRegionsListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the returned page, where one corresponds to the first page. As this
// collection doesn't support paging the result will always be `1`.
func (r *CloudRegionsListServerResponse) Page(value int) *CloudRegionsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Number of items that will be contained in the returned page. As this collection
// doesn't support paging or searching the result will always be the total number of
// regions of the provider.
func (r *CloudRegionsListServerResponse) Size(value int) *CloudRegionsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page. As this collection doesn't support paging or
// searching the result will always be the total number of regions of the provider.
func (r *CloudRegionsListServerResponse) Total(value int) *CloudRegionsListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *CloudRegionsListServerResponse) Status(value int) *CloudRegionsListServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *CloudRegionsListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(cloudRegionsListServerResponseData)
	data.Items, err = r.items.wrap()
	if err != nil {
		return err
	}
	data.Page = r.page
	data.Size = r.size
	data.Total = r.total
	err = encoder.Encode(data)
	return err
}

// cloudRegionsListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type cloudRegionsListServerResponseData struct {
	Items cloudRegionListData "json:\"items,omitempty\""
	Page  *int                "json:\"page,omitempty\""
	Size  *int                "json:\"size,omitempty\""
	Total *int                "json:\"total,omitempty\""
}

// CloudRegionsAdapter represents the structs that adapts Requests and Response to internal
// structs.
type CloudRegionsAdapter struct {
	server CloudRegionsServer
	router *mux.Router
}

func NewCloudRegionsAdapter(server CloudRegionsServer, router *mux.Router) *CloudRegionsAdapter {
	adapter := new(CloudRegionsAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}").HandlerFunc(adapter.regionHandler)
	adapter.router.Methods(http.MethodGet).Path("").HandlerFunc(adapter.handlerList)
	return adapter
}
func (a *CloudRegionsAdapter) regionHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.Region(id)
	targetAdapter := NewCloudRegionAdapter(target, a.router.PathPrefix("/{id}").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *CloudRegionsAdapter) readListRequest(r *http.Request) (*CloudRegionsListServerRequest, error) {
	var err error
	result := new(CloudRegionsListServerRequest)
	return result, err
}
func (a *CloudRegionsAdapter) writeListResponse(w http.ResponseWriter, r *CloudRegionsListServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *CloudRegionsAdapter) handlerList(w http.ResponseWriter, r *http.Request) {
	request, err := a.readListRequest(r)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to read request from client: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
		return
	}
	response := new(CloudRegionsListServerResponse)
	response.status = http.StatusOK
	err = a.server.List(r.Context(), request, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to run method List: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
	err = a.writeListResponse(w, response)
	if err != nil {
		reason := fmt.Sprintf(
			"An error occurred while trying to write response for client: %v",
			err,
		)
		body, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, body)
	}
}
func (a *CloudRegionsAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}