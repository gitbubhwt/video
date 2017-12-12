var img_dailog=document.getElementById("img-dialog");
//初始化数据
videoPageList(null);
//跳转至增加页面
function videoToAddHtml(){
	window.location.href="/admin/video/add";
}
//分页数据
function videoPageList(obj){
	var form=document.getElementById("video-add-form");
	var fd = new FormData();
	if(obj!=null){
		fd.append("pageNo",obj.pageNo);
	}else{
		fd.append("pageNo","0");
	}
	var url = "/admin/video/pageList";
	var xhr = new XMLHttpRequest();
	xhr.open("POST", url, true);
	xhr.onreadystatechange = function () {
	    if (xhr.readyState == 4 && xhr.status == 200) {
	        var response=xhr.responseText;
	        var data=JSON.parse(response);
	        if(data==null){
	        	return
	        }
	        var elem=document.getElementById("table_data_fill");
	        var list=data.list;
	        for(var i=0;i<list.length;i++){
	        	var tr="";
	        	tr+="<tr>";
	        	tr+="<td>"+list[i].name+"</td>";
	        	tr+="<td>"+list[i].type+"</td>";
	        	var cover='"'+list[i].cover+'"';
	        	tr+="<td><a href='javascript:showImg("+cover+");'>"+list[i].cover+"</a></td>";
	        	var unixTimestamp = new Date(list[i].createTime) ;
	        	tr+="<td>"+unixTimestamp.toLocaleString()+"</td>";
	        	tr+="</tr>";
	        	elem.innerHTML+=tr;
	        }
	    }
    };
    xhr.send(fd);
}

//展示图片
function showImg(data){
	logInfo("showImg");
	var img=document.getElementById("img-dailog-src");
	img.src=data;
	img_dailog.showModal();
}
