package main

import (
  "os"
  "flag"
  "net/http"
  "grafana/collector"
  "github.com/Blowfisher/log"
//  "github.com/prometheus/common/promlog"
//  "github.com/prometheus/common/promlog/flag"
//  "github.com/prometheus/common/version"

  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"
//  "gopkg.in/alecthomas/kingpin.v2"
)

var (
  Uri = flag.String("grafana.url","http://10.40.45.18:3000/","Grafana host's url address.")
  Username = flag.String("grafana.username","admin","Grafana service's account name")
  Password = flag.String("grafana.password","fisdfdfds","Grafana service's account password")
  listenAddress = flag.String("grafana.listen.address",":8090","Grafana exporter listen address")
)

func main(){
  flag.Parse()
  grafana := collector.Newgrafana(*Uri,*Username,*Password)
  prometheus.MustRegister(grafana)
  http.Handle("/metrics", promhttp.Handler())
  if err := http.ListenAndServe(*listenAddress, nil); err != nil {
                log.Info.Println("msg", "Error starting HTTP server", "err", err)
                os.Exit(1)
  }
}
