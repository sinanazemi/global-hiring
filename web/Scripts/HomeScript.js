var app = angular.module("myApp", ["ngRoute"]);
app.config(function ($routeProvider) {
    $routeProvider
    .when("/", {
        templateUrl: "Home2.html"
    })
    .when("/About", {
        templateUrl: "AboutUs.html"
    })
    .when("/FAQ", {
        templateUrl: "FAQ.html"
    })
    //.when("/profile", {
    //    templateUrl: "profile.html"
    //})

});
