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
    var dataToBeSend = {taskId:this.id};
      $.ajax({
        url: "/deleteTask/"+this.id,
        type: 'DELETE',
        data: dataToBeSend,
        traditional: true,
        success: getTaskLists,
    });
    getTaskLists();

}

var uploadCsv = function(){
    var formData = new FormData($(this)[0]);
     $.ajax({
            url: "uploadCsv",
            type: 'POST',
            data: formData,
            async: false,
            success: function (data) {
                alert("File has been uploaded")
            },
            contentType: false,
            processData: false
        });

}

$(document).ready(function(){
    $(".add").click(addTask);
    $("form#csvUploader").submit(uploadCsv)
    getTaskLists();
})