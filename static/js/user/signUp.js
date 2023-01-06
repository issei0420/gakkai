const userTypeButtons = document.getElementsByClassName("user-type-button");
const submitButton = document.getElementsByClassName("submit-button")[0];

for (let button of userTypeButtons) {
    button.addEventListener("mouseover",  () => {
        button.style.backgroundColor = "#BFBFBF";
    }, false)

    button.addEventListener("mouseleave", () => {
        button.style.backgroundColor = "#dcdcdc";
    }, false);

    button.addEventListener("click", changeStatus, false);
}

function changeStatus() {
    // 初期化
    for (let button of userTypeButtons) {
        button.classList.remove("user-type-button-clicked");
    }
    this.classList.add("user-type-button-clicked");   
}

submitButton.addEventListener("mouseover", () => {
    submitButton.style.backgroundColor = "#cf6e07";
}, false)

submitButton.addEventListener("mouseleave", () => {
    submitButton.style.backgroundColor = "#DB7307"
}, false)

