package netsuiteodbc

import (
	"testing"
)

func Test_connStringToParameterMap(t *testing.T) {
	params, err := connStringToParameterMap("DSN=dsn_value;CustomProperties=(property_a=property_a_value;property_b=property_b_value;);Driver=driver_value")
	if err != nil {
		t.Fatalf("expected no error, got: %s", err.Error())
	}

	if dsn, ok := params["DSN"]; !ok || dsn != "dsn_value" {
		t.Errorf("DSN value not expected.  Got: %s", dsn)
	}

	if props, ok := params["CustomProperties"]; !ok || props != "(property_a=property_a_value;property_b=property_b_value;)" {
		t.Errorf("CustomProperties value not expected.  Got: %s", props)
	}

	if driver, ok := params["Driver"]; !ok || driver != "driver_value" {
		t.Errorf("Driver value not expected.  Got: %s", driver)
	}
}
