// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package log

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	redactedValue  = "[redacted]"
	redactedFields = []string{
		"auth_",
		"password",
		"token",
	}
)

// RedactJSON removes sensitive data from JSON structures
func RedactJSON(data []byte) (res []byte, err error) {
	var jsonBlob any
	err = json.Unmarshal(data, &jsonBlob)
	if err != nil {
		return data, err
	}
	redactValue(&jsonBlob)

	return json.Marshal(jsonBlob)
}

// blatently copied from https://github.com/cloudfoundry/lager/blob/master/json_redacter.go#L45
func redactValue(data *any) any {
	if data == nil {
		return data
	}

	if a, ok := (*data).([]any); ok {
		redactArray(&a)
	} else if m, ok := (*data).(map[string]any); ok {
		redactObject(&m)
	} else if s, ok := (*data).(string); ok {
		for _, prohibited := range redactedFields {
			if strings.Contains(strings.ToLower(fmt.Sprintf("%v", s)), prohibited) {
				(*data) = redactedValue
				continue
			}
		}
	}
	return (*data)
}

func redactArray(data *[]any) {
	for i := range *data {
		redactValue(&((*data)[i]))
	}
}

func redactObject(data *map[string]any) {
	for k, v := range *data {
		for _, prohibited := range redactedFields {
			if strings.Contains(strings.ToLower(fmt.Sprintf("%v", k)), prohibited) {
				(*data)[k] = redactedValue
				continue
			}
		}

		if (*data)[k] != redactedValue {
			//TODO: refactor
			//nolint:gosec
			(*data)[k] = redactValue(&v)
		}
	}
}
