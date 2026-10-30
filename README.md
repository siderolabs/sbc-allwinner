# sbc-allwinner

This repo provides the overlay for Allwinner based Talos image.

## Supported Overlay

| Overlay Name           | Board              | Description                                                                                                                                 |
| ---------------------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------- |
| bananapi_m64           | BananaPi M64       | [Installation guide](https://docs.siderolabs.com/talos/v1.14/platform-specific-installations/single-board-computers/bananapi_m64)           |
| libretech_all_h3_cc_h5 | LibreTech H3 CC H5 | [Installation guide](https://docs.siderolabs.com/talos/v1.14/platform-specific-installations/single-board-computers/libretech_all_h3_cc_h5) |
| pine64                 | Pine64 A64         | [Installation guide](https://docs.siderolabs.com/talos/v1.14/platform-specific-installations/single-board-computers/pine64)                 |

U-Boot (`u-boot-sunxi-with-spl.bin`, built with Trusted Firmware-A) is written by the installer to the boot disk at an 8 KiB offset.
