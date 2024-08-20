# SIMD Optimizations for Montgomery Multiplication

This project is designed to explore and implement various SIMD (Single Instruction, Multiple Data) optimizations for Montgomery multiplication algorithms. The primary focus is on leveraging AVX2 instructions to enhance performance for cryptographic applications on x86 architecture. Particular optimizations may also leverage multithreading via goroutines.

## Optimized Montgomery Modular Multiplication

The new implementation for montgomery Modular multiplication `MUL` can be found in the `ecc/bls12-377/fp/element_ops_amd64.go` file. `MUL` implements a parallel radix-2^64 interleaved Montgomery multiplication algorithm described in Algorithm 4 of the paper ["Improved Montgomery Multiplication on SIMD Architectures"](https://eprint.iacr.org/2017/1057.pdf).

## Optimized Pairing and Hashing
The new implementation for paralleilzed hashing computations to $G_2$ can be found in the `ecc/bls12-377/hash_to_g2.go` file, namely in functions `MaptoG2` and `g2Isogeny`. 

The new implementation for parallelized pairing computations can be found in the `ecc/bls12-377/pairing.go` file, namely `MillerLoop` and `MillerLoopFixedQ`.

## Fork Information

This repository was forked from [https://github.com/Consensys/gnark-crypto/tree/master](https://github.com/Consensys/gnark-crypto/tree/master). `gnark-crypto` provides elliptic curve and pairing-based cryptography on BN, BLS12, BLS24, and BW6 curves. It also provides various algorithms (algebra, crypto) of particular interest to zero knowledge proof systems.

## Getting Started
### Install modified gnark-crypto

```bash
go get https://github.com/RGBmarya/gnark-crypto
```
## Running MSM Benchmarks
To benchmark the new implementation for montgomery Modular multiplication `MUL` run:

```bash
cd ecc/bls12-377
go test -bench=BenchmarkMultiExpG1 -cpu 16,32,64,128
```
To plot benchmark results for the new implementation run:

```bash
cd ecc/bls12-377/plots
python plot.py
```
## Running Hashing and Pairing Benchmarks

To benchmark the optimizations for hashing and pairing computations, run:

```bash
cd ecc/bls12-377
go test -v
```
This runs all tests in `ecc/bls12-377`. The most time-consuming 

To generate an XML report of the test results, run:
```bash
cd ecc/bls12-377
go test -v 2>&1 | go-junit-report > report.xml
```