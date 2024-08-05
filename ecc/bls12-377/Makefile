bench:
	mkdir -p /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version)
	go test -bench=BenchmarkMultiExpG1 -run=^$$ . -cpuprofile /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version)/BenchmarkMultiExpG1-cpu-profile-v$(version).out -memprofile /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version)/BenchmarkMultiExpG1-mem-profile-v$(version).out

test:
	@$(MAKE) bench version=$(version)

profile:
	@if [ ! -d /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version) ]; then echo "Directory /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version) does not exist"; exit 1; fi
	cd /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version) && \
	touch BenchmarkMultiExpG1-cpu-profile-v$(version).pdf && \
	touch BenchmarkMultiExpG1-mem-profile-v$(version).pdf && \
	go tool pprof -pdf BenchmarkMultiExpG1-cpu-profile-v$(version).out > BenchmarkMultiExpG1-cpu-profile-v$(version).pdf && \
	go tool pprof -pdf BenchmarkMultiExpG1-mem-profile-v$(version).out > BenchmarkMultiExpG1-mem-profile-v$(version).pdf
