// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/siderolabs/go-copy/copy"
	"github.com/siderolabs/talos/pkg/machinery/overlay"
	"github.com/siderolabs/talos/pkg/machinery/overlay/adapter"
	"golang.org/x/sys/unix"
)

const (
	off int64 = 1024 * 8
	dtb       = "allwinner/sun50i-h5-libretech-all-h3-cc.dtb"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	adapter.Execute(ctx, &LibretechAllH3CCH5Installer{})
}

// References:
//   - https://libre.computer/products/boards/all-h3-cc/
type LibretechAllH3CCH5Installer struct{}

type libretechAllH3CCH5ExtraOptions struct{}

func (i *LibretechAllH3CCH5Installer) GetOptions(_ context.Context, extra libretechAllH3CCH5ExtraOptions) (overlay.Options, error) {
	return overlay.Options{
		Name: "libretech_all_h3_cc_h5",
		KernelArgs: []string{
			"console=tty0",
			"console=ttyS0,115200",
			"sysctl.kernel.kexec_load_disabled=1",
			"talos.dashboard.disabled=1",
		},
		PartitionOptions: overlay.PartitionOptions{
			Offset: 2048,
		},
	}, nil
}

func (i *LibretechAllH3CCH5Installer) Install(_ context.Context, options overlay.InstallOptions[libretechAllH3CCH5ExtraOptions]) error {
	var f *os.File

	f, err := os.OpenFile(options.InstallDisk, os.O_RDWR|unix.O_CLOEXEC, 0o666)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", options.InstallDisk, err)
	}

	defer f.Close() //nolint:errcheck

	uboot, err := os.ReadFile(filepath.Join(options.ArtifactsPath, "arm64/u-boot/libretech_all_h3_cc_h5/u-boot-sunxi-with-spl.bin"))
	if err != nil {
		return err
	}

	if _, err = f.WriteAt(uboot, off); err != nil {
		return err
	}

	// NB: In the case that the block device is a loopback device, we sync here
	// to ensure that the file is written before the loopback device is
	// unmounted.
	err = f.Sync()
	if err != nil {
		return err
	}

	src := filepath.Join(options.ArtifactsPath, "arm64/dtb", dtb)
	dst := filepath.Join(options.MountPrefix, "/boot/EFI/dtb", dtb)

	err = os.MkdirAll(filepath.Dir(dst), 0o700)
	if err != nil {
		return err
	}

	return copy.File(src, dst)
}
