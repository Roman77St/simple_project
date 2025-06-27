
// сброс значений формы
// document.getElementById("resetBtn").addEventListener("click", () =>  {
//     let form = document.getElementById("form");
//     form.reset();
// });

// отправка формы
// document.getElementById("saveBtn").addEventListener("click", async () => {
//     const id = document.getElementById("userId").value;
//     const name = document.getElementById("userName").value;
//     const age = document.getElementById("userAge").value;
//     const email = document.getElementById("userEmail").value;

//     if (id === "")
//         await createUser(name, age, email);
//     else
//         await editUser(id, name, age, email);
//     reset();
// });

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
        const user_token = await response.json();
        // document.querySelector("tbody").append(row(user));
        document.querySelector("").innerText()

    }
    else {
        const error = await response.json();
        console.log(error.message);
    }
}

