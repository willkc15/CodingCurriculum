document.querySelector('#name-submit').addEventListener('click', () => {
    var value= document.querySelector('#names-list').value;
    names = value.split(',');
    names.forEach((name, index) => {
        if (name.charAt(0) === ' ') {
            names[index] = name.substring(1);
        }
    })
    names.sort();
    printNames(names);
})

function printNames(names) {
    var nameDisplay = document.querySelector(".name-display");
    names.forEach((name) => {
        var li = `<li>${name}</li>`;
        nameDisplay.insertAdjacentHTML('beforeend', li);
    })
    
}