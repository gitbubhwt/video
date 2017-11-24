var uploadProgress=document.getElementById("uploadProgress");
//上传文件
function uploadEvent(fileId){
	var elem=document.getElementById(fileId);
	if(elem.value==""){
		alert("请选择文件");
		return ;
	}
	var fileObj = document.getElementById(fileId).files[0];
	var name=fileObj.name;
	var path ="img/"
	if(name.indexOf('.')!=-1){
		var arr=name.split('.');
		name=Date.parse(new Date())+'.'+arr[1];
		if(arr[1].indexOf("mp4")!=-1){
			path="mp4/";
		}
	}
	uploadProgress.showModal();
	//创建xhr
    var xhr = new XMLHttpRequest();
    var url = "http://127.0.0.1:8080/admin/video/upload";
    //FormData对象
    var fd = new FormData();
    fd.append("name",name);
    fd.append("file", fileObj);
    fd.append("path",path);    //本人喜欢在参数中添加时间戳，防止缓存（--、）
    xhr.open("POST", url, true);
    //进度条部分
    xhr.upload.onprogress = function (evt) {
	    if (evt.lengthComputable) {
	        var percentComplete = Math.round(evt.loaded * 100 / evt.total);
	        document.getElementById('progress').value = percentComplete;
	    }
    };
    xhr.onreadystatechange = function () {
	    if (xhr.readyState == 4 && xhr.status == 200) {
	        uploadProgress.close();
	        document.getElementById('progress').value=0;
	        var response=xhr.responseText;
	        var data=JSON.parse(response);
	        if(data.code==-1){
	        	alert(data.msg);
	        }else if(data.code==0){
	        	document.getElementById(fileId+"_1").value=data.msg;
	        }
	    }
    };
    xhr.send(fd);
}

//保存
function saveEvent(){
	var form=document.getElementById("video-add-form");
	var fd = new FormData(form);
//	fd.append("name","hhhh");
	logInfo(fd);
	var url = "/admin/video/save";
	var xhr = new XMLHttpRequest();
	xhr.open("POST", url, true);
	xhr.onreadystatechange = function () {
	    if (xhr.readyState == 4 && xhr.status == 200) {
	        var response=xhr.responseText;
	        var data=JSON.parse(response);
	        if(data.code==-1){
	        	alert(data.msg);
	        }
	    }
    };
    xhr.send(fd);
}
