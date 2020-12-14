

# 图书管理系统api

## 返回值约定
示例
```json
{
  "code": 0,
  "msg": "get data successfully",
  "data": {
              "PublicationID": 2,
              "ISBN": "123412341234",
              ......
          }
}
```
* code为-1代表服务器错误；为0代表正确返回；为1代表数据错误，可以把返回的msg反馈给用户。

* msg为消息提示。

* Session

| 用户类型 | session项 |
|:----:|:----:|
| 系统管理员 | administratorID<br>UserName<br>Roll="administrator"|
| 图书管理员 | LibrarianID<br>UserName<br>Roll="librarian"|
| 借阅者 | UID<br>CardNO<br>Name<br>Major<br>BorrowerType |

| 接口url | 功能描述 | 接收值 | 返回值 |
|:------:|:------:|:-----:|:-----:|
|/api/create_administrator|新增系统管理员|"UserName"<br/>"Password"|null|
|/api/login_administrator|系统管理员登陆|"UserName"<br>"Password"|null|
|/api/create_librarian|新增图书管理员|"UserName"<br/>"Password"|null|
|/api/login_librarian|图书管理员登录|"UserName"<br/>"Password"|null|
|/api/create_borrower|借阅者注册|"UserName"<br/>"Password"|null|
|/api/login_librarian|借阅者登录|"UserName"<br/>"Password"|null|

