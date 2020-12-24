$(document).ready(function(){
    $.ajax({
        type: "POST",
        url: "/api/get_publication_type",
        contentType: "application/json;charset=utf-8",
        data: null,
        dataType: "json",
        success: function (data) {
            if (data["code"] === 0) {
                let h = `<select>`;
                $.each(data["data"]["PublicationType"], function (index, element) {
                    h += `<option>` + element["publication_type"] + `</option>`;
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
    // 提交书籍信息
    $("#Submit").click(function (){
        let Data = {
            "Name": $("#Name").val(),
            "ISBN": $("#ISBN").val(),
            "Price": parseInt($("#Price").val()),
            "Total": parseInt($("#Total").val()),
            "Author": $("#Author").val(),
            "PublicationType": $("#Type option:selected").val()
        }
        $.ajax({
            type: "POST",
            url: "/api/add_publication",
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