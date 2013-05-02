var dat;
function handle_dot(data) {
    dat = data;
    try {
        $("#content").html(Viz(data, "svg"));
    }
    finally {
    }
}

function get_dot() {
    $.ajax({
        url: "/dot",
    }).done(handle_dot);
    setTimeout(get_dot, 200);
}

$(document).ready(function(){
    get_dot();
});
