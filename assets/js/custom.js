window.sr = new ScrollReveal({ reset: true });

// Add class to <html> if ScrollReveal is supported
if (sr.isSupported()) {
    document.documentElement.classList.add('sr');
}

$(document).ready(function () {
    $('.scroll-down').click(function (event) {
        var scrollTo = $(this).offset().top - 5;
        if ((scrollTo > $(window).scrollTop() - 25) && (scrollTo < $(window).scrollTop() + 25)) {
            var scrollers = $('.scroll-down');
            if (scrollers.length > 1) {
                var target = scrollers.eq((scrollers.index($(this))+1) % scrollers.length);
                if (target.offset().top < scrollTo) {
                    scrollTo = $('footer').offset().top
                } else {
                    scrollTo = target.offset().top - 5 - target.height();
                }
            } else {
                scrollTo = $('footer').offset().top
            }
        }
        var duration = Math.sqrt(Math.abs($('html, body').scrollTop() - scrollTo)) * 16;
        $('html, body').animate({scrollTop: scrollTo}, duration);
    });

    $('.scroll-up').click(function (event) {
        $('html, body').animate({scrollTop: 0}, 500);
    });

    sr.reveal('.reveal');

    $('img').tilt({
        glare: true,
        maxGlare: 1,
        scale: 1.2
    });
});
