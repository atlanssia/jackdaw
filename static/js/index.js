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
                    $("brokers_table").html("<tr><th>Broker ID</th><th>Host</th><th>Port</th><th>Last Start Time</th><th>JMX Port</th><th>Version</th></tr>");
                    $.each( data, function( key, val ) {
                        var jsonObj = $.parseJSON(val);
                        $("#brokers_table").append(
                            "<tr><td>" + key + "</td><td>" + jsonObj.host + "</td><td>" + jsonObj.port + "</td><td>" + jsonObj.timestamp + "</td><td>" + jsonObj.jmx_port + "</td><td>" + jsonObj.version + "</td></tr>"
                        );
                    });
                }
            })
        })

        // map[groupName]map[topicName]map[partitionId]map[string]string
        $("[aria-controls='groups']").click(function (e) {
            $.ajax({
                url: "/groups",
                type: "GET",
                success: function(data) {
                    $(#groups_table).html("<tr><th>Group</th><th>Topic</th><th>Partition</th><th>offset</th><th>Log Size</th><th>Lag</th><th>Owner</th></tr>");
                    $.each( data, function( gName, val ) {
                        
                        $.each( val, function( topicName, val ) {
                            
                            $.each( val, function( partitionId, val ) {
                                
                                var offset = "", logSize = "", lag = "";
                                $.each( val, function( key, val ) {
                                    
                                    if (key === "offset") {
                                        offset = val;
                                    } else if (key === "logSize") {
                                        logSize = val;
                                    } else if (key === "lag") {
                                        lag = val;
                                    }
                                });

                                $("#groups_table").append(
                                    "<tr><td>" + gName + "</td><td>" + topicName + "</td><td>" + partitionId + "</td><td>" + offset + "</td><td>" + logSize + "</td><td>" + lag + "</td><td> - </td></tr>"
                                );
                                //$("#groups_table").html(val);
                            });
                            //$("#groups_table").html(val);
                        });
                        //$("#groups_table").html(val);
                    });
                }
            })
        })
    });
