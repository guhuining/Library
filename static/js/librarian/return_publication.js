$(document).ready(function (){
    $("#Search").click(function () {
        let Data = {
            CardNO: $("#CardNO").val()
        }
        $.ajax({
            type: "POST",
            url: "/api/get_borrow_item",
            contentType: "application/json;charset=utf-8",
            data: JSON.stringify(Data),
            dataType: "json",
            success: function (data) {
                let h = `
                    <table class="table table-hover">
                            <thead>
                                <tr>
                                    <th>#</th>
                                    <th>书名</th>
                                    <th>作者</th>
                                </tr>
                            </thead>
                            <tbody>
                    `;
                if (data["data"]["BorrowItems"] == null) {
                    h += `
                            </tbody>
                        </table>
                        `;
                    h += `<h3>暂无订单</h3>`
                } else {
                    // 填充表格
                    $.each(data["data"]["BorrowItems"], function (index, element){
                        h += `
                            <tr>
                                <input type="hidden" value="` + element["BorrowItemID"] + `">
                                <td>` + (index + 1) + `</td>
                                <td>` + element["Name"] + `</td>
                                <td>` + element["Author"] + `</td>
                                <td><a href="javascript:void(0)" onclick="returnPublication(this)">归还</a></td>
                                <td><a href="javascript:void(0)" onclick="lostPublication(this)">丢失</a></td>
                            </tr>
                            `
                    });
                    h += `
                            </tbody>
                        </table>
                        `;
                }
                $("#orderItems").html(h);
            },
            error: function (message) {
                alert("error");
            }
        });
    });



})

// 还书
function returnPublication(element) {
    let Data = {
        "BorrowItemID": parseInt($(element).parents("tr").children("input").val())
    };
    let is_out_of_time = false;
    let fine = 0;
    //检查是否超期
    $.ajax({
        type: "POST",
        url: "/api/is_out_of_time",
        async: false,
        contentType: "application/json;charset=utf-8",
        data: JSON.stringify(Data),
        dataType: "json",
        success: function (data) {
            if (data["code"] === -1) {
                is_out_of_time = true;
                fine = data["data"]["Fine"];
            }
        },
        error: function (message) {
            alert("error")
        }
    })
    if (is_out_of_time === true) {
        let fined = confirm("需缴纳超期罚款" + fine + "元")
        if (!fined) {
            return
        }
    }

    $.ajax({
        type: "POST",
        url: "/api/return_publication",
        async: false,
        contentType: "application/json;charset=utf-8",
        data: JSON.stringify(Data),
        dataType: "json",
        success: function (data){
            $(element).parents("tr").remove();
            alert(data["msg"]);
        },
        error: function (message){
            alert("error");
        }
    });
}

// 丢失
function lostPublication(element) {
    let Data = {
        "BorrowItemID": parseInt($(element).parents("tr").children("input").val())
    };
    let fine = 0;
    //查询图书价格
    $.ajax({
        type: "POST",
        url: "/api/get_price",
        async: false,
        contentType: "application/json;charset=utf-8",
        data: JSON.stringify(Data),
        dataType: "json",
        success: function (data) {
            if (data["code"] === 0) {
                fine = data["data"]["Price"];
            }
        },
        error: function (message) {
            alert("error")
        }
    })
    let fined = confirm("需缴纳丢失罚款" + fine + "元")
    if (!fined) {
        return
    }

    $.ajax({
        type: "POST",
        url: "/api/lost",
        async: false,
        contentType: "application/json;charset=utf-8",
        data: JSON.stringify(Data),
        dataType: "json",
        success: function (data){
            $(element).parents("tr").remove();
            alert(data["msg"]);
        },
        error: function (message){
            alert("error");
        }
    });
}