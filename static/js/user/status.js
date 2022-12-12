const menuLIs = document.getElementsByClassName("menu-li");
const menuIndex = document.getElementById("menu-index");

for (let li of menuLIs) {
    li.addEventListener("mouseover", () => {
        li.style.backgroundColor = "#f0ee8a";
    }, false);
    
    li.addEventListener("mouseleave", () => {
        li.style.backgroundColor = "#F9F790";
    }, false);
}

menuIndex.addEventListener("click", switchPage, false);

function switchPage() {
    window.location.href = 'http://localhost:8080/index'
}