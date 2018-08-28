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

            // hide others
            document.getElementById("passwordView").style.display = "none";
            document.getElementById("passwordEdit").style.display = "none";
            document.getElementById("passwordDelete").style.display = "none";
        });
    }
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

            // hide others
            document.getElementById("passwordView").style.display = "none";
            document.getElementById("passwordEdit").style.display = "none";
            document.getElementById("passwordDelete").style.display = "none";
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
                // fill element and hide others
                el.value = decrypted;
                document.getElementById("passwordEdit").style.display = "none";
                document.getElementById("passwordDelete").style.display = "none";

                // show
                el.style.display = "block";
            } else {
                // hide
                el.style.display = "none";
            }
        });
    }
}

// show password edit form
var editbtn = document.querySelectorAll('.editbtn');
if (editbtn != null) {
    for (i = 0; i < editbtn.length; i++) {
        // add click listener
        editbtn[i].addEventListener('click', function (event) {
            // get element
            var el = document.getElementById("passwordEdit");

            // get passwordid, password and cookie hash
            var passwordid = event.target.getAttribute("passwordid");
            var password = document.getElementById("pw-" + passwordid);
            var cookie = getCookie('secpass_hash');

            // get input elements
            var input = document.getElementById("passwordEditInput");
            var elid = document.getElementById("passwordEditID");
            var eltitle = document.getElementById("passwordEditTitle");

            // decrypt
            var decrypted = sjcl.decrypt(cookie, password.value);

            // check that is is not displayed
            if (el.style.display === "none" || input.value != decrypted || (elid.value != password.name && eltitle.value != password.title)) {
                // fill elements
                input.value = decrypted;
                elid.value = password.name;
                eltitle.value = password.title;

                // hide others
                document.getElementById("passwordView").style.display = "none";
                document.getElementById("passwordDelete").style.display = "none";

                // show
                el.style.display = "block";
            } else {
                // hide
                el.style.display = "none";
            }
        });
    }
}

// show password delete button
var deletebtn = document.querySelectorAll('.deletebtn');
if (deletebtn != null) {
    for (i = 0; i < deletebtn.length; i++) {
        // add click listener
        deletebtn[i].addEventListener('click', function (event) {
            // get element
            var el = document.getElementById("passwordDelete");

            // get passwordid, input and title element
            var passwordid = document.getElementById("pw-" + event.target.getAttribute("passwordid"));
            var input = document.getElementById("passwordDeleteInput");
            var title = document.getElementById("passwordDeleteTitle");

            // check that is is not displayed
            if (el.style.display === "none" || input.value != passwordid) {
                // hide others
                document.getElementById("passwordView").style.display = "none";
                document.getElementById("passwordEdit").style.display = "none";

                // set value to name and title and show
                input.value = passwordid.name;
                input.checked = false;
                title.value = passwordid.title;
                el.style.display = "block";
            } else {
                //hide
                el.style.display = "none";
            }
        });
    }
}

// encrypt password
function manipulateAddPassword() {
    // forward title and name
    document.getElementById("title").value = document.getElementById("titlePre").value;
    document.getElementById("name").value = document.getElementById("namePre").value;

    // encrypt password
    var password = document.getElementById("passwordPre").value;
    document.getElementById("password").value = sjcl.encrypt(getCookie('secpass_hash'), password, { ks: 256 });

    // submit and return
    document.getElementById("addPassword").submit();
    return false;
}

// encrypt password
function manipulateEditPassword() {
    // forward title and name
    document.getElementById("passwordEditTitleAfter").value = document.getElementById("passwordEditTitle").value;
    document.getElementById("passwordEditIDAfter").value = document.getElementById("passwordEditID").value;

    // encrypt password
    var password = document.getElementById("passwordEditInput").value;
    document.getElementById("passwordEditInputAfter").value = sjcl.encrypt(getCookie('secpass_hash'), password, { ks: 256 });

    // submit and return
    document.getElementById("passwordEditAfter").submit();
    return false;
}

// generate password function
function generatePassword() {
    // chars to use
    charset = "abcdefghijklmnopqrstuvwxyz!ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_abcdefghijklmnopqrstuvwxyz#ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@";
    password = "";

    // add random chars
    for (var i = 0, n = charset.length; i < 16; ++i) {
        password += charset.charAt(Math.floor(Math.random() * n));
    }

    // return
    return password;
}

// generate random password
document.getElementById("generatePassword").addEventListener('click', function (event) {
    // get element and modify
    var el = document.getElementById("passwordPre");
    el.value = generatePassword();
});

// show password
document.getElementById("showPassword").addEventListener('click', function (event) {
    // get element
    var el = document.getElementById("passwordPre");

    // show or hide password
    if (el.type === "password") {
        el.type = "text";
    } else {
        el.type = "password";
    }
});

// close pwned list
var closePwnedList = document.getElementById("closePwnedList")
if (closePwnedList != null) {
    closePwnedList.addEventListener('click', function (event) {
        // reload page
        window.location.replace("/");
    });
}