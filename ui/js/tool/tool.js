AjaxTool={
    serviceURL:window.location.origin.indexOf("http") == 0 ? window.location.origin : "https://192.168.51.12:8080",
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
        let query = window.location.search.substring(1);
        let vars = query.split("&");
        for (let i=0;i<vars.length;i++) {
            let pair = vars[i].split("=");
            if(pair[0] == variable){return pair[1];}
        }
        return(undefined);
    },
    setCookie:function(value, ip){
        value = "MXCookie" + "="+ encodeURIComponent(JSON.stringify(value));
        if (!(ip)){
            ip = window.location.host.slice(0, window.location.host.indexOf(":"));
        }
        document.cookie = value+"; path=/; domain="+ip;
    },
    getCookie:function(){
        let MXCookie = undefined;

        let arrStr = document.cookie.split("; ");
        for (let i = 0; i < arrStr.length; i++) {
            let temp = arrStr[i].split("=");
            if (temp[0] == "MXCookie"){
                MXCookie = decodeURIComponent(temp[1]);
            }
        }

        if (MXCookie){
            return JSON.parse(MXCookie);
        }
    },
    SetCookie:function(name, value){
        let cookieData = this.getCookie();
        if (!(cookieData)){
            cookieData = {};
        }
        cookieData[name] = value;
        this.setCookie(cookieData);
    },
    GetCookie:function(name){
        let cookieData = this.getCookie();
        if (!(cookieData)){
            cookieData = {};
        }
        return cookieData[name];
    },
    ClearCookie:function(){
        this.setCookie({});
    }
}
