

function jsonFormatHighlight() {
	var json={"name":"HTL","sex":"男","age":"24", };
	var str = formatJson(json);
	$(".result_json code").html(syntaxHighlight(str))
	// console.log(str);

	// var editor = ace.edit("editor");
 //    editor.setTheme("ace/theme/monokai");
 //    editor.getSession().setMode("ace/mode/javascript");
}



// 渲染执行结果
function renderRet(index) {
	var dataObj = {
        "node": "node1",
        "address": "127.0.0.1:3307",
        "type": "master",
        "status": "up",
        "laste_ping": "2016-09-24 17:17:52 +0800 CST",
        "max_conn": 32,
        "idle_conn": 8
    }
	var dataArr = [
	    {
	        "db": "kingshard",
	        "Table": {
	        	name: 'kathy',
	        	age: 18
	        },
	        "Key": "",
	        "Nodes": [
	            "node1",
	            "node2"
	        ],
	        "Locations": null,
	        "Type": "default",
	        "TableRowLimit": 0,
	        "DateRange": null
	    },
	    {
	        "db": "kingshard",
	        "Table": "test_shard_hash",
	        "Key": "id",
	        "Nodes": [
	            "node1",
	            "node2"
	        ],
	        "Locations": [
	            4,
	            4
	        ],
	        "Type": "hash",
	        "TableRowLimit": 0,
	        "DateRange": null
	    },
	    {
	        "db": "kingshard",
	        "Table": "test_shard_range",
	        "Key": "id",
	        "Nodes": [
	            "node1",
	            "node2"
	        ],
	        "Locations": [
	            4,
	            4
	        ],
	        "Type": "range",
	        "TableRowLimit": 10000,
	        "DateRange": null
	    },
	    {
	        "db": "kingshard",
	        "Table": "test_shard_time",
	        "Key": "id",
	        "Nodes": [
	            "node1",
	            "node2"
	        ],
	        "Locations": [
	            2,
	            2
	        ],
	        "Type": "hash",
	        "TableRowLimit": 0,
	        "DateRange": null
	    },
	    {
	        "db": "kingshard",
	        "Table": "test_shard_month",
	        "Key": "dtime",
	        "Nodes": [
	            "node1",
	            "node2"
	        ],
	        "Locations": null,
	        "Type": "date_month",
	        "TableRowLimit": 0,
	        "DateRange": [
	            "201603-201605",
	            "201609-201612"
	        ]
	    },
	    {
	        "db": "kingshard",
	        "Table": "test_shard_day",
	        "Key": "mtime",
	        "Nodes": [
	            "node1",
	            "node2"
	        ],
	        "Locations": null,
	        "Type": "date_day",
	        "TableRowLimit": 0,
	        "DateRange": [
	            "20160306-20160307",
	            "20160308-20160309"
	        ]
	    }
	]
	$("#js_ret_able").html(jsonToTable(dataObj));
	$("#js_table_wrapper").show();
}


// 渲染左侧菜单
function renderSideBar() {
	var htmlStr = ""
	for(var i=0,l=api_conf.length; i<l; i++) {
		htmlStr += '<li data-index="'+i+'">'+
			'<a href="javascript:void(0);">'+api_conf[i].name+'</a>'+
		'</li>';
	}
	$('.js_tab_view').html(htmlStr);
	$(".js_tab_view li").first().addClass('active_tab');
}


// 渲染接口说明
function renderActDesc(index) {
	var htmlStr = "";
	var apiItem = api_conf[index];
	// console.log(index, apiItem)
	for(key in apiItem) {
		if(key === 'action') {
			htmlStr += '<p><span class="key">Action</span>: <span class="string">'+apiItem.action+'</span></p>'
		}
		else if(key === 'url') {
			htmlStr += '<p><span class="key">URL</span>: <span class="string">'+apiItem.url+'</span></p>'
		}
		else if(key === 'param') {
			if($.type(apiItem.param) === 'array') {
				htmlStr += '<p><span class="key">参数</span>:</p>'
				for(var i=0,l=apiItem.param.length; i<l; i++) {
					htmlStr += '<p class="indent">- ' + 
						'<span class="key">'+
						apiItem.param[i].name + "</span>: " + 
						'<span class="string">'+
						apiItem.param[i].desc + '</span></p>';
				}
			}
			else {
				htmlStr += '<p><span class="key">参数</span>: <span class="string">'+apiItem.param+'</span></p>'
			}
		}
		else if(key === 'ret') {
			htmlStr += '<p><span class="key">返回结果</span>: <span class="string">'+apiItem.ret+'</span></p>'
		}
		else if(key === 'memo') {
			htmlStr += '<p><span class="key">注意</span>: <span class="string">'+apiItem.memo+'</span></p>'
		}
		else if(key === 'explain') {
			htmlStr += '<p><span class="key">说明</span>: <span class="string">'+apiItem.explain+'</span></p>'
		}
	}
	// console.log(htmlStr)
	$(".js_act_desc").html(htmlStr);
}


// 渲染接口参数
function renderActParam(index) {
	var param = api_conf[index].param;
	var htmlStr = "";
	if($.type(param) === 'array') {
		for(var i=0,l=param.length; i<l; i++) {
			htmlStr += '<div class="act_param">'+
				'<span class="label">'+param[i].name+':</span>'+
				'<input type="text">'+
			'</div>';
		}
	}
	$("#js_act_param_wrapper").html(htmlStr);
}


$(function() {
	// 渲染左侧菜单
	renderSideBar();

	// 渲染第一个接口说明以及参数
	var apiId = parseInt($(".js_tab_view li").first().data('index'));
	renderActDesc(apiId);
	renderActParam(apiId)

	// 点击左侧tab
	$(".js_tab_view li").on('click', function() {
		// tab样式切换
		$(this).addClass("active_tab").siblings().removeClass("active_tab");
		// 隐藏表格
		$("#js_table_wrapper").hide();
		// 渲染接口说明、参数
		var apiId = parseInt($(this).data('index'));
		renderActDesc(apiId);
		renderActParam(apiId);
	});

	// 点击执行
	$("#js_act_btn").on('click', function() {
		// 执行接口，展示结果
		renderRet(1);

	});

	// jsonFormatHighlight();

	
	// $("#js_ret_able").html(jsonToTable(dataArr));

	/*$.ajax({
		url: 'http://127.0.0.1:9797/api/v1/nodes/status',
		type: 'GET',
		success: function(data) {
			console.log(data)
		}
	})*/
});
