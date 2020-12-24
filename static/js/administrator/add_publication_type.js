$(document).ready(function(){
    // 提交出版物类型信息
    $("#Submit").click(function (){
        let Data = {
            "PublicationType": $("#Type").val(),
            "Fine": parseInt($("#Fine").val())
        }
        $.ajax({
            type: "POST",
            url: "/api/add_publication_type",
            contentType: "application/json;charset=utf-8",
            data: JSON.stringify(Data),
            dataType: "json",
            success: function (data) {
                alert(data["msg"]);
            },
            error: function (message) {
                alert("error");
            }
        });
    });
});