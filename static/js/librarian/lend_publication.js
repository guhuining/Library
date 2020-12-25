$(document).ready(function () {
    // 搜索出版物
    $("#Search").click(function (){
        let Data = {
            "Name": $("#Name").val()
        };

        $.ajax({
            type: "POST",
            url: "/api/librarian_get_publication_by_name",
            contentType: "application/json;charset=utf-8",
            data: JSON.stringify(Data),
            dataType: "json",
            success: function (data) {
                if (data["code"] !== 0) {
                    alert(data["msg"]);
                } else {
                    let h = ``;
                    $.each(data["data"]["Publications"], function (index, element) {
                        h += `<option value="` + element["publication_id"] + `" id="PublicationID">` + element["name"] + `</option>`
                    });
                    $("#PublicationList").html(h);
                }
            },
            error: function (message) {
                alert("error")
            }
        });
    });

    // 借阅
    $("#Submit").click(function (){
        let Data = {
            CardNO: $("#CardNO").val(),
            PublicationID: parseInt($("#PublicationID").val())
        }
        $.ajax({
            type: "POST",
            url: "/api/borrow_publication",
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
})