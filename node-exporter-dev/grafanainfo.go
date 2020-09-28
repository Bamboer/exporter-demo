package collector

import (
        "github.com/go-kit/kit/log/level"
        kingpin "gopkg.in/alecthomas/kingpin.v2"
        "encoding/json"
        "net/http"
        "net/url"
        "errors"
)

var (
        gd         gAdmin
        grafanaUri = kingpin.Flag(
                "grafana.address",
                "Grafana address.",
        ).Default("http://10.40.45.18:3000/").String()
        grafanaUsername = kingpin.Flag(
                "grafana.username",
                "Grafana controller username.",
        ).Default("admin").String()
        grafanaPassword = kingpin.Flag(
                "grafana.password",
                "Grafana controller password.",
        ).Default("f13dfd").String()
)

type (
        gAdmin      map[string]float64
        http_client struct {
                url        *url.URL
                username   string
                password   string
                httpClient *http.Client
        }
)

func newHTTPClient(uri, username, password string) (*http_client, error) {
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

func (c *http_client) admin() (gAdmin, error) {
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

func (g *grafanainfoCollector) getGrafanaInfo() (gAdmin, error) {
        if *grafanaUri == "" && *grafanaUsername == "" && *grafanaPassword == "" {
               return nil,errors.New("Grafana info not set")
        }
                client, err := newHTTPClient(*grafanaUri, *grafanaUsername, *grafanaPassword)
                if err != nil {
                        level.Debug(g.logger).Log("msg", "http client access", "grafana: ", err)
                        return nil, err
                }
        gadmin, err := client.admin()
        if err != nil {
                level.Debug(g.logger).Log("msg", "http client access", "grafana admin api: ", err)
                return nil, err
        }
        return gadmin, nil
}
