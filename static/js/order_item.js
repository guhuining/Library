$(document).ready(function (){
    $.ajax({
        type: "POST",
        url: "/api/borrower_get_order_item",
        contentType: "application/json;charset=utf-8",
        data: null,
        dataType: "json",
        success: function (data) {
            let h = `
                        <table class="table table-hover">
                            <thead>
                                <tr>
                                    <th>#</th>
                                    <th>书名</th>
                                    <th>作者</th>
                                    <th>库存/总量</th>
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
                                <td>` + element["Inventory"] + `/` + element["Total"] + `</td>
                                <td><a href="javascript:void(0)" onclick="cancelOrder(this)">取消预定</a></td>
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

})

// 取消订单
function cancelOrder(element) {
    let Data = {
        "OrderItemID": parseInt($(element).parents("tr").children("input").val())
    }
    $.ajax({
        type: "POST",
        url: "/api/cancel_order_item",
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