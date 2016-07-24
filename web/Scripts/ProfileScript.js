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

        
        //$scope.months = [{ id: 1, name: "January" }, { id: 2, name: "February" }, { id: 3, name: "March" }, { id: 4, name: "April" }, { id: 5, name: "May" }, { id: 6, name: "June" }, { id: 6, name: "July" }
        //    , { id: 7, name: "August" }, { id: 8, name: "September" }, { id: 10, name: "September" }, { id: 11, name: "November" }, { id: 12, name: "December" }];
        $scope.months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "September", "November", "December"];

        var degrees = $resource("/degrees")

        degrees.query(
          function (data) {
              $scope.degrees = data;
          } // function(data)
        ) // service.query

        //**************************
        // Work History
        //*************************
        $scope.locations = ["location1", "location2", "location3"];
        //$scope.roles = ["Intern", "Individual Contributor", "Lead", "Manager", "Executive", "Owner"];
        $scope.roles = [{ value: "I", name: "Intern" }, { value: "C", name: "Individual Contributor" }, { value: "L", name: "Lead" }, { value: "M", name: "Manager" }, { value: "E", name: "Executive" }, {value:"O",name:"Owner"}];

        var saveWhRes = $resource("/saveWork")
        $scope.addWorkHistory = function () {
            var saveWh = new saveWhRes();
            saveWh.id = $scope.whId;
            saveWh.company = $scope.whCompany;
            saveWh.location = $scope.whLocation;
            saveWh.title = $scope.whTitle;
            saveWh.role = $scope.whRole; /*$scope.whRole;*/
            saveWh.frommonth =/*$scope.whFromMonth.id;*/$scope.months.indexOf($scope.whFromMonth)+1;
            saveWh.fromyear = $scope.whFromYear;
            saveWh.tomonth = /*$scope.whToMonth.id;*/$scope.months.indexOf($scope.whToMonth)+1;
            saveWh.toyear = $scope.whToYear;
            saveWh.currently = $scope.whCurrently;
            saveWh.description = $scope.whDesc;
            
            var index;
            $scope.account.works.some(function (elem, i) {
                return elem.id === $scope.whId ? (index = i, true) : false;
            });

            saveWh.$save(function (wh) {
                if (index > 0)
                    $scope.account.works[index] = wh;
                else
                    $scope.account.works.push(wh);
                
                cleanHistoryInputs();
            });
        }

        $scope.editWorkHistory = function (wh) {
            
            $scope.whId = wh.id;
            $scope.whCompany = wh.company;
            $scope.whLocation = wh.location;
            $scope.whTitle = wh.title;
            $scope.whRole = wh.role.value;
            $scope.whFromMonth = wh.frommonth.name;
            $scope.whFromYear = wh.fromyear;
            $scope.whToMonth = wh.tomonth.name;
            $scope.whToYear = wh.toyear;
            $scope.whCurrently = wh.currently;
            $scope.whDesc = wh.description;
        }

        var delWhRes = $resource("/deleteWork")
        $scope.deleteWorkHistory = function (wh) {
            var delWh = new delWhRes();
            delWh.id = wh.id;
            delWh.$save(function (dwr) {
                $scope.account.works.splice($scope.account.works.indexOf(dwr), 1);
                cleanHistoryInputs();
            });
        }
        
        function cleanHistoryInputs()
        {
            $scope.whId = "";
            $scope.whCompany = "";
            $scope.whLocation = "";
            $scope.whTitle = "";
            $scope.whRole = [];
            $scope.whFromMonth = [];
            $scope.whFromYear = "";
            $scope.whToMonth = [];
            $scope.whToYear = "";
            $scope.whCurrently = "";
            $scope.whDesc = "";
        }

        //*************************
        // education
        //*************************
        var saveEduRes = $resource("/saveEducation")
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