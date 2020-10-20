IndexPage={
    pagingSize:10,
    tableMainCols:[
        { id:"TaskID",	    header:["任务ID",  {content:"textFilter" }],   fillspace:true},
        { id:"InstanceID",	header:["实例ID",  {content:"textFilter" }],   fillspace:true},
        { id:"ProjectID",	header:["项目ID",  {content:"textFilter" }],   fillspace:true},
        { id:"TaskType",	header:["类别",    {content:"selectFilter" }], fillspace:true},
        { id:"Status",		header:["状态",    {content:"selectFilter" }], fillspace:true},
        { id:"CreateTime",  header:["创建时间"],   fillspace:true, format:function(data){return IndexPage.timeStamp2str(data)}, sort:"string"}
    ],
    editorData:[
        { id:"query",   value:"我的任务", icon: "wxi-drag", data:[
            { id: "unfinishTask", value: "未完成任务", icon: "wxi-close"},
            { id: "finishTask",   value: "已完成任务", icon: "wxi-check"}
        ]},
        { id:"receive", value:"任务领取", icon: "wxi-folder"},
    ],
    managerMenu:[
        { id:"assign",  value:"任务指派", icon: "wxi-pencil"},
    ],
    show:function(groups){
        webix.ui({
            id:"index_page",
            rows: [
                { view: "toolbar", css:"webix_dark", padding:3, elements: [
                    { view: "button", type: "icon", icon: "wxi-angle-double-left",
                        width: 37, align: "left", css: "app_button", click: function(){
                            $$("sidebar").toggle();
                            $$("showTable").resize();
                        }
                    },
                    { view: "label", label: "导航", width:100},
                    { view: "label", label: "TaskDash", align:"center"},
                    { view: "icon",  icon:"wxi-user", click: function(){Login.reLogin();}},
                    { view: "label", label: VarTool.GetCookie("username"), width:130}
                ]},
                { cols:[
                    {
                        id:"sidebar",
                        view:"sidebar",
                        css:"webix_dark",
                        width:200,
                        data: this.getMenuData(groups),
                        on:{
                            onItemDblClick: function(id){
                                IndexPage.showTask(id)
                            },
                            onAfterSelect: function(id){
                                IndexPage.showTask(id)
                            },
                            onAfterOpen: function(id){
                                switch(id){
                                    case "query":
                                        $$("sidebar").select("unfinishTask");
                                        break;
                                    default:
                                        break;
                                }
                            }
                        }
                    },
                    { rows:[
                        { id:"task" },
                        { id:"paging", height:30 },
                    ]},
                ]}
            ]
        })
        $$("paging")._contentobj.innerHTML = `<div id="paging_here_too" style="width:400px; float:right;"/>`;
        $$("task")._contentobj.innerHTML = `<div id="testA" style="width:100%;height:100%"/>`;
        $$('sidebar').open("query");
    },
    showUnfinishTask:function(username){
        webix.ui({
            id:"showTable",
            container:"testA",
            view:"datatable", 
            css:"webix_header_border",
            columns:this.tableMainCols,
            select:"row",
            pager:{
                template:IndexPage.pagerTemplate,
                container:"paging_here_too",
                size:this.pagingSize,
                animate:true,
                group:5
            },
            on:{
                onBeforeLoad:function(){
                    this.showOverlay("Loading...");
                },
                onAfterLoad:function(){
                    this.hideOverlay();
                    if (!this.count())
                        this.showOverlay("Sorry, there is no data");
                },
                onItemDblClick:function(id, e, trg){
                    var record = $$("showTable").getItem(id.row);
                    Task.OpenEditPage(record);
                }
            },
            url:AjaxTool.serviceURL+"/rget/task/?operation=GetTasksByUser&username="+username,
        }).show(); 
    },
    showFinishTask:function(username){
        webix.ui({
            id:"showTable",
            container:"testA",
            view:"datatable",
            css:"webix_header_border",
            columns:this.tableMainCols,
            select:"row",
            pager:{
                template:IndexPage.pagerTemplate,
                container:"paging_here_too",
                size:this.pagingSize,
                animate:true,
                group:5
            },
            on:{
                onBeforeLoad:function(){
                    this.showOverlay("Loading...");
                },
                onAfterLoad:function(){
                    this.hideOverlay();
                    if (!this.count())
                        this.showOverlay("Sorry, there is no data");
                },
                onItemDblClick:function(id, e, trg){
                    webix.message("已完成任务无法编辑")
                }
            },
            url:AjaxTool.serviceURL+"/rget/task/?operation=GetCompleteTasksByUser&username="+username,
        }).show(); 
    },
    showNoUserTask:function(){
        webix.ui({
            id:"showTable",
            container:"testA",
            view:"datatable", 
            css:"webix_header_border",
            columns:this.tableMainCols.concat([
                { id:"receive", header:"", template:"<input class='delbtn' type='button' value='领取'>", fillspace:true}
            ]),
            select:"row",
            pager:{
                template:IndexPage.pagerTemplate,
                container:"paging_here_too",
                size:this.pagingSize,
                animate:true,
                group:5
            },
            on:{
                onBeforeLoad:function(){
                    this.showOverlay("Loading...");
                },
                onAfterLoad:function(){
                    this.hideOverlay();
                    if (!this.count())
                        this.showOverlay("Sorry, there is no data");
                },
                onItemClick:function(id, e, trg){
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
            url:AjaxTool.serviceURL+"/rget/task/?operation=GetTasksNoUser",
        }).show(); 
    },
    assign:function(){
        webix.ui({
            id:"showTable",
            container:"testA",
            view:"datatable", 
            css:"webix_header_border",
            columns:this.tableMainCols.concat([
                { id:"User", header:"指派", fillspace:true, editor:"select", options:Task.GetEditors()},//["sunmq","2"]
            ]),
            select:"row",
            multiselect:true,
			editable:true,
			editaction:"click",
            pager:{
                template:IndexPage.pagerTemplate,
                container:"paging_here_too",
                size:this.pagingSize,
                animate:true,
                group:5
            },
            on:{
                onBeforeLoad:function(){
                    this.showOverlay("Loading...");
                },
                onAfterLoad:function(){
                    this.hideOverlay();
                    if (!this.count())
                        this.showOverlay("Sorry, there is no data");
                },
                onEditorChange:function(id, value){
                    var editRow = this.getItem(id);
                    editRow.Status = "待编辑";
                    this.editStop();

                    var selectedItems = this.getSelectedItem();
                    for (var i=0; i<selectedItems.length; i++){
                        if (selectedItems[i].ID == editRow.ID){continue};
                        if (selectedItems[i].User == value){continue};

                        selectedItems[i].User = value;
                        selectedItems[i].Status = "待编辑";
                        this.updateItem(selectedItems[i].id, selectedItems[i]);
                    }
                }
            },
            url:AjaxTool.serviceURL+"/rget/task/?operation=GetTasksNoUser",
            save:{
                "update":AjaxTool.serviceURL+"/rget/task/?operation=updateTask",
            },
        }).show(); 
    },
    timeStamp2str:function(date) {
        date = new Date(date * 1000);
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
    },
    getMenuData:function(groups){
        var menuData = this.editorData.concat();
        if(groups.indexOf("manager") != -1){
            menuData = menuData.concat(this.managerMenu);
        }
        return menuData;
    },
    showTask:function(id){
        if($$("showTable")){
            $$("showTable").destructor();
            document.getElementById("paging_here_too").innerHTML = ""
        }
        let username = VarTool.GetCookie("username")
        switch(id){
            case "unfinishTask":
                IndexPage.showUnfinishTask(username);
                break;
            case "finishTask":
                IndexPage.showFinishTask(username);
                break;
            case "receive":
                IndexPage.showNoUserTask(username);
                break;
            case "assign":
                IndexPage.assign(username);
                break;
            default:
                break;
        }
    },
    pagerTemplate:'{common.first()} {common.prev()} {common.pages()} {common.next()} {common.last()}',
    // pagerTemplate:function(data, common){
    //     var html = "<div style='width:100px; text-align:right; line-height:20px; font-size:10pt; float:left'>第"+(data.page+1)+"页，共"+data.limit+"页</div> "
    //     return common.prev()+html+common.next();
    // },
    // pagerTemplate:function(data, common){
    //     var start = data.page * data.size;
    //     var end = start + data.size;
    //     var html = "<div style='width:600px; text-align:right; line-height:20px; font-size:10pt; float:left'> Rows "+(start+1)+" - "+end+" </div> ";
    //     return common.prev()+html+common.next();
    // },
}