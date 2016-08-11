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
        //var months = $resource("/month")
        //months.query(
        //  function (data) {
        //      $scope.months = data;
        //  } // function(data)
        //) // service.query

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

        //$scope.occupations = ["occupation 1", "occupation 2", "occupation 3"];
        var occupations = $resource("/occupations")
        occupations.query(
          function (data) {
              $scope.occupations = data;
          } // function(data)
        )
        //**************************
        // Work History
        //*************************
        $scope.locations = ["location1", "location2", "location3"];
        //$scope.roles = ["Intern", "Individual Contributor", "Lead", "Manager", "Executive", "Owner"];
        $scope.roles = [{ value: "I", name: "Intern" }, { value: "C", name: "Individual Contributor" }, { value: "L", name: "Lead" }, { value: "M", name: "Manager" }, { value: "E", name: "Executive" }, { value: "O", name: "Owner" }];

        var saveWhRes = $resource("/saveWork")
        $scope.saveWorkHistory = function () {
            if (saveWh())
                $('#addHistory').modal('hide');
        }
        $scope.saveWorkHistoryMore = function () {
            saveWh();
        }
        function saveWh()
        {
            if (checkWhValidation()) {                
                var saveWh = new saveWhRes();
                //if(checkWhValidation())
                saveWh.id = $scope.whId;
                saveWh.company = $scope.whCompany;
                saveWh.location = $scope.whLocation;
                saveWh.title = $scope.whTitle;
                saveWh.role = $scope.whRole; /*$scope.whRole;*/
                saveWh.frommonth = $scope.whFromMonth;/*$scope.months.indexOf($scope.whFromMonth)+1;*/
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
                    if (index >= 0)
                        $scope.account.works[index] = wh;
                    else
                        $scope.account.works.push(wh);

                });
                cleanHistoryInputs();
                return true;
            }
            return false;
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

        function cleanHistoryInputs() {
            //isValid = true;
            $scope.whForm.$setUntouched();
            $scope.vwhCompanyShow = false;
            $scope.vwhLocationShow = false;
            $scope.vwhTitleShow = false;
            $scope.vwhRoleShow = false;
            $scope.vwhFromMonthShow = false;
            $scope.vwhFromYearShow = false;
            $scope.vwhToMonthShow = false;
            $scope.vwhToYearShow = false;
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

        $scope.cleanWhInputs = function () {
            cleanHistoryInputs();
        }

        // ********** check validation ****************
        $scope.vwhCompanyShow = false;
        $scope.vwhLocationShow = false;
        $scope.vwhTitleShow = false;
        $scope.vwhRoleShow = false;
        $scope.vwhFromMonthShow = false;
        $scope.vwhFromYearShow = false;
        $scope.vwhToMonthShow = false;
        $scope.vwhToYearShow = false;

        function checkWhValidation() {
            var isValid = true;
            if ($scope.whCompany == "") {
                isValid = false;
                $scope.vwhCompanyShow = true;

            }
            if ($scope.whLocation == "") {
                isValid = false;
                $scope.vwhLocationShow = true;
            }
            if ($scope.whTitle == "") {
                isValid = false;
                $scope.vwhTitleShow = true;
            }
            if ($scope.whRole == "") {
                isValid = false;
                $scope.vwhRoleShow = true;
            }
            if ($scope.whFromMonth == "") {
                isValid = false;
                $scope.vwhFromMonthShow = true;
            }
            if ($scope.whFromYear == "") {
                isValid = false;
                $scope.vwhFromYearShow = true;
            }
            if (!$scope.whCurrently && ($scope.whToMonth == "")) {
                isValid = false;
                $scope.vwhToMonthShow = true;
            }
            if (!$scope.whCurrently &&($scope.whToYear == "")) {
                isValid = false;
                $scope.vwhToYearShow = true;
            }
            

            if (isValid)
                return true;
            else
                return false;
        }
        $scope.whCompanyChg = function () {
            $scope.vwhCompanyShow = false;
        }
        $scope.whLocationChg = function () {
            $scope.vwhLocationShow = false;
        }
        $scope.whTitleChg = function () {
            $scope.vwhTitleShow = false;
        }
        $scope.whRoleChg = function () {
            $scope.vwhRoleShow = false;
        }
        $scope.whFromMonthChg = function () {
            if ($scope.whFromMonth == "")
                $scope.vwhFromMonthShow = true;
            else
            $scope.vwhFromMonthShow = false;
        }
        $scope.whFromYearChg = function () {
            if ($scope.whFromYear == "")
                $scope.vwhFromYearShow = true;
            else
            $scope.vwhFromYearShow = false;
        }
        $scope.whToMonthChg = function () {
            if ($scope.whToMonth == "")
                $scope.vwhToMonthShow = true;
            else
            $scope.vwhToMonthShow = false;
        }
        $scope.whToYearChg = function () {
            if ($scope.whToYear == "")
                $scope.vwhToYearShow = true;
            else
            $scope.vwhToYearShow = false;
        }


        $scope.whCurrentlyChg = function () {
            if($scope.whCurrently)
            {
                $scope.vwhToMonthShow = false;
                $scope.vwhToYearShow = false;
                $scope.whToMonth = "";
                $scope.whToYear = "";
            }
        }

        // *********** for highlight and show the edit and delete buttons
        $scope.whMouseOver=function(context){
            context.popoverRemove = true;
            context.whHoverStyle = { 'background-color': '#b8e986' };
        }

        $scope.whMouseLeave = function (context) {
            context.popoverRemove = false;
            context.whHoverStyle = {};
        }
        //*************************
        // education
        //*************************
        var saveEduRes = $resource("/saveEducation")
        $scope.saveEducation = function () {
            if (saveEdu())
                $('#addEducation').modal('hide');
        }
        $scope.saveEducationMore = function () {
            saveEdu();
        }
        saveEdu = function () {
            if (checkEduValidation()) {  
                var saveEdu = new saveEduRes();
                saveEdu.id = $scope.eduId;
                saveEdu.school = $scope.eduSchool;
                saveEdu.fromdate = $scope.eduFromDate;
                saveEdu.todate = $scope.eduToDate;
                saveEdu.degree = $scope.eduDegree;
                saveEdu.field = $scope.eduField;
                saveEdu.grade = $scope.eduGrade;
                //saveEdu.desc = $scope.eduDesc;

                var index;
                $scope.account.educations.some(function (elem, i) {
                    return elem.id === $scope.eduId ? (index = i, true) : false;
                });

                saveEdu.$save(function (edu) {
                    if (index >= 0)
                        $scope.account.educations[index] = edu;
                    else
                        $scope.account.educations.push(edu);                    
                });
                cleanEducationInputs();
                return true;
            }
            return false;
        }

        $scope.editEducation = function (edu) {
            $scope.eduId = edu.id;
            $scope.eduSchool = edu.school;
            $scope.eduFromDate = edu.fromdate;
            $scope.eduToDate = edu.todate;
            $scope.eduDegree = edu.degree;
            $scope.eduField = edu.field;
            $scope.eduGrade = edu.grade;
            //$scpoe.eduDesc = edu.desc;
        }

        function cleanEducationInputs() {

            $scope.eduForm.$setUntouched();
            $scope.veduSchoolShow = false;
            $scope.veduFromDateShow = false;
            $scope.veduToDateShow = false;

            $scope.eduId = "";
            $scope.eduSchool = "";
            $scope.eduFromDate = "";
            $scope.eduToDate = "";
            $scope.eduDegree = [];
            $scope.eduField = "";
            $scope.eduGrade = "";
            $scope.eduDesc = "";
        }

        $scope.cleanEduInputs = function () {
            cleanEducationInputs();
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

        // ********** check validation ****************
        $scope.veduSchoolShow = false;
        $scope.veduFromDateShow = false;
        $scope.veduToDateShow = false;

        function checkEduValidation() {
            var isValid = true;
            if ($scope.eduSchool == "") {
                isValid = false;
                $scope.veduSchoolShow = true;
            }       
            if ($scope.eduFromDate == "") {
                isValid = false;
                $scope.veduFromDateShow = true;
            }
            if ($scope.eduToDate == "") {
                isValid = false;
                $scope.veduToDateShow = true;
            }

            if (isValid)
                return true;
            else
                return false;
            }

        $scope.eduSchoolChg = function () {
            $scope.veduSchoolShow = false;
        }        
        $scope.eduFromDateChg = function () {
            if ($scope.eduFromDate == "")
                $scope.veduFromDateShow = true;
            else
                $scope.veduFromDateShow = false;
        }
        $scope.eduToDateChg = function () {
            if ($scope.eduToDate == "")
                $scope.veduToDateShow = true;
            else
            $scope.veduToDateShow = false;
        }
        $scope.eduDegreeChg = function () {
            if ($scope.eduDegree == "")
                $scope.veduDegreeShow = true;
            else
                $scope.veduDegreeShow = false;
        }


        // ************* for highlight and show the edit and delete buttons
        $scope.eduMouseOver = function (context) {
            context.popoverRemove = true;
            context.eduHoverStyle = { 'background-color': '#b8e986' };
        }

        $scope.eduMouseLeave = function (context) {
            context.popoverRemove = false;
            context.eduHoverStyle = {};
        }

        //*************************
        // Volunteering 
        //*************************
        var saveVolRes = $resource("/saveVolunteering")
        $scope.saveVolunteering = function () {
            if (saveVol())
                $('#addVolunteering').modal('hide');
        }
        function saveVol() {
            var saveVol = new saveVolRes();
            if (checkVolValidation()) {
                saveVol.id = $scope.volId;
                saveVol.organization = $scope.volOrganization;
                saveVol.role = $scope.volRole;
                saveVol.cause = $scope.volCause;
                saveVol.frommonth = $scope.volFromMonth;
                saveVol.fromyear = $scope.volFromYear;
                saveVol.tomonth = $scope.volToMonth;
                saveVol.toyear = $scope.volToYear;
                saveVol.description = $scope.volDesc;


                var index;
                $scope.account.volunteerings.some(function (elem, i) {
                    return elem.id === $scope.volId ? (index = i, true) : false;
                });

                saveVol.$save(function (vol) {
                    if (index >= 0)
                        $scope.account.volunteerings[index] = vol;
                    else
                        $scope.account.volunteerings.push(vol);


                });
                cleanVolunteeringInputs();
                return true;
            }
            return false;
        }

        $scope.editVolunteering = function (vol) {
            $scope.volId = vol.id;
            $scope.volOrganization = vol.organization;
            $scope.volRole = vol.role;
            $scope.volCause = vol.cause;
            $scope.volFromMonth = vol.frommonth;
            $scope.volFromYear = vol.fromyear;
            $scope.volToMonth = vol.tomonth;
            $scope.volToYear = vol.toyear;
            $scope.volDesc = vol.description;
        }

        var delVolRes = $resource("/deleteVolunteering")
        $scope.deleteVolunteering = function (vol) {
            var delVol = new delVolRes();
            delVol.id = vol.id;
            delVol.$save(function (dv) {
                $scope.account.volunteerings.splice($scope.account.volunteerings.indexOf(dv), 1);
                //cleanHistoryInputs();
            });
        }

        function cleanVolunteeringInputs() {
            $scope.volForm.$setUntouched();
            $scope.vvolOrganizationShow = false;
            $scope.vvolRoleShow = false;
            $scope.vvolCauseShow = false;
            $scope.vvolFromMonthShow = false;
            $scope.vvolFromYearShow = false;

            $scope.volId = "";
            $scope.volOrganization = "";
            $scope.volRole = "";
            $scope.volCause = [];
            $scope.volFromMonth = [];
            $scope.volFromYear = "";
            $scope.volToMonth = [];
            $scope.volToYear = "";
            $scope.volDesc = "";
        }

        $scope.cleanVolInputs = function () {
            cleanVolunteeringInputs();
        }

        // ********** check validation ****************
        $scope.vvolOrganizationShow = false;
        $scope.vvolRoleShow = false;
        $scope.vvolCauseShow = false;
        $scope.vvolFromMonthShow = false;
        $scope.vvolFromYearShow = false;

        function checkVolValidation() {
            var isValid = true;
            if ($scope.volOrganization == "") {
                isValid = false;
                $scope.vvolOrganizationShow = true;
            }
            if ($scope.volRole == "") {
                isValid = false;
                $scope.vvolRoleShow = true;
            }
            if ($scope.volCause == "") {
                isValid = false;
                $scope.vvolCauseShow = true;
            }
            if ($scope.volFromMonth == "") {
                isValid = false;
                $scope.vvolFromMonthShow = true;
            }
            if ($scope.volFromYear == "") {
                isValid = false;
                $scope.vvolFromYearShow = true;
            }

            if (isValid)
                return true;
            else
                return false;
        }

        $scope.volOrganizationChg = function () {
            $scope.vvolOrganizationShow = false;
        }
        $scope.volRoleChg = function () {
            $scope.vvolRoleShow = false;
        }
        $scope.volCauseChg = function () {
            if ($scope.volCause == "")
                $scope.vvolCauseShow = true;
            else
                $scope.vvolCauseShow = false;
        }
        $scope.volFromMonthChg = function () {
            if ($scope.volFromMonth == "")
                $scope.vvolFromMonthShow = true;
            else
                $scope.vvolFromMonthShow = false;
        }
        $scope.volFromYearChg = function () {
            if ($scope.volFromYear == "")
                $scope.vvolFromYearShow = true;
            else
                $scope.vvolFromYearShow = false;
        }



        // for highlight and show the edit and delete buttons
        $scope.volMouseOver = function (context) {
            context.popoverRemove = true;
            context.volHoverStyle = { 'background-color': '#b8e986' };
        }

        $scope.volMouseLeave = function (context) {
            context.popoverRemove = false;
            context.volHoverStyle = {};
        }

        //*************************
        // Certifications 
        //*************************
        // **************** for highlight and show the edit and delete buttons
        $scope.crfMouseOver = function (context) {
            context.popoverRemove = true;
            context.crfHoverStyle = { 'background-color': '#b8e986' };
        }

        $scope.crfMouseLeave = function (context) {
            context.popoverRemove = false;
            context.crfHoverStyle = {};
        }

        //*************************
        // Test Scores 
        //*************************
        var saveTcRes = $resource("/saveTest")
        $scope.saveTestScores = function () {
            if (saveTc())
                $('#addScores').modal('hide');
        }
        $scope.saveTestScoresMore = function () {
            saveTc();
        }
        function saveTc() {
            if (checkTcValidation()) {
                var saveTc = new saveTcRes();
                saveTc.id = $scope.tcId;
                saveTc.name = $scope.tcName;
                saveTc.occupation = $scope.tcOccupation;
                saveTc.month = $scope.tcFromMonth;
                saveTc.year = $scope.tcFromYear;
                saveTc.score = $scope.tcScore;
                saveTc.description = $scope.tcDesc;

                var index;
                $scope.account.tests.some(function (elem, i) {
                    return elem.id === $scope.tcId ? (index = i, true) : false;
                });

                saveTc.$save(function (tc) {
                    if (index >= 0)
                        $scope.account.tests[index] = tc;
                    else
                        $scope.account.tests.push(tc);

                });
                cleanScoresInputs();
                return true;
            }
            return false;
        }

        $scope.editTestScore = function (tc) {

            $scope.tcId = tc.id;
            $scope.tcName = tc.name;
            $scope.tcOccupation = tc.occupation;
            $scope.tcFromMonth = tc.month;
            $scope.tcFromYear = tc.year;
            $scope.tcScore = tc.score;
            $scope.tcDesc = tc.description;
        }

        var delTcRes = $resource("/deleteTest")
        $scope.deleteTestScore = function (tc) {
            var delTc = new delTcRes();
            delTc.id = tc.id;
            delTc.$save(function (dtc) {
                $scope.account.tests.splice($scope.account.tests.indexOf(dtc), 1);
                //cleanHistoryInputs();
            });
        }

        function cleanScoresInputs() {
            //isValid = true;
            $scope.tcForm.$setUntouched();
            $scope.vtcNameShow = false;
            $scope.vtcOccupationShow = false;
            $scope.vtcFromMonthShow = false;
            $scope.vtcFromYearShow = false;
            $scope.vtcScoreShow = false;

            $scope.tcId = "";
            $scope.tcName = "";
            $scope.tcOccupation = [];
            $scope.tcFromMonth = [];
            $scope.tcFromYear = [];
            $scope.tcScore = "";
            $scope.tcDesc = "";
        }

        $scope.cleanTcInputs = function () {
            cleanScoresInputs();
        }

        // ********** check validation ****************
        $scope.vtcNameShow = false;
        $scope.vtcOccupationShow = false;
        $scope.vtcFromMonthShow = false;
        $scope.vtcFromYearShow = false;
        $scope.vtcScoreShow = false;

        function checkTcValidation() {
            var isValid = true;
            if ($scope.tcName == "") {
                isValid = false;
                $scope.vtcNameShow = true;
            }
            if ($scope.tcOccupation == "") {
                isValid = false;
                $scope.vtcOccupationShow = true;
            }
            if ($scope.tcFromMonth == "") {
                isValid = false;
                $scope.vtcFromMonthShow = true;
            }
            if ($scope.tcFromYear == "") {
                isValid = false;
                $scope.vtcFromYearShow = true;
            }
            if ($scope.tcScore == "") {
                isValid = false;
                $scope.vtcScoreShow = true;
            }


            if (isValid)
                return true;
            else
                return false;
        }
        $scope.tcNameChg = function () {
            $scope.vtcNameShow = false;
        }
        $scope.tcOccupationChg = function () {
            $scope.vtcOccupationShow = false;
        }        
        $scope.tcFromMonthChg = function () {
            if ($scope.tcFromMonth == "")
                $scope.vtcFromMonthShow = true;
            else
                $scope.vtcFromMonthShow = false;
        }
        $scope.tcFromYearChg = function () {
            if ($scope.tcFromYear == "")
                $scope.vtcFromYearShow = true;
            else
                $scope.vtcFromYearShow = false;
        }
        $scope.tcScoreChg = function () {
            $scope.vtcScoreShow = false;
        }

        // *********** for highlight and show the edit and delete buttons
        $scope.tcMouseOver = function (context) {
            context.popoverRemove = true;
            context.tcHoverStyle = { 'background-color': '#b8e986' };
        }

        $scope.tcMouseLeave = function (context) {
            context.popoverRemove = false;
            context.tcHoverStyle = {};
        }


        //*************************
        // Projects 
        //*************************
        var savePrjRes = $resource("/saveProject")
        $scope.saveProjects = function () {
            if (SavePrj())
                $('#addProjects').modal('hide');
        }
        $scope.saveProjectsMore = function () {
            SavePrj();
        }
        function SavePrj() {
            if (checkPrjValidation()) {
                var savePrj = new savePrjRes();
                savePrj.id = $scope.prjId;
                savePrj.name = $scope.prjName;
                savePrj.occupation = $scope.prjOccupation;
                savePrj.month = $scope.prjMonth;
                savePrj.year = $scope.prjYear;
                savePrj.url = $scope.prjUrl;
                savePrj.description = $scope.prjDesc;

                var index;
                $scope.account.projects.some(function (elem, i) {
                    return elem.id === $scope.prjId ? (index = i, true) : false;
                });

                savePrj.$save(function (prj) {
                    if (index >= 0)
                        $scope.account.projects[index] = prj;
                    else
                        $scope.account.projects.push(prj);

                });
                cleanProjectsInputs();
                return true;
            }
            return false;
        }

        $scope.editProject = function (prj) {

            $scope.prjId = prj.id;
            $scope.prjName = prj.name;
            $scope.prjOccupation = prj.occupation;
            $scope.prjMonth = prj.month;
            $scope.prjYear = prj.year;
            $scope.prjUrl = prj.url;
            $scope.prjDesc = prj.description;
        }

        var delPrjRes = $resource("/deleteProject")
        $scope.deleteProject = function (prj) {
            var delPrj = new delPrjRes();
            delPrj.id = prj.id;
            delPrj.$save(function (dprj) {
                $scope.account.projects.splice($scope.account.projects.indexOf(dprj), 1);
                //cleanHistoryInputs();
            });
        }

        function cleanProjectsInputs() {
            //isValid = true;
            $scope.prjForm.$setUntouched();
            $scope.vprjNameShow = false;
            $scope.vprjOccupationShow = false;

            $scope.prjId = "";
            $scope.prjName = "";
            $scope.prjOccupation = [];
            $scope.prjMonth = [];
            $scope.prjYear = [];
            $scope.prjUrl = "";
            $scope.prjDesc = "";
        }

        $scope.cleanPrjInputs = function () {
            cleanScoresInputs();
        }

        // ********** check validation ****************
        $scope.vprjNameShow = false;
        $scope.vprjOccupationShow = false;

        function checkPrjValidation() {
            var isValid = true;
            if ($scope.prjName == "") {
                isValid = false;
                $scope.vprjNameShow = true;
            }
            if ($scope.prjOccupation == "") {
                isValid = false;
                $scope.vprjOccupationShow = true;
            }

            if (isValid)
                return true;
            else
                return false;
        }
        $scope.prjNameChg = function () {
            $scope.vprjNameShow = false;
        }
        $scope.prjOccupationChg = function () {
            $scope.vprjOccupationShow = false;
        }
       

        // *********** for highlight and show the edit and delete buttons
        $scope.prjMouseOver = function (context) {
            context.popoverRemove = true;
            context.prjHoverStyle = { 'background-color': '#b8e986' };
        }

        $scope.prjMouseLeave = function (context) {
            context.popoverRemove = false;
            context.prjHoverStyle = {};
        }


        //*************************
        //  Honors & Awards 
        //*************************
        var saveHaRes = $resource("/saveHonor")
        $scope.saveAwards = function () {
            if (SaveHa())
                $('#addAwards').modal('hide');
        }
        $scope.saveAwardsMore = function () {
            SaveHa();
        }
        function SaveHa() {
            if (checkHaValidation()) {
                var saveHa = new saveHaRes();
                saveHa.id = $scope.haId;
                saveHa.title = $scope.haTitle;
                saveHa.occupation = $scope.haOccupation;
                saveHa.month = $scope.haMonth;
                saveHa.year = $scope.haYear;
                saveHa.description = $scope.haDesc;

                var index;
                $scope.account.honors.some(function (elem, i) {
                    return elem.id === $scope.haId ? (index = i, true) : false;
                });

                saveHa.$save(function (ha) {
                    if (index >= 0)
                        $scope.account.honors[index] = ha;
                    else
                        $scope.account.honors.push(ha);

                });
                cleanAwardsInputs();
                return true;
            }
            return false;
        }

        $scope.editAward = function (ha) {

            $scope.haId = ha.id;
            $scope.haTitle = ha.title;
            $scope.haOccupation = ha.occupation;
            $scope.haMonth = ha.month;
            $scope.haYear = ha.year;
            $scope.haDesc = ha.description;
        }

        var delHaRes = $resource("/deleteHonor")
        $scope.deleteAward = function (ha) {
            var delHa = new delHaRes();
            delHa.id = ha.id;
            delHa.$save(function (dha) {
                $scope.account.honors.splice($scope.account.honors.indexOf(dha), 1);
                //cleanHistoryInputs();
            });
        }

        function cleanAwardsInputs() {
            //isValid = true;
            $scope.haForm.$setUntouched();
            $scope.vhaTitleShow = false;
            $scope.vhaOccupationShow = false;
            $scope.vhaMonthShow = false;
            $scope.vhaYearShow = false;

            $scope.haId = "";
            $scope.haTitle = "";
            $scope.haOccupation = [];
            $scope.haMonth = [];
            $scope.haYear = [];
            $scope.haDesc = "";
        }

        $scope.cleanHaInputs = function () {
            cleanAwardsInputs();
        }

        // ********** check validation ****************
        $scope.vhaTitleShow = false;
        $scope.vhaOccupationShow = false;
        $scope.vhaMonthShow = false;
        $scope.vhaYearShow = false;

        function checkHaValidation() {
            var isValid = true;
            if ($scope.haTitle == "") {
                isValid = false;
                $scope.vhaTitleShow = true;
            }
            if ($scope.haOccupation == "") {
                isValid = false;
                $scope.vhaOccupationShow = true;
            }
            if ($scope.haMonth == "") {
                isValid = false;
                $scope.vhaMonthShow = true;
            }
            if ($scope.haYear == "") {
                isValid = false;
                $scope.vhaYearShow = true;
            }

            if (isValid)
                return true;
            else
                return false;
        }
        $scope.haTitleChg = function () {
            $scope.vhaTitleShow = false;
        }
        $scope.haOccupationChg = function () {
            $scope.vhaOccupationShow = false;
        }
        $scope.haMonthChg = function () {
            if ($scope.haMonth == "")
                $scope.vhaMonthShow = true;
            else
                $scope.vhaMonthShow = false;
        }
        $scope.haYearChg = function () {
            if ($scope.haYear == "")
                $scope.vhaYearShow = true;
            else
                $scope.vhaYearShow = false;
        }

        // *********** for highlight and show the edit and delete buttons
        $scope.haMouseOver = function (context) {
            context.popoverRemove = true;
            context.haHoverStyle = { 'background-color': '#b8e986' };
        }

        $scope.haMouseLeave = function (context) {
            context.popoverRemove = false;
            context.haHoverStyle = {};
        }


        //*************************
        //  Courses 
        //*************************
        var saveCrRes = $resource("/saveCourse")
        $scope.saveCourses = function () {
            if (SaveCr())
                $('#addCourses').modal('hide');
        }
        $scope.saveCoursesMore = function () {
            SaveCr();
        }
        function SaveCr() {
            if (checkCrValidation()) {
                var saveCr = new saveCrRes();
                saveCr.id = $scope.crId;
                saveCr.name = $scope.crName;
                saveCr.occupation = $scope.crOccupation;
                saveCr.number = $scope.crNumber;
                saveCr.description = $scope.crDesc;

                var index;
                $scope.account.courses.some(function (elem, i) {
                    return elem.id === $scope.crId ? (index = i, true) : false;
                });

                saveCr.$save(function (cr) {
                    if (index >= 0)
                        $scope.account.courses[index] = cr;
                    else
                        $scope.account.courses.push(cr);

                });
                cleanCoursesInputs();
                return true;
            }
            return false;
        }

        $scope.editCourse = function (cr) {

            $scope.crId = cr.id;
            $scope.crName = cr.name;
            $scope.crOccupation = cr.occupation;
            $scope.crNumber = cr.number;
            $scope.crDesc = cr.description;
        }

        var delCrRes = $resource("/deleteCourse")
        $scope.deleteCourse = function (cr) {
            var delCr = new delCrRes();
            delCr.id = cr.id;
            delCr.$save(function (dcr) {
                $scope.account.courses.splice($scope.account.courses.indexOf(dcr), 1);
                //cleanHistoryInputs();
            });
        }

        function cleanCoursesInputs() {
            //isValid = true;
            $scope.crForm.$setUntouched();
            $scope.vcrNameShow = false;
            $scope.vcrOccupationShow = false;
            $scope.vcrNumberShow = false;

            $scope.crId = "";
            $scope.crName = "";
            $scope.crOccupation = [];
            $scope.crNumber = "";
            $scope.crDesc = "";
        }

        $scope.cleanCrInputs = function () {
            cleanCoursesInputs();
        }

        // ********** check validation ****************
        $scope.vcrNameShow = false;
        $scope.vcrOccupationShow = false;
        $scope.vcrNumberShow = false;

        function checkCrValidation() {
            var isValid = true;
            if ($scope.crName == "") {
                isValid = false;
                $scope.vcrNameShow = true;
            }
            if ($scope.crOccupation == "") {
                isValid = false;
                $scope.vcrOccupationShow = true;
            }
            if ($scope.crNumber == "") {
                isValid = false;
                $scope.vcrNumberShow = true;
            }
            if (isValid)
                return true;
            else
                return false;
        }
        $scope.crNameChg = function () {
            $scope.vcrNameShow = false;
        }
        $scope.crOccupationChg = function () {
            $scope.vcrOccupationShow = false;
        }
        $scope.crNumberChg = function () {
            $scope.vcrNumberShow = false;
        }

        // *********** for highlight and show the edit and delete buttons
        $scope.crMouseOver = function (context) {
            context.popoverRemove = true;
            context.crHoverStyle = { 'background-color': '#b8e986' };
        }

        $scope.crMouseLeave = function (context) {
            context.popoverRemove = false;
            context.crHoverStyle = {};
        }

    }

  ]
)