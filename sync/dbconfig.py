#!/usr/bin/env python

import ConfigParser
import os

class Parser:
    def __init__(self): 
        config = ConfigParser.ConfigParser()
        path = os.path.split(os.path.realpath(__file__))[0] + '/db.conf'
        config.read(path)
        self.config = config

    def getConfig(self, section, key):
        return self.config.get(section, key)

DBConfig = Parser()
