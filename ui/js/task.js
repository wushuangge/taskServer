Task = {
    GetUnclaimedTasks:function(){
         var frameData = {
            "operation":"GetTasksNoUser",
            "status":"待领取"
        };
        var resData = AjaxTool.SendData(frameData,"/rpost/task");
        
        if (resData == ""){
            resData = "[]"
        }

        var tasks = JSON.parse(resData);
        tasks = this.initTasks(tasks);
        return tasks
        
    },
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
            "TaskID":record.TaskID,
            "InstanceID":record.InstanceID,
            "ProjectID":record.ProjectID,
            "TaskType":record.TaskType,
            "Status":record.Status,
            "URL":record.URL,
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
        window.open(record.URL+"?"+
            "taskType"+"="+record.TaskType+"&"+
            "taskID"+"="+record.TaskID+"&"+
            "instanceID"+"="+record.InstanceID+"&"+
            "projectID"+"="+record.ProjectID
        );
    }
}