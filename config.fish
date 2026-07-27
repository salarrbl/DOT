if status is-interactive

    set -e SSH_ASKPASS
    # Commands to run in interactive sessions can go here
    alias token 'cat ~/.config/token | xclip -selection cilpboard'
    alias cls 'clear'
    alias sik 'exit'
    alias nixins 'sudo nvim /etc/nixos/packages.nix'
    alias by 'poweroff'
    fish_vi_key_bindings
    alias ls 'eza --icons'
    alias l 'ls -al'
    alias .. 'cd ..'
    alias ... 'cd ../..'
    alias .... 'cd ../../..'
    alias burp 'java -jar  ~/Downloads/Burp.Suite.Professional.2025.12.5/BurpLoaderKeygen117.jar'
    alias fastfetch 'fastfetch --logo ~/.config/fastfetch/lain.txt' 
    alias salam 'espeak-ng -v en+f5  "salam salar , kefin"'
    alias kefin  'espeak-ng -v en+f5  "senin kefin"'
    alias salam-ver  "espeak-ng -v en+f5 'kefin selin jon' "
    alias hg "history | grep $1"
    alias fh 'find . -name '   
    alias .... 'cd ../..'
    alias gs 'git status'
    alias serverpass 'cat ~/.config/server | xclip -selection clipboard'
    alias cop ' xclip -selection clipboard'
    # Github abbr
    #
    abbr gs 'git status'
    abbr ga 'git add .'
    abbr gp 'xclip -selection clipboard ~/.config/token ; git push'
    abbr gpl 'git pull'
    abbr gc 'git commit -m ""'
    


    set  SC '/home/qarqa/Documents/Docs/Second-Brain'
    set bet '/home/qarqa/rebel/bett-usernames'
    #source <(/home/qarqa/go/bin/watchdogs completion fish)
    /home/qarqa/go/bin/watchdogs completion fish | source
    set -gx PATH /home/qarqa/.npm-global/bin $PATH
    set -gx PATH $HOME/.local/bin $PATH
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
