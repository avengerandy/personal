document.addEventListener("DOMContentLoaded", function() {
    document.getElementsByClassName("navigation-icon")[0].addEventListener("click", function() {
        this.classList.toggle('navigation-icon-x');
        document.getElementsByClassName("navigation")[0].classList.toggle('navigationOpen');
    });
});
