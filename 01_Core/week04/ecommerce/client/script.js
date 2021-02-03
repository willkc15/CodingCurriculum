import { queryStrings, getHtmlFromSpice, toggleMobileNav} from './utils.js'

//SlideArray is the array of slides, has a list of images, size, and current index
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

//Allows us to go to the next image, going either forward or backward
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

//Fetches "featured" sauces and displays them
const displayFeaturedSauces = () => {
    fetch('http://localhost:8000/sauces').then(function(response) {
        response.json().then(function(data) {
            let container = document.querySelector(queryStrings.homeProductsContainer)
            data.forEach((sauce) => {
                let spiceHtml = getHtmlFromSpice(sauce)
                if (sauce.featured === true) {
                    container.insertAdjacentHTML('beforeend', 
                    `<a href="/productDetails?id=${sauce.id}" class="product-list__card" id="${sauce.id}">
                    <img src="images/product-tall/${sauce.imgPath}" alt="${sauce.name}" class="product-list__image">
                    <h5 class="product-list__title">${sauce.brand} - ${sauce.name}</h5>
                    <div class="product-list__hotness--tall">
                        ${spiceHtml} 
                    </div>
                    <div class="product-list__volume">${sauce.size} fl oz.</div>
                    </a>`) 
                }
            })
        })
    })
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
}

addListeners();
displayFeaturedSauces();
let slidesArray = new SlideArray('images/slideshow/slideshow', 4);
