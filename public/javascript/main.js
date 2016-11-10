$(function(){
    $("form[name='login_form']").validate({
        rules:{
          "username":"required",
          "password":{
            "required":true,
            "minlength":6
         }
       },
       messages:{
         "username":"Please enter your username",
         "password":{
           "required":"Please enter your password",
           "minlength":"Your password length is less than 6!"
         }
       },
       submitHandler:function(){
         var data={}
         $("form[name='login_form']").serializeArray().map(function(x){data[x.name]=x.value});
         $.ajax({
            type:"POST",
            url:"/login",
            data:data,
            success:function(data){
                var result = JSON.parse(data)
                if(result&&result["token"]){
                    window.localStorage.setItem("token",result["token"])
                    window.location.replace("/admin")
                }else{
                    alert("Unknown Error");
                }
            },
            error:function(e){
                alert(e);
            }
         })
       }
    });

    $("form[name='signup_form']").validate({
        rules:{
          "username":"required",
          "email":{
            "required":true,
            "email":true
          },
          "password":{
            "required":true,
            "minlength":6
         },
         "password2":{
           "required":true,
           "minlength":6,
           "equalTo":"#InputPassword"
         }
       },
       messages:{
         "username":"Please enter your username",
         "email":{
           "required":"Please enter your email",
           "email":"Your email is invalid!"
         },
         "password":{
           "required":"Please enter your password",
           "minlength":"Your password length is less than 6!"
         },
         "password2":{
           "required":"Please enter your confirm password",
           "minlength":"Your password length is less than 6!",
           "equalTo":"please enter the same value again"
         }
       },
       submitHandler:function(form){
         form.submit()

       }
    });
})
