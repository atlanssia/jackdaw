/**
 * Created on 10/23/15.
 */

/* Define API endpoints once globally */
$.fn.api.settings.api = {
    'list topics' : '/topics',
    'list brokers' : '/brokers',
    'add user'      : '/add/{id}',
    'follow user'   : '/follow/{id}',
    'search'        : '//api.github.com/search/repositories?q=golang'
};

$(document)
    .ready(function() {

        // fix menu when passed
        $('.masthead')
            .visibility({
                once: false,
                onBottomPassed: function() {
                    $('.fixed.menu').transition('fade in');
                },
                onBottomPassedReverse: function() {
                    $('.fixed.menu').transition('fade out');
                }
            })
        ;

        // create sidebar and attach to menu open
        $('.ui.sidebar')
            .sidebar('attach events', '.toc.item')
        ;

        $('.follow.button')
            .api({
                action: 'follow user',
                //on: 'mouseenter',
                urlData: {
                    id: 22
                }
            })
        ;

        $('.item.topics')
            .api({
                action: 'list topics',
                //on: 'mouseenter',
                onResponse: function(response) {
                    // make some adjustments to response
                    alert('response goes here')
                    return response;
                },
                //successTest: function(response) {
                //    // test whether a json response is valid
                //    alert('successTest: ' + response.success)
                //    return response.success || false;
                //},
                onComplete: function(response) {
                    // always called after xhr complete
                    alert('onComplete goes here')
                },
                onSuccess: function(response) {
                    // valid response and response.success = true
                    //if ($.isArray(response)) {
                    //    response.map(function (topic) {
                    //        $('.segment').html(response.a)
                    //    });
                    //} else {
                    //    $('.segment').html(response.a)
                    //}

                    $(".segment").html("");
                    $.each(response.items, function(index, item) {

                    });
                },
                onFailure: function(response) {
                    // request failed, or valid response but response.success = false
                    alert('onFailure goes here')
                },
                onError: function(errorMessage) {
                    // invalid response
                    alert('onError goes here')
                },
                onAbort: function(errorMessage) {
                    // navigated to a new page, CORS issue, or user canceled request
                    alert('onAbort goes here')
                }
            })
            .state({
                button: {

                }
            })
        ;
    })
;
