// encrypt password and re-encrypt crypter
document.getElementById("submitbtn").addEventListener('click', function (event) {
    var currentPassword = document.getElementById("currentPassword");
    var newPassword = document.getElementById("newPassword");
    var repeatNewPassword = document.getElementById("repeatNewPassword");

    if (currentPassword.value != "" && newPassword.value != "" && repeatNewPassword.value != "") {
        document.getElementById("crypter").value = sjcl.encrypt(hashOne(newPassword.value), sessionStorage.getItem("crypter"), { ks: 256 });
        currentPassword.value = hash(currentPassword.value);
        newPassword.value = hash(newPassword.value);
        repeatNewPassword.value = hash(repeatNewPassword.value);
    }

    document.getElementById("settings").submit();
});