// const userTypeButtons = document.getElementsByClassName("user-type-button");
const submitButton = document.getElementsByClassName("submit-button")[0];
const mailInput = document.getElementById("mail-input");
const passwordInput = document.getElementById("password-input");
let userType = "";

// for (let button of userTypeButtons) {
//     button.addEventListener("mouseover",  () => {
//         button.style.backgroundColor = "#BFBFBF";
//     }, false)

//     button.addEventListener("mouseleave", () => {
//         button.style.backgroundColor = "#dcdcdc";
//     }, false);

//     button.addEventListener("click", changeStatus, false);
// }

submitButton.addEventListener("mouseover", () => {
    submitButton.style.backgroundColor = "#cf6e07";
}, false)

submitButton.addEventListener("mouseleave", () => {
    submitButton.style.backgroundColor = "#DB7307"
}, false)


submitButton.addEventListener("click", submit, false)

// function changeStatus() {
//     // 初期化
//     for (let button of userTypeButtons) {
//         button.classList.remove("user-type-button-clicked");
//     }
//     this.classList.add("user-type-button-clicked");
//     userType = this.innerHTML;
// } 

function submit() {
    const data = {
        "mail": mailInput.value,
        "password": passwordInput.value,
        "userType": userType
    }

    send(data)
}

async function send(data){
    url = 'http://localhost:8080/signUp'
    const res = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });
    const resObj = await res.json();

    if (resObj["invalid"]) {
        invalidMessage(resObj["resmap"]);
    } else {
        window.alert("登録が完了しました")
    }
}

