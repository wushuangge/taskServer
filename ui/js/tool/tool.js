AjaxTool={
	serviceURL:"http://192.168.51.12:8080",///v2/task
	getFormData:function(data){
        var form_data = new FormData();

        for (k in data){
            form_data.append(k,data[k]);
        }

        form_data.append("username",VarTool.GetCookie("username"));
        form_data.append("password",VarTool.GetCookie("password"));

        return form_data;
    },
    SendData:function(data, service){
        if(!(this.serviceURL)){
            webix.message({ type:"error", text:"Tool Ajax not init" });
            return;
        }

        var ajaxData = undefined;
        $.ajax({
            url: AjaxTool.serviceURL+service,
            type:'post',
            data: this.getFormData(data),
            processData: false,
            contentType: false,
            async:false,//同步
            success: function(returnData) {
                ajaxData = returnData;
            },
            error:function(err){
				
            }
        });

        return ajaxData;
    }
}

VarTool={
    GetQueryVariable:function(variable){
		var query = window.location.search.substring(1);
		var vars = query.split("&");
		for (var i=0;i<vars.length;i++) {
				var pair = vars[i].split("=");
				if(pair[0] == variable){return pair[1];}
		}
		return(undefined);
	},
	////////////////////////////////////////////////
	setCookie:function(value){
        value = "MXCookie" + "="+ encodeURIComponent(JSON.stringify(value));
        document.cookie = value;
    },
    getCookie:function(){
		var MXCookie = undefined;

        var arrStr = document.cookie.split("; "); 
        for (var i = 0; i < arrStr.length; i++) { 
            var temp = arrStr[i].split("="); 
            if (temp[0] == "MXCookie"){
                MXCookie = decodeURIComponent(temp[1]); 
            }
        } 

		if (MXCookie){
			return JSON.parse(MXCookie);
		}
	},
	SetCookie:function(name, value){
		var cookieData = this.getCookie();
		if (!(cookieData)){
			cookieData = {};
        }
        cookieData[name] = value;
		this.setCookie(cookieData);
	},
	GetCookie:function(name){
        var cookieData = this.getCookie();
		if (!(cookieData)){
			cookieData = {};
        }
		return cookieData[name];
	},
	ClearCookie:function(){
		this.setCookie({});
    }
}

var formatDateTime = function (date) {
    var y = date.getFullYear();
    var m = date.getMonth() + 1;
    m = m < 10 ? ('0' + m) : m;
    var d = date.getDate();
    d = d < 10 ? ('0' + d) : d;
    var h = date.getHours();
    h=h < 10 ? ('0' + h) : h;
    var minute = date.getMinutes();
    minute = minute < 10 ? ('0' + minute) : minute;
    var second=date.getSeconds();
    second=second < 10 ? ('0' + second) : second;
    return y + '-' + m + '-' + d+' '+h+':'+minute+':'+second;
};