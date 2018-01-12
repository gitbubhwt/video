//跳出iframe
if (window.parent.length > 0) {
    window.parent.location = location;
}
//后台登录
function adminLogin() {
    var form = document.getElementById("admin-login-form");
    var fd = new FormData(form);
    var b = validate_login_form(form);
    logInfo(b);
    if (!b) {
        return;
    }
    var url = "/admin/login";
    var xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.onreadystatechange = function() {
        var data = HandleAdminAjaxRequest(xhr);
        if (data == "") {
            return
        }
        if (data.code == -1) {
            alert(data.msg);
        } else {
            window.location.href = "/admin/index";
        }
    };
    xhr.send(fd);
    alert
}
//登录验证
function validate_login_form(thisform) {
    with(thisform) {
        if (validate_required(admin_userName, "请输入用户名") == false) {
            admin_userName.focus();
            return false;
        }
        if (validate_required(admin_pwd, "请输入密码") == false) {
            admin_pwd.focus();
            return false;
        }
        return true;
    }
}