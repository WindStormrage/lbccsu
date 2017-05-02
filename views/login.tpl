<html>
<body>
<form method="post">
    <input type="email" name="email" id="email">
    <button type="button" onclick="SendEmail()"></button>
    <input type="text" name="verify" id="verify">
    <button type="submit" onclick="Register()">
</form>

<script src="../static/js/jquery-1.11.1.min.js"></script>
<script>

   function Register() {
      username=$("#email").val();
      password=$("#verify").val();

      $.ajax({
            url:'{{urlfor "LoginController.Login"}}',
            type:'POST',
            data:{
               'username':username,
               'password':password
            },
            dataType:'json',
            contentType:'application/x-www-form-urlencoded;charset=UTF-8',
            cache:false,
            success:function (data) {
               if (data.status==10000){
                  location.href='{{urlfor "MainController.Main"}}'
               }
            },
            error:function () {
               console.log("出问题了");
            }

         }
      );
   }


   function SendEmail() {
      email=$("#email").val();

      $.ajax({
         url:'{{urlfor "LoginController.Register"}}',
         type:'POST',
         data:{
            email:email
         },
         dataType:'json',
         contentType:'application/x-www-form-urlencoded;charset=UTF-8',
         cache:false,
         success:function (data) {
            if (data.status==10000){
               console.log(data)
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