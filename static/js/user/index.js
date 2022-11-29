const menuLIs = document.getElementsByClassName("menu-li");

for (let li of menuLIs) {
    li.addEventListener("mouseover", () => {
        li.style.backgroundColor = "#f0ee8a";
    }, false);
    
    li.addEventListener("mouseleave", () => {
        li.style.backgroundColor = "#F9F790";
    }, false);

    li.addEventListener("click", switchPage, false);
}

function switchPage() {
    window.location.href = 'http://localhost:8080/status'
}


