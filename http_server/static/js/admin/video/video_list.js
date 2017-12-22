var img_dailog=document.getElementById("img-dialog");
//初始化数据
paging(1);
//跳转至增加页面
function videoToAddHtml(){
	window.location.href="/admin/video/add";
}
//分页数据
function paging(pageNo){
	var fd = new FormData();
	fd.append("pageNo",pageNo);
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
	        var elem=document.getElementById("page_elem");
	        var table=elem.children[0];
	        var list=data.list;
	        var tr="<tbody>";
	        for(var i=0;i<list.length;i++){
	        	tr+="<tr>";
	        	tr+="<td>"+list[i].name+"</td>";
	        	tr+="<td>"+list[i].type+"</td>";
	        	var cover='"'+list[i].cover+'"';
	        	tr+="<td><a href='javascript:showImg("+cover+");'>"+list[i].cover+"</a></td>";
	        	var unixTimestamp = new Date(list[i].createTime) ;
	        	tr+="<td>"+unixTimestamp.toLocaleString()+"</td>";
	        	tr+="</tr>";
	        }
	        tr+="</tbody>";
	        var tbody=table.children;
	        if(tbody!=null && tbody.length==2){
	        	tbody[1].remove();
	        }
	        table.innerHTML+=tr;
	        //展示分页信息
	        if(data.isShow){
	        	var pageInfo=document.getElementById("page_info");
	        	if(pageInfo!=null){
	        		pageInfo.remove();
	        	}
	        	var classArr=["page-info btn page","page-info btn page disable"];
	        	var html="<div class='page-info' id='page_info'>";
		        html+="<ul class='page-ul' unselectable='unselectable'>";
		        html+="<li class='page-info text' >"+data.pageText+"</li>";
		        if(data.isHome){
		        	html+="<li class='"+classArr[1]+"' >上一页</li>";
		        }else{
		        	var pageNoT=data.pageNo-1;
		        	html+="<li class='"+classArr[0]+"' onclick='paging("+pageNoT+")'>上一页</li>";
		        }
		        html+="<li class='page-info text' >"+data.pageSizeText+"</li>";
		        if(data.isEnd){	
		        	html+="<li class='"+classArr[1]+"'>下一页</li>";
		        }else{
		        	var pageNoT=data.pageNo+1;
		        	html+="<li class='"+classArr[0]+"' onclick='paging("+pageNoT+")'>下一页</li>";
		        }
		        html+="</div>";
                var pageElem=document.getElementById("page_elem");
                var pageNextElem=pageElem.nextElementSibling;
		        if(pageNextElem!=null){
		        	pageNextElem.outerHTML+=html;
		        }else{
		        	pageElem.outerHTML+=html;
		        }
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
