package flags

import "flag"

var (
	Container = flag.Bool("container", false, "flag to indicate application running in container mode")
	Instance  = flag.Bool("instance", false, "flag to indicate application running in instance mode")
)

func init() {
	flag.Parse()
}
