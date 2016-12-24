#!/bin/bash

# autostart
tee -a ~/.zshrc 2>/dev/null <<EOF
if [[ -z "\$TMUX" ]]; then
        tmux
fi
EOF

# plugin manager
git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm

# configuration
wget -P ~ https://raw.githubusercontent.com/jamrizzi/staz-ide/master/tmux/.tmux.conf
tmux source-file ~/.tmux.conf

git clone https://github.com/erikw/tmux-powerline.git ~/.tmux/
cp ~/.tmux/tmux-powerline/themes/default.sh ~/.tmux-theme.sh
$EDITOR ~/.tmux-theme.sh
~/.tmux/tmux-powerline/generate_rc.sh
mv ~/.tmux-powerlinerc.default ~/.tmux-powerlinerc
$EDITOR ~/.tmux-powerlinerc
