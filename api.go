package iarapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/publicsuffix"
)

var (
	dashboardHost = `https://dashboard.iamresponding.com`
	apiBase       = `https://coordinator.iamresponding.com/api`
	loginUrl      = `https://auth.iamresponding.com/login/member`
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
	token, err := c.fetchRequestToken()
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("Input.Agency", agency)
	form.Add("Input.Username", user)
	form.Add("Input.Password", password)
	form.Add("__RequestVerificationToken", token)
	form.Add("Input.RememberLogin", "false")
	form.Add("Input.button", "login")
	form.Add("Input.ReturnUrl", "")

	return c.doLogin(form)
}

// Get the request verification token from the hidden form field on the HTML login page
// Passed in login request along with Agency, Username, and Password
func (c *Client) fetchRequestToken() (string, error) {
	res, err := c.httpClient.Get(loginUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	var token string
	doc.FindMatcher(goquery.Single("form.iar-form__form")).ChildrenFiltered("input").Each(func(i int, s *goquery.Selection) {
		if val, ok := s.Attr("name"); ok {
			if val == "__RequestVerificationToken" {
				token, _ = s.Attr("value")
			}
		}
	})

	return token, nil
}

// POST login form values to user auth endpoint, then perform oauth authorization
func (c *Client) doLogin(data url.Values) error {
	var req *http.Request
	var res *http.Response
	var err error

	req, err = http.NewRequest(http.MethodPost, loginUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(&http.Cookie{Name: "CookieConsent", Value: "yes"})

	res, err = c.httpClient.Do(req)
	if err != nil {
		return err
	}
	res.Body.Close()

	res, err = c.httpClient.Get(dashboardHost + `/system/login?returnUrl=/`)
	if err != nil {
		return err
	}
	defer res.Body.Close()

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
	req.Header.Set("X-CSRF", "1")
	req.AddCookie(&http.Cookie{Name: "CookieConsent", Value: "yes"})
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
