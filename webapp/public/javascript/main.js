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
                var token=""
                if (data["Status"]==false){
                    $("#auth_alert").removeClass("alert-info hidden").addClass("alert-danger").html(data["data"]["message"]+":"+data["data"]["error"])
                }else{
                    try{
                        console.log(data)
                        token=data["data"]["data"]["token"]
                        window.localStorage.setItem("token",token)
                        window.location.replace("/admin")
                    }catch(e){
                         $("#auth_alert").removeClass("alert-info hidden").addClass("alert-danger").html("Retrive Token Error")
                    }
                }

             },
             error:function(e){
                $("#auth_alert").removeClass("alert-info").addClass("alert-danger").html("Please try again later")
                 console.log(e)
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
          var data={}
          $("form[name='signup_form']").serializeArray().map(function(x){data[x.name]=x.value});
          $.ajax({
             type:"POST",
             url:"/signup",
             data:data,
             success:function(data){
                console.log(data)
                if(data["Stauts"]==false){
                    $("#auth_alert").removeClass("alert-info hidden").addClass("alert-danger").html(data["data"]["message"])
                }else{
                    $("#auth_alert").removeClass("alert-danger hidden").addClass("alert-info").html(data["data"]["message"])
                }
             },
             error:function(e){
                $("#auth_alert").removeClass("alert-info hidden").addClass("alert-danger").html("Please try again later")
                 console.log(e)
             }

           });
        }
    });
});
