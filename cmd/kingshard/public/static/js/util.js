// Example usage: http://jsfiddle.net/q2gnX/
 

// json格式化
var formatJson = function(json, options) {
	var reg = null,
		formatted = '',
		pad = 0,
		PADDING = '    '; // one can also use '\t' or a different number of spaces
 
	// optional settings
	options = options || {};
	// remove newline where '{' or '[' follows ':'
	options.newlineAfterColonIfBeforeBraceOrBracket = (options.newlineAfterColonIfBeforeBraceOrBracket === true) ? true : false;
	// use a space after a colon
	options.spaceAfterColon = (options.spaceAfterColon === false) ? false : true;
 
	// begin formatting...
	if (typeof json !== 'string') {
		// make sure we start with the JSON as a string
		json = JSON.stringify(json);
	} else {
		// is already a string, so parse and re-stringify in order to remove extra whitespace
		json = JSON.parse(json);
		json = JSON.stringify(json);
	}
 
	// add newline before and after curly braces
	reg = /([\{\}])/g;
	json = json.replace(reg, '\r\n$1\r\n');
 
	// add newline before and after square brackets
	reg = /([\[\]])/g;
	json = json.replace(reg, '\r\n$1\r\n');
 
	// add newline after comma
	reg = /(\,)/g;
	json = json.replace(reg, '$1\r\n');
 
	// remove multiple newlines
	reg = /(\r\n\r\n)/g;
	json = json.replace(reg, '\r\n');
 
	// remove newlines before commas
	reg = /\r\n\,/g;
	json = json.replace(reg, ',');
 
	// optional formatting...
	if (!options.newlineAfterColonIfBeforeBraceOrBracket) {			
		reg = /\:\r\n\{/g;
		json = json.replace(reg, ':{');
		reg = /\:\r\n\[/g;
		json = json.replace(reg, ':[');
	}
	if (options.spaceAfterColon) {			
		reg = /\:/g;
		json = json.replace(reg, ':');
	}
 
	$.each(json.split('\r\n'), function(index, node) {
		var i = 0,
			indent = 0,
			padding = '';
 
		if (node.match(/\{$/) || node.match(/\[$/)) {
			indent = 1;
		} else if (node.match(/\}/) || node.match(/\]/)) {
			if (pad !== 0) {
				pad -= 1;
			}
		} else {
			indent = 0;
		}
 
		for (i = 0; i < pad; i++) {
			padding += PADDING;
		}
 
		formatted += padding + node + '\r\n';
		pad += indent;
	});
 
	return formatted;
};


// json高亮
function syntaxHighlight(json) {
    json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
    return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
        var cls = 'number';
        if (/^"/.test(match)) {
            if (/:$/.test(match)) {
                cls = 'key';
            } else {
                cls = 'string';
            }
        } else if (/true|false/.test(match)) {
            cls = 'boolean';
        } else if (/null/.test(match)) {
            cls = 'null';
        }
        return '<span class="' + cls + '">' + match + '</span>';
    });
}




/*
  data: 数组json，对象json
  return: 生成对象的表格html字符串
*/
function jsonToTable(data) {

	var htmlStr = "";
	var keyLen, objData;
	var dataType = $.type(data);


	if(dataType === 'array') {
		for(var i=0, l=data.length; i<l; i++) {
			keyLen = 0;
			objData = data[i];

			for(key in objData) {
				keyLen++;
			}
			
			var flag = true;

			for(key in objData) {

				var val = objData[key];
				if($.type(val) === 'array') {
					val = '[' + val + ']';
				}
				else if($.type(val) === 'object') {
					val = JSON.stringify(val);
				}
				else if($.type(val) === 'string') {
					val = '"'+val+'"';
				}

				if(flag) {
					flag = false;
					htmlStr += '<tr>'+
						'<td rowspan="' + keyLen + '">'+(i+1)+'</td>'+
						'<td>'+key+'</td>'+
						'<td>'+val+'</td>'+
					'</tr>'
				}
				else{
					htmlStr += '<tr>'+
						'<td>'+key+'</td>'+
						'<td>'+val+'</td>'+
					'</tr>'
				}
			}	
		}
	}
	else {
		for(key in data) {

			var val = data[key];
			if($.type(val) === 'array') {
				val = '[' + val + ']';
			}
			else if($.type(val) === 'object') {
				val = JSON.stringify(val);
			}
			else if($.type(val) === 'string') {
				val = '"'+val+'"';
			}

			htmlStr += '<tr>'+
				'<td>'+key+'</td>'+
				'<td>'+val+'</td>'+
			'</tr>'
		}
	}
	// console.log(htmlStr);
	htmlStr = '<thead>'+
		'<tr>'+
			(dataType === 'array' ? '<th>序号</th>' : '') + 
			'<th>键</th>'+
			'<th>值</th>'+
		'</tr>'+
	'</thead>'+
	'<tbody>'+ htmlStr
	'</tbody>';

	return htmlStr;
}



