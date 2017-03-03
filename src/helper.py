import sys
import os
import platform
from builtins import input
from pwd import getpwnam

class Helper:
    def prepare(self):
        if (platform.dist()[0] == 'centos'):
            os.system('''
            yum update -y
            yum install -y python-pip python-wheel
            ''')
        elif (platform.dist()[0] == 'Ubuntu'):
            os.system('''
            apt-get update -y
            apt-get install -y python-pip
            ''')
        else:
            print('Operating system not supported')
            sys.exit('Exiting installer')

    def is_root(self):
        if os.getuid() != 0:
            print('Requires root privileges')
            sys.exit('Exiting installer')

    def default_prompt(self, name, fallback):
        response = input(name + ' (' + fallback + '): ')
        assert isinstance(response, str)
        if (response):
            return response
        else:
            return fallback

    def boolean_prompt(self, name, fallback):
        default = 'Y|n'
        fallback = fallback.upper()
        if (fallback == 'N'):
            default = 'y|N'
        response = input(name + ' (' + default + '): ')
        assert isinstance(response, str)
        if (response):
            return response.upper()
        else:
            return fallback

    def append_to_user_file(self, path, content):
        user = os.environ['SUDO_USER'] if 'SUDO_USER' in os.environ else os.environ['USER']
        pid = os.fork()
        if pid == 0:
            try:
                os.setgid(getpwnam(user).pw_gid)
                os.setuid(getpwnam(user).pw_uid)
                f = open(filename, 'a+')
                for i in range(2):
                    f.write(content)
                f.close()
            finally:
                os._exit(0)

    def user_system(self, command):
        user = os.environ['SUDO_USER'] if 'SUDO_USER' in os.environ else os.environ['USER']
        pid = os.fork()
        if pid == 0:
            try:
                os.setgid(getpwnam(user).pw_gid)
                os.setuid(getpwnam(user).pw_uid)
                os.system(command)
            finally:
                os._exit(0)

    def find_replace(self, path, find, replace):
        filedata = None
        with open(path, 'r') as file:
            filedata = file.read()
            filedata = filedata.replace(find, replace)
        with open(path, 'w') as file:
            file.write(filedata)
