{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... } @inputs: flake-utils.lib.eachDefaultSystem(system: let
    pkgs = import nixpkgs { inherit system; };
  in {
    devShells.default = pkgs.mkShell {
      buildInputs = [pkgs.go];
    };
    packages.default = pkgs.buildGoModule {
      pname = "game";
      version = "0.1.0";
      src = ./.;
      vendorHash = "sha256-fV6qYXC4MJWzlLF46qoY0NoFS0W/Jy0YuMHv0aiQQyk=";
    };
  });
}
