Task = {
    GetTasksByUser:function(){
        var frameData = {
            "operation":"GetTasksByUser"
        };
        var resData = AjaxTool.SendData(frameData,"/rpost/task");

        if (resData == ""){
            resData = "[]"
        }
        var tasks = JSON.parse(resData);
        tasks = this.initTasks(tasks);
        return tasks
    },
    TaskBind:function(record){
        var frameData = {
            "operation":"TaskBind",
            "id":record.ID,
        };
        var resData = AjaxTool.SendData(frameData,"/rpost/task");
        return resData;
    },
    initTasks:function(tasks){
        for (var i=0; i<tasks.length; i++){
            tasks[i].CreateTime = formatDateTime(new Date((tasks[i].CreateTime*1000)));
        }        
        return tasks
    },
    OpenEditPage:function(record){
        let argDict = {};
        argDict["projectID"] = record.ProjectID;
        argDict["instanceID"] = record.InstanceID;
        argDict["taskID"] = record.TaskID;
        argDict["taskType"] = record.TaskType;
        let b64Str = window.btoa(JSON.stringify(argDict));
        window.open(record.URL+"?"+b64Str);
    },
    GetEditors:function(){
        var frameData = {
            "operation":"GetEditors",
        };
        return JSON.parse(AjaxTool.SendData(frameData,"/rpost/task"));
    }
}