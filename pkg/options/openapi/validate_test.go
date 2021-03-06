// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openapi

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-cluster-bundle/pkg/converter"
	"github.com/GoogleCloudPlatform/k8s-cluster-bundle/pkg/testutil"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		desc         string
		object       string
		schema       string
		expErrSubstr string
	}{
		{
			desc:   "success",
			object: "foo: derp",
			schema: `
properties:
  foo:
    type: string
`,
		},
		{
			desc: "success: extra params",
			object: `
foo: derp
bar: dorp
`,
			schema: `
properties:
  foo:
    type: string
`,
		},
		{
			desc:   "fail: bad type",
			object: "foo: derp",
			schema: `
properties:
  foo:
    type: int64
`,
			expErrSubstr: "validation failure",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			obj, err := converter.FromYAMLString(tc.object).ToJSONMap()
			if err != nil {
				t.Fatal(err)
			}
			schema := &apiextv1beta1.JSONSchemaProps{}
			err = converter.FromYAMLString(tc.schema).ToObject(schema)
			if err != nil {
				t.Fatal(err)
			}
			res, err := ValidateOptions(obj, schema)

			cerr := testutil.CheckErrorCases(err, tc.expErrSubstr)
			if cerr != nil {
				t.Fatal(cerr)
			}
			if err != nil {
				return
			}

			if res == nil {
				t.Error("result was nil in non-error case")
			}
		})
	}
}
