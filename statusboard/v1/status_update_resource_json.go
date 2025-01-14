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

package v1 // github.com/openshift-online/ocm-sdk-go/statusboard/v1

import (
	"io"
	"net/http"
)

func readStatusUpdateDeleteRequest(request *StatusUpdateDeleteServerRequest, r *http.Request) error {
	return nil
}
func writeStatusUpdateDeleteRequest(request *StatusUpdateDeleteRequest, writer io.Writer) error {
	return nil
}
func readStatusUpdateDeleteResponse(response *StatusUpdateDeleteResponse, reader io.Reader) error {
	return nil
}
func writeStatusUpdateDeleteResponse(response *StatusUpdateDeleteServerResponse, w http.ResponseWriter) error {
	return nil
}
func readStatusUpdateGetRequest(request *StatusUpdateGetServerRequest, r *http.Request) error {
	return nil
}
func writeStatusUpdateGetRequest(request *StatusUpdateGetRequest, writer io.Writer) error {
	return nil
}
func readStatusUpdateGetResponse(response *StatusUpdateGetResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalStatus(reader)
	return err
}
func writeStatusUpdateGetResponse(response *StatusUpdateGetServerResponse, w http.ResponseWriter) error {
	return MarshalStatus(response.body, w)
}
func readStatusUpdateUpdateRequest(request *StatusUpdateUpdateServerRequest, r *http.Request) error {
	var err error
	request.body, err = UnmarshalStatus(r.Body)
	return err
}
func writeStatusUpdateUpdateRequest(request *StatusUpdateUpdateRequest, writer io.Writer) error {
	return MarshalStatus(request.body, writer)
}
func readStatusUpdateUpdateResponse(response *StatusUpdateUpdateResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalStatus(reader)
	return err
}
func writeStatusUpdateUpdateResponse(response *StatusUpdateUpdateServerResponse, w http.ResponseWriter) error {
	return MarshalStatus(response.body, w)
}
