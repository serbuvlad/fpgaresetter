# fpgaresetter -- TUI for resetting AMD (Xilinx) FPGA boards

`fpgaresetter` is a small script to reset an FPGA board or component
using the `xsdb` utility.

What specific component should be reset to reset the whole board; or
whether this reset method will work will depend on the particular board.

This tool was tested on Pynq-Z2 where the first CPU should be selected.

## Prerequisites

`fpgaresetter` requires a Vivado instllation. Make sure the Vivado
enviornment settings are properly set by sourcing the appropriate file
in your `.bashrc` (or equivalent):

```
# For example (in ~/.bashrc)
source /tools/Xilinx/Vivado/2024.2/settings.sh
```

Make sure the cable drivers are installed. The following should be run once:

```
sudo /tools/Xilinx/Vivado/2024.2/data/xicom/cable_drivers/lin64/install_script/install_drivers/install_drivers
```

## Installing

To install, make sure `$(go env GOPATH)/bin` is in your path, then run

```
go install github.com/serbuvlad/fpgaresetter@latest
```

## Running

`fpgaresetter` takes no arguments and will provide you with a TUI to
select the board or component you want to restart.
