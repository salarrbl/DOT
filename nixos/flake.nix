{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    hermes-agent.url = "github:NousResearch/hermes-agent";

    # your other inputs...
  };

  outputs = { nixpkgs, hermes-agent, ... }:
    {
      nixosConfigurations.your-host = nixpkgs.lib.nixosSystem {
        system = "x86_64-linux";

        modules = [
          hermes-agent.nixosModules.default
          ./configuration.nix
        ];
      };
    };
}
