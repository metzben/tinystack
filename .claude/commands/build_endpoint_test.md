# Create a test for the following endpoint func $ARGUMENTS in the file $ARGUMENTS

## Please use the following test examples to create a test for the http endpoint:

### Test 1

```go
// go test -v '-run=^TestHandleUserName$'
func TestHandleUserName(t *testing.T) {
	mux := http.NewServeMux()
	router := testApp.BuildRoutes(mux)

	server := httptest.NewServer(mux)
	defer server.Close()
	url := "/v1/users/Mac"

	request, _ := http.NewRequest("GET", url, nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	expectedBody := "yo name is:  Mac\n"
	assert.Equal(t, expectedBody, writer.Body.String())
}
```

### Test 2

```go
// go test -v '-run=^TestHomeEndpoint$'
func TestHomeEndpoint(t *testing.T) {
	// Create a new router
	mux := http.NewServeMux()
	router := testApp.BuildRoutes(mux)

	// Create a test server
	server := httptest.NewServer(mux)
	defer server.Close()

	// Make a request to the home endpoint
	r, _ := http.NewRequest("GET", url.Home, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	// Check status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check response body
	expectedBody := "This is the home route!\n"
	assert.Equal(t, expectedBody, w.Body.String())
}

```

## Instructions:

* Please do not use any external libraries or packages.
* Use pkg/assert/assert.go for assertions. Do not try to use any other packages for assertions.
* Use the httptest package for testing http
* Please use the test examples above to create the test.
* Please verify the status code just like the examples above.
* Please verify the expected body just like the examples above.
