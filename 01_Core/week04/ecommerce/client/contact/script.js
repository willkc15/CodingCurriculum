import {queryStrings, getQueryParamById, toggleMobileNav} from '../utils.js'

const formSuccess = () => {
    let html = '<p class="contact-form--success"><i class="far fa-check-square fa-lg" style="margin-right: 12px"></i>Form was successfully submitted</p>'
    document.querySelector(queryStrings.contactFormIntro).style.display = 'none'
    document.querySelector(queryStrings.contactFormIntro).insertAdjacentHTML('afterend', html)
}

const addEventListeners = () => {
    document.querySelector(queryStrings.mobileNavIcon).addEventListener('click',  toggleMobileNav)
}

addEventListeners()
let isSubmitted = getQueryParamById("submitted")
if (isSubmitted == 'true') formSuccess()