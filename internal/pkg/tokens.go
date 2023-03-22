package pkg

import (
	"crypto/tls"
	"encoding/json"
	"github.com/gorilla/schema"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type TicketInput struct {
	Username string `schema:"username,required"`
	Password string `schema:"password,required"`
}

type TicketOutput struct {
	Data struct {
		CSRFPreventionToken string `json:"CSRFPreventionToken"`
		Ticket              string `json:"ticket"`
	} `json:"data"`
}

func (t *TicketInput) GetTokens(client *http.Client, tokenEndpoint string) (*http.Cookie, string) {
	rd := TicketOutput{}
	form := url.Values{}
	encoder := schema.NewEncoder()
	jar, _ := cookiejar.New(nil)
	client.Jar = jar
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	err := encoder.Encode(t, form)
	CheckErr(err)

	req, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(form.Encode()))
	CheckErr(err)

	res, err := client.Do(req)
	CheckErr(err)

	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &rd)
	CheckErr(err)

	cookie := &http.Cookie{
		Name:   "PVEAuthCookie",
		Value:  rd.Data.Ticket,
		MaxAge: 300,
	}

	return cookie, rd.Data.CSRFPreventionToken
}
