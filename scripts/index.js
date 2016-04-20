var addTask = function(){
    var task = $('#task').val();
    var priority = $('#priority').val();
    var data = "task="+task+"&priority="+priority;
    $.post("/addTask",data,function(data,status){
        if(status == "success"){
            $(".todoList").html(data)
        }
    })
}

var getTaskLists = function(player){
	$.get("/getAllTasks","getAllTasks",function(data,status){
		if(status == "success"){
			$(".todoList").html(data);
		}
	});
};

$(document).ready(function(){
    $(".add").click(addTask);
    getTaskLists();
})