/*
Copyright 2023 Nitro Agility S.r.l.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx := context.Background()

	// Start processing the expectations
	go checkExpectations()

	// Start the prometheus HTTP server and pass the exporter Collector to it
	go serveMetrics()

	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
	<-ctx.Done()
}

func checkExpectations() {
	for {
		log.Printf("Checking expectations")
		time.Sleep(2 * time.Second)
	}
}

func serveMetrics() {
	log.Printf("serving metrics at localhost:2225/metrics")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2225", nil)
	if err != nil {
		fmt.Printf("error serving http: %v", err)
		return
	}
}
