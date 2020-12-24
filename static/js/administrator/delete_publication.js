$(document).ready(function (){
    let postFlag = false
    // 搜索
    $("#Submit").click(function (){
        if (postFlag === true) {
            return;
        }
        postFlag = true
        let name = $("#Search").val();
        let Data = {
            "Name": name
        };
        $.ajax({
            type: "POST",
            url: "/api/administrator_get_publication_by_name",
            contentType: "application/json; charset=utf-8",
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
                                    <th>库存/总量</th>
                                    <th>种类</th>
                                </tr>
                            </thead>
                            <tbody>
                       `;
                if (data["data"]["Publications"] == null) {
                    h += `
                            </tbody>
                        </table>
                        `;
                    h += `<h3>未查找到相关书籍</h3>`
                } else {
                    // 填充表格
                    $.each(data["data"]["Publications"], function (index, element){
                        h += `
                            <tr>
                                <input type="hidden" value="` + element["publication_id"] + `">
                                <td>` + (index + 1) + `</td>
                                <td>` + element["name"] + `</td>
                                <td>` + element["author"] + `</td>
                                <td>` + element["inventory"] + `/` + element["total"] + `</td>
                                <td>` + element["publication_type"]["publication_type"] + `</td>
                                <td><a href="javascript:void(0)" onclick="deletePublication(this)">删除</a></td>
                            </tr>
                            `
                    });
                    h += `
                            </tbody>
                        </table>
                        `;
                }
                $("#PublicationList").html(h);
            },
            error: function (message) {
                alert("error");
            }
        });
        postFlag = false;
    });
});

//删除
function deletePublication(element) {
    let id = parseInt($(element).parents("tr").children("input").val());
    $.ajax({
        type: "POST",
        url: "/api/delete_publication",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            "PublicationID": id
        }),
        dataType: "json",
        success: function (data) {
            alert(data["msg"]);
        },
        error: function (message) {
            alert("error");
        }
    });
}