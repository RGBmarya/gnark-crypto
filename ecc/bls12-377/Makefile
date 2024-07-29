bench:
	mkdir -p /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version)
	go test -bench=BenchmarkG1AffineBatchScalarMultiplication -run=^$$ . -cpuprofile /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version)/cpu-profile-v$(version).out -memprofile /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version)/mem-profile-v$(version).out

test:
	@$(MAKE) bench version=$(version)

profile:
	@if [ ! -d /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version) ]; then echo "Directory /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version) does not exist"; exit 1; fi
	cd /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version) && \
	touch cpu-profile-v$(version).pdf && \
	touch mem-profile-v$(version).pdf && \
	go tool pprof -pdf cpu-profile-v$(version).out > cpu-profile-v$(version).pdf && \
	go tool pprof -pdf mem-profile-v$(version).out > mem-profile-v$(version).pdf