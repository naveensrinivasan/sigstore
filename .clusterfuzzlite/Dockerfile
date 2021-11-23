FROM gcr.io/oss-fuzz-base/base-builder-go:v1
RUN apt-get update && apt-get install -y make autoconf automake libtool pkg-config libpcsclite-dev
COPY . $SRC/sigstore
WORKDIR sigstore
COPY .clusterfuzzlite/build.sh $SRC/
