function updateRangeValue(value){
    const rangeValue = document.getElementById("range-value");
    rangeValue.textContent = value;
    const rangeInput = document.getElementById("range-slider");
    const percentage = ((value - rangeInput.min) / (rangeInput.max - rangeInput.min)) * 100;
    rangeValue.style.left = `calc(${percentage}%)`;
}