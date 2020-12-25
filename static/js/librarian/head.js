$(document).ready(function(){
    // 判断是否登录
    $.ajax({
        type: "POST",
        url: "/api/is_login",
        contentType: "application/json; charset=utf-8",
        data: null,
        dataType: "json",
        success: function (data) {
            if(data["code"] === 0) {
                let h = `
                <li><a id="Logout" href="javascript:void(0)" onclick="Logout()">退出</a></li>
                `
                $("#IsLogin").html(h)
            } else {
                let h = `
                <li><a id ="Login" href="/librarian/login.html">登录</a></li>
                `
                $("#IsLogin").html(h)
            }
        },
        error: function (message) {
            alert("error");
        }
    });

});

//退出登录
function Logout(){
    $.ajax({
        type: "POST",
        url: "/api/logout",
        contentType: "application/json; charset=utf-8",
        data: null,
        dataType: "json",
        success: function (data) {
            if(data["code"] === 0) {
                window.location.replace("/librarian/login.html")
            } else {
                alert(data["msg"])
            }
        },
        error: function (message) {
            alert("error")
        }
    })
}