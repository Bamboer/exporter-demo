package collector

import (
        "github.com/prometheus/client_golang/prometheus"
        //  "github.com/prometheus/client_golang/prometheus/promhttp"
        "github.com/Blowfisher/log"
)

func (g grafana) Describe(h chan<- *prometheus.Desc) {
        log.Info.Println("Describe function start...")
        for _, v := range g.admins {
                h <- v
        }
}

func (g grafana) Collect(h chan<- prometheus.Metric) {
        if err := g.adminUpdate(h); err != nil {
                log.Info.Println(err)
        }
}

func (g grafana)adminUpdate(h chan<- prometheus.Metric) error {
        client, err := NewHTTPClient(g.uri, g.username, g.password)
        if err != nil {
                log.Info.Println("Error happened: ", err)
                return err
        }
        gadmin, err := client.Admin()
        if err != nil {
                log.Info.Println("Error happened: ", err)
                return err
        }
        for k,v := range gadmin{
                h <- prometheus.MustNewConstMetric(g.admins[k], prometheus.GaugeValue, v)
        }
        return nil
}
