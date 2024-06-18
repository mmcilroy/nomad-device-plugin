// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os/exec"

	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/plugins"
	"github.com/mmcilroy/nomad-device-plugin/device"
)

func readTopology() {
	cmd := exec.Command("lstopo", "--of", "console", "--no-caches", "--no-smt")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

func main() {
	// Serve the plugin
	plugins.Serve(factory)
}

// factory returns a new instance of our example device plugin
func factory(log log.Logger) interface{} {
	log.Debug("factory")
	return device.NewPlugin(log)
}
