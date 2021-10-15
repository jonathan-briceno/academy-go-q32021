package repositories

import "testing"

func TestExternalApi(t *testing.T) {

	result, er := CallExternalApi()

	if result.Fact == "" || er != nil {
		t.Errorf("error in call")
	}
}
