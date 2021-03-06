package dbhelpers

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"net/url"
	"strings"
	"testing"
	"unicode"
)

func stripAllWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func TestBuildQuery(t *testing.T) {
	v := url.Values{}
	v.Set("param1", "queryParamv1")
	v.Set("param2", "queryParamv2")

	selectStmt := `SELECT
	t.col1,
	t.col2
FROM table t
`
	// Query Parameters to Database Query column mappings
	// see the fields mapped in the SQL query
	queryParamsToSQLCols := map[string]string{
		"param1": "t.col1",
		"param2": "t.col2",
	}
	query, queryValues := BuildQuery(v, selectStmt, queryParamsToSQLCols)

	actualQuery := stripAllWhitespace(query)

	if queryValues == nil {
		t.Errorf("expected: nil error, actual: %v", queryValues)
	}
	expectedV1 := v.Get("param1")
	actualV1 := queryValues["param1"]
	if expectedV1 != actualV1 {
		t.Errorf("expected: %v error, actual: %v", expectedV1, actualV1)
	}

	if strings.Contains(actualQuery, expectedV1) {
		t.Errorf("expected: %v error, actual: %v", actualQuery, expectedV1)
	}

	expectedV2 := v.Get("param2")
	if strings.Contains(actualQuery, expectedV2) {
		t.Errorf("expected: %v error, actual: %v", actualQuery, expectedV2)
	}

}
