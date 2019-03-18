<html lang="en">
  <head>
<script type="text/javascript" src="https://libs.baidu.com/jquery/1.9.0/jquery.js"></script>

<!-- <script type="text/javascript">
$(document).ready(function(){

$('#trmt').on('click', function() {

 alert("终止");

});

});

</script>
-->

<script type="text/javascript">
$(document).ready(function(){
$('#test').on('click', function() {
  var title=$('#test').val()
  var data={
                'title': "title" ,
                'content': 4567
            }
$.ajax({
    url: "/specify",
    type: "get",
  <!-- data: JSON.stringify(data),   # 转化为json字符串  post 使用-->
    data:data,  <!-- get方法使用 -->
    dataType: "json", <!-- 注意：这里是指希望服务端返回json格式的数据 -->
    success: function(data){
        alert( data.title );
    }
<!-- $(this).resetForm(); // 提交后重置表单 -->
});
     return false; <!-- # 阻止表单自动提交事件 -->
});

});
</script>


<script  type="text/javascript">

//window.onload = function () {
$(document).ready(function(){
    var conn;
    var log = document.getElementById("log");
    var trmt = document.getElementById("trmt")
    var query = document.getElementById("query")
   var containerid   = document.getElementById("containerid")

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    if (window["WebSocket"]) {

conn = new WebSocket("wss://" + document.location.host + "/ws");
<!-- 点击终止websocket -->
       trmt.onclick=function () {
       conn.close()
       conn = null;
       }
       <!-- 获取input中的值 -->
              query.onclick=function () {
              data1=$(" #containerid ").val()
              conn.send(data1)
              }

        conn.onclose = function (evt) {
            console.log('websocket 断开: ' + evt.code + ' ' + evt.reason + ' ' + evt.wasClean)
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var item = document.createElement("div");
                item.innerText = messages[i];
                appendLog(item);
                var obj = {};
                obj.message = 'still alive' + new Date().toLocaleString();
                console.log(obj);
                obj = JSON.stringify(obj);



               <!--  conn.send(obj); -->
            }
        };
       conn.onerror(evt) = function (etv){
        console.log("websocket error"); 
       };




    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }

});


</script>


<style type="text/css">
html {
    overflow: hidden;
}

body {
    overflow: hidden;
    padding: 0;
    margin: 0;
    width: 100%;
    height: 100%;
    background: gray;
}

#log {
    background: white;
    margin: 0;
    padding: 0.5em 0.5em 0.5em 0.5em;
    position: absolute;
    top: 0.5em;
    left: 0.5em;
    right: 0.5em;
    bottom: 8em;
    overflow: auto;
}

#log pre {
  margin: 0;
}

#form {
    padding: 0 0.5em 0 0.5em;
    margin: 0;
    position: absolute;
    bottom: 5em;
    left: 0px;
    width: 100%;
    overflow: hidden;
}

.button {
padding: 0 0.5em 0 0.5em;
position: absolute;
margin: 0;
}

#closed {
bottom: 4em;
left: 1px;
}

#open {
bottom: 2em;
left : 1px;
}


</style>
  </head>

  <body>
<div id="log"></div>
<form id="form">
  <!--  <input type="text" id="msg" size="32"/>-->
  <select>
              <option value="">select role</option>
           <option id="name1" name="name1"  value ="local">local</option>
  		 <option id="name2" name="name2"  value ="cluster">中信优享+测试</option>
          </select>
    <input type="text" id="containerid" size="32"/>

    <!-- submit 页面会刷新一次 -->
    <input id="query" type="button" value="Specify" /> <br/>
<button id="trmt" type="button" >terminate</button>
<button  type="button" id="open">open</button>

<input type="button"  value="显示警告框" />

<!--    <input type="text" id="msg" size="64"/> -->
</form>

<!--<button type="button" class="button" id="closed">closed</button>
<button type="button" class="button" id="open">open</button>
-->
  </body>
</html>
