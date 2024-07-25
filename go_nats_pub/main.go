package main

import "runtime"

var stat runtime.MemStats

func main() {
	// nc, _ := nats.Connect(nats.DefaultURL)
	// _ = nc.Publish("hello", []byte("Just do it"))
	// _ = nc.Drain()
	runtime.ReadMemStats(&stat)
	println(stat.HeapSys)
}
