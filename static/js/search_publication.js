$(document).ready(function (){
    $("#Submit").click(function (){
        let name = $("#Search").val();
        let Data = {
            "Name": name
        };
        $.ajax({
            type: "POST",
            url: "/api/get_publication_by_name",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(Data),
            dataType: "json",
            success: function (data) {
                let h =
                    `
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
                // 填充表格
                $.each(data["data"]["Publications"], function (index, element){
                    h += `
                    <tr>
                        <input style="display: none" value="` + element["publication_id"] + `">
                        <td>` + (index + 1) + `</td>
                        <td>` + element["name"] + `</td>
                        <td>` + element["author"] + `</td>
                        <td>` + element["inventory"] + `/` + element["total"] + `</td>
                        <td>` + element["publication_type"]["publication_type"] + `</td>
                        <td><a>预定</a></td>
                    </tr>
                    `
                })
                h += `
                    </tbody>
                </table>
                `
                $("#PublicationList").html(h);
            },
            error: function (message) {
                alert("error");
            }
        })
    });
});