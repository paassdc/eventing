/*
Copyright 2018 The Knative Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/knative/pkg/cloudevents"

	corev1 "k8s.io/api/core/v1"
)

func handler(ctx context.Context, e *corev1.Event) {
	metadata := cloudevents.FromContext(ctx)
	log.Printf("[%s] %s : %q", metadata.EventTime.Format(time.RFC3339), metadata.Source, e.Message)
}

func main() {
	log.Print("Ready and listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", cloudevents.Handler(handler)))
}
