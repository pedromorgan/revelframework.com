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
CLONES_DIR = ROOT_DIR + "/externals"

env.hosts = ['revel-puppy.daffodil.uk.com']
env.user = "revel"
env.password = "using-ssh-ssl-key"
env.use_ssh_config = True # this is using ~/.ssh/config = sshkey login

repos = [
	"github.com/revel/revel",
	"github.com/revel/modules",
	"github.com/revel/cmd",
	"github.com/revel/revel.github.io"
]


def init_clones():
	"""Clones  repositories to externals/*"""
	for r in repos:
		with lcd(CLONES_DIR):
			local("git clone https://%s.git" % r)


def update_clones():
	"""Update clones to latest in dev branch for now"""
	for r in repos:
		parts = r.split("/")
		pth = CLONES_DIR + "/%s" % parts[-1]
		print "path=", pth
		with lcd(pth):
			local("git fetch")
			local("git checkout develop")
			local("git pull origin develop")

