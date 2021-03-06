// Copyright 2016 CoreOS, Inc.
// Copyright 2018 Red Hat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package qemu

import (
	"github.com/coreos/pkg/capnslog"

	"github.com/coreos/mantle/platform"
	"github.com/coreos/mantle/platform/local"
)

const (
	Platform platform.Name = "qemu"
)

// Options contains QEMU-specific options for the flight.
type Options struct {
	// DiskImage is the full path to the disk image to boot in QEMU.
	DiskImage string
	Board     string

	// BIOSImage is name of the BIOS file to pass to QEMU.
	// It can be a plain name, or a full path.
	BIOSImage string

	//Option to create a temporary software TPM - true by default
	Swtpm bool

	*platform.Options
}

type flight struct {
	*local.LocalFlight
	opts *Options
}

var (
	plog = capnslog.NewPackageLogger("github.com/coreos/mantle", "platform/machine/qemu")
)

func NewFlight(opts *Options) (platform.Flight, error) {
	lf, err := local.NewLocalFlight(opts.Options, Platform)
	if err != nil {
		return nil, err
	}

	qf := &flight{
		LocalFlight: lf,
		opts:        opts,
	}

	return qf, nil
}

// NewCluster creates a Cluster instance, suitable for running virtual
// machines in QEMU.
func (qf *flight) NewCluster(rconf *platform.RuntimeConfig) (platform.Cluster, error) {
	lc, err := qf.LocalFlight.NewCluster(rconf)
	if err != nil {
		return nil, err
	}

	qc := &Cluster{
		flight:       qf,
		LocalCluster: lc,
	}

	qf.AddCluster(qc)

	return qc, nil
}
