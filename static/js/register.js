$(document).ready(function(){
   // 注册
   $("#Submit").click(function(){
      let UserName = $("#UserName").val();
      let Password = $("#Password").val();
      let RePassword = $("#RePassword").val();
      if (Password !== RePassword) {
         alert("两次输入的密码不相等");
         return;
      }

      let Data = {
         UserName: UserName,
         Password: Password
      }

      $.ajax({
         type: "POST",
         url: "/api/create_borrower",
         contentType: "application/json; charset=utf-8",
         data: JSON.stringify(Data),
         dataType: "json",
         success: function (data) {
            if(data["code"] === 0) {
               window.location.replace("/login.html");
            } else {
               alert(data["msg"])
            }
         },
         error: function (message) {
            alert("error");
         }
      });
   })
});