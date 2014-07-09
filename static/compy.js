;// learning jquery yay

// change to your web server... when compy is production-ready ^_^
var relativeurl = "localhost:4444"

commentform = "<form id='commentform'>" +
    "pageid: <input type='text' id='formpageid' name='pageid'></br>" +
    "author: <input type='text' id='formauthor' name='author'></br>" +
    "email: <input type='text' id='formemail' name='email'></br>" +
    "body: <input type='text' id='formbody' name='body'></br>" +
    "<input type='submit' value='submit'>" +
    "</form>";

$("#compy-comments").append(commentform);

$("#commentform").submit(function(evt) {
    evt.preventDefault();
    var pageid = $("#formpageid").val();
    var author = $("#formauthor").val();
    var email  = $("#formemail").val();
    var body   = $("#formbody").val();
    if (pageid == "" || author == "" || body == "") {
        $("compy-comments").append("<p>error D:</p>");
    } else {
        // assemble json data
        jsonobj = {
            "pageid" : pageid,
            "author" : author,
            "email"  : email,
            "body"   : body,
        };
        console.log("turn back now, this doesn't work")
        // $.ajax({
        //     url: relativeurl + "/comments/add",
        //     dataType: 'json',
        //     type: 'POST',
        //     data: jsonobj,
        //     error: function(xhr, status, err) {
        //         console.error("/comments/add", status, err.toString())
        //     }.bind(this)
        // });
    }
});
