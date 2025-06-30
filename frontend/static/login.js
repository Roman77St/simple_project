
// login
document.getElementById("confirmBtn").addEventListener("click", async () =>  {
    const email = document.getElementById("user-Email").value;
    const password = document.getElementById("user-Password").value;
    await loginUser(email, password)
});

// login пользователя
async function loginUser(userEmail, userPassword) {
    const response = await fetch('http://127.0.0.1:8082/api/auth/login', {
        method: "POST",
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json",
        },
        mode: "cors",
        body: JSON.stringify({
            email: userEmail,
            password: userPassword,
        })
    });
    if (response.ok === true) {
        const user_token_json = await response.json();
        const {token, userEmail} = user_token_json;
        document.getElementById("token").innerHTML = userEmail + " Вы успешно авторизованы";
        setJwtCookie(token);
    }
    else {
        const error = await response.json();
        console.log(error.message);
    }
}

function setJwtCookie(token) {
    const d = new Date();
    // Устанавливаем срок действия cookie, например, на 1 день
    d.setTime(d.getTime() + (1 * 24 * 60 * 60 * 1000)); // 1 день в миллисекундах
    const expires = "expires=" + d.toUTCString();

    // Устанавливаем cookie
    // 'jwtToken' - имя вашего cookie
    // token - значение JWT токена
    // expires - срок действия
    // path=/ - делает cookie доступным для всего домена
    document.cookie = "jwtToken=" + token + ";" + expires + ";path=/;SameSite=Lax;Secure";
    // Рекомендуется использовать SameSite=Lax (или Strict) для защиты от CSRF
    // Secure - отправлять cookie только по HTTPS (обязательно для продакшена)
}
