var app = angular.module("myApp", ['ngRoute', 'ngResource', 'ui.bootstrap']);
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

app.controller("controller", ["$scope", "$resource",
    function ($scope, $resource) {
        var accountRes = $resource("/account")
        accountRes.get(
          function (data) {
              $scope.account = data;
              
          }, function (err) {
              if ($scope.account == null) {
                  ;
              }
          }
        );


        $scope.userImage = "/profilePicture"; //"images/Chrysanthemum.jpg";

        /*    logout     */
        var logout = $resource("/logout");
        $scope.btnLogout = function () {
            logout.get(function (data) {
                window.location = "/HomePage.html"; // Logout and move to home page
            })
        };
    }
]
)
