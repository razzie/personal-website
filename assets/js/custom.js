$(document).ready(function () {
    $('.scroll-down').click(function (event) {
        var scrollTo = $(this).offset().top - 5;
        if ((scrollTo > $(window).scrollTop() - 5) && (scrollTo < $(window).scrollTop() + 5)) {
            var scrollers = $('.scroll-down');
            var target = scrollers.eq((scrollers.index($(this))+1) % scrollers.length);
            scrollTo = target.offset().top - 5 - target.height();
        }
        $('html, body').animate({scrollTop: scrollTo}, 500);
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
