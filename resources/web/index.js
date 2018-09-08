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

// loop through entries and hide/show
function showCategory(category) {
    var entries = document.querySelectorAll('.passwordentry');
    if (entries != null) {
        for (i = 0; i < entries.length; i++) {
            var entry = entries[i];

            if (entry.getAttribute("category") == category) {
                // show
                entry.style.display = "block";
            } else {
                // hide
                entry.style.display = "none";
            }
        }
    }
}

// show only uncategorized
showCategory(document.getElementById("showall").getAttribute("name"));

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
                el.value = "";
                el.style.display = "none";
            }
        });
    }
}

// show category create
document.getElementById("createcategory").addEventListener('click', function (event) {
    // get element
    var el = document.getElementById("createCategory");

    // check that it is not displayed
    if (el.style.display === "none") {
        // show and hide others
        el.style.display = "block";
        document.getElementById("editCategory").style.display = "none";
        document.getElementById("deleteCategory").style.display = "none";
    } else {
        // hide
        el.value = "";
        el.style.display = "none";
    }
});

// show category edit
var editcategory = document.querySelectorAll('.editcategory');
if (editcategory != null) {
    for (i = 0; i < editcategory.length; i++) {
        // add click listener
        editcategory[i].addEventListener('click', function (event) {
            // get category id and category edit
            var id = event.target.getAttribute("id");
            var name = event.target.getAttribute("name");
            var el = document.getElementById("editCategory");

            // categoryID and categoryName elements
            var categoryID = document.getElementById("categoryID");
            var categoryName = document.getElementById("categoryName");

            // check that it is not displayed
            if (el.style.display === "none" || categoryID.value != id) {
                // fill element and show
                categoryID.value = id;
                categoryName.value = name;
                el.style.display = "block";

                // hide others
                document.getElementById("createCategory").style.display = "none";
                document.getElementById("deleteCategory").style.display = "none";
            } else {
                // hide
                el.value = "";
                el.style.display = "none";
            }
        });
    }
}

// show category delete
var deletecategory = document.querySelectorAll('.deletecategory');
if (deletecategory != null) {
    for (i = 0; i < deletecategory.length; i++) {
        // add click listener
        deletecategory[i].addEventListener('click', function (event) {
            // get category id and category delete elements
            var id = event.target.getAttribute("id");
            var el = document.getElementById("deleteCategory");
            var categoryDelete = document.getElementById("categoryDelete");

            // check that it is not displayed
            if (el.style.display === "none" || categoryDelete.value != id) {
                // fill element and show
                categoryDelete.value = id;
                el.style.display = "block";

                // hide others
                document.getElementById("createCategory").style.display = "none";
                document.getElementById("editCategory").style.display = "none";
            } else {
                // hide
                categoryDelete.value = "";
                el.style.display = "none";
            }
        });
    }
}

// view password with category
var viewcategory = document.querySelectorAll('.viewcategory');
if (viewcategory != null) {
    for (i = 0; i < viewcategory.length; i++) {
        // add click listener
        viewcategory[i].addEventListener('click', function (event) {
            // get category and show/hide
            var id = event.target.getAttribute("id");
            showCategory(id);
        });
    }
}

// show uncategorized entries
document.getElementById("showall").addEventListener('click', function (event) {
    showCategory(document.getElementById("showall").getAttribute("name"));
});

// close pwned list
var closePwnedList = document.getElementById("closePwnedList")
if (closePwnedList != null) {
    closePwnedList.addEventListener('click', function (event) {
        // reload page
        window.location.replace("/");
    });
}
