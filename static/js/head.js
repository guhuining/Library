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
                <div class="dropdown">
                    <button class="btn btn-default dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
                        用户
                        <span class="caret"></span>
                    </button>
                    <ul class="dropdown-menu" aria-labelledby="dropdownMenu1">
                        <li><a href="/user_info.html">个人中心</a></li>
                        <li><a id="Logout" href="javascript:void(0)" onclick="Logout()">退出</a></li>
                    </ul>
                </div>
                `
                $("#IsLogin").html(h)
            } else {
                let h = `
                <a id ="Login" href="/login.html">登录</a>
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
                window.location.replace("/login.html")
            } else {
                alert(data["msg"])
            }
        },
        error: function (message) {
            alert("error")
        }
    })
}