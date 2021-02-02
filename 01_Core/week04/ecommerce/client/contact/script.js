import {queryStrings, getQueryParamById, toggleMobileNav} from '../utils.js'

const formSuccess = () => {
    let html = '<p class="contact-form--success">Form was submitted</p>'
    document.querySelector(queryStrings.contactFormIntro).style.display = 'none'
    document.querySelector(queryStrings.contactFormIntro).insertAdjacentHTML('afterend', html)
}

const addEventListeners = () => {
    document.querySelector(queryStrings.mobileNavIcon).addEventListener('click',  toggleMobileNav)
}

const init = () => {
    addEventListeners()
}

init()
let isSubmitted = getQueryParamById("submitted")
if (isSubmitted == 'true') formSuccess()