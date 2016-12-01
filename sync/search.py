#!/usr/bin/env python
import os
import cjson
import tornado
import tornado.ioloop
import tornado.web
import mysql
import qiniu
import datetime
from tornado.httputil import url_concat
from tornado.httpclient import AsyncHTTPClient
import tornado.web
import tornado.gen
import urllib
import string
import random
import json


class SupportHandler(tornado.web.RequestHandler):
    def get(self):
        result = dict()
        liveid = self.get_arguments("liveid", strip=True)
        userid = self.get_arguments("userid", strip=True)
        try:
            result = mysql.Support(liveid[0], userid[0])
        except:
            result['message'] = "parameter error"
            result['code'] = 1

        self.write(cjson.encode(result))

class ProductHandler(tornado.web.RequestHandler):
    @tornado.web.asynchronous
    @tornado.gen.coroutine
    def get(self):
        URL = "http://webapi/app/supplier.php?act=get_goods_detail"
        result = dict()
        liveid = self.get_arguments("liveid", strip=True)
        try:
            result = mysql.Products(liveid[0])
        except:
            result['message'] = "parameter error"
            result['code'] = 1
            self.write(cjson.encode(result))
            return

        result['goods_detail'] = list()
        for good in result['goods']:
            client = tornado.httpclient.AsyncHTTPClient()
            temp = dict()
            temp['goods_id'] = good
            data = urllib.urlencode(temp)
            response = yield client.fetch(URL, method='POST', body=data)
            response = json.loads(response.body)
            result['goods_detail'].append(response['result'])

        self.write(cjson.encode(result))


application = tornado.web.Application([
    (r"/querymy", QueryMyHandler),
    (r"/delete", DeleteHandler),
])

if __name__ == "__main__":
    application.listen(6666)
    tornado.ioloop.IOLoop.instance().start()

