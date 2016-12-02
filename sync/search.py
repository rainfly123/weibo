#!/usr/bin/env python
#-*- coding: utf-8 -*- 
import os
import cjson
import tornado
import tornado.ioloop
import tornado.web
import mysql
import datetime
from tornado.httputil import url_concat
from tornado.httpclient import AsyncHTTPClient
import tornado.web
import tornado.gen
import urllib
import string
import random
import json


class searchHandler(tornado.web.RequestHandler):
    def get(self):
        result = dict()
        key = self.get_argument("key")
        key = key.encode("utf-8")
        result = mysql.Query(key)
        self.write(cjson.encode(result))

class syncHandler(tornado.web.RequestHandler):
    def get(self):
        result = dict()
        weiboid = self.get_argument("weiboid")
        mysql.SyncWeibo(weiboid)
        self.write("OK")


application = tornado.web.Application([
    (r"/search", searchHandler),
    (r"/syncweibo", syncHandler),
])

if __name__ == "__main__":
    application.listen(6666)
    tornado.ioloop.IOLoop.instance().start()

