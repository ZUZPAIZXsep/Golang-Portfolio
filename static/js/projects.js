const modal = document.getElementById("projectModal");
const closeButton = document.getElementById("projectModalClose");

const projectButtons = document.querySelectorAll(".show-detail");

const modalProjectName =
    document.getElementById("modalProjectName");

const modalProjectDescription =
    document.getElementById("modalProjectDescription");


projectButtons.forEach(button => {

    button.addEventListener("click", function () {

        const name = this.dataset.name;
        const description = this.dataset.description;

        const projectItem = this.closest(".project-item");

        const bullets =
            projectItem.querySelectorAll(
                ".project-bullets-data li"
            );

        const technologies =
            projectItem.querySelectorAll(
                ".project-technologies-data span"
            );


        // Project name
        modalProjectName.textContent = name;


        // Description
        modalProjectDescription.textContent = description;


        // Bullets
        const modalBullets =
            document.getElementById("modalProjectBullets");

        modalBullets.innerHTML = "";

        bullets.forEach(bullet => {

            const li = document.createElement("li");

            li.textContent = bullet.textContent;

            modalBullets.appendChild(li);

        });


        // Technologies
        const modalTechnologies =
            document.getElementById(
                "modalProjectTechnologies"
            );

        modalTechnologies.innerHTML = "";

        technologies.forEach(technology => {

            const span = document.createElement("span");

            span.textContent = technology.textContent;

            modalTechnologies.appendChild(span);

        });


        // Open modal
        modal.classList.add("show");

    });

});


closeButton.addEventListener("click", function () {

    modal.classList.remove("show");

});


modal.addEventListener("click", function (event) {

    if (event.target === modal) {
        modal.classList.remove("show");
    }

});