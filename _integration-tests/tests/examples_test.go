// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2015 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package tests

import (
	"net/http"

	"launchpad.net/snappy/_integration-tests/testutils/cli"
	"launchpad.net/snappy/_integration-tests/testutils/common"
	"launchpad.net/snappy/_integration-tests/testutils/wait"

	"gopkg.in/check.v1"
)

var _ = check.Suite(&webserverExampleSuite{})

type webserverExampleSuite struct {
	common.SnappySuite
}

func (s *webserverExampleSuite) TestNetworkingServiceMustBeStarted(c *check.C) {
	baseAppName := "xkcd-webserver"
	appName := baseAppName + ".canonical"
	common.InstallSnap(c, appName)
	defer common.RemoveSnap(c, appName)

	err := wait.ForServerOnPort(c, 80)
	c.Assert(err, check.IsNil)

	resp, err := http.Get("http://localhost")
	c.Assert(err, check.IsNil)
	c.Check(resp.Status, check.Equals, "200 OK")
	c.Assert(resp.Proto, check.Equals, "HTTP/1.0")
}

var _ = check.Suite(&frameworkExampleSuite{})

type frameworkExampleSuite struct {
	common.SnappySuite
}

func (s *frameworkExampleSuite) TestFrameworkClient(c *check.C) {
	common.InstallSnap(c, "hello-dbus-fwk.canonical")
	defer common.RemoveSnap(c, "hello-dbus-fwk.canonical")

	common.InstallSnap(c, "hello-dbus-app.canonical")
	defer common.RemoveSnap(c, "hello-dbus-app.canonical")

	output := cli.ExecCommand(c, "hello-dbus-app.client")

	expected := "PASS\n"

	c.Assert(output, check.Equals, expected,
		check.Commentf("Expected output %s not found, %s", expected, output))
}