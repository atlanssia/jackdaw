/**
 * Created on 10/23/15.
 */

/* Define API endpoints once globally */
$.fn.api.settings.api = {
    'get followers' : '/followers/{id}?results={/count}',
    'create user'   : '/create',
    'add user'      : '/add/{id}',
    'follow user'   : '/follow/{id}',
    'search'        : '/search/?query={value}'
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
                action: 'add user',
                //on: 'mouseenter',
                urlData: {
                    id: 23
                },
                onResponse: function(response) {
                    // make some adjustments to response
                    return response;
                },
                successTest: function(response) {
                    // test whether a json response is valid
                    return response.success || false;
                },
                onComplete: function(response) {
                    // always called after xhr complete
                },
                onSuccess: function(response) {
                    // valid response and response.success = true
                },
                onFailure: function(response) {
                    // request failed, or valid response but response.success = false
                },
                onError: function(errorMessage) {
                    // invalid response
                },
                onAbort: function(errorMessage) {
                    // navigated to a new page, CORS issue, or user canceled request
                }
            })
            .state({
                button: {

                }
            })
        ;
    })
;
