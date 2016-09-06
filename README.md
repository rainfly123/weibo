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
`
返回评论列表JSON
[
    {
        "Author": {
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        },
        "Comment": "“地方的”"
    },
    {
        "Author": {
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        },
        "Comment": "“地方的”"
    },
    {
        "Author": {
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        },
        "Comment": "“地方的”"
    },
    {
        "Author": {
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        },
        "Comment": "“地方的”"
    },
    {
        "Author": {
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        },
        "Comment": "楼主傻逼"
    },
    {
        "Author": {
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        },
        "Comment": "楼主儿傻逼"
    },
    {
        "Author": {
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        },
        "Comment": "楼主儿傻逼"
    }
]`

5 写微博
http://192.168.1.251:8888/write?author=2&msg=%E4%BD%A0%E5%A5%BD%E6%88%91%E6%98%AF2%E5%8F%B7%E7%AC%A8%E8%9B%8B&pic=a.jpg,b.jpg
5.1 写微博
http://192.168.1.251:8888/writev2?
author=2&msg=%E4%BD%A0%E5%A5%BD%E6%88%91%E6%98%AF2%E5%8F%B7%E7%AC%A8%E8%9B%8B
Multipart-Form   file0,file1....

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
{
    "Userid": "3",
    "nickname": "”雨飞飞“",
    "Gender": "男",
    "Location": "北京",
    "Signature": "我是圣人",
    "Portrait": "http://192.168.1.251:8888/c086a66d2d3925ae2c015f5647200761.jpg",
    "Follower": [
        {
            "Userid": "1",
            "nickname": "长城",
            "Gender": "男",
            "Location": "北京",
            "Signature": "万里长城永不倒",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        },
        {
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        }
    ],
    "Following": [
        {
            "Userid": "1",
            "nickname": "长城",
            "Gender": "男",
            "Location": "北京",
            "Signature": "万里长城永不倒",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        }
    ],
    "Recommend": [
        {
            "Userid": "6",
            "nickname": "红嘴鸥",
            "Gender": "女",
            "Location": "昆明",
            "Signature": "滇池喂鸟",
            "Portrait": "http://192.168.1.251:8888/52b0d4fa3eae0814dae50d0d7ac3700a.jpg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        },
        {
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        }
    ]
}

15 查询广场
http://192.168.1.251:8888/square?login_user=xx
返回：
[
    {
        "Weiboid": 12,
        "Msg": "我是1号",
        "Author": "1",
        "Creatime": "2016-08-25 10:27:22",
        "Supports": 0,
        "Resent": 0,
        "Pictures": [
            "498115ec3419eb15140873d1bf1fdcb4.jpg"
        ],
        "Comments": 0,
        "Origin": null,
        "Userinfo": {
            "Userid": "1",
            "nickname": "长城",
            "Gender": "男",
            "Location": "北京",
            "Signature": "万里长城永不倒",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        }
    },
    {
        "Weiboid": 11,
        "Msg": "我是6号",
        "Author": "6",
        "Creatime": "2016-08-25 09:56:49",
        "Supports": 0,
        "Resent": 0,
        "Pictures": [
            "498115ec3419eb15140873d1bf1fdcb4.jpg"
        ],
        "Comments": 0,
        "Origin": null,
        "Userinfo": {
            "Userid": "6",
            "nickname": "红嘴鸥",
            "Gender": "女",
            "Location": "昆明",
            "Signature": "滇池喂鸟",
            "Portrait": "http://192.168.1.251:8888/52b0d4fa3eae0814dae50d0d7ac3700a.jpg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        }
    },
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
                "Userid": "2",
                "nickname": "蓝天",
                "Gender": "太监",
                "Location": "广州",
                "Signature": "爱蓝天",
                "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
                "Follower": null,
                "Following": null,
                "Recommend": null
            }
        },
        "Userinfo": {
            "Userid": "3",
            "nickname": "”雨飞飞“",
            "Gender": "男",
            "Location": "北京",
            "Signature": "我是圣人",
            "Portrait": "http://192.168.1.251:8888/c086a66d2d3925ae2c015f5647200761.jpg",
            "Follower": null,
            "Following": null,
            "Recommend": null
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
                "Userid": "2",
                "nickname": "蓝天",
                "Gender": "太监",
                "Location": "广州",
                "Signature": "爱蓝天",
                "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
                "Follower": null,
                "Following": null,
                "Recommend": null
            }
        },
        "Userinfo": {
            "Userid": "3",
            "nickname": "”雨飞飞“",
            "Gender": "男",
            "Location": "北京",
            "Signature": "我是圣人",
            "Portrait": "http://192.168.1.251:8888/c086a66d2d3925ae2c015f5647200761.jpg",
            "Follower": null,
            "Following": null,
            "Recommend": null
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
                "Userid": "2",
                "nickname": "蓝天",
                "Gender": "太监",
                "Location": "广州",
                "Signature": "爱蓝天",
                "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
                "Follower": null,
                "Following": null,
                "Recommend": null
            }
        },
        "Userinfo": {
            "Userid": "3",
            "nickname": "”雨飞飞“",
            "Gender": "男",
            "Location": "北京",
            "Signature": "我是圣人",
            "Portrait": "http://192.168.1.251:8888/c086a66d2d3925ae2c015f5647200761.jpg",
            "Follower": null,
            "Following": null,
            "Recommend": null
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
                "Userid": "2",
                "nickname": "蓝天",
                "Gender": "太监",
                "Location": "广州",
                "Signature": "爱蓝天",
                "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
                "Follower": null,
                "Following": null,
                "Recommend": null
            }
        },
        "Userinfo": {
            "Userid": "3",
            "nickname": "”雨飞飞“",
            "Gender": "男",
            "Location": "北京",
            "Signature": "我是圣人",
            "Portrait": "http://192.168.1.251:8888/c086a66d2d3925ae2c015f5647200761.jpg",
            "Follower": null,
            "Following": null,
            "Recommend": null
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
                "Userid": "2",
                "nickname": "蓝天",
                "Gender": "太监",
                "Location": "广州",
                "Signature": "爱蓝天",
                "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
                "Follower": null,
                "Following": null,
                "Recommend": null
            }
        },
        "Userinfo": {
            "Userid": "3",
            "nickname": "”雨飞飞“",
            "Gender": "男",
            "Location": "北京",
            "Signature": "我是圣人",
            "Portrait": "http://192.168.1.251:8888/c086a66d2d3925ae2c015f5647200761.jpg",
            "Follower": null,
            "Following": null,
            "Recommend": null
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
            "Userid": "1",
            "nickname": "长城",
            "Gender": "男",
            "Location": "北京",
            "Signature": "万里长城永不倒",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
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
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
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
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
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
            "Userid": "2",
            "nickname": "蓝天",
            "Gender": "太监",
            "Location": "广州",
            "Signature": "爱蓝天",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
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
            "Userid": "1",
            "nickname": "长城",
            "Gender": "男",
            "Location": "北京",
            "Signature": "万里长城永不倒",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
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
            "Userid": "1",
            "nickname": "长城",
            "Gender": "男",
            "Location": "北京",
            "Signature": "万里长城永不倒",
            "Portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "Follower": null,
            "Following": null,
            "Recommend": null
        }
    }
]
