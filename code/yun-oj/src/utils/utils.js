import moment from "moment";

export function autoTextarea(elem, extra, maxHeight) {
    extra = extra || 0;
    let isFirefox = !!document.getBoxObjectFor || 'mozInnerScreenX' in window,
        isOpera = !!window.opera && !!window.opera.toString().indexOf('Opera'),
        addEvent = function (type, callback) {
            elem.addEventListener ? elem.addEventListener(type, callback, false) : elem.attachEvent('on' + type, callback);
        },
        getStyle = elem.currentStyle ? function (name) {
            let val = elem.currentStyle[name];
            if (name === 'height' && val.search(/px/i) !== 1) {
                let rect = elem.getBoundingClientRect();
                return rect.bottom - rect.top - parseFloat(getStyle('paddingTop')) - parseFloat(getStyle('paddingBottom')) + 'px';
            }
            return val;
        } : function (name) {
            return getComputedStyle(elem, null)[name];
        },
        minHeight = parseFloat(getStyle('height'));
    elem.style.resize = 'none';
    let change = function () {
        let scrollTop, height,
            padding = 0,
            style = elem.style;
        if (elem._length === elem.value.length) return;
        elem._length = elem.value.length;
        if (!isFirefox && !isOpera) {
            padding = parseInt(getStyle('paddingTop')) + parseInt(getStyle('paddingBottom'));
        }
        scrollTop = document.body.scrollTop || document.documentElement.scrollTop;
        elem.style.height = minHeight + 'px';
        if (elem.scrollHeight > minHeight) {
            if (maxHeight && elem.scrollHeight > maxHeight) {
                height = maxHeight - padding;
                style.overflowY = 'auto';
            } else {
                height = elem.scrollHeight - padding;
                style.overflowY = 'hidden';
            }
            style.height = height + extra + 'px';
            scrollTop += parseInt(style.height) - elem.currHeight;
            document.body.scrollTop = scrollTop;
            document.documentElement.scrollTop = scrollTop;
            elem.currHeight = parseInt(style.height);
        }
    };
    addEvent('propertychange', change);
    addEvent('input', change);
    addEvent('focus', change);
    change();
}


const debounce = (fn, delay) => {
    let timer = null;
    return function () {
        let context = this;
        let args = arguments;
        clearTimeout(timer);
        timer = setTimeout(function () {
            fn.apply(context, args);
        }, delay);
    }
}

export const _ResizeObserver = window.ResizeObserver;
window.ResizeObserver = class ResizeObserver extends _ResizeObserver {
    constructor(callback) {
        callback = debounce(callback, 16);
        super(callback);
    }
}


export function getLoginStatus() {
    return localStorage.getItem('U-Token') !== null
    // return localStorage.getItem('userInfo') !== null
}

export function formDate(time) {
    return moment(time * 1000).format('YYYY-MM-DD HH:mm')
}

export function isCoach() {
    const userInfo = JSON.parse(localStorage.getItem('userInfo'))
    return userInfo && userInfo.userType && userInfo.userType === 1
}