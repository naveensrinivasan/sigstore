#!/bin/bash -eu
fuzzer=fuzz_loadCertificates
cd "test/fuzz/pem"

go-fuzz -tags go-fuzz -func FuzzLoadCertificates -o $fuzzer  .


$CXX $CXXFLAGS $LIB_FUZZING_ENGINE -v $fuzzer -o $fuzzer.a
