{
  description = "A simple sensor management API in Go";

  inputs = {
    devenv.url = "github:cachix/devenv";
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  # nixConfig = {
  #   extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
  #   extra-substituters = "https://devenv.cachix.org";
  # };

  outputs = { self, nixpkgs, devenv, ... } @ inputs: 
    let
      inherit (nixpkgs) lib;
      supportedSystems = [ "x86_64-linux" ]; 
      forEachSupportedSystem = f:
        lib.genAttrs supportedSystems (system:
          f {
            inherit system;
            pkgs = import nixpkgs { inherit system; config.allowUnfree = true; };
          });
    in rec {
      devShells = forEachSupportedSystem ({ pkgs, ... }: with pkgs; {
        default = devenv.lib.mkShell {
          inherit inputs pkgs;
          modules = [
            ({ pkgs, config, ... }: {
              languages.javascript.enable = true;
              languages.javascript.npm.enable = true;
              languages.typescript.enable = true;

              packages = with pkgs; [ pulumi-bin ];
            })
          ];
        };
      });
    };
}
