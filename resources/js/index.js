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

// copy password to clipboard
document.querySelector('.copybtn').addEventListener('click',
    function (event) {
        // encryption comment
        // event.srcElement.id = sjcl.encrypt(getCookie('secpass_hash'), "Test Passwort", { ks: 256 })

        // define textarea to copy
        var el = document.createElement("textarea")

        // get passwordid and password
        var passwordid = event.srcElement.getAttribute("passwordid")
        var password = document.getElementById(passwordid).value

        // cookie hash
        var cookie = getCookie('secpass_hash')

        // decrypt and fill element
        el.value = sjcl.decrypt(cookie, password);

        // add element to page and select
        document.body.appendChild(el);
        el.select();

        // copy and remove element
        document.execCommand('copy');
        document.body.removeChild(el);

        // hide others
        document.getElementById("passwordView").style.display = "none"
        document.getElementById("passwordEdit").style.display = "none"
        document.getElementById("passwordDelete").style.display = "none"
    }
);

// show password view
document.querySelector('.viewbtn').addEventListener('click',
    function (event) {
        // get element
        var el = document.getElementById("passwordView")

        // check that it is not displayed
        if (el.style.display === "none") {
            // get passwordid and password
            var passwordid = event.srcElement.getAttribute("passwordid")
            var password = document.getElementById(passwordid).value

            // cookie hash
            var cookie = getCookie('secpass_hash')

            // decrypt and fill element
            el.value = sjcl.decrypt(cookie, password);

            // hide other password elements
            document.getElementById("passwordEdit").style.display = "none"
            document.getElementById("passwordDelete").style.display = "none"

            // show
            el.style.display = "block";
        } else {
            // hide
            el.style.display = "none";
        }
    }
);

// show password edit form
document.querySelector('.editbtn').addEventListener('click',
    function (event) {
        // get element
        var el = document.getElementById("passwordEdit")

        // check that is is not displayed
        if (el.style.display === "none") {
            // get passwordid and password
            var passwordid = event.srcElement.getAttribute("passwordid")
            var password = document.getElementById(passwordid).value

            // cookie hash
            var cookie = getCookie('secpass_hash')

            // get input element
            var input = document.getElementById("passwordEditInput")

            // decrypt and fill element
            input.value = sjcl.decrypt(cookie, password);

            // hide others
            document.getElementById("passwordView").style.display = "none"
            document.getElementById("passwordDelete").style.display = "none"

            // show
            el.style.display = "block";
        } else {
            // hide
            el.style.display = "none";
        }
    }
);

// show password delete button
document.querySelector('.deletebtn').addEventListener('click',
    function (event) {
        // get element
        var el = document.getElementById("passwordDelete")

        // check that is is not displayed
        if (el.style.display === "none") {
            // hide others
            document.getElementById("passwordView").style.display = "none"
            document.getElementById("passwordEdit").style.display = "none"

            // get passwordid
            var passwordid = event.srcElement.getAttribute("passwordid")

            // get input element
            var input = document.getElementById("passwordDeleteInput")

            // set value to passwordid
            input.value = passwordid

            // show
            el.style.display = "block";
        } else {
            //hide
            el.style.display = "none";
        }
    }
);