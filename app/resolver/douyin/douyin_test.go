package douyin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStreamData_ValidURL_ReturnsStreamData(t *testing.T) {
	// Mock server to simulate a valid response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<script>var state = {"FULL_HD1": "stream_url"}</script>`))
	}))
	defer server.Close()

	result, err := GetStreamData(server.URL)
	assert.NoError(t, err)
	assert.Equal(t, ` "stream_url"`, result)
}

func TestGetStreamData_InvalidURL_ReturnsError(t *testing.T) {
	_, err := GetStreamData("http://invalid.url")
	assert.Error(t, err)
}

func TestGetStreamData_NoStateInScript_ReturnsEmpty(t *testing.T) {
	// Mock server to simulate a response without "state"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<script>var data = {"NO_STATE": "no_state_data"}</script>`))
	}))
	defer server.Close()

	result, err := GetStreamData(server.URL)
	assert.NoError(t, err)
	assert.Equal(t, "", result)
}

func TestGetStreamData_NoFullHD1InState_ReturnsEmpty(t *testing.T) {
	// Mock server to simulate a response with "state" but no "FULL_HD1"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<script>var state = {"NO_FULL_HD1": "no_full_hd1_data"}</script>`))
	}))
	defer server.Close()

	result, err := GetStreamData(server.URL)
	assert.NoError(t, err)
	assert.Equal(t, "", result)
}
