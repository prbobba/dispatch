///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package main

// NO TEST

import (
	"os"

	"github.com/vmware/dispatch/pkg/event-driver/sidecar"
)

func main() {
	cli := sidecar.NewCmd(os.Stdin, os.Stdout, os.Stderr)
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
