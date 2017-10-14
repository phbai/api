## 接口一览

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
http://todayapp.tv/Movie/Play **需要登录**
- MovieID: 5040679

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
