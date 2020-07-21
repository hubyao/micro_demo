package main

import (
	//"github.com/micro/go-plugins//micro/cors"
	//"github.com/micro/micro/plugin"
	//"github.com/opentracing/opentracing-go"
	//"github.com/micro/micro/cmd"
	"github.com/micro/micro/v2/cmd"
	)

func init() {
	//plugin.Register(cors.NewPlugin())

	//plugin.Register(plugin.NewPlugin(
	//	plugin.WithName("tracer"),
	//	plugin.WithHandler(
	//		stdhttp.TracerWrapper,
		//),
	//))
}

const name = "API gateway"

func main() {
	//stdhttp.SetSamplingFrequency(50)
	//t, io, err := tracer.NewTracer(name, "")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer io.Close()
	//opentracing.SetGlobalTracer(t)

	cmd.Init()
}