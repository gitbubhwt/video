var uploadProgress=document.getElementById("uploadProgress");
var fileName =null;

//上传文件
function upload(ele){
	var parentElem=ele.parentNode;
	filterElem(parentElem);
	var file=parentElem.childNodes[1];//上传文件组件
	var btn=parentElem.childNodes[2];//按键组件
	var text=btn.innerText;
	if(text=="上传"){
		uploadFile(file,btn);
	}else if(text=="取消"){
//		delFile(file,btn);
	}
}

//上传文件
function uploadFile(file,btn){
	if(file.value==""){
		alert("请选择文件");
		return ;
	}
	var fileObj = file.files[0];
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
    var url = "/admin/video/upload";
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
	        	file.setAttribute("type","text");
	        	file.setAttribute("readonly","readonly");
                file.removeAttribute("accept");
	        	file.value=data.msg;
	        	
	        	btn.innerText="取消";
	        }
	    }
    };
    xhr.send(fd);
}

//删除文件
function delFile(file,btn){
	//创建xhr
    var xhr = new XMLHttpRequest();
    var url = "/admin/video/del";
    //FormData对象
    var fd = new FormData();
    fd.append("path",file.value);
    xhr.open("POST", url, true);
    xhr.onreadystatechange = function () {
	    if (xhr.readyState == 4 && xhr.status == 200) {
	        var response=xhr.responseText;
	        var data=JSON.parse(response);
	        if(data.code==-1){
	        	alert(data.msg);
	        }else if(data.code==0){
	        	file.setAttribute("type","file");
	        	file.removeAttribute("value");
	        	
	        	btn.innerText="上传";
	        }
	    }
    };
    xhr.send(fd);
}

//添加上传文件组件
function addUploadFile(ev){
	var elem=document.getElementById("muli-file");
	var div="<div><div class='video-add-elem'><span>文件</span><input type='file' class='file' name='video_child_file' accept='.mp4'/>";
	div+="<div class='video-btn' onclick='upload(this)'><span />上传</span></div>";
	div+="<div class='video-img-btn' onclick='removeUploadFile(this)'><img src='/img/ajx.png'/></div></div></div>";
	elem.outerHTML=div+elem.outerHTML;
}
//删除上传文件组件
function removeUploadFile(ev){
	var elem=ev.parentNode;
	elem.remove();
}
//保存
function saveEvent(){
	var form=document.getElementById("video-add-form");
	var fd = new FormData(form);
	var b=validate_form(form);
	logInfo(b);
	if(!b){
		return ;
	}
	var url = "/admin/video/save";
	var xhr = new XMLHttpRequest();
	xhr.open("POST", url, true);
	xhr.onreadystatechange = function () {
	    if (xhr.readyState == 4 && xhr.status == 200) {
	        var response=xhr.responseText;
	        var data=JSON.parse(response);
	        if(data.code==-1){
	        	alert(data.msg);
	        }else{
	        	alert("保存成功");
	        }
	    }
    };
    xhr.send(fd);
}

function validate_form(thisform)
{
with (thisform)
  {
    if (validate_required(video_name,"请输入名称")==false)
    {
    	video_name.focus();
    	return false;
    }
    if (validate_required(video_type,"请选择类型")==false)
    {
    	video_type.focus();
    	return false;
    }
    if (validate_required(video_cover,"请上传封面")==false)
    {
    	return false;
    }
    if (validate_required(video_file,"请上传文件")==false)
    {
    	return false;
    }
    return true;
  }
}