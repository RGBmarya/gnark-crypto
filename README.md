# SIMD Optimizations for Montgomery Multiplication

This project is designed to explore and implement various SIMD (Single Instruction, Multiple Data) optimizations for Montgomery multiplication algorithms. The primary focus is on leveraging AVX2 instructions to enhance performance for cryptographic applications on x86 architecture.

## Optimized Montgomery Modular Multiplication

The new implementation for montgomery Modular multiplication `MUL` can be found in the `ecc/bls12-377/fp/element_ops_amd64.go` file. `MUL` implements a parallel radix-2^64 interleaved Montgomery multiplication algorithm described in Algorithm 4 of the paper ["Improved Montgomery Multiplication on SIMD Architectures"](https://eprint.iacr.org/2017/1057.pdf).

## Fork Information

This repository was forked from [https://github.com/Consensys/gnark-crypto/tree/master](https://github.com/Consensys/gnark-crypto/tree/master). `gnark-crypto` provides elliptic curve and pairing-based cryptography on BN, BLS12, BLS24, and BW6 curves. It also provides various algorithms (algebra, crypto) of particular interest to zero knowledge proof systems.
