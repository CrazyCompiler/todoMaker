var $scope = {};

var addTask = function(){
    var task = $('#task').val();
    var priority = $('#priority').val();
    var data = "task="+task+"&priority="+priority;

    $.post("/addTask",data,function(data,status){
        if(status == "success"){
             $scope.gridOptions.api.refreshView();
        }
    })
}


$scope.gridOptions = {
    debug: true,
    rowData: null,
    groupHeaders: true,
    enableSorting: true,
    enableFilter: true,
    enableColResize: true,
    rowHeight:40
};



var getTaskLists = function(player){
	$.get("/getAllTasks","getAllTasks",function(data,status){
		if(status == "success"){

            var columnDefs = [
                    {headerName: "Task_Id", field: "TASKID"},
                    {headerName: "Task Description", field: "TASK"},
                    {headerName: "Priority" , field: "PRIORITY"},
                    {headerName: "" , field: "delete"}
                ];

            data = JSON.parse(data);
            data.forEach(function(each){
                each.delete = " <td><div class='deleteTask' id="+each.TASKID+ "> ✗ </div>"
            })

            $scope.gridOptions.columnDefs =  columnDefs;

            var eGridDiv = document.querySelector('.todoList');
            new agGrid.Grid(eGridDiv, $scope.gridOptions);
            $scope.gridOptions.api.setRowData(data);
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
        success:$scope.gridOptions.api.refreshView()
    });
      $scope.gridOptions.api.refreshView();

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