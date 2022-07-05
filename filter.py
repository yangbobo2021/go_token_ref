
import sys

with open(sys.argv[1], 'r') as f:
    lines = f.readlines(-1)

state = 0
packages = {}

for line in lines:
    if line.find('----<<end>>----') >= 0:
        state = 0

    if state == 1:
        pkg = line.strip()
        if pkg.find('github.com/merico-dev') >= 0:
            pkg = 'github.com/merico-dev'
        if pkg not in packages:
            packages[pkg] = 0
        packages[pkg] += 1
    
    if line.find('----<<begin2>>----') >= 0:
        state = 1

print(packages)