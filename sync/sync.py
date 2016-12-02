#!/usr/bin/env python
#-*- coding: utf-8 -*- 

import MySQLdb
import time
import redis
from DBUtils.PooledDB import PooledDB
import dbconfig


class DbManager():
    def __init__(self):
        kwargs = {}
        kwargs['host'] =  dbconfig.DBConfig.getConfig('database', 'dbhost')
        kwargs['port'] =  int(dbconfig.DBConfig.getConfig('database', 'dbport'))
        kwargs['user'] =  dbconfig.DBConfig.getConfig('database', 'dbuser')
        kwargs['passwd'] =  dbconfig.DBConfig.getConfig('database', 'dbpassword')
        kwargs['db'] =  dbconfig.DBConfig.getConfig('database', 'dbname')
        kwargs['charset'] =  dbconfig.DBConfig.getConfig('database', 'dbcharset')
        self._pool = PooledDB(MySQLdb, mincached=1, maxcached=15, maxshared=10, maxusage=10000, **kwargs)

    def getConn(self):
        return self._pool.connection()

_dbManager = DbManager()

def getConn():
    return _dbManager.getConn()

def Write(cur, weiboid, msg):
    msg =  MySQLdb.escape_string(msg)
    sql = """insert into weibo values('{0}', '{1}')""".format(weiboid, msg)
    try:
        cur.execute(sql)
    except:
        pass

def Users(cur, userid, name):
    name =  MySQLdb.escape_string(name)
    sql = """insert into user values('{0}', '{1}')""".format(userid, name)
    try:
        cur.execute(sql)
    except:
        pass

if __name__ ==  '__main__':
    r = redis.Redis(host='localhost',port=6379,db=0)
    #string
    globalid = r.get('globalID')
    globalid = int(globalid)
    con = getConn()
    cur =  con.cursor()
    cur.execute("SET NAMES utf8mb4");
    con.commit()

    for i in xrange(globalid):
        key = "weibo_" + str(i)
        weibo = r.hgetall(key)
        if not weibo.has_key("weiboid"):
            continue
        Write(cur, weibo['weiboid'], weibo['msg'])

    con.commit()


    allusers = r.smembers('all_users')

    for user in allusers:
        key = "user_" + user + "_profile"
        userinfo = r.hget(key, "nickname")
        if userinfo is None:
            continue
        Users(cur, user, userinfo)

    con.commit()

    cur.close()
    con.close()
