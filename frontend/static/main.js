// Получение всех пользователей
async function getUsers() {
    const response = await fetch("http://127.0.0.1:8080/api/users", {
        method: "GET",
        headers: {
            "Accept": "application/json",
         },
         mode: "cors"
    });
    if (response.ok === true) {

        const users = await response.json();
        const rows = document.querySelector("tbody");
        // добавляем полученные элементы в таблицу
        users.forEach(user => rows.append(row(user)));
        console.log(users)
    }
}

// Получение одного пользователя
async function getUser(id) {
    const response = await fetch(`http://127.0.0.1:8080/api/users/${id}`, {
        method: "GET",
        headers: { "Accept": "application/json" },
        mode: "cors"
    });
    if (response.ok === true) {
        const user = await response.json();
        document.getElementById("userId").value = user.id;
        document.getElementById("userName").value = user.name;
        document.getElementById("userAge").value = user.age;
        document.getElementById("userEmail").value = user.email;
    }
    else {
        // если произошла ошибка, получаем сообщение об ошибке
        const error = await response.json();
        console.log(error.message); // и выводим его на консоль
    }
}

// Добавление пользователя
async function createUser(userName, userAge, userEmail, userPassword) {
    const response = await fetch('http://127.0.0.1:8080/api/users', {
        method: "POST",
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json",
        },
        mode: "cors",
        body: JSON.stringify({
            name: userName,
            age: parseInt(userAge, 10),
            email: userEmail,
            password: userPassword,
        })
    });
    // if (response.ok === true) {
    //     const user = await response.json();
    //     document.querySelector("tbody").append(row(user));
    // }
    // else {
    //     const error = await response.json();
    //     console.log(error.message);
    // }
}

async function editUser(userId, userName, userAge, userEmail, userPassword) {
    const response = await fetch(`http://127.0.0.1:8080/api/users/${userId}`, {
        method: "PUT",
        headers: { "Accept": "application/json", "Content-Type": "application/json"},
        body: JSON.stringify({
            id: userId,
            name: userName,
            age: parseInt(userAge, 10),
            email: userEmail,
            password: userPassword,
        }),
        mode:"cors",
    });
    // if (response.ok === true) {
    //     const user = await response.json();
    //     document.querySelector(`tr[data-rowid='${user.id}']`).replaceWith(row(user));
    // }
    // else {
    //     const error = await response.json();
    //     console.log(error.message);
    // }
}

async function deleteUser(id) {
    const response = await fetch(`http://127.0.0.1:8080/api/users/${id}`, {
        method: "DELETE",
        headers: { "Accept": "application/json" },
        mode: "cors",
    });
    if (response.ok === true) {
        const user = await response.json();
        document.querySelector(`tr[data-rowid='${user.id}']`).remove();
    }
    else {
        const error = await response.json();
        console.log(error.message);
    }
}


function row(user) {

    const tr = document.createElement("tr");
    tr.setAttribute("data-rowid", user.id);

    const nameTd = document.createElement("td");
    nameTd.append(user.name);
    tr.append(nameTd);

    const ageTd = document.createElement("td");
    ageTd.append(user.age);
    tr.append(ageTd);

    const emailTd = document.createElement("td");
    emailTd.append(user.email);
    tr.append(emailTd);

    const linksTd = document.createElement("td");

    const editLink = document.createElement("button");
    editLink.append("Изменить");
    editLink.addEventListener("click", async() => await getUser(user.id));
    linksTd.append(editLink);

    const removeLink = document.createElement("button");
    removeLink.append("Удалить");
    removeLink.addEventListener("click", async () => await deleteUser(user.id));

    linksTd.append(removeLink);
    tr.appendChild(linksTd);

    return tr;
}
// сброс значений формы
document.getElementById("resetBtn").addEventListener("click", () =>  {
    let form = document.getElementById("form");
    form.reset();
});

// отправка формы
document.getElementById("saveBtn").addEventListener("click", async () => {
    const id = document.getElementById("userId").value;
    const name = document.getElementById("userName").value;
    const age = document.getElementById("userAge").value;
    const email = document.getElementById("userEmail").value;
    const password = document.getElementById("userPassword").value;
    let form = document.getElementById("form");
    form.reset();

    if (id === "")
        await createUser(name, age, email, password);
    else
        await editUser(id, name, age, email, password);
});

getUsers()