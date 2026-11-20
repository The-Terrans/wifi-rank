# Pinned to nixos-26.05 so everyone gets the same Go toolchain.
{ pkgs ? import (fetchTarball {
    url = "https://github.com/NixOS/nixpkgs/archive/c56a87d8e41ce69bd2af57eb35dcc117a9dfca33.tar.gz";
    sha256 = "06g8y8y9bysnpxrl6bfqmn63l8w2lx2yhj0618r34j5jfqy5vjpx";
  }) { } }:

pkgs.mkShell {
  packages = [ pkgs.go ];
}
