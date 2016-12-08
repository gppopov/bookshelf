$(function () {
    $.ajax({
        type: "GET",
        url: "/books"
    }).done(function (data) {
        for (var i = 0; i < data.length; i++) {
            console.log(data[i]);
            $("#items").append("<li>" + data[i].name + "</li>");
        }
    }).fail(function (x, e, message) {
        alert(message);
    });
});