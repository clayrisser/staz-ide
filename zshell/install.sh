#!/bin/bash
sh -c "$(wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)" &
sleep 10s
git clone https://github.com/bhilburn/powerlevel9k.git ~/.oh-my-zsh/custom/themes/powerlevel9k/
sed -i 's#ZSH_THEME="robbyrussell"#ZSH_THEME="powerlevel9k/powerlevel9k"#g' ~/.zshrc
mkdir ~/.fonts/
mkdir -p ~/.config/fontconfig/conf.d/
wget -P ~/.fonts/ https://github.com/powerline/powerline/raw/develop/font/PowerlineSymbols.otf
wget -P ~/.config/fontconfig/conf.d/ https://github.com/powerline/powerline/raw/develop/font/10-powerline-symbols.conf
fc-cache -vf ~/.fonts/
exec zsh
