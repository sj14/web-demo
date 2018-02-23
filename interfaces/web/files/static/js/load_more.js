// https://stackoverflow.com/questions/41136000/creating-load-more-button-in-golang-with-templates
function loadMore() {
    var e = document.getElementById("nextBatch");
    var xhr = new XMLHttpRequest();
    var txt_no_more = document.getElementById("txt-no_more")
    txt_no_more.style.visibility = "hidden"

    xhr.open("GET", window.url + "&limit=" + window.limit + "&offset=" + window.offset, true);

    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
            e.outerHTML = xhr.responseText;
        }
        else if (xhr.status == 204) {
            // hide more button
            var btn_more = document.getElementById("btn-more");
            var txt_no_more = document.getElementById("txt-no_more")
            btn_more.style.visibility = "hidden";
            txt_no_more.style.visibility = "visible"
        }
    }
    window.offset += window.limit;
    try { xhr.send(); } catch (err) { /* handle error */ }
}

window.onload = function () {
    loadMore();
};