const menuLIs = document.getElementsByClassName("menu-li");

for (let li of menuLIs) {
    li.addEventListener("mouseover", () => {
        li.style.backgroundColor = "#f0ee8a";
    }, false);
    
    li.addEventListener("mouseleave", () => {
        li.style.backgroundColor = "#F9F790";
    }, false)
}


