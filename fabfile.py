# -*- coding: utf-8 -*-


"""
This is a workaround, later this stuff need to be in a revel job
for now its a quick way to get git clones of external repos
"""

import os
import sys



from fabric.api import env, local, run, cd, lcd, sudo, prompt
from fabric import colors


ROOT_DIR = os.path.dirname(os.path.realpath(__file__))
WORK_DIR = ROOT_DIR + "/workspace"

env.hosts = ['revel-www.daffodil.uk.com']
env.user = "revel"
env.password = "using-ssh-ssl-key"
env.use_ssh_config = True # this is using ~/.ssh/config = sshkey login

## repositories with url / branch
repos = [
	["github.com/revel/revel", "develop"],
	["github.com/revel/modules", "develop"],
	["github.com/revel/cmd", "develop"],
	#["github.com/revel/samples", "master"],
	#//["github.com/revel/revel.github.io", "develop"]
	["github.com/pedromorgan/revel.github.io", "www_dev"]
]


def goget():
	"""Read go_deps.txt and runs go get.. wtf"""
	f = open(ROOT_DIR + "/go_deps.txt", "r")
	s = f.read()
	f.close()
	for line in s.split("\n"):
		naked = line.strip()
		if naked != "":
			local("go get %s" % naked)

def init_clones():
	"""Clones  repositories to externals/*"""
	print "#>>>> Initialising clones"
	for r in repos:
		with lcd(WORK_DIR):
			parts = r[0].split("/")
			print "repo=", WORK_DIR + "/" + parts[-1]
			if not os.path.exists(WORK_DIR + "/" + parts[-1]):
				local("git clone https://%s.git" % r[0])


def update_clones():
	"""Update clones to latest in dev branch for now"""
	for r in repos:
		parts = r[0].split("/")
		pth = WORK_DIR + "/%s" % parts[-1]
		print "path=", pth
		with lcd(pth):
			local("git fetch")
			local("git checkout %s" % r[1])
			local("git pull origin %s" % r[1])

