$(document).ready(function (){
    // 搜索
    $.ajax({
        type: "POST",
        url: "/api/get_publication_type",
        contentType: "application/json; charset=utf-8",
        data: null,
        dataType: "json",
        success: function (data) {
            let h = `
                    <table class="table table-hover">
                            <thead>
                                <tr>
                                    <th>#</th>
                                    <th>出版物类型</th>
                                    <th>罚金</th>
                                </tr>
                            </thead>
                            <tbody>
                    `;
            if (data["data"]["PublicationType"] == null) {
                h += `
                         </tbody>
                     </table>   
                     `;
                h += `<h3>暂无出版物类型</h3>`
            } else {
                // 填充表格
                $.each(data["data"]["PublicationType"], function (index, element){
                    h += `
                            <tr>
                                <input type="hidden" value="` + element["publication_type"] + `">
                                <td>` + (index + 1) + `</td>
                                <td>` + element["publication_type"] + `</td>
                                <td>` + element["fine"] + `</td>
                                <td><a href="javascript:void(0)" onclick="deletePublicationType(this)">删除</a></td>
                            </tr>
                            `
                });
                h += `
                            </tbody>
                        </table>
                        `;
            }
            $("#PublicationTypeList").html(h);
        },
        error: function (message) {
            alert("error");
        }
    });
});

//删除
function deletePublicationType(element) {
    let publicationType = $(element).parents("tr").children("input").val();
    $.ajax({
        type: "POST",
        url: "/api/delete_publication_type",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            "PublicationType": publicationType
        }),
        dataType: "json",
        success: function (data) {
            $(element).parents("tr").remove();
            alert(data["msg"]);
        },
        error: function (message) {
            alert("error");
        }
    });
}