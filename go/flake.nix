{
  description = "A simple Webfinger server in Go";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = { nixpkgs, ... }: 
    let
      inherit (nixpkgs) lib;
      supportedSystems = [ "x86_64-linux" ]; 
      forEachSupportedSystem = f:
        lib.genAttrs supportedSystems (system:
          f {
            inherit system;
            pkgs = import nixpkgs { inherit system; };
          });
    in rec {
      packages = forEachSupportedSystem ({ pkgs, ... }: {  
        default = pkgs.buildGoModule {
          pname = "webfinger";
          version = "0.0.1";

          src = lib.cleanSource ./.;

          vendorHash = null;

          doCheck = false;

          postInstall = ''
            mkdir -p $out/bin/resources
            cp -r $src/resources $out/bin
          '';
        };
      });

      devShells = forEachSupportedSystem ({ pkgs, ... }: with pkgs; {
        default = mkShell {
          buildInputs = [ go ];
        };
      });

      apps = forEachSupportedSystem ({ pkgs, system }: {
        default = {
          type = "app";
          program = "${packages.${system}.default}/bin/webfinger";
        };
      });
    };
}
