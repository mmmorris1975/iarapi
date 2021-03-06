package iarapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"golang.org/x/net/publicsuffix"
)

var (
	apiBase  = `https://coordinator.iamresponding.com/api`
	loginUrl = `https://iamresponding.com/v3/Pages/memberlogin.aspx/ValidateLoginInfo`
)

func NewClient(agency, user, password string) (*Client, error) {
	c := new(Client)
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	c.httpClient = http.Client{Jar: jar}

	if err := c.login(agency, user, password); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) login(agency, user, password string) error {
	login := &LoginRequest{
		MemberLogin: true,
		Agency:      agency,
		User:        user,
		Password:    password,
	}

	msg := new(LoginReply)
	if err := c.apiPost(loginUrl, login, msg); err != nil {
		return err
	}

	if !strings.Contains(msg.Message, "iamresponding.com/") {
		return errors.New(msg.Message)
	}

	return nil
}

func (c *Client) Subscriber() (*SubscriberInfo, error) {
	si := new(SubscriberInfo)
	return si, c.apiGet("/Subscriber", si)
}

func (c *Client) Member() (*MemberInfo, error) {
	mi := new(MemberInfo)
	return mi, c.apiGet("/Member", mi)
}

func (c *Client) Incidents() (*IncidentList, error) {
	il := new(IncidentList)
	return il, c.apiGet("/IncidentList", il)
}

func (c *Client) Messages() (*MessageList, error) {
	ml := new(MessageList)
	return ml, c.apiGet("/MessageList", ml)
}

func (c *Client) Dispatchers() (*Dispatchers, error) {
	dl := new(Dispatchers)
	return dl, c.apiGet("/DispatcherContent/AssociatedDispatchers", dl)
}

func (c *Client) ResponderCodes() (*ResponderCodes, error) {
	rc := new(ResponderCodes)
	return rc, c.apiGet("/ResponderCodes", rc)
}

func (c *Client) OnDutyAtCodes() (*OnDutyAtCodeList, error) {
	cl := new(OnDutyAtCodeList)
	return cl, c.apiGet("/OnDutyAtCodes", cl)
}

func (c *Client) ResponderList() (*ResponderList, error) {
	rl := new(ResponderList)
	return rl, c.apiGet("/ResponderList", rl)
}

func (c *Client) ApparatusList() (*ApparatusList, error) {
	al := new(ApparatusList)
	return al, c.apiGet("/ApparatusList", al)
}

func (c *Client) SearchIncidents(isr *IncidentSearchRequest) (*IncidentList, error) {
	il := new(IncidentList)
	return il, c.apiPost(apiBase+"/SearchIncidents", isr, il)
}

func (c *Client) apiGet(path string, t interface{}) error {
	return c.apiGetWithContext(context.Background(), path, t)
}

func (c *Client) apiGetWithContext(ctx context.Context, path string, t interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiBase+path, http.NoBody)
	if err != nil {
		return err
	}

	return c.doApiRequest(req, t)
}

func (c *Client) apiPost(url string, input, output interface{}) error {
	return c.apiPostWithContext(context.Background(), url, input, output)
}

func (c *Client) apiPostWithContext(ctx context.Context, url string, input, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	return c.doApiRequest(req, output)
}

func (c *Client) doApiRequest(req *http.Request, t interface{}) error {
	req.Header.Add("Accept", "text/plain,application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Status %d", res.StatusCode)
	}

	var b []byte
	b, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, t)
}
