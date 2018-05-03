$(document).ready(function () {
    $('.scroll-down').click(function (event) {
        $('html, body').animate({
            scrollTop: $(this).offset().top + 1
        }, 500);
    });

    $('.scroll-up').click(function (event) {
        $('html, body').animate({
            scrollTop: 0
        }, 500);
    });

    window.sr = ScrollReveal({ reset: true });
    sr.reveal('.reveal');

    $('img').tilt({
        glare: true,
        maxGlare: 1,
        scale: 1.2
    });
});
