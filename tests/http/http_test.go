package http

import (
	"github.com/0xIsRookie/common/http"
	stHttp "net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	//http.UseDNSCache()
	headers := map[string]string{"Content-Type": "application/json"}
	timeout := 10
	ip := ""

	// Test GET request
	resp, err := http.Get("https://www.example.com", headers, timeout)
	if err != nil {
		t.Errorf("Error sending GET request: %v", err)
	}
	if resp.StatusCode != stHttp.StatusOK {
		t.Errorf("Unexpected status code for GET request: %d", resp.StatusCode)
	}

	// Test GET request with local IP
	resp, err = http.GetWithLocalIP("https://www.example.com", headers, timeout, ip)
	if err != nil {
		t.Errorf("Error sending GET request with local IP: %v", err)
	}
	if resp.StatusCode != stHttp.StatusOK {
		t.Errorf("Unexpected status code for GET request with local IP: %d", resp.StatusCode)
	}

	// Test POST request
	data := `{"name": "John Doe", "age": 30}`
	resp, err = http.Post("https://www.example.com", data, headers, timeout)
	if err != nil {
		t.Errorf("Error sending POST request: %v", err)
	}
	if resp.StatusCode != stHttp.StatusOK {
		t.Errorf("Unexpected status code for POST request: %d", resp.StatusCode)
	}

	// Test POST request with local IP
	resp, err = http.PostWithLocalIP("https://www.example.com", data, headers, timeout, ip)
	if err != nil {
		t.Errorf("Error sending POST request with local IP: %v", err)
	}
	if resp.StatusCode != stHttp.StatusOK {
		t.Errorf("Unexpected status code for POST request with local IP: %d", resp.StatusCode)
	}
}
