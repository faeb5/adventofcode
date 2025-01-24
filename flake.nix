{
  description = "Advent of Code 2025 Flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-24.11";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    pkgs = nixpkgs.legacyPackages.x86_64-linux;
  in {
    formatter.x86_64-linux = pkgs.alejandra;
    devShells.x86_64-linux.janet = pkgs.mkShellNoCC {
      packages = with pkgs; [go gopls gotools];
      shellHook = ''echo "Welcome to the Advent of Code 2024 environment!'';
    };
    devShells.x86_64-linux.default = self.devShells.x86_64-linux.janet;
  };
}
