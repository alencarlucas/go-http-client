package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")

	builder := &clientBuilder{
		headers: commonHeaders,
	}
	client := &httpClient{
		builder: builder,
	}

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("Wen expect 3 headers")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Invalid content type received")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("Invalid user agent received")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("Invalid request id received")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("No error expected when passing a nil body")
		}

		if body != nil {
			t.Error("No body expected when passing a nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		type user struct {
			Name string `json:"name"`
			Age  uint   `json:"age"`
		}

		requestBody := user{
			Name: "Jair Inácio",
			Age:  65,
		}
		body, err := client.getRequestBody("application/json", requestBody)

		if err != nil {
			t.Error("No error expected when marshaling struct as json")
		}

		if string(body) != `{"name":"Jair Inácio","age":65}` {
			t.Error("Invalid json body received")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		type user struct {
			Name string `xml:"name"`
			Age  uint   `xml:"age"`
		}

		requestBody := user{
			Name: "Jair Inácio",
			Age:  65,
		}

		body, err := client.getRequestBody("application/xml", requestBody)

		if err != nil {
			t.Error("No error expected when marshaling struct as xml")
		}

		if string(body) != `<user><name>Jair Inácio</name><age>65</age></user>` {
			t.Error("Invalid xml body received")
		}
	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
		type user struct {
			Name string `json:"name"`
			Age  uint   `json:"age"`
		}

		requestBody := user{
			Name: "Jair Inácio",
			Age:  65,
		}
		body, err := client.getRequestBody("invalid-application-type", requestBody)

		if err != nil {
			t.Error("No error expected when marshaling struct as json")
		}

		if string(body) != `{"name":"Jair Inácio","age":65}` {
			t.Error("Invalid json body received")
		}
	})

}
