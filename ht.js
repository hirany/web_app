function tbox1(){
	var str1=document.js.txtb.value;
	alert("ようこそ"+str1+"さん"); 
}

$(function(){
	var socket=null;
	var msgbox=$("#chatbox textarea");
	var messages=$("#messages");

	$("#chatbox").submit(function(){
		if(!msgbox.val()) return false;
		if(!socket){
			alert("エラー：Websocket接続なし");
			return false;
		}
		socket.send(msgbox.val());
		msgbox.val("");
		return false;
	});
	if (!window["WebSocket"]) {
		alert("エラー：対応していないブラウザです")
	} else {
		socket = new WebSocket("ws://localhost:8000/room");
		socket.onclose = function(){
			alert("接続が終了しました");
		}
		socket.onmessages = function(e){
			messages.append($("<li>").text(e.data));
		}   
	}
});