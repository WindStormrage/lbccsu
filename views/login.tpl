<html>

<body>
<form method="post">
用户名：<input type="text" name="username" id="username" placeholder="用户名唯一">
<!-- <button type="button" onclick="SendEmail()" >-->
    密码：<input type="password" name="password" id="password">
    确认密码：<input type="password" name="password1" id="password1">
    <button type="submit" onclick="Register()">注册</button>
</form>
用户名：<input type="text" name="username" id="username" >
确认密码：<input type="password" name="password" id="password">
<button type="submit" onclick="Login()">登陆</button>

<script src="/static/js/jquery-1.11.1.min.js"></script>
<script>

   function Register() {
      username=$("#username").val();
      password=$("#password").val();
      password1=$("#password1").val();


      if (username==""||password==""){
           alert("内容填写不完整");

        }
      if (password!=password1){
         alert("两次密码输入不一致");
         return
      }
      $.ajax({
            url:'{{urlfor "LoginController.Register"}}',
            type:'POST',
            data:{
               'username':username,
               'password':password
            },
            dataType:'json',
            contentType:'application/x-www-form-urlencoded;charset=UTF-8',
            cache:true,
            success:function (data) {
               if (data.status == 10000){
                  location.href='{{urlfor "MainController.Main"}}';

               } else if (data.status == 10001) {
                  alert("用户名已经存在");
               }
            }
         }
      );
   }

   function Login() {
      username=$("#username").val();
      password=$("#password").val();


      $.ajax({
            url:'{{urlfor "LoginController.LoginSubmit"}}',
            type:'POST',
            data:{
               email:username,
               password:password
            },
            dataType:'json',
            contentType:'application/x-www-form-urlencoded;charset=UTF-8',
            cache:true,
            success:function (data) {
               if (data.status==10000){
                  alert("登陆成功");
               }
            }
         }
      );

   }


   function SendEmail() {
      email=$("#email").val();
      verify=$("#verify").val();
      if (verify==""||email==""){
         alert("内容填写不完整");
         return
      }
      $.ajax({
         url:'{{urlfor "LoginController.SendEmail"}}',
         type:'POST',
         data:{
            email:email
         },
         dataType:'json',
         contentType:'application/x-www-form-urlencoded;charset=UTF-8',
         cache:false,
         success:function (data) {
            if (data.status==10000){
               code=data.message;
            }
         },
         error:function () {
            console.log("出问题了");
         }
         }
      );
   }
</script>

</body>
</html>