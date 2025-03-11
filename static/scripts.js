function updateRangeValue(value) {

    const rangeInput = document.getElementById("range-slider");
    const rangeValue = document.getElementById("range-value");

    const min = parseInt(rangeInput.min);
    const max = parseInt(rangeInput.max);
    const percentage = ((value - min) / (max - min)) * 100;

    rangeValue.textContent = value;

    rangeInput.style.background = `linear-gradient(to right, var(--yellow-bios) 
                                ${percentage}%, var(--white-bios) ${percentage}%)`;
}



function toggleMode(){

    let modeTitle = document.getElementById("mode-title");
    let rulesContent = document.getElementById("rules-content");

    if(modeTitle.textContent === "HELP"){
        modeTitle.textContent = "OUTPUT";
        rulesContent.innerHTML = `
                <p class="output__description">Your generated password:</p>
                <div class="output__box">********</div>`;
    }else{
        modeTitle.textContent = "HELP";
        rulesContent.innerHTML = `
                <p>Rules for creating a good password:</p>
                <ol>
                    <li>The password must consist of characters, numbers, and letters.</li>
                    <li>The more characters in the password, the better.</li>
                    <li>The password should not contain the name and date of birth.</li>
                    <li>Do not use your personal information (full name, date of birth) as a keyword.</li>
                </ol>`;
    }
}



document.addEventListener("DOMContentLoaded", function () {
    
    document.getElementById("password-form").addEventListener("submit", function (event) {
        event.preventDefault();

        const formData = new FormData(this);
        const jsonData = Object.fromEntries(formData.entries());

        console.log("Отправляемые данные:", jsonData);

    });
});