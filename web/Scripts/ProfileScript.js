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

        
        $scope.months = [{ value: 1, name: "January" }, { value: 2, name: "February" }, { value: 3, name: "March" }, { value: 4, name: "April" }, { value: 5, name: "May" }, { value: 6, name: "June" }, { value: 7, name: "July" }
            , { value: 8, name: "August" }, { value: 9, name: "September" }, { value: 10, name: "October" }, { value: 11, name: "November" }, { value: 12, name: "December" }];
        //$scope.months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];

        var degrees = $resource("/degrees")

        degrees.query(
          function (data) {
              $scope.degrees = data;
          } // function(data)
        ) // service.query

        var volCauses = $resource("/volunteeringCauses")
        volCauses.query(
         function (data) {
             $scope.volCauses = data;
         } // function(data)
       ) // service.query
        //**************************
        // Work History
        //*************************
        $scope.locations = ["location1", "location2", "location3"];
        //$scope.roles = ["Intern", "Individual Contributor", "Lead", "Manager", "Executive", "Owner"];
        $scope.roles = [{ value: "I", name: "Intern" }, { value: "C", name: "Individual Contributor" }, { value: "L", name: "Lead" }, { value: "M", name: "Manager" }, { value: "E", name: "Executive" }, {value:"O",name:"Owner"}];

        var saveWhRes = $resource("/saveWork")
        $scope.saveWorkHistory = function () {
            var saveWh = new saveWhRes();
            //if(checkWhValidation())
            saveWh.id = $scope.whId;
            saveWh.company = $scope.whCompany;
            saveWh.location = $scope.whLocation;
            saveWh.title = $scope.whTitle;
            saveWh.role = $scope.whRole; /*$scope.whRole;*/
            saveWh.frommonth =$scope.whFromMonth;/*$scope.months.indexOf($scope.whFromMonth)+1;*/
            saveWh.fromyear = $scope.whFromYear;
            saveWh.tomonth = $scope.whToMonth;/*$scope.months.indexOf($scope.whToMonth)+1;*/
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
            $scope.whFromMonth = wh.frommonth;
            $scope.whFromYear = wh.fromyear;
            $scope.whToMonth = wh.tomonth;//wh.tomonth.name;
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
                //cleanHistoryInputs();
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

        //function checkWhValidation()
        //{
        //    if($scope.whCompany='')
        //    {
        //        $scope.vwhCompany.
        //    }
        //}

        //*************************
        // education
        //*************************
        var saveEduRes = $resource("/saveEducation")
        $scope.saveEducation = function () {
            var saveEdu = new saveEduRes();
            saveEdu.id = $scope.eduId;
            saveEdu.school = $scope.eduSchool;
            saveEdu.fromdate = $scope.eduFromDate;
            saveEdu.todate = $scope.eduToDate;
            saveEdu.degree= $scope.eduDegree;
            saveEdu.field = $scope.eduField;
            saveEdu.grade = $scope.eduGrade;

            var index;
            $scope.account.educations.some(function (elem, i) {
                return elem.id === $scope.eduId ? (index = i, true) : false;
            });

            saveEdu.$save(function (edu) {
                if (index > 0)
                    $scope.account.educations[index] = edu;
                else
                    $scope.account.educations.push(edu);

                cleanEducationInputs();
            });
        }

        $scope.editEducation = function (edu)
        {
            $scope.eduId = edu.id;
            $scope.eduSchool = edu.school;
            $scope.eduFromDate = edu.fromdate;
            $scope.eduToDate = edu.todate;
            $scope.eduDegree = edu.degree;
            $scope.eduField = edu.field;
            $scope.eduGrade = edu.grade;
        }

        function cleanEducationInputs() {
            $scope.eduId = "";
            $scope.eduSchool = "";
            $scope.eduFromDate = "";
            $scope.eduToDate = "";
            $scope.eduDegree = [];
            $scope.eduField = "";
            $scope.eduGrade = "";
        }

        var delEduRes = $resource("/deleteEducation")
        $scope.deleteEducation = function (edu) {
            var delEdu = new delEduRes();
            delEdu.id = edu.id;
            delEdu.$save(function (dEdu) {
                $scope.account.educations.splice($scope.account.educations.indexOf(dEdu), 1);
                //cleanEducationInputs();
            });
        }

        //*************************
        // Volunteering 
        //*************************
        var saveVolRes = $resource("/saveVolunteering")
        $scope.saveVolunteering = function () {
            var saveVol = new saveVolRes();
            saveVol.id = $scope.whId;
            saveVol.organization = $scope.volOrganization;
            saveVol.role = $scope.volRole;
            saveVol.cause = $scope.volCause;
            saveVol.frommonth = $scope.months.indexOf($scope.volFromMonth) + 1;
            saveVol.fromyear = $scope.volFromYear;
            saveVol.tomonth = $scope.months.indexOf($scope.volToMonth) + 1;
            saveVol.toyear = $scope.volToYear;
            saveVol.description = $scope.volDesc;
            

            //var index;
            //$scope.account.works.some(function (elem, i) {
            //    return elem.id === $scope.whId ? (index = i, true) : false;
            //});

            //saveWh.$save(function (wh) {
            //    if (index > 0)
            //        $scope.account.works[index] = wh;
            //    else
            //        $scope.account.works.push(wh);

            //    cleanHistoryInputs();
            //});
        }

        //$scope.editWorkHistory = function (wh) {

        //    $scope.whId = wh.id;
        //    $scope.whCompany = wh.company;
        //    $scope.whLocation = wh.location;
        //    $scope.whTitle = wh.title;
        //    $scope.whRole = wh.role.value;
        //    $scope.whFromMonth = wh.frommonth.name;
        //    $scope.whFromYear = wh.fromyear;
        //    $scope.whToMonth = wh.tomonth.name;
        //    $scope.whToYear = wh.toyear;
        //    $scope.whCurrently = wh.currently;
        //    $scope.whDesc = wh.description;
        //}

        //var delWhRes = $resource("/deleteWork")
        //$scope.deleteWorkHistory = function (wh) {
        //    var delWh = new delWhRes();
        //    delWh.id = wh.id;
        //    delWh.$save(function (dwr) {
        //        $scope.account.works.splice($scope.account.works.indexOf(dwr), 1);
        //        //cleanHistoryInputs();
        //    });
        //}

        //function cleanHistoryInputs() {
        //    $scope.whId = "";
        //    $scope.whCompany = "";
        //    $scope.whLocation = "";
        //    $scope.whTitle = "";
        //    $scope.whRole = [];
        //    $scope.whFromMonth = [];
        //    $scope.whFromYear = "";
        //    $scope.whToMonth = [];
        //    $scope.whToYear = "";
        //    $scope.whCurrently = "";
        //    $scope.whDesc = "";
        //}

        }

  ]
)