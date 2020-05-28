# 用户接口
* [前端网站](https://account.ys1994.nl/)
* [前端项目](https://github.com/digi1874/account.ys1994)

## 构建
> 依赖
> 1. go 1.13+
> 2. mysql (4.1+，本项目开发时使用5.7；库需要设置为utf8mb4)
> 3. 项目根目录下创建文件db.json连接mysql数据库
>> ./db.json
>> ```
>> {
>>   "user": "用户名",
>>   "password": "密码",
>>   "localhost": "地址",
>>   "databaseName": "库名"
>> }
>> ```
>> #
> ```
> # 开发，开启http://localhost:8031/
> $ go run main.go -env=dev
>
> # 生产程序
> $ go build
> ```
> #

## 接口说明
> * response status code:
>> 1. 200: 确定
>> 2. 400: 错误
>> 3. 401: 无权限，token无效
>> 4. 404: 不存在
> * response data msg: 回应说明
> #
> 1. 获取
