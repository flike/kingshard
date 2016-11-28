var api_conf = [
	{
		// id: 0,
		name: '查看node的状态',
		action: 'GET',
		url: 'http://127.0.0.1:9797/api/v1/nodes/status',
		param: '无',
		ret: 'node数组，node中包含Master和Slave信息'
	},{
		// id: 1,
		name: '添加slave',
		action: 'POST',
		url: 'http://127.0.0.1:9797/api/v1/nodes/slaves',
		// param: '- node：节点名字 - addr：slave的IP和端口',
		param: [{
			name: 'node',
			desc: '节点名字',
		},{
			name: 'addr',
			desc: 'slave的IP和端口',
		}],
		ret: 'node数组，node中包含Master和Slave信息'
	},{
		// id: 2,
		name: '删除slave',
		action: 'DELETE',
		url: 'http://127.0.0.1:9797/api/v1/nodes/slaves',
		param: [{
			name: 'node',
			desc: '节点名字',
		},{
			name: 'addr',
			desc: 'slave的IP和端口',
		}],
		ret: '成功:"ok"，失败："error message"'
	},{
		// id: 3,
		name: '设置slave状态',
		action: 'PUT',
		url: 'http://127.0.0.1:9797/api/v1/nodes/slaves/status',
		param: [{
			name: 'opt',
			desc: '"up" or "down"，上线或下线slave',
		},{
			name: 'node',
			desc: '节点名字',
		},{
			name: 'addr',
			desc: 'slave的IP和端口',
		}],
		ret: '成功:"ok"，失败："error message"'
	},{
		// id: 4,
		name: '设置master状态',
		action: 'PUT',
		url: 'http://127.0.0.1:9797/api/v1/nodes/masters/status',
		param: [{
			name: 'opt',
			desc: '"up" or "down"，上线或下线slave',
		},{
			name: 'node',
			desc: '节点名字',
		},{
			name: 'addr',
			desc: 'slave的IP和端口',
		}],
		ret: '成功:"ok"，失败："error message"'
	},{
		// id: 5,
		name: '查看proxy状态',
		action: 'GET',
		url: 'http://127.0.0.1:9797/api/v1/proxy/status',
		param: '无',
		ret: '"online"或者"offline"',
		memo: '该API主要用于配合LVS平滑下线'
	},{
		// id: 6,
		name: '设置proxy状态',
		action: 'PUT',
		url: 'http://127.0.0.1:9797/api/v1/proxy/status',
		param: [{
			name: 'opt',
			desc: '"online"或者"offline"',
		}],
		ret: '成功:"ok",失败："error message"',
		memo: '该API主要用于配合LVS平滑下线'
	},{
		// id: 7,
		name: '查看proxy的schema',
		action: 'GET',
		url: 'http://127.0.0.1:9797/api/v1/proxy/schema',
		param: '无',
		ret: 'schema数组，但只有一项。字段意思参考配置文件',
	},{
		// id: 8,
		name: '查看proxy的客户端白名单',
		action: 'GET',
		url: 'http://127.0.0.1:9797/api/v1/proxy/allow_ips',
		param: [{
			name: 'allow_ips',
			desc: 'ip列表数组',
		}],
		ret: 'IP列表',
	},{
		// id: 9,
		name: '添加proxy的客户端白名单',
		action: 'POST',
		url: 'http://127.0.0.1:9797/api/v1/proxy/allow_ips',
		param: [{
			name: 'allow_ips',
			desc: 'ip列表数组',
		}],
		ret: '成功:"ok",失败："error message"',
	},{
		// id: 10,
		name: '删除proxy的客户端白名单',
		action: 'DELETE',
		url: 'http://127.0.0.1:9797/api/v1/proxy/allow_ips',
		param: [{
			name: 'allow_ips',
			desc: 'ip列表数组',
		}],
		ret: '成功:"ok",失败："error message"',
	},{
		// id: 11,
		name: '查看proxy的black_sql',
		action: 'GET',
		url: 'http://127.0.0.1：9797/api/v1/proxy/black_sqls',
		param: '无',
		ret: 'sql列表',
	},{
		// id: 12,
		name: '增加proxy的black_sql',
		action: 'POST',
		url: 'http://127.0.0.1:9797/api/v1/proxy/black_sqls',
		param: [{
			name: 'sql',
			desc: '注意：一次只能添加一条SQL'
		}],
		ret: '成功:"ok",失败："error message"',
	},{
		// id: 13,
		name: '删除proxy的black_sql',
		action: 'DELETE',
		url: 'http:// 127.0.0.1:9797/api/v1/proxy/black_sqls',
		param: [{
			name: 'sql',
			desc: '注意：一次只能删除一条SQL'
		}],
		ret: '成功:"ok",失败："error message"',
	},{
		// id: 14,
		name: '设置proxy的slow sql开关',
		action: 'PUT',
		url: 'http://127.0.0.1:9797/api/v1/proxy/slow_sql/status',
		param: [{
			name: 'opt',
			desc: '可选值："on","off"'
		}],
		ret: '成功:"ok",失败："error message"',
	},{
		// id: 15,
		name: '查看proxy的slow sql的时间',
		action: 'GET',
		url: 'http://127.0.0.1:9797/api/v1/proxy/slow_sql/time',
		param: '无',
		ret: 'slow_log 时间，单位ms',
	},{
		// id: 16,
		name: '设置proxy的slow sql的时间',
		action: 'PUT',
		url: 'http://127.0.0.1:9797/api/v1/proxy/slow_sql/time',
		param: [{
			name: 'slow_time',
			desc: '慢日志时间，单位ms。执行时间超过该值的SQL会输出到慢日志文件'
		}],
		ret: '成功:"ok",失败："error message"',
	},{
		// id: 16,
		name: '保存proxy的配置',
		action: 'PUT',
		url: 'http://127.0.0.1:9797/api/v1/proxy/config/save',
		param: '无',
		ret: '成功:"ok",失败："error message"',
		explain: '将kingshard的配置写入文件'
	}
]


