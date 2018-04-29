var ddlog;
function reset_cron() {
	if (confirm("确定要重新生成计划任务吗?") === false) {
		return;
	}
	var dlog = art.dialog({id:'id22', icon: 'face-smile',left:'50%',top:'20%',background:'#FF6347'});
	ddlog = dlog;
	dlog.content('正在执行...');
	$.ajax({
		type : 'post',
		url : '?c=func&a=resetCrontab',
		dataType : 'json',
		success : function(ret) {
			var msg = ret['code']==200 ? "重置计划任务成功" : ret['msg'] ? ret['msg'] : "重置失败,请手动删除计划任务后再执行本操作";
			dlog.content(msg);
			setTimeout(function(){
				dlog.close();
			},2000);
		}
	});
}
function del_cache()
{
	//close_piao('msg');
	if (confirm("当easypanel升级后或出现foot被重复执行的提示时可运行?") == false) {
		return;
	}
	var dlog = art.dialog({id:'id22', icon: 'face-smile',left:'50%',top:'20%',background:'#FF6347'});
	ddlog = dlog;
	dlog.content('正在执行...');
	$.ajax({
		type : 'post',
		url : '?c=func&a=delSmartyCache',
		dataType : 'json',
		success : function(ret) {
			var msg = ret['code'] == 200 ? "清理缓存成功" : ret['msg'] ? ret['msg'] : "清理失败,请手动清理";
			dlog.content(msg);
			setTimeout(function(){
				dlog.close();
			},2000);
		}
	});
}
function piao_ftpusl(obj)
{
	//close_piao('msg');
	if (confirm("这将修改所有网站的设置,确定要继续吗?") === false) {
		return;
	}
	var dlog = art.dialog({id:'id22', icon: 'face-smile',left:'50%',top:'20%',background:'#FF6347'});
	ddlog = dlog;
	dlog.follow("#"+obj.id);
	var html = "<form action='' method='post'>" ;
		html += "限制:<input name='ftpusl' size='5'>KB";
		html += "<input type='button' onclick='change_ftpusl(ftpusl.value)' value='确定'>";
		html += "</form>";
	//piao_msg(id,html,'160px','60px');
	dlog.content(html);
}
function change_ftpusl(val)
{
	if (val != '') {
		$.ajax({
			type:'get',
			url :'?c=func&a=changeFtpusl',
			data:'ftp_usl=' + val,
			dataType:'json',
			success:function(ret){
				var m = ret['code'] == 200 ? "成功" : ret['msg'] ? ret['msg'] : "失败";
				//close_piao('msg');
				ddlog.close();
				alert(m);
			}
		});
	}
}
function piao_ftpdsl(obj)
{
	//close_piao('msg');
	if (confirm("这将修改所有网站的设置,确定要继续吗?") === false) {
		return;
	}
	var dlog = art.dialog({id:'id22', icon: 'face-smile',left:'50%',top:'20%',background:'#FF6347'});
	ddlog = dlog;
	dlog.follow("#"+obj.id);
	var html = "<form action='' method='post'>" ;
		html += "限制数:<input name='ftp_dsl'size='5'>KB";
		html += "&nbsp;&nbsp;<input type='button' onclick='change_ftpdsl(ftp_dsl.value)' value='确定'>";
		html += "</form>";
	//piao_msg(id,html,'160px','60px');
	dlog.content(html);
}
function change_ftpdsl(val)
{
	//close_piao('msg');
	if (val != '') {
		$.ajax({
			type:'get',
			url :'?c=func&a=changeFtpdsl',
			data:'ftp_dsl=' + val,
			dataType:'json',
			success:function(ret){
				var m = ret['code'] == 200 ? "成功" : ret['msg'] ? ret['msg'] : "失败";
				//close_piao('msg');
				ddlog.close();
				alert(m);
				
			}
		});
	}
}
function piao_htaccess(obj)
{
	var dlog = art.dialog({id:'id22', icon: 'face-smile',left:'50%',top:'20%',background:'#FF6347'});
	dlog.follow("#"+obj.id);
	ddlog = dlog;
	//close_piao('msg');
	//var id = obj.id;
	//use_msg_id = id;
	var html = "<div class='piao_ul'><li><a href='javascript:change_htaccess(1)'>开启</></li><br><li><a href='javascript:change_htaccess(0)'>关闭</a></li></div>"; 
	//piao_msg(id,html);
	dlog.content(html);
}
function change_htaccess(val)
{
	var p = val == 1 ? '开启' : '关闭';
	if (confirm("给所有网站" + p + "htaccess(url重写)功能?") === false) {
		return;
	}
	$.ajax({
		type : 'get',
		url : '?c=func&a=changeHtaccess',
		data : 'htaccess=' + val,
		dataType : 'json',
		success : function(ret) {
			var m = ret['code'] == 200 ? p + "成功" : ret['msg'] ? ret['msg'] : p + "失败";
			//close_piao('msg');
			ddlog.close();
			alert(m);
		}
	});
}
function piao_ftpconnect(obj)
{
	//close_piao('msg');
	if (confirm("这将修改所有网站的设置,确定要继续吗?") === false) {
		return;
	}
	var dlog = art.dialog({id:'id22', icon: 'face-smile',left:'50%',top:'20%',background:'#FF6347'});
	dlog.follow("#"+obj.id);
	ddlog = dlog;
	var html = "<form action='' method='post'>" ;
		html += "限制数:<input name='ftp_connect' size='5'>个";
		html += "&nbsp;&nbsp;<input type='button' onclick='change_ftpconnect(ftp_connect.value)' value='确定'>";
		html += "</form>";
	//piao_msg(id,html,'160px','60px');
	dlog.content(html);
}
function get_event(event)
{
	return window.event || event;
}
function change_ftpconnect(val)
{
	if (val != '') {
		$.ajax({
			type:'get',
			url :'?c=func&a=changeFtpconnect',
			data:'ftp_connect=' + val,
			dataType:'json',
			success:function(ret){
				var m = ret['code'] == 200 ? "成功" : ret['msg'] ? ret['msg'] : "失败";
				//close_piao('msg');
				ddlog.close();
				alert(m);
				
			}
		
		});
	}
}
function piao_msg(id,html,w,h)
{
	//var event = get_event(event);
	//event.cancelBubble = true;
	var offset = $("#" + id).offset();
	var msg = document.getElementById('msg');
	msg.style.top = offset.top + 54 + 'px';
	msg.style.left = offset.left + 90 + 'px';
	msg.style.height = h ? h : '40px';
	msg.style.width = w ? w :'60px';
	msg.innerHTML = html;
	msg.style.display = 'block';
}

function piao_ftp(obj)
{
	//close_piao('msg');
	//var id = obj.id;
	//use_msg_id = id;
	var dlog = art.dialog({id:'id22', icon: 'face-smile',left:'50%',top:'20%',background:'#FF6347'});
	dlog.follow("#"+obj.id);
	ddlog = dlog;
	var html = "<div class='piao_ul'><li><a href='javascript:change_ftp(1)'>开启</></li><br><li><a href='javascript:change_ftp(0)'>关闭</a></li></div>"; 
	//piao_msg(id,html);
	dlog.content(html);
}
function change_ftp(val)
{
	var p = val == 1 ? '开启' : '关闭';
	if (confirm("给所有网站" + p + "ftp功能?") === false) {
		return;
	}
	$.ajax({
		type : 'get',
		url : '?c=func&a=changeFtp',
		data : 'ftp=' + val,
		dataType : 'json',
		success : function(ret) {
			var m = ret['code'] == 200 ? p + "成功" : ret['msg'] ? ret['msg'] : p + "失败";
			//close_piao('msg');
			ddlog.close();
			alert(m);
		}
	});
}
function piao_loghanle(obj)
{
	//close_piao('msg');
	//var id = obj.id;
	//use_msg_id = id;
	var dlog = art.dialog({id:'id22', icon: 'face-smile',left:'50%',top:'20%',background:'#FF6347'});
	dlog.follow("#"+obj.id);
	ddlog = dlog;
	var html = "<div class='piao_ul'><li><a href='javascript:change_loghanle(1)'>开启</></li><br><li><a href='javascript:change_loghanle(0)'>关闭</a></li></div>";
	//piao_msg(id,html);
	dlog.content(html);
}
function change_loghanle(val) {
	var p = val == 1 ? '开启' : '关闭';
	if (confirm("给所有网站" + p + "日志压缩功能?") === false) {
		return;
	}
	$.ajax({
		type : 'get',
		url : '?c=func&a=changeLoghandle',
		data : 'log_handle=' + val,
		dataType : 'json',
		success : function(ret) {
			var m = ret['code'] == 200 ? p + "成功" : ret['msg'] ? ret['msg'] : p
					+ "失败";
			//close_piao('msg');
			ddlog.close();
			alert(m);
		}
	});
}