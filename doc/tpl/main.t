package main

import (
    "base/util"
)

func main() {
    // dump project configs
    gutil.Dump("gutil.GetEnv()", gutil.GetEnv())
    gutil.Dump("gutil.GetRootPath()", gutil.GetRootPath())
}
