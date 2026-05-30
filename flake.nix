{
  description = "My rebel NixOS system";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.05";
  };

  outputs = { self, nixpkgs, ... }@inputs: {

    nixosConfigurations.nixos = nixpkgs.lib.nixosSystem {

      system = "x86_64-linux";
      specialArgs = { inherit inputs; };
      modules = [
        ./configuration.nix
      ];

    };

  };
}
