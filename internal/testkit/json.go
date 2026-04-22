package testkit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// AssertJSONEqual decodes expected and actual JSON and compares structures.
func AssertJSONEqual(t *testing.T, expectedJSON string, actual interface{}) {
	t.Helper()

	actualJSON, err := json.Marshal(actual)
	if err != nil {
		t.Fatalf("marshal actual json: %v", err)
	}

	var expectedObj interface{}
	if err := json.NewDecoder(bytes.NewBufferString(expectedJSON)).Decode(&expectedObj); err != nil {
		t.Fatalf("decode expected json: %v", err)
	}

	var actualObj interface{}
	if err := json.NewDecoder(bytes.NewBuffer(actualJSON)).Decode(&actualObj); err != nil {
		t.Fatalf("decode actual json: %v", err)
	}

	if err := compareJSON(expectedObj, actualObj, "$"); err != nil {
		t.Fatalf("json mismatch: %v", err)
	}
}

func compareJSON(expected interface{}, actual interface{}, path string) error {
	if expected == nil {
		return nil
	}

	expectedKind := reflect.TypeOf(expected).Kind()
	actualKind := reflect.TypeOf(actual).Kind()
	if expectedKind != actualKind {
		return fmt.Errorf("%s type mismatch: expected=%s actual=%s", path, expectedKind, actualKind)
	}

	switch exp := expected.(type) {
	case map[string]interface{}:
		act := actual.(map[string]interface{})
		for key, expVal := range exp {
			actVal, ok := act[key]
			if !ok {
				return fmt.Errorf("%s.%s key missing", path, key)
			}
			if err := compareJSON(expVal, actVal, path+"."+key); err != nil {
				return err
			}
		}
	case []interface{}:
		act := actual.([]interface{})
		if len(exp) > len(act) {
			return fmt.Errorf("%s length mismatch: expected>actual (%d>%d)", path, len(exp), len(act))
		}
		for i := range exp {
			if err := compareJSON(exp[i], act[i], fmt.Sprintf("%s[%d]", path, i)); err != nil {
				return err
			}
		}
	default:
		if !reflect.DeepEqual(expected, actual) {
			return fmt.Errorf("%s value mismatch: expected=%v actual=%v", path, expected, actual)
		}
	}

	return nil
}
