// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

// #cgo pkg-config: hwloc
// #include <hwloc.h>
import "C"
import (
	"fmt"

	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/plugins"
	"github.com/mmcilroy/nomad-device-plugin/device"
)

type Hwloc struct {
	topology C.hwloc_topology_t
}

func ReadTopology() {
	hwloc := Hwloc{}
	returncode := C.hwloc_topology_init(&hwloc.topology)
	if returncode != 0 {
		fmt.Printf("error initialising hwlog topology: %d\n", returncode)
	}

	C.hwloc_topology_load(hwloc.topology)
	numa_nodes := C.hwloc_get_nbobjs_by_type(hwloc.topology, C.HWLOC_OBJ_NUMANODE)
	fmt.Printf("num nodes: %d\n", numa_nodes)

	numa_node := C.hwloc_get_obj_by_type(hwloc.topology, C.HWLOC_OBJ_NUMANODE, 0)
	core_count := C.hwloc_get_nbobjs_inside_cpuset_by_type(hwloc.topology, numa_node.cpuset, C.HWLOC_OBJ_CORE)
	fmt.Printf("num cores per node: %d\n", core_count)
}

func main() {
	ReadTopology()
	// Serve the plugin
	plugins.Serve(factory)
}

// factory returns a new instance of our example device plugin
func factory(log log.Logger) interface{} {
	return device.NewPlugin(log)
}
