package controllers

import "testing"

func TestValidationOfEmptyPostTitle(t *testing.T) {
	var bodyPost postNewAdBody = postNewAdBody{
		"",
		"Body",
		[]string{"Dolor", "site"},
	}

	expectedResult := "field `title` cannot be empty"
	
	result := bodyPost.Validate()
	
	if result.Error() !=  expectedResult {
		t.Errorf("Wrong value, you must spesify the value like %v", expectedResult)
	}
}

func TestValidationOfEmptyPostBody(t *testing.T) {
	var bodyPost postNewAdBody = postNewAdBody{
		"Title",
		"",
		[]string{"Dolor", "site"},
	}

	expectedResult := "field `body` cannot be empty"
	
	result := bodyPost.Validate()
	
	if result.Error() !=  expectedResult {
		t.Errorf("Wrong value, you must spesify the value like %v", expectedResult)
	}
}