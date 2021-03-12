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

package gateway

import (
	"odfe-cli/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidEndpoint(t *testing.T) {
	t.Run("valid endpoint", func(t *testing.T) {

		profile := entity.Profile{
			Name:     "test1",
			Endpoint: "https://localhost:9200",
			UserName: "foo",
			Password: "bar",
		}
		url, err := GetValidEndpoint(&profile)
		assert.NoError(t, err)
		assert.EqualValues(t, "https://localhost:9200", url.String())
	})
	t.Run("empty endpoint", func(t *testing.T) {
		profile := entity.Profile{
			Name:     "test1",
			Endpoint: "",
			UserName: "foo",
			Password: "bar",
		}
		_, err := GetValidEndpoint(&profile)
		assert.EqualErrorf(t, err, "invalid endpoint:  due to parse \"\": empty url", "failed to get expected error")
	})
}
