$("#show-pass-check").on("click", function() {
    var show = this.checked;
    toggle_password("pass-input", show);
});

function toggle_password(target, show){
    var d = document;
    var tag = d.getElementById(target);

    if (show){
        tag.setAttribute('type', 'text');
    } else {
        tag.setAttribute('type', 'password');
    }
}