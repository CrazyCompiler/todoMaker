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

var setSelectionOptions = ['High','Medium','Low'];

var update = function(taskId,dataToBeUpDated){
    data = "taskId="+taskId+"&priority="+dataToBeUpDated;
    $.post("/updatePriority",data,function(data,status){
        if(status == "success"){
            gridOptions.api.refreshView();
        }
    })
}

var customEditor = function(params) {
    var editing = false;
    var eCell = document.createElement('span');
    var eLabel = document.createTextNode(params.value);
    eCell.appendChild(eLabel);

    var eSelect = document.createElement("select");

    setSelectionOptions.forEach(function(item) {
        var eOption = document.createElement("option");
        eOption.setAttribute("value", item);
        eOption.innerHTML = item;
        eSelect.appendChild(eOption);
    });
    eSelect.value = params.value;

    eCell.addEventListener('click', function () {
        if (!editing) {
            eCell.removeChild(eLabel);
            eCell.appendChild(eSelect);
            eSelect.focus();
            editing = true;
        }
    });

    eSelect.addEventListener('blur', function () {
        if (editing) {
            editing = false;
            eCell.removeChild(eSelect);
            eCell.appendChild(eLabel);
        }
    });

    eSelect.addEventListener('change', function () {
        if (editing) {
            editing = false;
            var newValue = eSelect.value;
            update(params.data.TASKID,newValue);
            params.data[params.colDef.field] = newValue;
            eLabel.nodeValue = newValue;
            eCell.removeChild(eSelect);
            eCell.appendChild(eLabel);
        }
    });

    return eCell;
}

var gridOptions = {
    debug: true,
    rowData: null,
    groupHeaders: true,
    enableSorting: true,
    enableFilter: true,
    enableColResize: true,
    rowHeight:40
};

var rowData;

var displayData = function(data){
            $('.todoList').html("");
             var columnDefs = [
                                {headerName: "Task_Id", field: "TASKID", width:100},
                                {headerName: "Task Description", field: "TASK",width:600},
                                {headerName: "Priority" , field: "PRIORITY", cellRenderer: customEditor,width : 100},
                                {headerName: "" , field: "delete" , width:100 ,onCellClicked : deleteTask}
                            ];

            gridOptions.columnDefs =  columnDefs;
            var eGridDiv = document.querySelector('.todoList');
            new agGrid.Grid(eGridDiv, gridOptions);
            gridOptions.api.setRowData(rowData);
            gridOptions.api.sizeColumnsToFit();
}

var getTaskLists = function(player){
	$.get("/getAllTasks","getAllTasks",function(data,status){
		if(status == "success"){
            rowData = JSON.parse(data);
            rowData.forEach(function(each){
                each.delete = " <td><div class='deleteTask' id="+each.TASKID+ "> ✗ </div>"
            })
            displayData(data);
		}
	});
};

var deleteTask = function(params){
    var dataToBeSend = {taskId:params.data.TASKID};
      $.ajax({
        url: "/deleteTask/"+params.data.TASKID,
        type: 'DELETE',
        data: dataToBeSend,
        traditional: true,
        success: function() {
            rowData.splice(rowData.indexOf(params.data),1);
            displayData(rowData);
        }
    });
}

var uploadCsv = function(){
    var formData = new FormData($(this)[0]);
     $.ajax({
            url: "uploadCsv",
            type: 'POST',
            data: formData,
            async: false,
            success: function (data) {
                if(data){
                    alert(data)
                }
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