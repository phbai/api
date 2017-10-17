## 接口一览

-------
## 原始接口
### 获取影片
http://todayapp.tv/Movie/GetMovies
- PageIndex: 1
- PageSize: 100
- Type: 1   // 其他分类默认为 2
- ID: -1    // 分类，在影片分类处返回

### 搜索影片
http://todayapp.tv/Movie/GetMovies
- PageIndex: 1
- PageSize: 50
- Type: 1
- ID: -1
- Data: ももか

### 获取影片详情
http://todayapp.tv/Movie/GetMovieInfo
- MovieID: 5040679

### 播放影片
~~http://todayapp.tv/Movie/Play **需要登录**~~
~~- MovieID: 5040679~~
http://todayapi.com/vodapi.com
- data: {"Action":"PlayFreeMovie", "Message":{"UID": "62ME6IA5-BF46-WO84-LK7Q-AG3DMHMDPUL3", "Token": "F3207E3BD9E04A6DA69D3D5934132697", "MovieID": "5041971"}}

### 获取影片分类
http://todayapp.tv/Movie/GetClass
- PageIndex: 1
- PageSize: 50


### 获取演员分类
http://todayapp.tv/Movie/GetActor
- PageIndex: 1
- PageSize: 60

### 频道分类
http://todayapp.tv/Movie/GetChannel

### 获取Token
http://todayapi.com/vodapi.com
- data: {"Action": "CreateToken", "Message": {"UID": "2925541B-A946-4D6A-848A-9C11F7148E71"}}

### 登录用户
http://todayapp.tv/User/Login
- Token: A117482F65EA4240BE5C861084C682D3

------------
- /getChannel    获取频道              不用传参数
- /getActor      获取演员              size = 100, index = 1
- /getClass      获取类别              size = 100, index = 1
- /getMovieInfo  获取影片详细信息      id = 5041972
- /getMovies     获取所有影片          type = 1, index = 1, size = 100, keyword = テスト, id = -1
- /getToken      获取token             不用传参数(已过时 直接调/play即可)
- /play          获取详细影片m3u8地址  id = 5041972