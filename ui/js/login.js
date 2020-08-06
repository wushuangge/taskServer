Login = {
    webixData:{
        view:"form",
        borderless:true,
        elements: [
            { view:"text", id: "username", label:'用户名', name:"username", on:{
                onEnter: function(ev){
                    this.blur();
                    $$("password").focus();
                }
            }},
            { view:"text", id: "password", label:'密码', name:"password", type:"password", on:{
                onEnter: function(ev){ 
                    this.blur();
					$$("queren").focus();
                }
            }},
			{ view:"button",id:"queren", value: "确认", click:function(){
                let parent   = $$("LoginForm").getChildViews()[1]
                let username = $$("username").getValue();
                let password = $$("password").getValue();
				if (parent.validate()){
                    VarTool.SetCookie("username", username);
                    VarTool.SetCookie("password", password);
                    Login.login();
                }else
                    webix.message({ type:"error", text:"Form data is empty" });
            }} 
        ],
        rules:{
            "username":webix.rules.isNotEmpty,
            "password":webix.rules.isNotEmpty
        },
        elementsConfig:{
            labelPosition:"top",
        }
    },
    show:function(){
        if ($$("LoginForm")){
            $$("LoginForm").destructor();
        }

        webix.ui({
            view:"window",
            id:"LoginForm",
            width:500,
            position:"center",
            modal:true,
            head:"登录",
            body:webix.copy(Login.webixData)
        }).show();
    },
    login:function(){
        var frameData = {
            "operation":"UserLogin",
            "username":VarTool.GetCookie("username"),
            "password":VarTool.GetCookie("password")
        };
        var resData = AjaxTool.SendData(frameData,"/rpost/task");
        if (resData == "success"){
            if ($$("LoginForm")){
                $$("LoginForm").hide();
            }
            IndexPage.show();
        }else{
            this.show();
        }
    },
    reLogin:function(){
        //清除cookie
        VarTool.SetCookie("username","");
        VarTool.SetCookie("password","");
        //刷新网页
        location.href = location.href;
    },
}