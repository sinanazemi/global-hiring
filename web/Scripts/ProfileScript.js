// new dependency: ngResource is included just above
var myapp = new angular.module("app", ['ngAnimate', 'ui.bootstrap', 'selectize', 'ngResource', 'ngFileUpload', 'ngImgCrop']);

// inject the $resource dependency here
myapp.controller("controller",
  ["$scope", "$window", "$resource", "$document",
    function ($scope, $window, $resource, $document) {


        $scope.shAddOverview = true;

        var accountRes = $resource("/account")
        accountRes.get(
          function (data) {
              $scope.account = data;
              if ($scope.account == null) {
                  window.location = "/HomePage.html";
              }
              if ($scope.account.description != "")
                  $scope.shAddOverview = false;
          }, function (err) {
              if ($scope.account == null) {
                  window.location = "/HomePage.html";
              }
          }
        );


        $scope.userImage = "/profilePicture"; //"images/Chrysanthemum.jpg";

        var profilePictureRes = $resource("/saveProfilePicture")
        $scope.savePrfPic = function () {

          var profilePicture = new profilePictureRes();
          profilePicture.profilepicture = $scope.prfImgUrl;

          profilePicture.$save(
            function (data) {
              $scope.userImage = $scope.prfImgUrl;
            } // function(data)
          ); // profilePicture.$save
        } // $scope.savePrfPic = function ()

        
            var prfStrengthNum = 0;
            var prfStrengthRes = $resource("/accountStrength");
            //prfStrengthRes.query(
            prfStrengthRes.get(
               function (data) {
                   prfStrengthNum = data.value;
                   if (prfStrengthNum < 20) {
                       $scope.prfStrengthImg = "/images/beginner.png";
                   }
                   else if (prfStrengthNum < 40)
                       $scope.prfStrengthImg = "/images/Intermediate.png";
                   else if (prfStrengthNum < 60)
                       $scope.prfStrengthImg = "/images/Advanced.png";
                   else if (prfStrengthNum < 80)
                       $scope.prfStrengthImg = "/images/Expert.png";
                   else
                       $scope.prfStrengthImg = "/images/AllStar.png";
               }, function (err) {
                   prfStrengthNum = 0;
               })
       
       

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
        function setDegrees() {
        degrees.query(
          function (data) {
              $scope.degrees = data;
          } // function(data)
        ) // service.query
        }

        var volCauses = $resource("/volunteeringCauses")
        function setVolCauses() {
            volCauses.query(
                 function (data) {
                 $scope.volCauses = data;
                 } // function(data)
            ) // service.query
        }

        var occupations = $resource("/occupations")
        function setOccupations() {
            occupations.query(
              function (data) {
                  $scope.occupations = data;
              } // function(data)
            )
        }

        function checkInvDate(fMonth, fYear, tMonth, tYear) {
            if (tYear < fYear)
                return true;
            if (tYear == fYear)
                if (tMonth < fMonth)
                    return true;

            return false;
        }

        /*    logout     */
        var logout = $resource("/logout");
        $scope.btnLogout = function () {
            logout.get(function (data) {
                window.location = "/HomePage.html"; // Logout and move to home page
            })
        };

        //********************** Auto Suggestion ************************
        //************************************
        $scope.nameConfig = {
            valueField: 'name',
            labelField: 'name',
            searchField: 'name',
            maxItems: 1,
            maxOptions: 5,
        };

        $scope.nameConfig2 = {
            valueField: 'name',
            labelField: 'name',
            searchField: 'name',
            maxItems: 1,
            maxOptions: 5,
            create: 'true',
        };


        $scope.idConfig = {
            valueField: 'id',
            labelField: 'name',
            searchField: 'name',
            maxItems: 1,
            maxOptions: 5,
        };

        $scope.valueConfig = {
            valueField: 'value',
            labelField: 'name',
            searchField: 'name',
            maxItems: 1,
            maxOptions: 5,
        };


        //******************************
        // Job title
        //******************************
        //$scope.jtTitle = "Experienced Web and Mobile Developer";

        $scope.editJobtitle = function () {
            $scope.jtTitle = $scope.account.jobtitle;
        }

        var saveJtRes = $resource("/saveJobTitle")
        $scope.saveJobTitle = function()
        {
            var saveJt = new saveJtRes();
            saveJt.jobtitle = $scope.jtTitle;
            saveJt.$save();
            $('#addJobtitle').modal('hide');
            $scope.account.jobtitle = $scope.jtTitle;
        }

        // ************* for highlight and show the edit and delete buttons
        $scope.jtMouseOver = function (context) {
            context.jtPopoverRemove = true;
            context.jtHoverStyle = { 'background-color': '#b8e986'};
        }

        $scope.jtMouseLeave = function (context) {
            context.jtPopoverRemove = false;
            context.jtHoverStyle = {};
        }



        //**********************************
        // Add Skill
        //**********************************
        // Main Services
        $scope.toggleSelectedSRV = function (service) {
            service.isselected = !service.isselected;
        };

        var service = $resource("/mainServices")
        $scope.getMainServices = function () {
            service.query(
              function (data) {
                  $scope.mainServices = data;
                  for(var i=0;i<$scope.account.skills.length;i++)
                  {
                      for(var j=0;j<$scope.mainServices.length;j++)
                      {
                          if ($scope.account.skills[i].mainserviceid == $scope.mainServices[j].id) {
                              $scope.mainServices[j].isselected = true;
                              for (var z = 0; z < $scope.mainServices[j].skills.length; z++)
                              {
                                  if ($scope.mainServices[j].skills[z].id == $scope.account.skills[i].id) {
                                      $scope.mainServices[j].skills[z].isselected = true;
                                      break;
                                  }
                              }
                              break;
                          }
                      }
                  }
              } // function(data)
            ) // service.query
        };

        $scope.getIcon = function (serviceSelect) {
            if (serviceSelect.isselected) return serviceSelect.selectimageurl;
            else return serviceSelect.unselectimageurl;
        }

        // Skill
        $scope.skillClick = function (selectedSkill) {
            selectedSkill.isselected = !selectedSkill.isselected;
            if (selectedSkill.isselected) {
                $scope.selSkill = selectedSkill;
                $("#addSkillPrf").modal("show");
            }
            else
            {
                selectedSkill.profeciency = "";
            }
        }


        $scope.skillProfeciencies = [
           { text: "Student/Fresh Graduate", value: "S" },
           { text: "Junior Professional", value: "J" },
           { text: "Experienced Professional", value: "E" },
           { text: "Manager", value: "M" }
        ]


        var saveSkillRes = $resource("/saveSkill")
        $scope.saveSkill=function()
        {
            var saveSk = new saveSkillRes();
            for (var i = 0; i < $scope.mainServices.length; i++)
            {
                if($scope.mainServices[i].isselected)
                {
                    for(var j=0; j<$scope.mainServices[i].skills.length;j++)
                    {
                        if ($scope.mainServices[i].skills[j].isselected) {
                            saveSk.skillid = $scope.mainServices[i].skills[j].id;
                            saveSk.profeciency = $scope.mainServices[i].skills[j].profeciency;
                            saveSk.$save();
                            $('#addSkill').modal('hide');
                        }
                    }
                }

            }
        }

        var delSkillRes = $resource("/deleteSkill")
        $scope.removeSkill = function (skill) {
            var delSkill = new delSkillRes();
            delSkill.id = skill.id;
            delSkill.$save(function (dsk) {
                $scope.account.skills.splice($scope.account.skills.indexOf(dsk), 1);
                //cleanHistoryInputs();
            });
        }

        $scope.getSkills = function () {
            $scope.skills = $scope.selectedService.skills;
        }; // function()

        //******************************
        // Overview
        //******************************
        //$scope.jtTitle = "Experienced Web and Mobile Developer";

        $scope.editOverview = function () {
            $scope.ovOverview = $scope.account.description;
        }

        var saveOvRes = $resource("/saveDescription")
        $scope.saveOverview = function () {
            var saveOv = new saveOvRes();
            saveOv.description = $scope.ovOverview;
            saveOv.$save();
            $('#addOverview').modal('hide');
            $scope.account.description = $scope.ovOverview;
            $scope.shAddOverview = false;
        }

        // ************* for highlight and show the edit and delete buttons
        $scope.ovMouseOver = function (context) {
            if ($scope.account.description != "")
                context.ovPopoverRemove = true;
            else
                context.ovAddSign = true;
            context.ovHoverStyle = { 'background-color': '#b8e986' };
        }

        $scope.ovMouseLeave = function (context) {
            if ($scope.account.description != "")
                context.ovPopoverRemove = false;
            else
                context.ovAddSign = false;
            context.ovHoverStyle = {};
        }

        //**************************
        // language
        //**************************
        $scope.lblLgName = false;
        $scope.lgInvalid = false;
        $scope.lgPrfInvalid = false;

        $scope.addLanguage = function () {
            cleanLanguageInputs();
        }
        var saveLgRes = $resource("/saveLanguage")
        function saveLg() {
            if (checkLgValidation()) {
                if (!checkLgExist()) {
                    var saveLg = new saveLgRes();
                    saveLg.id = $scope.lgId;
                    saveLg.name = $scope.lgName;
                    saveLg.profeciency = $scope.lgProfeciency;
                    $scope.lblLgName = false;
                    var index;
                    $scope.account.languages.some(function (elem, i) {
                        return elem.id === $scope.lgId ? (index = i, true) : false;
                    });
                    saveLg.$save(function (lg) {
                        if (index >= 0)
                            $scope.account.languages[index] = lg;
                        else
                            $scope.account.languages.push(lg);

                    });
                    return true;
                }
                else
                    $scope.lgInvalid = true;
                return true;
            }
            else
                return false;
        }

        $scope.saveLanguage = function () {
            if (saveLg())
                $('#addLanguage').modal('hide');
        }
        $scope.saveLanguageMore = function () {
            saveLg();
            cleanLanguageInputs();
        }
        $scope.editLanguage = function (lg) {
            $scope.lgId = lg.id;
            $scope.lgName = lg.name;
            $scope.lgProfeciency = lg.profeciency.value;
            $scope.lblLgName = true;

        }

        var delLgRes = $resource("/deleteLanguage")
        $scope.deleteLanguage = function (lg) {
            var delLg = new delLgRes();
            delLg.id = lg.id;
            delLg.$save(function (dlg) {
                $scope.account.languages.splice($scope.account.languages.indexOf(dlg), 1);
                //cleanHistoryInputs();
            });
        }
        function checkLgValidation() {
            var isValid = true;

            if ($scope.lgProfeciency == "") {
                isValid = false;
                $scope.lgPrfInvalid = true;
            }

            if ($scope.lgName == "" || $scope.lgName == null) {
                isValid = false;
                $scope.vlgNameShow = true;
            }

            if (isValid)
                return true;
            else
                return false;
        }

        function cleanLanguageInputs() {
            $scope.lgId = '';
            $scope.lgName = '';
            $scope.lgProfeciency = '';
            $scope.lblLgName = false;
        }

        function checkLgExist() {
            var data = $scope.account.languages;
            for (var i=0; i < data.length; i++) {
                if (angular.lowercase(data[i].name) == angular.lowercase($scope.lgName))
                    return true;
            }

            return false;
        }

        $scope.lgNameChg = function () {
            if ($scope.lgName != "")
                $scope.vlgNameShow = false;
            if (checkLgExist())
                $scope.lgInvalid = true;
            else
                $scope.lgInvalid = false;
        }
        $scope.lgPrfChg = function()
        {
            if ($scope.lgProfeciency == "")
                $scope.lgPrfInvalid = true;
            else
                $scope.lgPrfInvalid = false;

        }
        // ************* for highlight and show the edit and delete buttons
        $scope.lgMouseOver = function (context) {
            context.popoverRemove = true;
            context.lgHoverStyle = { 'background-color': '#b8e986' };
        }

        $scope.lgMouseLeave = function (context) {
            context.popoverRemove = false;
            context.lgHoverStyle = {};
        }



        //**************************
        // Work History
        //*************************
        function setLocation() {
            var cities = $resource("/cities");
            cities.query(
          function (data) {
              $scope.locations = data;
          } // function(data)
        ) // service.query
        //$scope.locations = [{ id: 1, name: "location 1" }, { id: 2, name: "location 2" }, { id: 3, name: "location 3" }, { id: 4, name: "location 11" }, { id: 5, name: "location 12" }, { id: 6, name: "location 13" }
        //, { id: 7, name: "location 21" }, { id: 8, name: "location 22" }, { id: 9, name: "location 23" }, { id: 10, name: "location 31" }, { id: 11, name: "location 32" }, { id: 12, name: "location 33" }];
        }
        function setRoles() {
            $scope.roles = [{ value: "I", name: "Intern" }, { value: "C", name: "Individual Contributor" }, { value: "L", name: "Lead" }, { value: "M", name: "Manager" }, { value: "E", name: "Executive" }, { value: "O", name: "Owner" }];
        }

        $scope.addWorkHistory = function () {
            if ($scope.locations == null || $scope.locations.length==0) {
                setLocation();
            }
            if ($scope.roles == null || $scope.roles.length==0) {
                setRoles();
            }
            cleanHistoryInputs();
        }

        var saveWhRes = $resource("/saveWork")
        $scope.saveWorkHistory = function () {
            if (saveWh())
                $('#addHistory').modal('hide');
        }
        $scope.saveWorkHistoryMore = function () {
            saveWh();
        }
        function saveWh() {
            if (checkWhValidation()) {
                var saveWh = new saveWhRes();
                //if(checkWhValidation())
                saveWh.id = $scope.whId;
                saveWh.company = $scope.whCompany;
                saveWh.location = $scope.whLocation;
                saveWh.title = $scope.whTitle;
                saveWh.role = $scope.whRole;
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
            if ($scope.locations == null || $scope.locations.length==0) {
                setLocation();
            }
            if ($scope.roles == null || $scope.roles.length==0) {
                setRoles();
            }
            $scope.whId = wh.id;
            $scope.whCompany = wh.company;
            $scope.whLocation = wh.location;
            $scope.whTitle = wh.title;
            $scope.whRole = wh.role.value;
            $scope.whRoleValue = wh.role.value;
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
            if ($scope.whLocation == "" || $scope.whLocation == null) {
                isValid = false;
                $scope.vwhLocationShow = true;
            }
            if ($scope.whTitle == "") {
                isValid = false;
                $scope.vwhTitleShow = true;
            }
            if ($scope.whRole == "" || $scope.whRole == null) {
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
            if (!$scope.whCurrently && ($scope.whToYear == "")) {
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
        $scope.whLocationChg = function (data) {
            $scope.vwhLocationShow = false;     // for validation
        }
        $scope.whTitleChg = function () {
            $scope.vwhTitleShow = false;
        }
        $scope.whRoleChg = function (data) {
            $scope.vwhRoleShow = false;             // for validation
        }
        $scope.whFromMonthChg = function () {
            if ($scope.whFromMonth == "")
                $scope.vwhFromMonthShow = true;
            else {
                $scope.vwhFromMonthShow = false;
                if (!$scope.whCurrently && $scope.whToYear != "" && $scope.whToMonth != "")
                    $scope.whInvalidDate = checkInvDate($scope.whFromMonth.value, $scope.whFromYear, $scope.whToMonth.value, $scope.whToYear);
            }
        }
        $scope.whFromYearChg = function () {
            if ($scope.whFromYear == "")
                $scope.vwhFromYearShow = true;
            else {
                $scope.vwhFromYearShow = false;
                if (!$scope.whCurrently && $scope.whToMonth != "" && $scope.whToYear != "")
                    $scope.whInvalidDate = checkInvDate($scope.whFromMonth.value, $scope.whFromYear, $scope.whToMonth.value, $scope.whToYear);

            }
        }
        $scope.whToMonthChg = function () {
            if ($scope.whToMonth == "")
                $scope.vwhToMonthShow = true;
            else {
                $scope.vwhToMonthShow = false;
                if ($scope.whToYear != "")
                    $scope.whInvalidDate = checkInvDate($scope.whFromMonth.value, $scope.whFromYear, $scope.whToMonth.value, $scope.whToYear);
            }
        }
        $scope.whToYearChg = function () {
            if ($scope.whToYear == "")
                $scope.vwhToYearShow = true;
            else {
                $scope.vwhToYearShow = false;
                if ($scope.whToMonth != "")
                    $scope.whInvalidDate = checkInvDate($scope.whFromMonth.value, $scope.whFromYear, $scope.whToMonth.value, $scope.whToYear);
            }
        }

        $scope.whCurrentlyChg = function () {
            if ($scope.whCurrently) {
                $scope.vwhToMonthShow = false;
                $scope.vwhToYearShow = false;
                $scope.whToMonth = "";
                $scope.whToYear = "";
            }
        }

        // *********** for highlight and show the edit and delete buttons
        $scope.whMouseOver = function (context) {
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
        //function setEduFiels() {
        //    $scope.eduFields = [{ id: 1, name: "Computer science" }, { id: 2, name: "Biology" }, { id: 3, name: "Medical" }, { id: 4, name: "Architecture" }, { id: 5, name: "Civil engineering" }, { id: 6, name: "Linguistic" }, { id: 7, name: "Physics" }];
        //}
        function setEduFiels() {
            $scope.eduFields = [{ name: "Computer science" }, { name: "Biology" }, { name: "Medical" }, { name: "Architecture" }, { name: "Civil engineering" }, { name: "Linguistic" }, { name: "Physics" }];
        }

        function setEduItems() {
            if ($scope.degrees == null || $scope.degrees.length == 0)
                setDegrees();
            if ($scope.eduFields == null || $scope.eduFields.length == 0)
                setEduFiels();
        }
        $scope.addEducation = function () {
            setEduItems();
            cleanEducationInputs();
        }

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
                saveEdu.degree = parseInt($scope.eduDegree);
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
            setEduItems();
            $scope.eduId = edu.id;
            $scope.eduSchool = edu.school;
            $scope.eduFromDate = edu.fromdate;
            $scope.eduToDate = edu.todate;
            $scope.eduDegree = edu.degree.id;
            checkNewEduField(edu.field);
            $scope.eduField = edu.field;
            $scope.eduGrade = edu.grade;
            //$scpoe.eduDesc = edu.desc;
        }

        function checkNewEduField(field) {
            for(var i=0; i< $scope.eduFields.length;i++)
            {
                if (field == $scope.eduFields[i])
                    return;
            }
            $scope.eduFields.push({ name: field });
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
        $scope.eduDegreeChg = function (data) {
            if ($scope.eduDegree == "")
                $scope.veduDegreeShow = true;
            else
                $scope.veduDegreeShow = false;

            search(data);
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
        // Projects
        //*************************
        function setPrjItems() {
            if ($scope.occupations == null || $scope.occupations.length == 0)
                setOccupations();
        }
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
                savePrj.occupation = parseInt($scope.prjOccupation);
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
            occupations.query(
           function (data) {
               $scope.occupations = data;
           } // function(data)
           );
            $scope.prjId = prj.id;
            $scope.prjName = prj.name;
            $scope.prjOccupation = prj.occupation.id;
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
            occupations.query(
           function (data) {
               $scope.occupations = data;
           } // function(data)
           );
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
            cleanProjectsInputs();
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
        // Volunteering
        //*************************
        function setVolItems(){
            if($scope.volCauses == null || $scope.volCauses.length==0)
                setVolCauses();
        }
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
                saveVol.cause = parseInt($scope.volCause);
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
            setVolItems();
            $scope.volId = vol.id;
            $scope.volOrganization = vol.organization;
            $scope.volRole = vol.role;
            $scope.volCause = vol.cause.id;
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
            setVolItems();
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
            else {
                $scope.vvolFromMonthShow = false;
                if ($scope.volToMonth != "" && $scope.volToYear != "")
                    $scope.volInvalidDate = checkInvDate($scope.volFromMonth.value, $scope.volFromYear, $scope.volToMonth.value, $scope.volToYear);
            }
        }
        $scope.volFromYearChg = function () {
            if ($scope.volFromYear == "")
                $scope.vvolFromYearShow = true;
            else {
                $scope.vvolFromYearShow = false;
                if ($scope.volToMonth != "" && $scope.volToYear != "")
                    $scope.volInvalidDate = checkInvDate($scope.volFromMonth.value, $scope.volFromYear, $scope.volToMonth.value, $scope.volToYear);
            }
        }
        $scope.volToMonthChg = function () {
            if ($scope.volToMonth != "" && $scope.volToYear != "")
                $scope.volInvalidDate = checkInvDate($scope.volFromMonth.value, $scope.volFromYear, $scope.volToMonth.value, $scope.volToYear);
        }
        $scope.volToYearChg = function () {
            if ($scope.volToMonth != "" && $scope.volToYear != "")
                $scope.volInvalidDate = checkInvDate($scope.volFromMonth.value, $scope.volFromYear, $scope.volToMonth.value, $scope.volToYear);
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
        function setTcItems()
        {
            if ($scope.occupations == null || $scope.occupations.length == 0)
                setOccupations();
        }
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
                saveTc.occupation = parseInt($scope.tcOccupation);
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
            setTcItems();
            $scope.tcId = tc.id;
            $scope.tcName = tc.name;
            $scope.tcOccupation = tc.occupation.id;
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
            isValid = true;
            setTcItems();
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
        //  Honors & Awards
        //*************************
        function setHaItems() {
            if ($scope.occupations == null || $scope.occupations.length == 0)
                setOccupations();
        }
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
                saveHa.occupation = parseInt($scope.haOccupation);
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
            setHaItems();
            $scope.haId = ha.id;
            $scope.haTitle = ha.title;
            $scope.haOccupation = ha.occupation.id;
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
            setHaItems();
            isValid = true;
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
        function setCrItems() {
            if ($scope.occupations == null || $scope.occupations.length == 0)
                setOccupations();
        }
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
                saveCr.occupation = parseInt($scope.crOccupation);
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
            setCrItems();
            $scope.crId = cr.id;
            $scope.crName = cr.name;
            $scope.crOccupation = cr.occupation.id;
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
            setCrItems();
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
