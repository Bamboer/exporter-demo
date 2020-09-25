package collector

import (
        "github.com/prometheus/client_golang/prometheus"
        //      "github.com/Blowfisher/log"
)

const (
        namespace = "grafana"
)

type (
        gAlerts  map[string] string
        gAdmin map[string] float64
)

type grafana struct {
        admins map[string]*prometheus.Desc
        uri  string
        username string
        password string
}

func Newgrafana(uri,username,password string) grafana {
        return grafana{admins: newadmins(),
          uri: uri,
          username: username,
          password: password,
        }
}

func newadmins() map[string]*prometheus.Desc {
        return map[string]*prometheus.Desc{
                "orgs": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "orgs"),
                        "orgnazation number",
                        nil, nil,
                ),
                "dashboards": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "dashboards"),
                        "dashboards number",
                        nil, nil,
                ),

                "snapshots": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "snapshots"),
                        "snapshots number",
                        nil, nil,
                ),

                "tags": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "tags"),
                        "tags number",
                        nil, nil,
                ),

                "datasources": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "datasources"),
                        "datasources number",
                        nil, nil,
                ),

                "playlists": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "playlists"),
                        "playlists number",
                        nil, nil,
                ),

                "stars": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "stars"),
                        "stars number",
                        nil, nil,
                ),

                "alerts": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "alerts"),
                        "alerts number",
                        nil, nil,
                ),

                "users": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "users"),
                        "users number",
                        nil, nil,
                ),

                "admins": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "admins"),
                        "admins number",
                        nil, nil,
                ),

                "editors": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "editors"),
                        "editors number",
                        nil, nil,
                ),

                "viewers": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "viewers"),
                        "viewers number",
                        nil, nil,
                ),

                "activeUsers": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "activeusers"),
                        "activeusers number",
                        nil, nil,
                ),

                "activeAdmins": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "activeadmins"),
                        "activeadmins number",
                        nil, nil,
                ),

                "activeEditors": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "activeeditors"),
                        "activeeditors number",
                        nil, nil,
                ),

                "activeViewers": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "activeviewers"),
                        "activeviewers number",
                        nil, nil,
                ),

                "activeSessions": prometheus.NewDesc(
                        prometheus.BuildFQName(namespace, "export", "activesessions"),
                        "activesessions number",
                        nil, nil,
                ),
        }
}
