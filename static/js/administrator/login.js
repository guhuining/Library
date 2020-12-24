$(document).ready(function(){
    // 登录
    $("#Submit").click(function(){
        let UserName = $("#UserName").val();
        let Password = $("#Password").val();

        let Data = {
            UserName: UserName,
            Password: Password
        }

        $.ajax({
            type: "POST",
            url: "/api/login_administrator",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(Data),
            dataType: "json",
            success: function (data) {
                if(data["code"] === 0) {
                    // window.location.replace("/search_publication.html");
                } else {
                    alert(data["msg"])
                }
            },
            error: function (message) {
                alert(message);
            }
        });
    })
});