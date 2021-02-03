import {queryStrings, getHtmlFromSpice, toggleMobileNav} from '../utils.js'

//Fetches sauces from api and returns a promise
const fetchSauces = async () => {
    const response = await fetch('http://localhost:8000/sauces')
    return response.json()
}

//Take promise as an argument where promise will resolve into an array of sauces
const displaySauces = (promise) => {
    let container = document.querySelector(queryStrings.productsPageContainer)
    promise.then(data => {
        data.forEach((sauce) => {
            let spiceHtml = getHtmlFromSpice(sauce)
            container.insertAdjacentHTML('beforeend',` 
            <a href="/productDetails?id=${sauce.id}" class="product-list__item">
                <img src="images/product-square/${sauce.imgPath}" alt="${sauce.name}" class="product-list__image">
                <h5 class="product-list__title">${sauce.brand} | ${sauce.name}</h5>
                <div class="product-list__hotness">
                    ${spiceHtml}
                </div>
                <div class="product-list__price">$${sauce.price}</div>
            </a>`)
        })
    })
}

// Adds hidden products to search bar which can be shown/hidden based off user input
const activateSearchBar = (promise) => {
    let searchValues = document.querySelector(queryStrings.searchValuesContainer)
    promise.then(data => {
        data.forEach((sauce) => {
            searchValues.insertAdjacentHTML('beforeend', `
            <a class="search__value" href="productDetails?id=${sauce.id}" >
                <img src="images/product-square/${sauce.imgPath}" class="search__value__image" alt="${sauce.name}">
                <div class="search__value__name">${sauce.brand} | ${sauce.name}</div>
            </a>   
            `)
        })
    })
}

//Handles the search bar functionality, called every time there is a new value in the search bar
const filterSearch = (searchString) => {
    let searchValues = document.querySelectorAll(queryStrings.searchValue)
    for (let value of searchValues) {
        if (value.children[1].innerHTML.toUpperCase().includes(searchString.toUpperCase())) {
            value.style.display = 'flex'
        } else {
            value.style.display = 'none'
        }
    }
}

const addEventListeners = () => {
    //If input changes from blank, show the search bar values pane. If it goes back to blank, hide it
    document.querySelector(queryStrings.searchBar).addEventListener('input', (event) => {
        if (event.target.value != '') {
            document.querySelector(queryStrings.searchValuesContainer).style.display = "block"
        } else {
            document.querySelector(queryStrings.searchValuesContainer).style.display = "none"
        }
        filterSearch(event.target.value)
    })

    document.querySelector(queryStrings.mobileNavIcon).addEventListener('click', () => {
        toggleMobileNav();
    })

}

addEventListeners()
let promise = fetchSauces()
displaySauces(promise)
activateSearchBar(promise)
