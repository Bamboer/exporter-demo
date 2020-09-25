package collector

import (
        "encoding/json"
        "net/http"
        "net/url"
)

var (
        gd gAdmin
        gr []gAlerts
)

type http_client struct {
        url        *url.URL
        username   string
        password   string
        httpClient *http.Client
}

func NewHTTPClient(uri, username, password string) (*http_client, error) {
        grafanaurl, err := url.Parse(uri)
        if err != nil {
                return nil, err
        }
        return &http_client{
                url:        grafanaurl,
                username:   username,
                password:   password,
                httpClient: &http.Client{},
        }, nil
}

func (c *http_client) Alerts() ([]gAlerts, error) {
        uri := c.url
        uri.Path = "/api/alerts"
        req, err := http.NewRequest("GET", uri.String(), nil)
        if err != nil {
                return nil, err
        }
        req.SetBasicAuth(c.username, c.password)
        req.Header.Add("Content-Type", "application/json")
        req.Header.Add("Accept", "application/json")
        resp, err := c.httpClient.Do(req)
        if err != nil {
                return nil, err
        }
        defer resp.Body.Close()
        err = json.NewDecoder(resp.Body).Decode(&gr)
        if err != nil {
                return nil, err
        }
        return gr, nil
}

func (c *http_client) Admin() (gAdmin, error) {
        uri := c.url
        uri.Path = "/api/admin/stats"
        req, err := http.NewRequest("GET", uri.String(), nil)
        if err != nil {
                return gd, err
        }
        if c.username != "" && c.password != "" {
                req.SetBasicAuth(c.username, c.password)
        }
        req.Header.Add("Content-Type", "application/json")
        req.Header.Add("Accept", "application/json")
        resp, err := c.httpClient.Do(req)
        if err != nil {
                return gd, err
        }
        defer resp.Body.Close()
        err = json.NewDecoder(resp.Body).Decode(&gd)
        if err != nil {
                return gd, err
        }
        return gd, nil
}
