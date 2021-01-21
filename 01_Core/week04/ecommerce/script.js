const queryStrings = {
    slideshowImage: '.slideshow__image',
    prevArrow: '.slider--prev',
    nextArrow: '.slider--next',
    mobileNavIcon: '#mobile-nav-icon',
    navList: '.nav__list',
    advancedSearchButton: '#advanced-search-button',
    advancedSearchContainer: '.advanced-search',
    advancedSearchClose: '.advanced-search__close'
}

function SlideArray (base, size) {
    this.cache = [];
    this.size = size;
    this.index = 0; 
    for (let i = 1; i <= size; i++) {
        let img = new Image();
        img.src = `${base}_${i}.jpg`
        this.cache.push(img);
    }
}

SlideArray.prototype.nextImage = function(element, direction='forward') {
    
    var element = document.querySelector(queryStrings.slideshowImage)
    var dots = document.querySelectorAll('.dot');
    dots[this.index].classList.remove('dot--selected');

    //Index correctly based on direction
    if (direction === 'forward') {
        this.index++;
    } else if (direction === 'backward') {
        this.index--; 
    } else {
        console.error("That is not a valid direction");
    }

    //If index has surpassed the size of the array, reset to 0
    if (this.index >= this.size) {
        this.index = 0;
    }

    //If index has gone negative, go to the last index in the array
    if (this.index < 0) {
        this.index = this.size - 1;
    }

    
    element.src = this.cache[this.index].src;
    dots[this.index].classList.add('dot--selected');

}

const toggleMobileNav = () => {
    let navMenu = document.querySelector(queryStrings.navList);
    let navIcon = document.querySelector(queryStrings.mobileNavIcon);
    navMenu.classList.toggle("display-mobile-nav");
}

const toggleAdvancedSearch = () => {
    let advancedSearchContainer = document.querySelector(queryStrings.advancedSearchContainer);
    advancedSearchContainer.style.display = 'block';
}

const closeWindow = (el) => {
    el.style.display = 'none';
}

const addListeners = () => {
    if (document.querySelector(queryStrings.nextArrow)) {
        document.querySelector(queryStrings.nextArrow).addEventListener('click', () => {
            slidesArray.nextImage(queryStrings.slideshowImageID);
        });
    }

    if (document.querySelector(queryStrings.prevArrow)) {
        document.querySelector(queryStrings.prevArrow).addEventListener('click', () => {
            slidesArray.nextImage(queryStrings.slideshowImageID, 'backward');
        });
    }
    if (document.querySelector(queryStrings.mobileNavIcon)) {
        document.querySelector(queryStrings.mobileNavIcon).addEventListener('click', () => {
            toggleMobileNav();
        })
    }

    if (document.querySelector(queryStrings.advancedSearchButton)) {
        document.querySelector(queryStrings.advancedSearchButton).addEventListener('click', () => {
            toggleAdvancedSearch()
        });
    }

    if(document.querySelector(queryStrings.advancedSearchClose)) {
        document.querySelector(queryStrings.advancedSearchClose).addEventListener('click', (event) => {
            closeWindow(event.target.parentNode);
        });
    }

}

const init = () => {

    addListeners();
 
}

let slidesArray = new SlideArray('images/slideshow/slideshow', 4);
init();

lsdfsdalks