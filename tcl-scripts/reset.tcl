if { $argc != 1 } {
    puts stderr "Required 1 parameter: target number"
    exit 1
}

connect
target [lindex $argv 0]
rst
