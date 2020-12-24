$(document).ready(function (){
    $.ajax({
        type: "POST",
        url: "/api/get_borrower_message",
        contentType: "application/json;charset=utf-8",
        data: null,
        dataType: "json",
        success: function (data) {
            if (data["code"] !== 0) {
                alert(data["msg"]);
            } else {
                let h;
                if (data["data"]["CardNO"] === undefined) {
                    h = `
                    <div class="panel panel-primary">
                        <div class="panel-heading">UID</div>
                        <div class="panel-body">` + data["data"]["UID"] + `</div>
                    </div>
                    <div class="panel panel-primary">
                        <div class="panel-heading">CardNO</div>
                        <div class="panel-body">暂未绑定阅读证</div>
                    </div>
                    `;
                } else {
                    h = `
                    <div class="panel panel-primary">
                        <div class="panel-heading">UID</div>
                        <div class="panel-body">` + data["data"]["UID"] + `</div>
                    </div>
                    <div class="panel panel-primary">
                        <div class="panel-heading">借阅证号</div>
                        <div class="panel-body">` + data["data"]["CardNO"] + `</div>
                    </div>
                    <div class="panel panel-primary">
                        <div class="panel-heading">姓名</div>
                        <div class="panel-body">` + data["data"]["Name"] + `</div>
                    </div>
                    <div class="panel panel-primary">
                        <div class="panel-heading">专业</div>
                        <div class="panel-body">` + data["data"]["Major"] + `</div>
                    </div>
                    <div class="panel panel-primary">
                        <div class="panel-heading">借阅者类型</div>
                        <div class="panel-body">` + data["data"]["BorrowerType"] + `</div>
                    </div>
                    <div class="panel panel-primary">
                        <div class="panel-heading">最长借阅时间</div>
                        <div class="panel-body">` + data["data"]["Period"] + `</div>
                    </div>
                    <div class="panel panel-primary">
                        <div class="panel-heading">最大同时借阅数量</div>
                        <div class="panel-body">` + data["data"]["MaxBorrowerNumber"] + `</div>
                    </div>
                    `;
                }
                $("#Message").html(h);
            }
        }
    });
});