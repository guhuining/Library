$(document).ready(function (){
    $("#Search").click(function () {
        let Data = {
            CardNO: $("#CardNO").val()
        }
        $.ajax({
            type: "POST",
            url: "/api/librarian_get_order_item",
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
                if (data["data"]["OrderItem"] == null) {
                    h += `
                            </tbody>
                        </table>
                        `;
                    h += `<h3>暂无订单</h3>`
                } else {
                    // 填充表格
                    $.each(data["data"]["OrderItem"], function (index, element){
                        h += `
                            <tr>
                                <input type="hidden" value="` + element["OrderItemID"] + `">
                                <td>` + (index + 1) + `</td>
                                <td>` + element["Name"] + `</td>
                                <td>` + element["Author"] + `</td>
                                <td><a href="javascript:void(0)" onclick="order(this)">借阅</a></td>
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

// 兑现订单
function order(element) {
    let Data = {
        "OrderItemID": parseInt($(element).parents("tr").children("input").val())
    };
    $.ajax({
        type: "POST",
        url: "/api/order_borrow",
        async: JSON.stringify(Data),
        contentType: "application/json;charset=utf-8",
        data: JSON.stringify(Data),
        dataType: "json",
        success: function (data){
            if (data["code"] === 0) {
                $(element).parents("tr").remove();
            }
            alert(data["msg"]);
        },
        error: function (message){
            alert("error");
        }
    });
}