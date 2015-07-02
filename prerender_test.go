package prerender

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var _MAC_USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12"

func Test_BotRequest(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "https://www.google.com/", nil)
	req.Header.Set("User-Agent", "twitterbot")

	NewOptions().NewPrerender().ServeHTTP(res, req, nil)

	if len(res.Body.Bytes()) == 0 {
		t.Error("Error, prerender.io not called")
	}
}

func Test_NonBotRequest(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "https://www.google.com/", nil)
	req.Header.Set("User-Agent", _MAC_USER_AGENT)

	NewOptions().NewPrerender().ServeHTTP(res, req, nil)
	if len(res.Body.Bytes()) > 0 {
		t.Error("Error, prerender.io called for non-proxy request")
	}
}

func Test_RootEscapedFragment(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "https://www.google.com/?_escaped_fragment_=", nil)
	req.Header.Set("User-Agent", _MAC_USER_AGENT)

	NewOptions().NewPrerender().ServeHTTP(res, req, nil)
	if len(res.Body.Bytes()) == 0 {
		t.Error("Error, prerender.io not called for empty _escaped_fragment_")
	}
}

func Test_EscapedFragment(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "https://www.google.com/?_escaped_fragment_=something", nil)
	req.Header.Set("User-Agent", _MAC_USER_AGENT)

	NewOptions().NewPrerender().ServeHTTP(res, req, nil)
	if len(res.Body.Bytes()) == 0 {
		t.Error("Error, prerender.io not called for _escaped_fragment_")
	}
}
