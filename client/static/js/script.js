function redirect($url) {
    window.location = $url;
}

function deleteTodo(id) {
    $.ajax({
        url: "/todo/" + id,
        type: "DELETE",
    })
    .done(function(data) {
        window.location = "/";
    })
    .fail(function(jqXHR, textStatus, errorThrown) {
        console.log(jqXHR)
        console.log(textStatus)
        console.log(errorThrown)
        alert(errorThrown)
    })
}