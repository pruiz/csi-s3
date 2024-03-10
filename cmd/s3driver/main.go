/*
Copyright 2017 The Kubernetes Authors.

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
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ctrox/csi-s3/pkg/driver"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Func("segment", "A key/value tuple to add as segment topology for this node. (like rack=xyz)", func(s string) error {
		if !strings.Contains(s, ":") {
			return fmt.Errorf("Segments must be a key/value pair separated by ':': %s", s)
		}
		parts := strings.Split(s, ":")
		if len(parts) != 2 {
			return fmt.Errorf("Segments must be a key/value pair separated by ':': %s", s)
		}
		segments[parts[0]] = parts[1]
		return nil
	})
}

var (
	endpoint = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	nodeID   = flag.String("nodeid", "", "node id")
	segments = make(map[string]string)
)

func main() {
	flag.Parse()

	driver, err := driver.New(*nodeID, *endpoint, segments)
	if err != nil {
		log.Fatal(err)
	}
	driver.Run()
	os.Exit(0)
}
