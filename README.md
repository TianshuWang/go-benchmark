# Benchmark test
-  go test -bench=. -cpu=1 -benchtime=1x

# pprof
- go test -bench . -test.blockprofile cpu.profile 
- go test -bench . -test.cpuprofile cpu.profile 
- go test -bench . -test.memprofile cpu.profile 
- go test -bench . -test.mutexprofile cpu.profile 
- go tool pprof -http :8080 cpu.profile
- go tool pprof cpu.profile
- list [cpu]
