varLoginClearError=function(a){
    if(a){
        varb=a.parent(".login-item");b.removeClass("error");varc=$(".tooltips[data-origin='"+a.attr("id")+"']");c.stop(!0,
        !0).animate({
            left: "-=40"
        },
        100,
        function(){
            c.fadeOut(100)
        }),
        a.timer&&clearTimeout(a.timer)
    }
},
LoginShowError=function(a,
b){
    if(a){
        varc=a.parent(".login-item");c.addClass("error"),
        a.timer=null;vard=$(".tooltips[data-origin='"+a.attr("id")+"']");d.find(".tooltips-content").text(b).end().show().stop(!0,
        !0).animate({
            left: d.data("currentleft")
        },
        200),
        "1"==a.data("autohide")&&(a.timer&&clearTimeout(a.timer),
        a.timer=setTimeout(function(){
            LoginClearError(a)
        },
        1e3))
    }
};$(function(){
    $(".tooltips").each(function(){
        var a=$(this),
        b=a.data("origin"),
        c=$("#"+b),
        d=c.offset(),
        e={
            height: c.outerHeight(!0),
            width: c.outerWidth(!0)
        };a.css({
            //top: d.top+e.height/2-a.outerHeight(!0)/2,
            //left: d.left+e.width-23 d.left+
        }).data("currentleft",
        e.width+17),
        a.mouseenter(function(){
            a.stop(!0,
            !0).animate({
                left: "-=40"
            },
            100,
            function(){
                a.fadeOut(100)
            })
        })
    }),
    $(".j-clearIpt").click(function(){
        return$(this).siblings("input").val("").focus(),
        !1
    })
});