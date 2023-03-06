package config

import (
	"flag"
)

type args struct {
	Port        string
	Host        string
	Environment string
}

var Args args = args{}

func ParseArgs() {
	portFlag := flag.String("p", "9090", "Port to listen on")
	hostFlag := flag.String("h", "0.0.0.0", "Interface to listen on")
	environmentFlag := flag.String("environment", "production", "Port to listen on")
	flag.Parse()
	Args.Port = *portFlag
	Args.Host = *hostFlag
	Args.Environment = *environmentFlag
}
