var globalHiring = angular.module("globalHiring", ["ngRoute","ngResource"]);
globalHiring.config(function ($routeProvider) {
    $routeProvider
    .when("/", {
        templateUrl: "home2.html"
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
    .when("/steps",{
        templateUrl : 'steps.html',
        controller  : 'stepsController'
    })
    .otherwise({redirectTo: '/'});

});
