1. 初始化：id生成器 初始化redis
2. 定义session结构体，要存什么信息
3. 给新用户分配身份牌
4. 根据id获取已存在Session，判断有没有失效
5. 核心操作
   1. 获取Session id
   2. 重置Session过期时间
   3. 从Redis中查找Session
   4. 给Session设置值，存入redis中
   5. 删除Session里的字段
   6. 用户退出登录销毁Session

