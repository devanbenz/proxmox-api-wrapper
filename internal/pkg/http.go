package pkg

import (
	"io"
	"net/http"
)

func SendApiGetRequest(cookie *http.Cookie, csrf string, client http.Client, endpoint string) []byte {
	req, err := http.NewRequest("GET", endpoint, nil)
	CheckErr(err)

	req.Header.Set("CSRFPreventionToken", csrf)
	req.AddCookie(cookie)
	res, err := client.Do(req)
	CheckErr(err)

	body, err := io.ReadAll(res.Body)
	CheckErr(err)

	return body
}

func SendApiPostRequest(cookie *http.Cookie, csrf string, client http.Client, endpoint string, b io.Reader) []byte {
	req, err := http.NewRequest("POST", endpoint, b)
	CheckErr(err)

	req.Header.Set("CSRFPreventionToken", csrf)
	req.AddCookie(cookie)
	res, err := client.Do(req)
	CheckErr(err)

	body, err := io.ReadAll(res.Body)
	CheckErr(err)

	return body
}
