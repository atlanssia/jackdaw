/**
 * Created on 10/23/15.
 */
$(document)
    .ready(function() {
//        $("#topics").click(function(e) {
//            $('.nav li').removeClass('active');
//            $.ajax({
//                url: "/topics",
//                type: "GET",
//                success: function(data) {
//                    alert(data);
//                    $.each( data, function( key, val ) {
//                        alert(key + " : " + val);
//                    });
//                }
//            })
//        })

        $("[aria-controls='topics']").click(function (e) {
            $.ajax({
                url: "/topics",
                type: "GET",
                success: function(data) {
                    alert(data);
                    $.each( data, function( key, val ) {
                        alert(key + " : " + val);
                        $("#topics").html(val);
                    });
                }
            })
        })
    });
