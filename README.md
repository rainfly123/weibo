weibo
1 点赞
http://192.168.1.251:8888/support?login_user=2&weiboid=3

2 查询点赞
http://192.168.1.251:8888/checksupport?&weiboid=3
返回点赞者列表JSON["1","1","1","1","1","1","1","1","1","2"]

3 评论
http://192.168.1.251:8888/comment?login_user=2&weiboid=3&comment=%E6%A5%BC%E4%B8%BB%E5%84%BF%E5%82%BB%E9%80%BC

4 查询评论
http://192.168.1.251:8888/checkcomment?weiboid=3
返回评论列表JSON
[{"Author":"2","Comment":"“地方的”"},{"Author":"2","Comment":"“地方的”"},{"Author":"2","Comment":"“地方的”"},
{"Author":"2","Comment":"“地方的”"},{"Author":"2","Comment":"楼主傻逼"},{"Author":"2","Comment":"楼主儿傻逼"}]

5 写微博
http://192.168.1.251:8888/write?author=2&msg=%E4%BD%A0%E5%A5%BD%E6%88%91%E6%98%AF2%E5%8F%B7%E7%AC%A8%E8%9B%8B&pic=a.jpg,b.jpg

6 关注
http://192.168.1.251:8888/concern?login_user=1&concern=2
用户1，关注用户2，成为其粉丝

7 取消关注
http://192.168.1.251:8888/cancelconcern?login_user=1&cancel=2
用户1，取消关注用户2，不在是其粉丝

8 查询自己的微博
http://192.168.1.251:8888/checkmy?login_user=1
返回列表JSON[{"Weiboid":2,"Msg":"你好我是1号","Author":"1","Creatime":"2016-08-18 13:04:21","Supports":0,"Resent":0,"Pictures":["a.jpg","b.jpg"],"Comments":0},
{"Weiboid":1,"Msg":"你好","Author":"1","Creatime":"2016-08-18 13:04:07","Supports":0,"Resent":0,"Pictures":["a.jpg","b.jpg"],"Comments":0}]

9 查询自己关注的人的微博
http://192.168.1.251:8888/check?login_user=1
返回自己关注的人的最新微博
[{"Weiboid":5,"Msg":"你好我是2号笨蛋","Author":"2","Creatime":"2016-08-18 14:23:46","Supports":0,"Resent":0,"Pictures":["a.jpg","b.jpg"],"Comments":0},
{"Weiboid":4,"Msg":"你好我是2号笨蛋","Author":"2","Creatime":"2016-08-18 13:04:41","Supports":0,"Resent":0,"Pictures":["a.jpg","b.jpg"],"Comments":0},
{"Weiboid":3,"Msg":"你好我是2号","Author":"2","Creatime":"2016-08-18 13:04:31","Supports":10,"Resent":0,"Pictures":["a.jpg","b.jpg"],"Comments":4}]

10 上传文件
http://192.168.1.251:8888/upload

11 更新用户资料
http://192.168.1.251:8888/profile?login_user=xx&nickname=xxx&gender=xx&location=xx&signature=xxx

12 更新头像
http://192.168.1.251:8888/portrait?login_user=xxx    FORM multi-part 传输 ,name=file
13 转发微博
http://192.168.1.251:8888/forwared?login_user=xxx&msg=你好&origin=xxx  （origin原微博ID）

14 查询用户信息
http://192.168.1.251:8888/userinfo?userid=3  含推荐
返回用户信息
{"nickname":"”雨飞飞“","Gender":"男","Location":"北京","Signature":"我是圣人","Portrait":"http://192.168.1.251:8888/de90b859c68296bbd9f27c9d69187738.jpg",
"Follower":["1","2"],"Following":["1"],"Recommend":["1","6","15","5","3","7","9","14","11","12"]}

15 查询广场
http://192.168.1.251:8888/square?login_user=xx
返回：
[
    {
        "Weiboid": 10,
        "Msg": "我转发的2号的5号微博",
        "Author": "3",
        "Creatime": "2016-08-20 11:20:19",
        "Supports": 0,
        "Resent": 0,
        "Pictures": [
            ""
        ],
        "Comments": 0,
        "Origin": {
            "Weiboid": 5,
            "Msg": "你好我是2号笨蛋",
            "Author": "2",
            "Creatime": "2016-08-18 14:23:46",
            "Supports": 0,
            "Resent": 3,
            "Pictures": [
                "a.jpg",
                "b.jpg"
            ],
            "Comments": 1,
            "Origin": null,
            "Userinfo": {
                "nickname": "蓝天",
                "Gender": "太监",
                "Location": "广州",
                "Signature": "爱蓝天",
                "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
                "Follower": [
                    "1"
                ],
                "Following": [
                    "1",
                    "3"
                ],
                "Recommend": [
                    "1",
                    "6",
                    "15",
                    "5",
                    "13",
                    "7",
                    "9",
                    "14",
                    "11",
                    "12"
                ]
            }
        },
        "Userinfo": {
            "nickname": "”雨飞飞“",
            "Gender": "男",
            "Location": "北京",
            "Signature": "我是圣人",
            "Portrait": "http://192.168.1.251:8888/236118c52ae0a9403cd9e87041b07426.jpg",
            "Follower": [
                "1",
                "2"
            ],
            "Following": [
                "1"
            ],
            "Recommend": [
                "6",
                "15",
                "5",
                "13",
                "3",
                "7",
                "9",
                "14",
                "11",
                "12"
            ]
        }
    },
    {
        "Weiboid": 9,
        "Msg": "我转发的2号的5号微博",
        "Author": "3",
        "Creatime": "2016-08-20 11:20:18",
        "Supports": 0,
        "Resent": 0,
        "Pictures": [
            ""
        ],
        "Comments": 0,
        "Origin": {
            "Weiboid": 5,
            "Msg": "你好我是2号笨蛋",
            "Author": "2",
            "Creatime": "2016-08-18 14:23:46",
            "Supports": 0,
            "Resent": 3,
            "Pictures": [
                "a.jpg",
                "b.jpg"
            ],
            "Comments": 1,
            "Origin": null,
            "Userinfo": {
                "nickname": "蓝天",
                "Gender": "太监",
                "Location": "广州",
                "Signature": "爱蓝天",
                "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
                "Follower": [
                    "1"
                ],
                "Following": [
                    "1",
                    "3"
                ],
                "Recommend": [
                    "1",
                    "6",
                    "15",
                    "5",
                    "13",
                    "3",
                    "7",
                    "14",
                    "11",
                    "12"
                ]
            }
        },
        "Userinfo": {
            "nickname": "”雨飞飞“",
            "Gender": "男",
            "Location": "北京",
            "Signature": "我是圣人",
            "Portrait": "http://192.168.1.251:8888/236118c52ae0a9403cd9e87041b07426.jpg",
            "Follower": [
                "1",
                "2"
            ],
            "Following": [
                "1"
            ],
            "Recommend": [
                "1",
                "6",
                "15",
                "5",
                "13",
                "7",
                "9",
                "14",
                "11",
                "12"
            ]
        }
    },
    {
        "Weiboid": 8,
        "Msg": "我转发的2号的5号微博",
        "Author": "3",
        "Creatime": "2016-08-20 11:20:06",
        "Supports": 0,
        "Resent": 0,
        "Pictures": [
            ""
        ],
        "Comments": 0,
        "Origin": {
            "Weiboid": 5,
            "Msg": "你好我是2号笨蛋",
            "Author": "2",
            "Creatime": "2016-08-18 14:23:46",
            "Supports": 0,
            "Resent": 3,
            "Pictures": [
                "a.jpg",
                "b.jpg"
            ],
            "Comments": 1,
            "Origin": null,
            "Userinfo": {
                "nickname": "蓝天",
                "Gender": "太监",
                "Location": "广州",
                "Signature": "爱蓝天",
                "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
                "Follower": [
                    "1"
                ],
                "Following": [
                    "1",
                    "3"
                ],
                "Recommend": [
                    "1",
                    "6",
                    "15",
                    "5",
                    "13",
                    "7",
                    "9",
                    "14",
                    "11",
                    "12"
                ]
            }
        },
        "Userinfo": {
            "nickname": "”雨飞飞“",
            "Gender": "男",
            "Location": "北京",
            "Signature": "我是圣人",
            "Portrait": "http://192.168.1.251:8888/236118c52ae0a9403cd9e87041b07426.jpg",
            "Follower": [
                "1",
                "2"
            ],
            "Following": [
                "1"
            ],
            "Recommend": [
                "1",
                "15",
                "5",
                "13",
                "3",
                "7",
                "9",
                "14",
                "11",
                "12"
            ]
        }
    },
    {
        "Weiboid": 7,
        "Msg": "我转发的2号的5号微博",
        "Author": "3",
        "Creatime": "2016-08-20 10:34:21",
        "Supports": 0,
        "Resent": 0,
        "Pictures": [
            ""
        ],
        "Comments": 0,
        "Origin": {
            "Weiboid": 5,
            "Msg": "你好我是2号笨蛋",
            "Author": "2",
            "Creatime": "2016-08-18 14:23:46",
            "Supports": 0,
            "Resent": 3,
            "Pictures": [
                "a.jpg",
                "b.jpg"
            ],
            "Comments": 1,
            "Origin": null,
            "Userinfo": {
                "nickname": "蓝天",
                "Gender": "太监",
                "Location": "广州",
                "Signature": "爱蓝天",
                "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
                "Follower": [
                    "1"
                ],
                "Following": [
                    "1",
                    "3"
                ],
                "Recommend": [
                    "6",
                    "15",
                    "5",
                    "13",
                    "3",
                    "7",
                    "9",
                    "14",
                    "11",
                    "12"
                ]
            }
        },
        "Userinfo": {
            "nickname": "”雨飞飞“",
            "Gender": "男",
            "Location": "北京",
            "Signature": "我是圣人",
            "Portrait": "http://192.168.1.251:8888/236118c52ae0a9403cd9e87041b07426.jpg",
            "Follower": [
                "1",
                "2"
            ],
            "Following": [
                "1"
            ],
            "Recommend": [
                "1",
                "6",
                "5",
                "13",
                "3",
                "7",
                "9",
                "14",
                "11",
                "12"
            ]
        }
    },
    {
        "Weiboid": 7,
        "Msg": "我转发的2号的5号微博",
        "Author": "3",
        "Creatime": "2016-08-20 10:34:21",
        "Supports": 0,
        "Resent": 0,
        "Pictures": [
            ""
        ],
        "Comments": 0,
        "Origin": {
            "Weiboid": 5,
            "Msg": "你好我是2号笨蛋",
            "Author": "2",
            "Creatime": "2016-08-18 14:23:46",
            "Supports": 0,
            "Resent": 3,
            "Pictures": [
                "a.jpg",
                "b.jpg"
            ],
            "Comments": 1,
            "Origin": null,
            "Userinfo": {
                "nickname": "蓝天",
                "Gender": "太监",
                "Location": "广州",
                "Signature": "爱蓝天",
                "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
                "Follower": [
                    "1"
                ],
                "Following": [
                    "1",
                    "3"
                ],
                "Recommend": [
                    "1",
                    "6",
                    "15",
                    "5",
                    "13",
                    "3",
                    "7",
                    "14",
                    "11",
                    "12"
                ]
            }
        },
        "Userinfo": {
            "nickname": "”雨飞飞“",
            "Gender": "男",
            "Location": "北京",
            "Signature": "我是圣人",
            "Portrait": "http://192.168.1.251:8888/236118c52ae0a9403cd9e87041b07426.jpg",
            "Follower": [
                "1",
                "2"
            ],
            "Following": [
                "1"
            ],
            "Recommend": [
                "1",
                "6",
                "15",
                "5",
                "13",
                "3",
                "7",
                "14",
                "11",
                "12"
            ]
        }
    },
    {
        "Weiboid": 6,
        "Msg": "微博图片URl",
        "Author": "1",
        "Creatime": "2016-08-20 10:06:00",
        "Supports": 0,
        "Resent": 0,
        "Pictures": [
            "http://192.168.1.251:8888/de90b859c68296bbd9f27c9d69187738.jpg"
        ],
        "Comments": 0,
        "Origin": null,
        "Userinfo": {
            "nickname": "长城",
            "Gender": "男",
            "Location": "北京",
            "Signature": "万里长城永不倒",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": [
                "2",
                "3"
            ],
            "Following": [
                "2",
                "3"
            ],
            "Recommend": [
                "1",
                "15",
                "5",
                "13",
                "3",
                "7",
                "9",
                "14",
                "11",
                "12"
            ]
        }
    },
    {
        "Weiboid": 5,
        "Msg": "你好我是2号笨蛋",
        "Author": "2",
        "Creatime": "2016-08-18 14:23:46",
        "Supports": 0,
        "Resent": 3,
        "Pictures": [
            "a.jpg",
            "b.jpg"
        ],
        "Comments": 1,
        "Origin": null,
        "Userinfo": {
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": [
                "1"
            ],
            "Following": [
                "1",
                "3"
            ],
            "Recommend": [
                "1",
                "6",
                "15",
                "5",
                "13",
                "7",
                "9",
                "14",
                "11",
                "12"
            ]
        }
    },
    {
        "Weiboid": 4,
        "Msg": "你好我是2号笨蛋",
        "Author": "2",
        "Creatime": "2016-08-18 13:04:41",
        "Supports": 0,
        "Resent": 0,
        "Pictures": [
            "a.jpg",
            "b.jpg"
        ],
        "Comments": 0,
        "Origin": null,
        "Userinfo": {
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": [
                "1"
            ],
            "Following": [
                "1",
                "3"
            ],
            "Recommend": [
                "6",
                "15",
                "5",
                "13",
                "3",
                "7",
                "9",
                "14",
                "11",
                "12"
            ]
        }
    },
    {
        "Weiboid": 3,
        "Msg": "你好我是2号",
        "Author": "2",
        "Creatime": "2016-08-18 13:04:31",
        "Supports": 11,
        "Resent": 0,
        "Pictures": [
            "a.jpg",
            "b.jpg"
        ],
        "Comments": 5,
        "Origin": null,
        "Userinfo": {
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": [
                "1"
            ],
            "Following": [
                "1",
                "3"
            ],
            "Recommend": [
                "1",
                "6",
                "15",
                "5",
                "13",
                "3",
                "7",
                "9",
                "11",
                "12"
            ]
        }
    },
    {
        "Weiboid": 2,
        "Msg": "你好我是1号",
        "Author": "1",
        "Creatime": "2016-08-18 13:04:21",
        "Supports": 0,
        "Resent": 0,
        "Pictures": [
            "a.jpg",
            "b.jpg"
        ],
        "Comments": 1,
        "Origin": null,
        "Userinfo": {
            "nickname": "长城",
            "Gender": "男",
            "Location": "北京",
            "Signature": "万里长城永不倒",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": [
                "2",
                "3"
            ],
            "Following": [
                "2",
                "3"
            ],
            "Recommend": [
                "1",
                "6",
                "15",
                "5",
                "13",
                "3",
                "7",
                "9",
                "11",
                "12"
            ]
        }
    },
    {
        "Weiboid": 1,
        "Msg": "你好",
        "Author": "1",
        "Creatime": "2016-08-18 13:04:07",
        "Supports": 1,
        "Resent": 0,
        "Pictures": [
            "a.jpg",
            "b.jpg"
        ],
        "Comments": 0,
        "Origin": null,
        "Userinfo": {
            "nickname": "长城",
            "Gender": "男",
            "Location": "北京",
            "Signature": "万里长城永不倒",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": [
                "2",
                "3"
            ],
            "Following": [
                "2",
                "3"
            ],
            "Recommend": [
                "1",
                "6",
                "15",
                "5",
                "13",
                "3",
                "7",
                "9",
                "11",
                "12"
            ]
        }
    }
]
