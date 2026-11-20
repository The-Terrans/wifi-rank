# Pinned to the same nixpkgs rev as home-manager so Go comes from the local store.
{ pkgs ? import (fetchTarball {
    url = "https://github.com/NixOS/nixpkgs/archive/04906337b6d53e8bb0eaf0470ae9f070a4d4eb55.tar.gz";
    sha256 = "1ixmpdkxv55py03w9gj4qn4p90jn494a5zzx3j90c703gaa3dhcl";
  }) { } }:

pkgs.mkShell {
  packages = [ pkgs.go ];
}
