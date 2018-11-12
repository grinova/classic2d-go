#!/bin/bash
go test -c github.com/grinova/classic2d-server/dynamic -o ./profile/world.test
./profile/world.test -test.run=NONE -test.benchmem -test.bench=World -test.cpuprofile=./profile/world.cpu.log dynamic
go tool pprof -text -nodecount=10 ./profile/world.test ./profile/world.cpu.log
