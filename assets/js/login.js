function getCookie(name){
	var r = document.cookie.match("\\b" + name + "=([^;]*)\\b");
	return r ? r[1] : undefined;
}
$(document).ready(function(){
	$('.form-login').submit(function(e){
		e.preventDefault();
		// 表单数据
		var data = {}
		$(this).serializeArray().map(function(x){data[x.name] = x.value;});
		var jsonData = JSON.stringify(data)

		$.ajax({
			url:"/login",
			type:"post",
			data:jsonData,
			contentType:"application/json",
			dataType:"json",
			headers:{
				"X-CSRFTOKEN":getCookie("csrf_token"),
			},
			success:function(data){
				if ("0"== data.errno){
					console.log(data)
					//alert("yes")
					location.href = "/admin/index";
					return;
				}else{
					console.log(data)
					$("#pwd-err span").html(data.errmsg)
					$("#pwd-err").show();
				}
			}
		})
	})
})
