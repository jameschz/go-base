package main

import (
    "base/util"
)

func main() {
    // dump project configs
    util.Dump("util.GetEnv()", util.GetEnv())
    util.Dump("util.GetRootPath()", util.GetRootPath())
}
