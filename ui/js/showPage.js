IndexPage={
    menu_data:function(){
        var menu_data = [
            { id:"query", value:"我的任务", icon: "wxi-drag"},
            { id:"receive", value:"任务领取", icon: "wxi-folder"},
            { id:"assign", value:"任务指派", icon: "wxi-drag"},
        ];
        return menu_data;
    },
    show:function(){
        webix.ui({
            id:"index_page",
            rows: [
                { view: "toolbar", css:"webix_dark", padding:3, elements: [
                    { view: "button", type: "icon", icon: "wxi-angle-double-left",
                        width: 37, align: "left", css: "app_button", click: function(){
                            $$("sidebar").toggle();
                        }
                    },
                    { view: "label", label: "导航",width:300},
                    {},
                    { view:"icon", icon:"wxi-user", click: function(){Login.reLogin();}},
                    { view: "label", label: VarTool.GetCookie("username"), width:130}
                    ]
                },
                {cols:[
                    {
                        id:"sidebar",
                        view:"sidebar",
                        css:"webix_dark",
                        width:200,
                        data: IndexPage.menu_data(),
                        on:{
                            onAfterSelect: function(id){
                                $$("task").reconstruct();
                                switch(id){
                                    case "query":
                                        $$("task").addView(IndexPage.query());
                                        break;
                                    case "receive":
                                        $$("task").addView(IndexPage.receive());
                                        break;
                                    case "assign":
                                        $$("task").addView(IndexPage.assign());
                                        break;
                                }
                            }
                        }
                    },
                    {id:"task", cols:[]}
                ]}
            ]
        })
        $$('sidebar').select("query");
    },
    query:function(){
        var userTask = Task.GetTasksByUser();
        var webixData = {
            id:"showProject",
            rows:[  
                {
                    id:"showTable",
                    view:"datatable", 
                    css:"webix_header_border",
                    columns:[
                        { id:"TaskID",	    header:["任务ID",  {content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
					    { id:"InstanceID",	header:["实例ID",  {content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
					    { id:"ProjectID",	header:["项目ID",  {content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
                        { id:"TaskType",	header:["类别",    {content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
                        { id:"Status",		header:["状态",    {content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
                        { id:"CreateTime",  header:["创建时间", {content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
                        // { id:"votes",   header:"", template:"<input class='delbtn' type='button' value='跳转'>",fillspace:true}
                    ],
                    scrollAlignY:true,
                    on:{
                        onBeforeLoad:function(){
                            this.showOverlay("Loading...");
                        },
                        onAfterLoad:function(){
                            this.hideOverlay();
                            if (!this.count())
                                this.showOverlay("Sorry, there is no data");
                        },
                        "onItemClick":function(id, e, trg){
                            var record = $$("showTable").getItem(id.row);
                            Task.OpenEditPage(record);
                        }
                    },
                    data: userTask,
                },
            ]
        }
        return webixData;
    },
    receive:function(){
        var unclaimedTasks = Task.GetUnclaimedTasks();
        var webixData = {
            rows: [
                {
                    view:"datatable",
                    id:"showTable",
                    css:"webix_header_border",
                    columns:[
					    { id:"TaskID",	    header:["任务ID",{content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
					    { id:"InstanceID",	header:["实例ID",{content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
					    { id:"ProjectID",	header:["项目ID",{content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
                        { id:"TaskType",	header:["类别",  {content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
                        { id:"Status",		header:["状态",  {content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
                        { id:"CreateTime",	header:["创建时间",  {content:"textFilter" , compare:IndexPage.caseNotSensitive}], fillspace:true},
                        { id:"receive",     header:"", template:"<input class='delbtn' type='button' value='领取'>",       fillspace:true},
                    ],
                    //autowidth:true,
                    scrollAlignY:true,
                    on:{
                        onBeforeLoad:function(){
                            this.showOverlay("Loading...");
                        },
                        onAfterLoad:function(){
                            this.hideOverlay();
                            if (!this.count())
                                this.showOverlay("Sorry, there is no data");
                        },
                        "onItemClick":function(id, e, trg){
                            var record = $$("showTable").getItem(id.row);
                            switch(id.column) {
                                case "receive":
                                    if (Task.TaskBind(record) == "success"){
                                        $$("showTable").remove(id);
                                        Task.OpenEditPage(record);
                                    }
                                    break;
                                default:

                            } 
                        }
                    },
                    data:unclaimedTasks
                },
            ]
        }
        return webixData;
    },
    assign:function(){
        var webixData = {
        }
        return webixData;
    },
    caseNotSensitive:function(value, filter){
        value = value.toString().toLowerCase();
        filter = filter.toString().toLowerCase();

        return value.indexOf(filter) !== -1;
    }
}