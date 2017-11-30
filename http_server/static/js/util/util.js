function logInfo(){
    var count = arguments.length    //把参数的长度保存的count的变量中
    var sum =""                      //声名一个总和的变量
    for(var i =0;i<count;i++){      //使用for循环把所有参数的长度遍历出来
        sum +=arguments[i]           //求出每个arguments所对应下标的参数的值相加
    }
    console.log(getNowFormatDate(),sum);
}

function getNowFormatDate() {
    var date = new Date();
    var seperator1 = "-";
    var seperator2 = ":";
    var month = date.getMonth() + 1;
    var strDate = date.getDate();
    if (month >= 1 && month <= 9) {
        month = "0" + month;
    }
    if (strDate >= 0 && strDate <= 9) {
        strDate = "0" + strDate;
    }
    var currentdate = date.getFullYear() + seperator1 + month + seperator1 + strDate
        + " " + date.getHours() + seperator2 + date.getMinutes()
        + seperator2 + date.getSeconds();
    return currentdate;
}
//表单验证
function validate_required(field,alerttxt)
{
  with (field){
	  if (value==null||value=="")
	    { 
	    	alert(alerttxt);
	    	return false;
	    }
	  else {
	  	return true;
	  }
  }
}