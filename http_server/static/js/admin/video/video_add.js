var uploadProgress=document.getElementById("uploadProgress");
//上传文件
function uploadEvent(){
	uploadProgress.showModal();
	setTimeout(function(){
		uploadProgress.close();
	},5000);
	var upNode=uploadProgress.previousSibling;
}
