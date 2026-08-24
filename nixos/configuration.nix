# Edit this configuration file to define what should be installed on
# your system.  Help is available in the configuration.nix(5) man page
# and in the NixOS manual (accessible by running ‘nixos-help’).

{ config, pkgs, ... }:

{
  imports =
    [ # Include the results of the hardware scan.
      ./hardware-configuration.nix
      ./packages.nix
       hermes-agent.homeManagerModules.default
  
    ];

  # Bootloader.
  boot.loader.systemd-boot.enable = true;
  boot.loader.efi.canTouchEfiVariables = true;

  networking.hostName = "nixos"; # Define your hostname.
  # networking.wireless.enable = true;  # Enables wireless support via wpa_supplicant.

  # Configure network proxy if necessary
  #networking.proxy.default = "socks5h://127.0.0.1:9050";
  # networking.proxy.noProxy = "127.0.0.1,localhost,internal.domain";

  # Enable networking
  networking.networkmanager.enable = true;

  # Set your time zone.
  time.timeZone = "Asia/Tehran";

  # Select internationalisation properties.
  i18n.defaultLocale = "en_US.UTF-8";

# monogo db
services.mongodb= {
 enable = true;
 package = pkgs.mongodb-ce;
};



# Docker
  # Enable the Docker service
virtualisation.docker.enable = true;

services.privoxy.enable = true;
#nix-ld
programs.nix-ld.enable = true;
  # Enable the X11 windowing system.
  # You can disable this if you're only using the Wayland session.
  services.xserver.enable = true;

  # Enable the KDE Plasma Desktop Environment.
services.displayManager.sddm = {
  enable = true;
  theme = "catppuccin-mocha-mauve";
  wayland.enable = true;
};
services.desktopManager.plasma6.enable = true;

  # Configure keymap in X11
  services.xserver.xkb = {
    layout = "us";
    variant = "";
  };

  # Enable CUPS to print documents.
#  services.printing.enable = true;

  # Enable sound with pipewire.
  #services.pulseaudio.enable = false;
  security.rtkit.enable = true;
  services.pipewire = {
    enable = true;
    alsa.enable = true;
    alsa.support32Bit = true;
    pulse.enable = true;
    # If you want to use JACK applications, uncomment this
    #jack.enable = true;

    # use the example session manager (no others are packaged yet so this is enabled by default,
    # no need to redefine it in your config for now)
    #media-session.enable = true;
  };

  # Enable touchpad support (enabled default in most desktopManager).
  # services.xserver.libinput.enable = true;

  # Define a user account. Don't forget to set a password with ‘passwd’.
  users.users.qarqa = {
    isNormalUser = true;
    description = "rebel";
    extraGroups = [ "networkmanager" "wheel" "docker" ];
    packages = with pkgs; [
      kdePackages.kate
    ];

    shell = pkgs.fish;
  };

programs.fish.enable = true;
# zoxide


  # Install firefox.
  programs.firefox.enable = true;

  # Allow unfree packages
  nixpkgs.config.allowUnfree = true;

  # List packages installed in system profile. To search, run:
  # $ nix search wget


# packges bulumide
# System auto upgrading

system.autoUpgrade.enable = false;
system.autoUpgrade.allowReboot = false;



# me serivce 
services.flatpak.enable = true;


## tor service
#services.tor = {
 # enable = true;
 # client.enable = true;
  #settings = {
      # Tor connects through your HTTP proxy
   #   HTTPProxy = "127.0.0.1:8085";
    #  HTTPSProxy = "127.0.0.1:8085";
   # };

#};
programs.ssh.extraConfig = ''
    Host *
      ProxyCommand nc -X connect -x 127.0.0.1:8085 %h %p
  '';

# Flakes
nix.settings.experimental-features = [ "nix-command" "flakes" ];

# XFCE

# ... rest of your configuration ...

#services.xserver.desktopManager.xfce.enable = true;
#services.displayManager.defaultSession = "xfce";

# Niri

#programs.niri = {
#  enable = true;
#  package = pkgs.niri;
#};

#services.displayManager.sessionPackages = [
#  pkgs.niri
#];
networking.nameservers=["208.67.222.222" "208.67.220.220"];

# polkit
security.polkit = {
  enable = true;
  enablePkexecWrapper = true;
};

# ollama
services.ollama = {
  enable = true;
};
# Hermes Agent

  programs.hermes-agent = {
    enable = true;
    desktop.enable = true;   # gives you the desktop app + launcher entry
  };
  services.hermes-agent = {
    enable = true;
    gateway.enable = true;
    settings.model.default = "anthropic/claude-sonnet-4";
    environmentFiles = [ config.sops.secrets."hermes-env".path ]; # or a plain file with your key
  };# Defualt shell
# Some programs need SUID wrappers, can be configured further or are

  # started in user sessions.
  # programs.mtr.enable = true;
  # programs.gnupg.agent = {
  #   enable = true;
  #   enableSSHSupport = true;
  # };

  # List services that you want to enable:

  # Enable the OpenSSH daemon.
  services.openssh.enable = true;

  # Open ports in the firewall.
  # networking.firewall.allowedTCPPorts = [ ... ];
  # networking.firewall.allowedUDPPorts = [ ... ];
  # Or disable the firewall altogether.
  # networking.firewall.enable = false;

  # This value determines the NixOS release from which the default
  # settings for stateful data, like file locations and database versions
  # on your system were taken. It‘s perfectly fine and recommended to leave
  # this value at the release version of the first install of this system.
  # Before changing this value read the documentation for this option
  # (e.g. man configuration.nix or on https://nixos.org/nixos/options.html).
  system.stateVersion = "25.11"; # Did you read the comment?


}
