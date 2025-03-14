let currentLang = localStorage.getItem("lang") || "en";



function updateRangeValue(value) {

    const rangeInput = document.getElementById("range-slider");
    const rangeValue = document.getElementById("range-value");

    const min = parseInt(rangeInput.min);
    const max = parseInt(rangeInput.max);
    const percentage = ((value - min) / (max - min)) * 100;

    rangeValue.textContent = value;

    rangeInput.style.background = `linear-gradient(to right, var(--yellow-bios) 
                                ${percentage}%, var(--white-bios) ${percentage}%)`;

    rangeInput.value = value;
}



function toggleMode() {
    let modeTitle = document.getElementById("mode-title");
    let rulesContent = document.getElementById("rules-content");
    let passwordContent = document.getElementById("password-content");

    if (modeTitle.getAttribute("data-key") === "rules") {
        modeTitle.setAttribute("data-key", "output");
        rulesContent.classList.add("hidden");
        passwordContent.classList.remove("hidden");
    } else {
        modeTitle.setAttribute("data-key", "rules");
        passwordContent.classList.add("hidden");
        rulesContent.classList.remove("hidden");
    }

    changeLanguage(currentLang);
}


// function onInput(el){
//     const length = el.value.length;

//     if(length >= 7){
//         updateRangeValue(length);
//     }
// }


function copyFunc(){
    const outputWord = document.getElementById("output-word");
    const text = outputWord.textContent || outputWord.innerText;
    const copyMessage = document.getElementById("copy-message");

    navigator.clipboard.writeText(text)
        .then(() =>{
            copyMessage.classList.add("active")

            setTimeout(() =>{
                copyMessage.classList.remove("active");
            }, 2000);
        })
        .catch(err => {
            alert("Ошибка копирования!");
        });
}



function changeLanguage(lang) {
    currentLang = lang;
    localStorage.setItem("lang", lang);

    document.querySelectorAll("[data-key]").forEach((el) => {
        let key = el.getAttribute("data-key");
        if (translations[key] && translations[key][lang]) {
            el.textContent = translations[key][lang];
        }
    });

    document.querySelectorAll("[data-placeholder]").forEach(input => {
         key = input.getAttribute("data-placeholder");
        if (translations[key] && translations[key][lang]) {
            input.setAttribute("placeholder", translations[key][lang]);
        }
    });
}



function toggleSettings() {
    document.getElementById("settings-panel").classList.toggle("active");
}



function changeTheme(theme){
    if(theme === 'light'){
        document.body.classList.remove("dark-theme");
        document.body.classList.add("light-theme");
        localStorage.setItem("theme", "light");
    } else {
        document.body.classList.remove("light-theme");
        document.body.classList.add("dark-theme");
        localStorage.setItem("theme", "dark");
    }
}




document.addEventListener("click", function (event) {
    const panel = document.getElementById("settings-panel");
    const button = document.querySelector(".header__settings-btn");

    if (!panel.contains(event.target) && !button.contains(event.target)) {
        panel.classList.remove("active");
    }
});


document.addEventListener("DOMContentLoaded", function () {
    changeLanguage(currentLang);
    const savedTheme = localStorage.getItem("theme") || "light";
    changeTheme(savedTheme);
});


document.addEventListener("DOMContentLoaded", function () {
    document.getElementById("password-form").addEventListener("submit", function (event) {
        event.preventDefault(); // Отключаем стандартную отправку формы

        const formData = new FormData(this);
        const queryParams = new URLSearchParams(formData).toString();

        fetch(`/generate?${queryParams}`)
            .then(response => response.json())
            .then(data => {
                console.log("Полученный пароль:", data.password);

                // Обновляем UI
                document.getElementById("mode-title").textContent = "OUTPUT";
                document.getElementById("rules-content").innerHTML = `
                    <p class="output__description">Your generated password:</p>
                    <div class="output__box">
                        <div class="output__container-output-word" id="output-word">
                            ${data.password}
                        </div>
                        <div class="output__container-copy-button">
                            <button class="output__copy-button" id="copy-button" onclick="copyFunc()">COPY</button>
                        </div>
                    </div>`;
            })
            .catch(error => {
                console.error("Ошибка:", error);
            });
    });
});
