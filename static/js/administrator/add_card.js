$(document).ready(function(){
    $.ajax({
        type: "POST",
        url: "/api/get_borrower_type",
        contentType: "application/json;charset=utf-8",
        data: null,
        dataType: "json",
        success: function (data) {
            if (data["code"] === 0) {
                let h = `<select>`;
                $.each(data["data"]["BorrowerType"], function (index, element) {
                    h += `<option>` + element["borrower_type"] + `</option>`;
                });
                h += "</select>";
                $("#Type").html(h);
            } else {
                alert(data["msg"]);
            }
        },
        error: function (message) {
            alert("error")
        }
    });
    // 提交绑定信息
    $("#Submit").click(function (){
        let Data = {
            "UID": parseInt($("#UID").val()),
            "CardNO": $("#CardNO").val(),
            "Name": $("#Name").val(),
            "Major": $("#Major").val(),
            "BorrowerType": $("#Type option:selected").val()
        }
        $.ajax({
            type: "POST",
            url: "/api/bind_card",
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