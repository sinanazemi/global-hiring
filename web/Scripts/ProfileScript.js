// new dependency: ngResource is included just above
var myapp = new angular.module("app", ["ngResource"]);

// inject the $resource dependency here
myapp.controller("controller",
  ["$scope", "$window", "$resource",
    function ($scope, $window, $resource) {

        var accountRes = $resource("/account")
        accountRes.get(
          function (data) {
              $scope.account = data
          }
        );

        $scope.getSkills = function () {
            $scope.skills = $scope.selectedService.skills;
        }; // function()

        $scope.dates = [];
        for (var i = 1970; i <= 2020; i++) {
            $scope.dates.push(i);
        }

        var degrees = $resource("/degrees")

        degrees.query(
          function (data) {
              $scope.degrees = data;
          } // function(data)
        ) // service.query

        
        var saveEduRes = $resource("/SaveEducation")
        $scope.addEducation = function () {
            var saveEdu = new saveEduRes();
            saveEdu.school = $scope.School;
            saveEdu.fromdate = $scope.FromDate;
            saveEdu.todate = $scope.ToDate;
            saveEdu.degree = $scope.SelectedDegree;
            saveEdu.field = $scope.Field;
            saveEdu.grade = $scope.Grade;

            saveEdu.$save(function (edu) {
                $scope.account.educations.push(edu);
                });
            }
        }

  ]
)