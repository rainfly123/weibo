#!/usr/bin/env python
#-*- coding: utf-8 -*- 

import MySQLdb
import time
import redis
import string
from DBUtils.PooledDB import PooledDB
import dbconfig


ERROR = {0:"OK", 1:"Parameter Error", 2:"Database Error", 3:u"您已赞", 4:u"你无权开通直播"}
Default_Snapshot = "http://7xvsyw.com2.z0.glb.qiniucdn.com/n.jpg"

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

def Query(key):
    con = getConn()
    cur =  con.cursor()

    users = []
    weibos = []
    sql = "select userid from user where name like '%{}%'".format(key)
    cur.execute(sql)
    res = cur.fetchall()

    for temp in res:
        users.append(temp[0])

    sql = "select weiboid from weibo where msg like '%{}%'".format(key)
    cur.execute(sql)
    res = cur.fetchall()

    for temp in res:
        weibos.append(temp[0])

    cur.close()
    con.close()
    return {"users":users, "weibos":weibos}

def SyncWeibo(weiboid):
    r = redis.Redis(host='localhost',port=6379,db=0)
    con = getConn()
    cur =  con.cursor()
    cur.execute("SET NAMES utf8mb4");
    con.commit()

    key = "weibo_" + weiboid
    weibo = r.hgetall(key)
    if weibo.has_key("weiboid"):
        Write(cur, weibo['weiboid'], weibo['msg'])

    con.commit()
    cur.close()
    con.close()


if __name__ ==  '__main__':
    #SyncWeibo("3220")
    print Query("9527测试")
