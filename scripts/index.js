var addTask = function(){
    var task = $('#task').val();
    var priority = $('#priority').val();
    var data = "task="+task+"&priority="+priority;

    $.post("/addTask",data,function(data,status){
        if(status == "success"){
            getTaskLists();
        }
    })
}

var getTaskLists = function(player){
	$.get("/getAllTasks","getAllTasks",function(data,status){
		if(status == "success"){
		    var tasksTable = "<table><tr><td>Task id</td><td> Task Description</td><td> Priority </td></tr>";
		    data = JSON.parse(data);
		    h = data
		    for(var key in data){
		         tasksTable += "<tr><td>"
		         + data[key].TASKID + "</td><td>"
		         + data[key].TASK + "</td><td>"
		         + data[key].PRIORITY + "</td>"
		         + " <td><div class='deleteTask' id="+data[key].TASKID+ "> ✗ </div></td> </tr>"
		    }
			$(".todoList").html(tasksTable);
		}
		$(".deleteTask").click(deleteTask);
	});
};

var deleteTask = function(){
    var data = "taskId="+this.id;
    $.post("/deleteTask",data,function(data,status){
            if(status == "success"){
                getTaskLists();
            }
    })
}

$(document).ready(function(){
    $(".add").click(addTask);
    getTaskLists();
})