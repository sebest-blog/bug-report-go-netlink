This repository demonstrates a bug in go-netlink when there is 2 default gateway with the same weight

To demonstrate the bug you need Docker to be able to reproduce the network setup with the 2 gateways

When you run `run.sh` it will start a Docker container and execute `bug.sh` inside.
