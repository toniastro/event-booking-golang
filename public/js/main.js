"use strict"
$(window).on("load", function() {
    function countDown(){
        var today = new Date();
        var eventDate = new Date("Nov 28,2019 00:00:00");
        var currentTime = today.getTime();
        var eventTime = eventDate.getTime();
        var remTime = eventTime - currentTime;

        var sec = Math.floor(remTime/1000);
        var min = Math.floor(sec/60);
        var hrs = Math.floor(min/60);
        var days = Math.floor(hrs/24);

        hrs %= 24;
        min %= 60;
        sec %= 60;

        days = (days<10) ? "0"+days : days;
        hrs = (hrs<10) ? "0"+hrs : hrs;
        min = (min<10) ? "0"+min : min;
        sec = (sec<10) ? "0"+sec : sec;

        var elTimeCounter = $('.time-counter');
        var elDays = $('.days', elTimeCounter);
        var elHours = $('.hours', elTimeCounter);
        var elMinutes = $('.minutes', elTimeCounter);
        var elSeconds = $('.seconds', elTimeCounter);

        $('.num', elDays).html(days);
        $('.num', elHours).html(hrs);
        $('.num', elMinutes).html(min);
        $('.num', elSeconds).html(sec);

        setTimeout(countDown, 1000);
    }
    countDown();
});