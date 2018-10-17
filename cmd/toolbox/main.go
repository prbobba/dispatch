///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// +build linux,amd64

package main

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/coreos/go-systemd/daemon"
	"github.com/sirupsen/logrus"

	"github.com/vmware/govmomi/toolbox"
	"github.com/vmware/govmomi/toolbox/vix"
)

const (
	keepEnvVars = false
	sdReady     = "READY=1"
)

func main() {
	log := logrus.New().WithField("app", "toolbox")

	in := toolbox.NewBackdoorChannelIn()
	out := toolbox.NewBackdoorChannelOut()

	service := toolbox.NewService(in, out)

	if os.Getuid() == 0 {
		service.Power.Halt.Handler = toolbox.Halt
		service.Power.Reboot.Handler = toolbox.Reboot
	}

	// Disable all guest operations
	service.Command.Authenticate = func(_ vix.CommandRequestHeader, _ []byte) error {
		return errors.New("not authorized")
	}

	err := service.Start()
	if err != nil {
		log.Fatal(err)
	}

	daemon.SdNotify(keepEnvVars, sdReady)

	// handle the signals and gracefully shutdown the service
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("signal %s received", <-sig)
		service.Stop()
	}()

	service.Wait()
}
