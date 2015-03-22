/**
 * Launches as a sidekick unit to monitor and announce the master unit is ready.
 */
package main

import (
	"github.com/bmorton/deployster/sidekick"
	"github.com/bmorton/deployster/util"

	"github.com/codegangsta/cli"
	"log"
	"os"
	"strings"
)

// supportedUpstreams stores the available upstreams implemented by this package.
var supportedUpstreams = map[string]sidekick.Presence{
	"vulcan-08": sidekick.Vulcan08Presence{},
}

var (
	healthRouteFlag = cli.StringFlag{
		Name:   "health-route,hr",
		Value:  "/",
		Usage:  "A route to call which determines if this service is online. It should return a 200 OK",
		EnvVar: "HEALTH_ROUTE",
	}
	upstreamTypeFlag = cli.StringFlag{
		Name:   "upstream,u",
		Value:  "vulcan-08",
		Usage:  "The upstream proxy to configure, options: [vulcan-07, vulcan-08]",
		EnvVar: "UPSTREAM",
	}
	EtcdURLFlag = cli.StringFlag{
		Name:   "etcd-url",
		Value:  "http://127.0.0.1:4001",
		Usage:  "Etcd URL",
		EnvVar: "ETCD_URL",
	}
)

func main() {
	// log.SetFlags(0)
	log.SetOutput(os.Stderr)
	log.SetPrefix("deployster-sidekick: ")

	app := cli.NewApp()
	app.Name = "deployster-sidekick"
	app.Usage = "Runs alongside your main unit and reports when it's ready to the load balancer"
	app.Flags = []cli.Flag{healthRouteFlag, upstreamTypeFlag}
	app.Version = version.GetCommit()

	app.Action = func(c *cli.Context) {
		SidekickMain(c)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func SidekickMain(c *cli.Context) {
	log.Println("Starting...")
	upstream := c.String("upstream")
	if upstreamService, ok := supportedUpstreams[upstream]; ok {
		log.Printf("Upstream selected: %s\n", upstream)

		upstreamService.EtcdEndpoint = c.String("etcd-url")

		err := sidekick.AnnounceServiceWasAdded(upstreamService)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatalln("Unknown upstream selected, availabe options are:", strings.Join(getAvailableUpstreamNames(), ", "))
		os.Exit(1)
	}
}

func getAvailableUpstreamNames() (names []string) {
	for name, _ := range supportedUpstreams {
		names = append(names, name)
	}
	return
}
