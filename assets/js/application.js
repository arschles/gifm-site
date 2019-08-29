require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
var Turbolinks = require("turbolinks")

$(() => {
    // TODO: there are gotchas with this, check out the following:
    // https://www.honeybadger.io/blog/turbolinks/
    Turbolinks.start()
});
