// get cookie
function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');

    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];

        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }

        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }

    return "";
}

// clear clipboard
var clearbtn = document.querySelectorAll('.clearbtn');
if (clearbtn != null) {
    for (i = 0; i < clearbtn.length; i++) {
        // add click listener
        clearbtn[i].addEventListener('click', function (event) {
            // define textarea and fill
            var el = document.createElement("textarea");
            el.value = " ";

            // add element to page and select
            document.body.appendChild(el);
            el.select();

            // copy and remove element
            document.execCommand('copy');
            document.body.removeChild(el);
        });
    }
}

// copy password to clipboard
var copybtn = document.querySelectorAll('.copybtn');
if (copybtn != null) {
    for (i = 0; i < copybtn.length; i++) {
        // add click listener
        copybtn[i].addEventListener('click', function (event) {
            // define textarea to copy
            var el = document.createElement("textarea");

            // get passwordid, password and cookie hash
            var passwordid = event.target.getAttribute("passwordid");
            var password = document.getElementById("pw-" + passwordid).value;
            var cookie = getCookie('secpass_hash');

            // decrypt and fill element
            el.value = sjcl.decrypt(cookie, password);

            // add element to page and select
            document.body.appendChild(el);
            el.select();

            // copy and remove element
            document.execCommand('copy');
            document.body.removeChild(el);
        });
    }
}

// show password view
var viewbtn = document.querySelectorAll('.viewbtn');
if (viewbtn != null) {
    for (i = 0; i < viewbtn.length; i++) {
        // add click listener
        viewbtn[i].addEventListener('click', function (event) {
            // get passwordid, password and cookie hash
            var passwordid = event.target.getAttribute("passwordid");
            var password = document.getElementById("pw-" + passwordid);
            var cookie = getCookie('secpass_hash');

            // get element and decrypt
            var el = document.getElementById("passwordView");
            var decrypted = sjcl.decrypt(cookie, password.value);

            // check that it is not displayed
            if (el.style.display === "none" || el.value != decrypted) {
                // fill element and show
                el.value = decrypted;
                el.style.display = "block";
            } else {
                // hide
                el.style.display = "none";
            }
        });
    }
}

// close pwned list
var closePwnedList = document.getElementById("closePwnedList")
if (closePwnedList != null) {
    closePwnedList.addEventListener('click', function (event) {
        // reload page
        window.location.replace("/");
    });
}
