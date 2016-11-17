#!/bin/bash
cp /home/jamrizzi/Projects/devos-user/tmux/.tmux.conf ~/.tmux.conf
tee -a ~/.zshrc 2>/dev/null <<EOF
if [[ -z "\$TMUX" ]]; then
        tmux
fi
EOF
tmux source-file ~/.tmux.conf
