$(document).ready(function (){
    $.ajax({
        type: "POST",
        url: "/api/borrower_get_borrowed_item",
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
                                    <th>借阅时间</th>
                                </tr>
                            </thead>
                            <tbody>
                       `;
            if (data["code"] !== 0){
                alert(data["msg"])
            } else if (data["data"]["BorrowItems"] == null) {
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
                                <td>` + element["BorrowDate"] + `</td>
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