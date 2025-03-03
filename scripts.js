function updateRangeValue(value){
    const rangeValue = document.getElementById("range-value");
    rangeValue.textContent = value;
    const rangeInput = document.getElementById("range-slider");
    const percentage = ((value - rangeInput.min) / (rangeInput.max - rangeInput.min)) * 100;
    rangeValue.style.left = `calc(${percentage}%)`;
}

function toggleMode(){
    let modeTitle = document.getElementById("mode-title");
    let rulesContent = document.getElementById("rules-content");

    if(modeTitle.textContent === "HELP"){
        modeTitle.textContent = "OUTPUT";
        rulesContent.innerHTML = `
                <p class="output-description">Your generated password:</p>
                <div class="output-box">********</div>`;
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