package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/coreos/pkg/flagutil"
)

var options struct {
	Port int
	Peers string
	UpScript string
	DownScript string
}

var electdFlags = flag.NewFlagSet("electd", flag.ExitOnError)

func init() {
	electdFlags.IntVar(&options.Port, "port", 2339, "use '--port' option to specify the port to listen on")
	electdFlags.StringVar(&options.Peers, "peers", "127.0.0.1:2339", "use '--peers' option to specify the peers")
	electdFlags.StringVar(&options.UpScript, "up-script", "/etc/electd/up.sh", "run this script when current node is elected as leader")
	electdFlags.StringVar(&options.DownScript, "down-script", "/etc/electd/down.sh", "run this script when current node becomes backup")
	electdFlags.Parse(os.Args[1:])
}

func main() {
	if electdFlags.Arg(0) == "version" {
		fmt.Printf("%s/%s\n", path.Base(os.Args[0]), "0.0.1")
		return
	}

	flagutil.SetFlagsFromEnv(electdFlags, "ELECTD")

	fmt.Println(options)
}
