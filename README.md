# Weibo API Summary
**1 点赞**
http://192.168.1.251:8888/support?login_user=2&weiboid=3


**1.1 取消点赞**
http://192.168.1.251:8888/unsupport?login_user=2&weiboid=3

**2 查询点赞**
http://192.168.1.251:8888/checksupport?&weiboid=3
返回点赞者列表JSON
```
{
    "code": 0,
    "message": "Succeeded",
    "data": [
        {
            "userid": "2",
            "nickname": "女神就是我",
            "gender": "女人",
            "location": "广州啊",
            "signature": "爱花城",
            "portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "follower": null,
            "following": null,
            "recommend": null
        },
        {
            "userid": "1",
            "nickname": "长城长",
            "gender": "男",
            "location": "北京",
            "signature": "万里长城永不倒",
            "portrait": "http://192.168.1.251:8888/44f4c56e25508cfa2909918e599a590b.jpg",
            "follower": null,
            "following": null,
            "recommend": null
        },
        {
            "userid": "2",
            "nickname": "女神就是我",
            "gender": "女人",
            "location": "广州啊",
            "signature": "爱花城",
            "portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "follower": null,
            "following": null,
            "recommend": null
        },
        {
            "userid": "2",
            "nickname": "女神就是我",
            "gender": "女人",
            "location": "广州啊",
            "signature": "爱花城",
            "portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "follower": null,
            "following": null,
            "recommend": null
        },
        {
            "userid": "4",
            "nickname": "山炮",
            "gender": "女",
            "location": "沈阳",
            "signature": "爱东北城",
            "portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "follower": null,
            "following": null,
            "recommend": null
        },
        {
            "userid": "2",
            "nickname": "女神就是我",
            "gender": "女人",
            "location": "广州啊",
            "signature": "爱花城",
            "portrait": "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg",
            "follower": null,
            "following": null,
            "recommend": null
        }
    ]
}
```

**3 评论**
http://192.168.1.251:8888/comment?login_user=2&weiboid=3&comment=%E6%A5%BC%E4%B8%BB%E5%84%BF%E5%82%BB%E9%80%BC

**4 查询评论**
http://192.168.1.251:8888/checkcomment?weiboid=3
`
返回评论列表JSON																	
```
{
    "Code: 0,
    "Message": "Succeeded" 
    "Data": 
	[{
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
    ]
}
```
**5 写微博**
http://192.168.1.251:8888/writev2?
author=2&msg=%E4%BD%A0%E5%A5%BD%E6%88%91%E6%98%AF2%E5%8F%B7%E7%AC%A8%E8%9B%8B
Multipart-Form   name="file0",name="file1"....name="file8"


**6 关注**
http://192.168.1.251:8888/concern?login_user=1&concern=2
用户1，关注用户2，成为其粉丝

**7 取消关注**
http://192.168.1.251:8888/cancelconcern?login_user=1&cancel=2
用户1，取消关注用户2，不在是其粉丝

**8 查询自己的微博**
http://192.168.1.251:8888/checkmy?login_user=1
返回列表JSON
```
{
    "code": 0,
    "message": "Succeeded",
    "data": [
        {
            "weiboid": 15,
            "msg": "你好",
            "author": "1",
            "creatime": "2016-09-10 13:38:40",
            "supports": 0,
            "resent": 0,
            "pictures": [
                "a.jpg",
                "b.jpg"
            ],
            "comments": 0,
            "origin": null,
            "user": {
                "userid": "1",
                "nickname": "长城长",
                "gender": "男",
                "location": "北京",
                "signature": "万里长城永不倒",
                "portrait": "http://192.168.1.251:8888/44f4c56e25508cfa2909918e599a590b.jpg",
                "follower": null,
                "following": null,
                "recommend": null
            }
        },
        {
            "weiboid": 10,
            "msg": "我是1号作者我要写文波",
            "author": "1",
            "creatime": "2016-09-07 17:17:39",
            "supports": 0,
            "resent": 1,
            "pictures": [
                "http://192.168.1.251:8888/18856e189b82d9c5d00421c498f7ce61.jpg"
            ],
            "comments": 0,
            "origin": null,
            "user": {
                "userid": "1",
                "nickname": "长城长",
                "gender": "男",
                "location": "北京",
                "signature": "万里长城永不倒",
                "portrait": "http://192.168.1.251:8888/44f4c56e25508cfa2909918e599a590b.jpg",
                "follower": null,
                "following": null,
                "recommend": null
            }
        },
        ]
     }
```

**9 查询自己关注的人的微博**
http://192.168.1.251:8888/check?login_user=1

返回自己关注的人的最新微博

```
   {
    "code": 0,
    "message": "Succeeded",
    "data": [
        {
            "weiboid": 15,
            "supported":true,
            "msg": "你好",
            "author": "1",
            "creatime": "2016-09-10 13:38:40",
            "supports": 0,
            "resent": 0,
            "pictures": [
                "a.jpg",
                "b.jpg"
            ],
            "comments": 0,
            "origin": null,
            "user": {
                "userid": "1",
                "nickname": "长城长",
                "gender": "男",
                "location": "北京",
                "signature": "万里长城永不倒",
                "portrait": "http://192.168.1.251:8888/44f4c56e25508cfa2909918e599a590b.jpg",
                "follower": null,
                "following": null,
                "recommend": null
            }
        },
        {
            "weiboid": 10,
            "msg": "我是1号作者我要写文波",
            "author": "1",
            "creatime": "2016-09-07 17:17:39",
            "supports": 0,
            "resent": 1,
            "pictures": [
                "http://192.168.1.251:8888/18856e189b82d9c5d00421c498f7ce61.jpg"
            ],
            "comments": 0,
            "origin": null,
            "user": {
                "userid": "1",
                "nickname": "长城长",
                "gender": "男",
                "location": "北京",
                "signature": "万里长城永不倒",
                "portrait": "http://192.168.1.251:8888/44f4c56e25508cfa2909918e599a590b.jpg",
                "follower": null,
                "following": null,
                "recommend": null
            }
        },
        ]
      }
```

**10 更新用户资料**
http://192.168.1.251:8888/profile?login_user=xx&nickname=xxx&gender=xx&location=xx&signature=xxx

**11 更新头像**
http://192.168.1.251:8888/portrait?login_user=xxx    FORM multi-part 传输 ,name="file"

**12 转发微博**
http://192.168.1.251:8888/forwared?login_user=xxx&msg=你好&origin=xxx  （origin原微博ID）

**13 查询用户信息**
http://192.168.1.251:8888/userinfo?userid=x&login_user=xx  含推荐
```
返回用户信息
{
    "code: 0,
    "message": "Succeeded" 
    "data": {
    "userid": "3",
    "concerned":true;
    "nickname": "”雨飞飞“",
    "gender": "男",
    "location": "北京",
    "signature": "我是圣人",
    "portrait": "http://192.168.1.251:8888/c086a66d2d3925ae2c015f5647200761.jpg",
    "follower": [
        {
            "Userid": "1",
            "concerned":false;
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
            "concerned":false;
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
    "following": [
        {
            "Userid": "1",
            "nickname": "长城",
            "concerned":true;
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
            "concerned":true;
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
}
```

**14 查询广场**
http://192.168.1.251:8888/square?login_user=xx
返回：
```
{
    "Code: 0,
    "Message": "Succeeded" 
    "Data": 
    [{
        "Weiboid": 12,
        "Concerned": true,
        "Supported": false,
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
   
]}
```
