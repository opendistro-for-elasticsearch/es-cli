/*
 * Copyright 2021 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package es

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"odfe-cli/entity/es"
	"strings"
)

const (
	HeaderSeparator           = ":"
	MultipleHeaderSeparator   = ";"
	QueryParamSeparator       = "&"
	FileNameIdentifier        = "@"
	PrettyPrintQueryParameter = "pretty=true"
)

//CommandToCurlRequestParameter map user input to Elasticsearch request
func CommandToCurlRequestParameter(request es.CurlCommandRequest) (result es.CurlRequest, err error) {

	if result.Action, err = toHTTPAction(request.Action); err != nil {
		return es.CurlRequest{}, err
	}
	if result.Headers, err = toHTTPHeaders(request.Headers); err != nil {
		return es.CurlRequest{}, err
	}
	if result.Data, err = toCurlPayload(request.Data); err != nil {
		return es.CurlRequest{}, err
	}
	if !isEmpty(request.Path) {
		result.Path = request.Path
	}
	result.QueryParams = request.QueryParams
	if !request.Pretty {
		return
	}
	result.QueryParams = addPrettyToQueryParameter(result.QueryParams)
	return
}

func addPrettyToQueryParameter(path string) string {
	splitValues := strings.Split(path, QueryParamSeparator)
	splitValues = append(splitValues, PrettyPrintQueryParameter)
	return strings.Join(splitValues, QueryParamSeparator)
}

func getSupportedHTTPAction() []string {
	return []string{
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodDelete,
	}
}

func isEmpty(input string) bool {
	if len(input) == 0 {
		return true
	}
	trimSpaceAction := strings.TrimSpace(input)
	return trimSpaceAction == ""
}

func toHTTPAction(action string) (string, error) {
	if isEmpty(action) {
		return "", errors.New("action cannot be empty")
	}
	upperAction := strings.ToUpper(strings.TrimSpace(action))
	for _, verb := range getSupportedHTTPAction() {
		if verb == upperAction {
			return verb, nil
		}
	}
	return "", fmt.Errorf("action: %s is not supported. Supported values are: %v", action, getSupportedHTTPAction())
}

func processHeader(header string) (name string, value string, err error) {
	if isEmpty(header) { // ignore any empty header
		return
	}
	values := strings.Split(header, HeaderSeparator)
	if len(values) != 2 {
		return name, value, fmt.Errorf("invalid header format, received %s but expected is 'name: value'", header)
	}
	name = strings.ToLower(strings.TrimSpace(values[0]))
	value = strings.ToLower(strings.TrimSpace(values[1]))
	return
}

func toHTTPHeaders(headers string) (map[string]string, error) {
	if isEmpty(headers) {
		return nil, nil
	}
	splitHeaders := strings.Split(strings.TrimSpace(headers), MultipleHeaderSeparator)
	httpHeaders := map[string]string{}
	for _, header := range splitHeaders {
		name, value, err := processHeader(header)
		if err != nil {
			return nil, err
		}
		if len(name) > 0 && len(value) > 0 { // will ignore empty header
			httpHeaders[name] = value
		}
	}
	return httpHeaders, nil
}

func toCurlPayload(data string) (payload []byte, err error) {
	if isEmpty(data) {
		return
	}
	// if data is file name, read file contents
	if strings.HasPrefix(data, FileNameIdentifier) && !isEmpty(strings.TrimPrefix(data, FileNameIdentifier)) {
		return ioutil.ReadFile(data[1:])
	}
	// if data is invalid json string
	if !json.Valid([]byte(data)) {
		return nil, fmt.Errorf("invalid data: %s, data can be either valid json or filename with prefix '@'", data)
	}
	return []byte(data), nil
}
