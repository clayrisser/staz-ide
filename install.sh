#!/bin/bash

curl --silent --show-error --retry 5 https://bootstrap.pypa.io/get-pip.py | sudo python2.7
pip install future
git clone https://github.com/jamrizzi/staz-ide.git
python ./staz-ide/src/install.py
rm -r ./staz-ide
