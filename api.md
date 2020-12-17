



# 图书管理系统api

## 返回值约定
示例
```json
{
  "code": 0,
  "msg": "get data successfully",
  "data": {
              "PublicationID": 2,
              "ISBN": "123412341234"
          }
}
```
* code为0代表正确返回；为1代表数据错误，可以把返回的msg反馈给用户。

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
|/api/add_publication_type|添加出版物类型|"PublicationType" ```出版物类别```<br/>"Fine" ```超期罚款```|null|
|/api/delete_publication_type|删除出版物类型|"PublicationType" ```出版物类别```|null|
|/api/add_publication|添加出版物|"Name" ```书名```<br/>"ISBN"```ISBN```<br/>"Price"```出版物价格```<br/>"Total"```总数```<br/>"PublicationType"```出版物类别```<br/>"Author"```作者```<br/>|null|
|/api/delete_publication|删除出版物|"PublicationID"```出版物ID```|null|
|/api/delete_librarian|删除图书管理员|"LibrarianID"```管理员ID```|null|
|/api/bind_card|绑定借阅证|"UID"```借阅者ID```<br/>"CardNO"```借阅证号码```<br/>"Name"```姓名```<br/>"Major"```专业```<br/>"BorrowerType"```借阅者类型```|null|
|/api/delete_card|删除借阅证|"CardNO"```借阅证号码```|null|

* /api/get_publication_by_name

* 通过书名查找出版物

* ```json
  {
      "Name": "书名"
  }
  ```

* ```json
  {
      "Publications": [
          {
              "publication_id": 1,
              "name": "C Primer Plus",
              "isbn": "978-7-115-13022-8",
              "price": 35,
              "total": 20,
              "inventory": 20,
              "publication_Type": {
                  "publication_type": "图书",
                  "fine": 0
              },
              "author": "Stephen Prata"
          },
          {
              "publication_id": 2,
              "name": "C Primer",
              "isbn": "978-7-115-13022-7",
              "price": 35,
              "total": 20,
              "inventory": 20,
              "publication_Type": {
                  "publication_type": "图书",
                  "fine": 0
              },
              "author": "Stephen Prata"
          }
      ]
      
  }
  ```
  
  

| 接口url                 | 功能描述                         | 接收值                                                 | 返回值                                     |
| ----------------------- | -------------------------------- | ------------------------------------------------------ | ------------------------------------------ |
| /api/borrow_publication | 借书                             | CardNO```借阅证号码```<br/>PublicationID```出版物ID``` | null                                       |
| /api/is_out_of_time     | 检查是否有图书逾期未还并返回罚金 | BorrowItemID```借阅订单ID```                           | Fine```罚金```(code=0)<br/>或null(code=-1) |
| /api/return_publication | 还书                             | BorrowItemID```借阅订单ID```                           | null                                       |

