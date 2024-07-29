bench:
	mkdir -p /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version)
	go test -bench=BenchmarkG1AffineBatchScalarMultiplication -run=^$$ . -cpuprofile /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version)/cpu-profile-v$(version).out -memprofile /home1/msarya/Dev/gnark-crypto/ecc/bls12-377/profiles/$(version)/mem-profile-v$(version).out

test:
	@$(MAKE) bench version=$(version)
