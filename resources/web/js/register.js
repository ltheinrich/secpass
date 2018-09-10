// hash password, generate crypter and submit
document.getElementById("submitbtn").addEventListener('click', function (event) {
    var name = document.getElementById("name");
    var password = document.getElementById("password");
    var repeat = document.getElementById("repeat");

    if (name.value != "" && password.value != "" && repeat.value != "") {
        document.getElementById("crypter").value = sjcl.encrypt(hashOne(password.value), uuid(), { ks: 256 });
        password.value = hash(password.value);
        repeat.value = hash(repeat.value);
        document.getElementById("register").submit();
    }
});