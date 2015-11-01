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

        $("[aria-controls='brokers']").click(function (e) {
            $.ajax({
                url: "/brokers",
                type: "GET",
                success: function(data) {
                    alert(data);
                    $.each( data, function( key, val ) {
                        var jsonObj = $.parseJSON(val)
                        $("#brokers_table").append(
                            "<tr><td>" + key + "</td><td>" + jsonObj.host + "</td><td>" + jsonObj.port + "</td><td>" + jsonObj.timestamp + "</td><td>" + jsonObj.jmx_port + "</td><td>" + jsonObj.version + "</td></tr>"
                        );
                    });
                }
            })
        })

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
