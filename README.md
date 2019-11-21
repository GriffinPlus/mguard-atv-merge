# mguard-atv-merge - A Tool for merging mGuard(TM) configuration files

[![Build Status](https://dev.azure.com/griffinplus/mGuard-ATV-Merge/_apis/build/status/mguard-atv-merge?branchName=master)](https://dev.azure.com/griffinplus/mGuard-ATV-Merge/_build/latest?definitionId=16&branchName=master)
[![Release](https://img.shields.io/github/release/griffinplus/mguard-atv-merge.svg?logo=github)](https://github.com/GriffinPlus/mguard-atv-merge/releases)

-----

## Status

**This project is under active development and should not be used in production, yet.**

At the moment only the application frame is there, no business logic. Stay tuned!

-----

## Overview and Motivation

The *mGuard* security router series is a family of firewall/router devices that is produced by the PHOENIX CONTACT
Cyber Security AG, a member of the [PHOENIX CONTACT group](https://www.phoenixcontact.com/). An *mGuard* in connection
with the [mGuard Secure Cloud](https://us.cloud.mguard.com/) can be used for remote servicing machines in the field
via IPSec-VPN.

Setting up an *mGuard* in the *mGuard Secure Cloud* generates a configuration file containing everything that is needed
to connect to the cloud. The configuration file can then be loaded into the *mGuard* to set it up properly. As a
configuration file is a snapshot of all settings, loading a configuration file replaces existing settings in the
*mGuard*. Therefore it is not possible to have custom settings along with the configuration file generated by the
cloud.

The idea of *mGuard-ATV-Merge* is to merge two configuration files (file extension: `.atv`) into one containing
everything - the static custom settings and dynamic settings like parameters specific to the VPN setup. The merge
process is configurable, so *mGuard-ATV-Merge* can be used to merge *mGuard* configuration files in general.

## Releases

*mGuard-ATV-Merge* is written in GO which makes it highly portable.

[Downloads](https://github.com/GriffinPlus/mguard-atv-merge/releases) are provided for the following combinations of
popular target operating systems and platforms:

- Linux
  - Intel x86 Platform (`386`)
  - Intel x64 Platform (`amd64`)
  - ARM 32-bit Platform (`arm`)
  - ARM 64-bit Platform (`arm64`)
- Windows
  - Intel x86 Platform (`386`)
  - Intel x64 Platform (`amd64`)

If any other target operating system and/or platform is needed and the combination is supported by GO, please open an
issue and we'll add support for it.

## Usage

TODO

## Issues and Contributions

As *mGuard-ATV-Merge* is **not** a project of PHOENIX CONTACT, please don't request help for it there!

If you encounter problems using *mGuard-ATV-Merge*, please file an [issue](https://github.com/GriffinPlus/mguard-atv-merge/issues).

If you have an idea on how to improve *mGuard-ATV-Merge*, please also file an [issue](https://github.com/GriffinPlus/mguard-atv-merge/issues).
Pull requests are also very appreciated. In case of major changes, please open an issue before to discuss the changes.
This helps to coordinate development and avoids wasting time.
