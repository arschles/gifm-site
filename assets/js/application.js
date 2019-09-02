var $ = require("jquery");
require("bootstrap");
var Turbolinks = require("turbolinks")

$(() => {
    // TODO: there are gotchas with this, check out the following:
    // https://www.honeybadger.io/blog/turbolinks/
    // Turbolinks.start();

    // $("#gifm-rotating-topics").Morphext({
    //     // The [in] animation type. Refer to Animate.css for a list of available animations.
    //     animation: "bounceIn",
    //     // An array of phrases to rotate are created based on this separator. Change it if you wish to separate the phrases differently (e.g. So Simple | Very Doge | Much Wow | Such Cool).
    //     separator: ",",
    //     // The delay between the changing of each phrase in milliseconds.
    //     speed: 2500,
    //     complete: function () {
    //         // Called after the entrance animation is executed.
    //     }
    // });

    $(function () {
        $('[data-toggle="tooltip"]').tooltip()
    })

});
