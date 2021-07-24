package main

import (
	"fmt"
	"hpilo_exporter/collector"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func metrichandler(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8080/static/test.json")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(data)
	}

	//defer config.GOFISH.Logout()

	fmt.Println(" Connect successful")

	mhandler := promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
		})
	mhandler.ServeHTTP(w, r)
}

func main() {
	const PORT = "9000"
	fmt.Println("Server listening at ", PORT)

	// Listen all interfaces at port 9000
	const IP_ADDRESS = ":" + PORT

	system := collector.SystemCollector{}
	fmt.Printf("%v", system)
	//prometheus.Register(system)

	// chassis := chassis.Chassis{}
	// prometheus.Register(chassis)

	// Starting server
	http.HandleFunc("/metrics", metrichandler)
	http.ListenAndServe(IP_ADDRESS, nil)
}
