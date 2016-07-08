// hides the slickbox as soon as the DOM is ready (a little sooner that page load)
// $('#slickbox').hide();
$('#slickbox2').hide();
$('#slickbox3').hide();
var counter = 0;
$('#slick-slidetoggle').click(function () {
    if (counter == 0) {

        $('#slickbox').hide('slide', 4000);
        $('#slickbox2').show('slide', {
            direction: 'right'
        }, 4000);
        counter += 1;
    } else if (counter == 1) {

        $('#slickbox2').hide('slide', 4000);
        $('#slickbox3').show('slide', {
            direction: 'right'
        }, 4000);
    }
    return false;
});