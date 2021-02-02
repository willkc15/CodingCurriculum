export const queryStrings = {
    slideshowImage: '.slideshow__image',
    slideshowContainer: '.slideshow-container',
    prevArrow: '.slider--prev',
    nextArrow: '.slider--next',
    mobileNavIcon: '#mobile-nav-icon',
    navList: '.nav__list',
    homeProductsContainer: '#home-page-products',
    productDetailsContainer: '#product-details-container',
    productsPageContainer: '#product-page-container',
    searchBar: '#search-bar',
    searchValuesContainer: '.search__values__container',
    searchValue: '.search__value',
    contactForm: '#contact-form',
    contactFormContainer: '.contact-form',
    contactFormIntro: '#form-intro'
}

//Takes a sauce, and returns the html according to its spice level
export const getHtmlFromSpice = (sauce) => {
    let html = ''
    for (var i = 0; i < sauce.spice; i++) {
            html += '<i class="fas fa-pepper-hot fa-md"></i>'
    }
    return html
}

//Toggles the mobile nav menu on and off
export const toggleMobileNav = () => {
    let navMenu = document.querySelector(queryStrings.navList);
    navMenu.classList.toggle("display-mobile-nav");
}

export const getQueryParamById = (param) => {
    const params = new URLSearchParams(window.location.search)
    return params.get(param)
}