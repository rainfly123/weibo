#!/usr/bin/env python
#-*- coding: utf-8 -*- 

import MySQLdb
import time
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

def Query(ownerid):
    con = getConn()
    cur =  con.cursor()

    inited = []
    live = []
    stopped = []
    result = {'ownerid': ownerid}

    sql = "select live.liveid, title, snapshot, persons, date_format(startime, '%Y-%m-%d %H:%i') from live, owner \
          where live.liveid = owner.liveid and owner.ownerid = {0} and live.state = 0".format(ownerid)
    cur.execute(sql)
    res = cur.fetchall()

    for channel in res:
        temp = dict()
        temp['liveid'] = channel[0]
        temp['title'] = channel[1]
        temp['snapshot'] = channel[2]
        temp['persons'] = channel[3]
        temp['startime'] = channel[4]
        inited.append(temp)

    sql = "select live.liveid, title, snapshot, persons, rtmp_live_url from live, owner \
          where live.liveid = owner.liveid and owner.ownerid = {0} and live.state = 1".format(ownerid)
    cur.execute(sql)
    res = cur.fetchall()

    for channel in res:
        temp = dict()
        temp['liveid'] = channel[0]
        temp['title'] = channel[1]
        temp['snapshot'] = channel[2]
        temp['persons'] = channel[3]
        temp['rtmp_live_url'] = channel[4]
        live.append(temp)

    sql = "select live.liveid, title, snapshot, persons, playback_hls_url from live, owner\
          where live.liveid = owner.liveid and owner.ownerid = {0} and live.state = 2".format(ownerid)
    cur.execute(sql)
    res = cur.fetchall()

    for channel in res:
        temp = dict()
        temp['liveid'] = channel[0]
        temp['title'] = channel[1]
        temp['snapshot'] = channel[2]
        temp['supports'] = channel[3]
        temp['playback_hls_url'] = channel[4]
        stopped.append(temp)

    result['inited'] = inited
    result['live'] = live
    result['stopped'] =  stopped


    cur.close()
    con.close()
    return result

def QueryMy(ownerid):
    con = getConn()
    cur =  con.cursor()

    inited = list()

    sql = "select live.liveid, state, title, snapshot, tojson, supports, date_format(startime,'%Y-%m-%d %H:%i'),playback_hls_url \
          from live, owner where live.liveid = owner.liveid and owner.ownerid = {0} and live.state != 1 \
          order by live.startime desc".format(ownerid)
    cur.execute(sql)
    res = cur.fetchall()

    for channel in res:
        temp = dict()
        temp['liveid'] = channel[0]
        temp['state'] = channel[1]
        temp['title'] = channel[2]
        temp['snapshot'] = channel[3]
        temp['tojson'] = channel[4]
        temp['supports'] = channel[5]
        temp['startime'] = channel[6]
        temp['playback_hls_url'] = channel[7]
        inited.append(temp)
    cur.close()
    con.close()
    return inited


if __name__ ==  '__main__':
    #print Support("z1.mycs.xiechc",11233)
    import cjson
#    print cjson.encode(Products("13"))
#    print cjson.encode(Products("1"))
#    print Products("b")
    #SaveProducts("z1.66boss.xiechc",['1','2','3'])
    print HasLive('00')
    a = u"王"
    b =u"女"
    ApplyLive("1234", a.encode('utf-8'), b.encode('utf-8'), "133333333333", "afadf.jpg")
