$(document).ready(function(){
    // 删除借阅证
    $("#Submit").click(function(){
        let Data = {
            CardNO: $("#CardNO").val()
        }

        $.ajax({
            type: "POST",
            url: "/api/delete_card",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(Data),
            dataType: "json",
            success: function (data) {
                alert(data["msg"]);
            },
            error: function (message) {
                alert(message);
            }
        });
    })
});