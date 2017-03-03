#!/bin/bash

curl -L https://bootstrap.pypa.io/get-pip.py | python2.7
pip install future
git clone https://github.com/jamrizzi/staz-ide.git
python ./staz-ide/src/install.py
