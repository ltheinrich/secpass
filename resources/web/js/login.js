// set crypter
var crypterElement = document.getElementById("crypter");
if (crypterElement.value != "") {
    var crypter = sjcl.decrypt(sessionStorage.getItem("password"), crypterElement.value);
    sessionStorage.removeItem("password");
    sessionStorage.setItem("crypter", crypter);
    window.location.replace("/");
}

// hash password and submit
document.getElementById("submitbtn").addEventListener('click', function (event) {
    var name = document.getElementById("name");
    var password = document.getElementById("password");

    if (name.value != "" && password.value != "") {
        sessionStorage.setItem("password", hashOne(password.value));
        password.value = hash(password.value);
        document.getElementById("login").submit();
    }
});