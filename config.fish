if status is-interactive

    set -e SSH_ASKPASS
    # Commands to run in interactive sessions can go here
    alias token 'cat ~/.config/token | xclip -selection cilpboard'
    alias fastfetch 'fastfetch --logo ~/.config/fastfetch/gallaxy2'
    alias cls clear
    alias sik exit
    alias nixins 'sudo nvim /etc/nixos/packages.nix'
    alias by poweroff
    fish_vi_key_bindings
    alias l 'ls -al'
    alias .. 'cd ..'
    alias ... 'cd ../..'
    alias .... 'cd ../../..'
    alias burp 'java -jar  ~/Downloads/Burp.Suite.Professional.2025.12.5/BurpLoaderKeygen117.jar'
    alias salam 'espeak-ng -v en+f5  "salam salar , kefin"'
    alias kefin 'espeak-ng -v en+f5  "senin kefin"'
    alias salam-ver "espeak-ng -v en+f5 'kefin selin jon' "
    alias hg "history | grep $1"
    alias fh 'find . -name '
    alias .... 'cd ../..'
    alias gs 'git status'
    alias serverpass 'cat ~/.config/server | xclip -selection clipboard'
    alias cop ' xclip -selection clipboard'
    alias dhermes '~/.hermes/hermes-agent/apps/desktop/release/linux-unpacked/Hermes'
    alias cd z
    # Github abbr
    #
    abbr gs 'git status'
    abbr ga 'git add .'
    abbr gp 'xclip -selection clipboard ~/.config/token ; git push'
    abbr gpl 'git pull'
    abbr gc 'git commit -m ""'

    set SC /home/qarqa/Documents/Docs/Second-Brain
    set bet /home/qarqa/rebel/bett-usernames
    #source <(/home/qarqa/go/bin/watchdogs completion fish)
    /home/qarqa/go/bin/watchdogs completion fish | source
    set -gx PATH /home/qarqa/.npm-global/bin $PATH
    set -gx PATH $HOME/.local/bin $PATH
    set -gx PATH /home/qarqa/rebel/Tools/Huntool-scr $PATH
    # Tmux
    # # Start tmux automatically if not already inside it
    if status is-interactive
        if not set -q TMUX
            exec tmux
        end

        # zoxide
        zoxide init fish | source

    end
end

# opencode
fish_add_path /home/qarqa/.opencode/bin

# Qwen Code PATH block begin
set -gx PATH '/home/qarqa/.local/bin' $PATH
# Qwen Code PATH block end

# tabtab source for electron-forge package
# uninstall by removing these lines or running `tabtab uninstall electron-forge`
[ -f /home/qarqa/.npm_cache/_npx/6913fdfd1ea7a741/node_modules/tabtab/.completions/electron-forge.fish ]; and . /home/qarqa/.npm_cache/_npx/6913fdfd1ea7a741/node_modules/tabtab/.completions/electron-forge.fish
