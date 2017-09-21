///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdGet(t *testing.T) {
	var buf bytes.Buffer

	cli := NewVSCLI(os.Stdin, &buf, &buf)
	cli.SetOutput(&buf)
	cli.SetArgs([]string{"get", "--help"})
	err := cli.Execute()
	assert.Nil(t, err)
	assert.True(t, strings.Contains(buf.String(), "Display one or many resources"))
}