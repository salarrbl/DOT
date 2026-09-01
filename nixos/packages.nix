{ pkgs, ... }:



{
  environment.systemPackages = with pkgs; [
	#vimHugeX
	vim-full
	wget
	btop
	curl
	kitty
	python3
	gcc
	bat
	plank
	fastfetch
	lm_sensors
	xarchiver
	zip
	unzip
	unrar
	ranger
	dmenu
	htop
	zsh 
	fish
	tmux
	tor
	jdk
	rar
	torsocks
	wine
	telegram-desktop
	fzf
	eza
        lm_sensors
	rustc
        cargo
        rust-analyzer
        clippy
        rustfmt
	subfinder
	httpx
	nuclei
	jq
	alacritty
	yazi 
	zoxide
	xclip
        mpg123
        ddcutil
        ddcui
        ddccontrol
        gtk2
        glib
        gdk-pixbuf
        chromium
        vlc
        v2ray
        v2raya
        connect
        corkscrew
        tree
        proxychains
        espeak
        # discord
        chromium
        cacert
        nodejs
        gtk3
        nss
        neovim
        git
        go
        cava
        termius
        obs-studio
	xmind
        rofi
        mousepad
        xfce4-terminal
        spotify-player
        google-chrome
        whois
        ncdu
        ffmpeg
        ntfs3g
        opencode
        lazyssh
        bind
        docker
        openssl
        bind
        catppuccin-sddm
        sddm-astronaut
        exiftool
        playwright
        uv
        agent-browser
        sshpass
	obsidian
	git-lfs
	discord
	chafa
	mongosh
	mongodb-tools
	mongodb-ce
	duff
	iw
	aircrack-ng
	ayugram-desktop
	kotatogram-desktop
	geph
	gephgui-wry
	gnumake
	nodejs
	electron
	pkg-config
        super-productivity
  	sqlite
	vscode
  ];
fonts.packages = with pkgs; [
  noto-fonts
  noto-fonts-cjk-sans
  noto-fonts-color-emoji
  liberation_ttf
  fira-code
  fira-code-symbols
  mplus-outline-fonts.githubRelease
  dina-font
  proggyfonts
];

programs.nix-ld.libraries = with pkgs; [
   gtk4
  libxml2
  gdk-pixbuf
  freetype
  fontconfig
  harfbuzz
  libpng
  libjpeg_turbo
  libwebp
  # GLib / GTK
  glib
  gobject-introspection
  gtk3
  atk
  at-spi2-atk
  cairo
  pango

  # Chromium / Electron
  nss
  nspr
  dbus
  cups
  expat

  # X11
   libX11
   libXcomposite
   libXdamage
   libXext
   libXfixes
   libXrandr
   libxcb
   libxkbcommon
   libXinerama
   libXi
   libXcomposite
   libXdamage
   libXcursor
   zlib
   pkg-config

  # Graphics
  libgbm
  mesa
  mesa.drivers
  libglvnd

  # System / hardware
  systemd
  alsa-lib
  gtk3
  webkitgtk_4_1
  gobject-introspection
  libsoup_3


];
}
