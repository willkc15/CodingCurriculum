import {queryStrings, getQueryParamById, toggleMobileNav} from '../utils.js'

const getSingleSauce = (id) => {
    fetch(`http://localhost:8000/sauces/${id}`)
    .then(response => response.json())
    .then(data => {
        let sauce = data[0]
        let productDetailsContainer = document.querySelector(queryStrings.productDetailsContainer);
        productDetailsContainer.insertAdjacentHTML('beforeend', `
        <div class="product-details__image">
            <img src="/images/product-square/${sauce.imgPath}" alt="${sauce.name}">
        </div>
        <div class="product-details__container">
            <h2>${sauce.brand} | ${sauce.name}</h2>
            <div class="product-details__reviews">
                <i class="fas fa-star"></i>
                <i class="fas fa-star"></i>
                <i class="fas fa-star"></i>
                <i class="fas fa-star"></i>
                <i class="fas fa-star-half-alt"></i>
                <span>127 Reviews</span>
            </div>
            <h3>$${sauce.price}</h3>
            <div class="button btn-md btn-cart">
                <i class="fas fa-shopping-cart"></i>
                <span>Add to Cart</span>
            </div>
            <p class="product-details__description">${sauce.description}</p>
            <p class="product-details__ingredients"><b>Ingredients:</b> ${sauce.ingredients}</p>
            <p class="product-details__size"><b>Size:</b> ${sauce.size} fl oz.</p>
        </div>
        `)
    })
}



const addEventListeners = () => {
    if (document.querySelector(queryStrings.mobileNavIcon)) {
        document.querySelector(queryStrings.mobileNavIcon).addEventListener('click', () => {
            toggleMobileNav();
        })
    }
}

const init = () => {
    addEventListeners()
    let paramId = getQueryParamById('id')
    getSingleSauce(paramId);
}

init();