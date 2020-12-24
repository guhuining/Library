$(document).ready(function(){
    // 判断是否登录
    $.ajax({
        type: "POST",
        url: "/api/is_login",
        contentType: "application/json; charset=utf-8",
        data: null,
        dataType: "json",
        success: function (data) {
            if(data["code"] === -1) {
                window.location.replace("/login.html");
            }
        },
        error: function (message) {
            alert("error");
        }
    });

});