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
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx := context.Background()

	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	// Start processing the expectations
	go func() {
		defer wg.Done()
		for {
			select {
				case <- done:
					return
				default:
					checkExpectations()
				}
		}
	}()

	// Start the prometheus HTTP server and pass the exporter Collector to it
	http.Handle("/metrics", promhttp.Handler())
	server := http.Server{Addr: ":2234"}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println("http server error:", err)
			}
		}
	}()

	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
	<-ctx.Done()

	log.Println("shutting down expectations")
	close(done)

	log.Println("shutting down http server")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown failed:", err)
	}

	wg.Wait()
}

func checkExpectations() {
	log.Printf("Checking expectations")
	time.Sleep(1 * time.Second)
}
