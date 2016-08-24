var globalHiring = angular.module('globalHiring',['ngResource','ngAnimate', 'ui.bootstrap']);
globalHiring.controller('stepsController',['$scope', '$resource','$uibModal','$location', function($scope, $resource,$uibModal,$location, $http, $window, $log) {

//initializing page
var accountUser = $resource("/account")
$scope.init = function () {

    if($location.absUrl().indexOf('id') > -1){
      //$scope.id= $location.absUrl().split('=')[1];

      accountUser.get(
        function (data) {
            $scope.userAccount = data
            if(!$scope.userAccount.iscompleted)
            {
              $scope.step1Hide=true;
              if($scope.userAccount.isstudent){
                $scope.stepLangHide=false;
                $scope.langBarImage="images/step_2_bar.png";
                $scope.langPrevBtn=true;
              }
              else {
                $scope.stepMainSRVHide=false;
                $scope.mainSRVBarImage="images/step_2_bar.png";
                $scope.mainSrvPrevBtn=true;
                $scope.getMainServices();
              }
            }//if iscompleted
        },
        function(errorMsg){
          $scope.errorMsg="There is a problem in Loading your account, Please Try Again."
          $scope.errorModal();
        }
      );
    }
}
// finish initializing

  $scope.phoneRegex='[0-9 ]+';
  $scope.userAccount='';

  $scope.langPrevBtn=false;
  $scope.mainSrvPrevBtn=false;
  // Steps controller
  //$scope.authenticateHide=false;
  $scope.step1Hide=false;
  $scope.stepLangHide=true;
  $scope.stepEduHide=true;
  $scope.stepMainSRVHide=true;
  $scope.stepSkillHide=true;
  $scope.isStudent=true;
  $scope.isStudentShow=false;

  $scope.fullname='';
  $scope.termsOfServiceShow= false;
  $scope.checkboxValue=false;
  $scope.password='';
  $scope.email='';

  $scope.fullnamePlaceholder="";
  $scope.emailPlaceholder="";
  $scope.passwordPlaceholder="min 6 characters";
  $scope.requiedStyleFullname='';
  $scope.requiedStylePassword='';
  $scope.requiedStyleEmail='';
  $scope.cityPlaceholder="";
  $scope.requiedSelectedCity='';
  $scope.selectColorStyle="";
  $scope.requiedStyleisStudent="";

  $scope.errorMsg='';

  $scope.skillsBTNLabel="";
  $scope.eduBTNLabel="";
//Changes for proper style
$scope.fullNameChange= function(){
  if($scope.fullname=='' )
  {

    $scope.requiedStyleFullname={
      "border-color": "red",
      "border-style": "solid",
      "border-width" : "1px"
    }
    $scope.fullnamePlaceholder="Fullname is required";
  }
  else {
    $scope.requiedStyleFullname='';
    $scope.fullnamePlaceholder="";
  }
}

$scope.emailChange=function(){
  if($scope.email=='' )
  {

    $scope.requiedStyleEmail={
      "border-color": "red",
      "border-style": "solid",
      "border-width" : "1px"
    }

    $scope.emailPlaceholder="Email is required";
  }
  else {
    $scope.requiedStyleEmail='';
    $scope.emailPlaceholder="";
  }

}

$scope.cityChange=function(){
  if( $scope.selectedCity=="" )
  {
      $scope.requiedSelectedCity={
      "border-color": "red",
      "border-style": "solid",
      "border-width" : "1px",
      "background-image":"url('../images/combo-red.png')",
      "background-color":"#f9f9f9",
      "background-repeat":"no-repeat"
    }

    $scope.cityPlaceholder="City is required";
    $scope.selectColorStyle={"color":"red"};
  }
  else {
    $scope.requiedSelectedCity='';
    $scope.cityPlaceholder="";
    $scope.selectColorStyle={"color":"#8a8888"};
  }

}
$scope.passwordChange=function(){
  if($scope.password=='' )
  {

    $scope.requiedStylePassword={
      "border-color": "red",
      "border-style": "solid",
      "border-width" : "1px"
    }

    $scope.passwordPlaceholder="Password is required";
  }
  else {
    $scope.requiedStylePassword='';
    $scope.passwordPlaceholder="";
  }

}

  // Next buttons click
  $scope.step1NextClick = function(){
    $scope.requiedStyleFullname='';
    $scope.requiedStylePassword='';
    $scope.requiedStyleEmail='';


    if($scope.fullname=='' || $scope.email=='' || !$scope.checkboxValue || $scope.password=='' || $scope.selectedCity=='' || (!($scope.isStudent=="yes" || $scope.isStudent=="no"))  )
    {
      if(!($scope.isStudent=="yes" || $scope.isStudent=="no"))
      {
        $scope.requiedStyleisStudent={
          "border-color":"red"
        }
        $scope.isStudentShow=true;
      }
      else{
          $scope.requiedStyleisStudent="";
          $scope.isStudentShow=false;
      }

      if($scope.fullname=='' )
      {

        $scope.requiedStyleFullname={
          "border-color": "red",
          "border-style": "solid",
          "border-width" : "1px"
        }
        $scope.fullnamePlaceholder="Fullname is required";
      }
      else {
        $scope.requiedStyleFullname='';
        $scope.fullnamePlaceholder="";
      }

      if($scope.email=='' )
      {

        $scope.requiedStyleEmail={
          "border-color": "red",
          "border-style": "solid",
          "border-width" : "1px"
        }

        $scope.emailPlaceholder="Email is required";
      }
      else {
        $scope.requiedStyleEmail='';
        $scope.emailPlaceholder="";
      }

      if( $scope.selectedCity=="" )
      {
          $scope.requiedSelectedCity={
          "border-color": "red",
          "border-style": "solid",
          "border-width" : "1px",
          "background-image":"url('../images/combo-red.png')",
          "background-color":"#f9f9f9",
          "background-repeat":"no-repeat"
        }

        $scope.cityPlaceholder="City is required";
        $scope.selectColorStyle={"color":"red"};
      }
      else {
        $scope.requiedSelectedCity='';
        $scope.cityPlaceholder="";
        $scope.selectColorStyle={"color":"#8a8888"};
      }

      if($scope.password=='' )
      {

        $scope.requiedStylePassword={
          "border-color": "red",
          "border-style": "solid",
          "border-width" : "1px"
        }

        $scope.passwordPlaceholder="Password is required";
      }
      else {
        $scope.requiedStylePassword='';
        $scope.passwordPlaceholder="";
      }

      if(!$scope.checkboxValue){
        $scope.termsOfServiceShow= true;
      }
      else{
        $scope.termsOfServiceShow= false;
      }

    }
    else {
      $scope.step1Hide=true;
      $scope.createNewAccount();
      if($scope.isStudent=="yes"){$scope.isStudent=true;}else{$scope.isStudent=false;}

      if($scope.isStudent){
        $scope.stepLangHide=false;
        $scope.langBarImage="images/step_2_bar.png";
      }
      else {
        $scope.stepMainSRVHide=false;
        $scope.mainSRVBarImage="images/step_2_bar.png";
        $scope.getMainServices();
    }
  }
  };

  $scope.langNextClick = function(){

    if (($scope.langs.length==0 && $scope.langToAdd.name!='')||($scope.langs.length>0 && $scope.langs[$scope.langs.length-1].name!=$scope.langToAdd.name))
    {
      $scope.addMoreLangClick($scope.langToAdd);
    }
    if($scope.langs.length>0)
    {
      hasEnglish=false;
      for(i=0;i<$scope.langs.length;i++)
      {
        if($scope.langs[i].name=="English")
        {
          hasEnglish=true;
        }
      }
      if(!hasEnglish)
      {
        $scope.errorMsg="NO ENGLISH"
        $scope.forceEnglishModal();
      }
    }
    if($scope.langs.length==0)
    {
      $scope.errorMsg="NO ENGLISH"
      $scope.forceEnglishModal();
    }

    $scope.stepLangHide=true;
    if($scope.isStudent){
      $scope.stepEduHide=false;
      $scope.eduBTNLabel="Next";
      $scope.eduBarImage="images/step_3_bar.png";
    }
    else{
      $scope.stepEduHide=false;
      $scope.eduBTNLabel="Finish";
      $scope.eduBarImage="images/step_5_bar.png";

    }
  };
  $scope.eduNextClick = function(){

    if (($scope.educations.length==0 && $scope.educationToAdd.school!='')||($scope.educations.length>0 && $scope.educations[$scope.educations.length-1].school!=$scope.educationToAdd.school))
    {
      $scope.addMoreEduClick($scope.educationToAdd);
    }

      $scope.stepEduHide=true;
    if($scope.isStudent){
      $scope.stepMainSRVHide=false;
      $scope.mainSRVBarImage="images/step_4_bar.png";
      $scope.getMainServices();
    }
    else{
      //$scope.stepSkillHide=false;
      $scope.finishClick();
    }

  };
  $scope.mainSRVNextClick = function(){

    $scope.stepMainSRVHide=true;
    if($scope.isStudent){
      $scope.stepSkillHide=false;
      $scope.skillsBTNLabel="Finish";
      $scope.skillBarImage="images/step_5_bar.png";
    }
    else{
      $scope.stepSkillHide=false;
      $scope.skillsBTNLabel="Next";
      $scope.skillBarImage="images/step_3_bar.png";
    }
  };

  $scope.skillNextClick=function(){
    if($scope.isStudent){
      $scope.finishClick();
    }
    else{
      $scope.stepSkillHide=true;
      $scope.stepLangHide=false;
      $scope.langBarImage="images/step_4_bar.png";

    }
  }

  // Previous buttons click
  $scope.langPreClick=function(){
    $scope.stepLangHide=true;
    if($scope.isStudent){
      $scope.step1Hide=false;
    }
    else {
      $scope.stepSkillHide=false;
      $scope.skillBarImage="images/step_3_bar.png";
    }
  }
  $scope.eduPreClick=function(){
    $scope.stepEduHide=true;
    $scope.stepLangHide=false;
    if($scope.isStudent){
      $scope.langBarImage="images/step_2_bar.png";
    }
    else {
      $scope.langBarImage="images/step_4_bar.png";
    }
  }
  $scope.mainSRVPreClick=function(){
    $scope.stepMainSRVHide=true;
    if($scope.isStudent){
      $scope.stepEduHide=false;
      $scope.eduBarImage="images/step_3_bar.png";
    }
    else {
      $scope.step1Hide=false;
    }
  }
  $scope.skillPreClick=function(){
    $scope.stepMainSRVHide=false;
    $scope.stepSkillHide=true;
    if($scope.isStudent){
      $scope.mainSRVBarImage="images/step_5_bar.png";
    }
    else {
      $scope.mainSRVBarImage="images/step_3_bar.png";
    }
  }

  $scope.selectedCity='';

  $scope.dates=[];
  for(var i=1970;i<=2020;i++) {
    $scope.dates.push(i);
  }

  var cities = $resource("/cities")
    cities.query(
      function(data){
        $scope.cities = data;
      } // function(data)
    ) // service.query

    var degrees = $resource("/degrees")
    degrees.query(
      function(data){
        $scope.degrees = data;
      } // function(data)
    ) // service.query


  //Add more Language
  $scope.LanguageProfeciencies = [
    {text: "Elementary", value: "E", desc: "Discription"},
    {text: "Basic", value: "B", desc: "Discription"},
    {text: "Conversational", value: "C", desc: "Discription"},
    {text: "Fluent", value: "F", desc: "Discription"},
    {text: "Native", value: "N", desc: "Discription"}
  ]
  $scope.langs = [];
  $scope.langToAdd={
    name:'',
    profeciency:''
  };

  var index =0;
  $scope.addMoreLangClick=function(langToAdd){
    if( langToAdd.name!='' && langToAdd.profeciency!='')
    {
      $scope.langs.push(angular.copy(langToAdd));
      $scope.langToAdd.name='';
      $scope.langToAdd.profeciency='';
    }
  }


  //Force Englisg Modal

  $scope.animationsEnabled = true;
  $scope.forceEnglishModal = function () {

    var modalInstance = $uibModal.open({
      animation: $scope.animationsEnabled,
      templateUrl: 'forceEnglishModalContent.html',
      controller: 'ForceEnglishModalInstanceCtrl',
      resolve: {
        langToAdd: function () {
          return $scope.langToAdd;
        },
        LanguageProfeciencies: function () {
          return $scope.LanguageProfeciencies;
        }
      }
    });

    modalInstance.result.then(function (langToAdd) {
      $scope.addMoreLangClick(langToAdd);
    }, function () {
      $log.info('Modal dismissed at: ' + new Date());
    });
  };
  //Modal Force Englisg is finished

  //Add Education
  $scope.educations = [];
  $scope.educationToAdd = {
    school: '',
    fromdate: '',
    todate: '',
    degree:'',
    field:'',
    grade:''
  };

  var indexEdu =0;
  $scope.addMoreEduClick=function(educationToAdd){
    if($scope.educationToAdd.school!='')
    {
      $scope.educations.push(angular.copy(educationToAdd));
      $scope.educationToAdd.school='';
      $scope.educationToAdd.fromdate='';
      $scope.educationToAdd.todate='';
      $scope.educationToAdd.degree='';
      $scope.educationToAdd.field='';
      $scope.educationToAdd.grade='';
    }
  }

  // Main Services
  $scope.toggleSelectedSRV = function(service) {
     service.isselected = !service.isselected;
  };

  var service = $resource("/mainServices")
  $scope.getMainServices = function(){
    service.query(
      function(data){
        $scope.mainServices = data;
      } // function(data)
    ) // service.query
  };

  $scope.getIcon = function(serviceSelect){
    if (serviceSelect.icon) {
      if (serviceSelect.isChecked) return serviceSelect.icon.on;
      else return serviceSelect.icon.off;
    }
  }

  $scope.skillProfeciencies = [
    {text: "Student/Fresh Graduate", value: "S"},
    {text: "Junior Professional", value: "J"},
    {text: "Experienced Professional", value: "E"},
    {text: "Manager", value: "M"}
  ]
  //Profeciency Modal

  $scope.animationsEnabled = true;
  $scope.openSkillProf = function (selectedSkill) {

    var modalInstance = $uibModal.open({
      animation: $scope.animationsEnabled,
      templateUrl: 'myModalContent.html',
      controller: 'ModalInstanceCtrl',
      resolve: {
        skillProfeciencies: function () {
          return $scope.skillProfeciencies;
        },
        selectedSkill : function (){
          return selectedSkill;
        }
      }
    });

    modalInstance.result.then(function (selectedItem) {
      //TODO: should be checked with Back-end
      selectedSkill.profeciency=selectedItem.profeciency.value;
    }, function () {
      $log.info('Modal dismissed at: ' + new Date());
    });
  };
  //Modal Profeciency is finished


  $scope.skillClick=function(selectedSkill){
    selectedSkill.isselected=!selectedSkill.isselected;
    if(selectedSkill.isselected)
    {
      $scope.openSkillProf(selectedSkill);

    }
  }

  //ErrorMsg Modal

  $scope.animationsEnabled = true;
  $scope.errorModal = function () {

    var errorModalInstance = $uibModal.open({
      animation: $scope.animationsEnabled,
      templateUrl: 'errorModal.html',
      controller: 'ErrormodalInstanceCtrl',
      resolve: {
        errorMsg: function () {
          return $scope.errorMsg;
        }
      }
    });

    errorModalInstance.result.then(function () {
      var test="";
        window.location = "/steps.html";
    }, function () {
      $log.info('Modal dismissed at: ' + new Date());
    });
  };
  //Modal ErrorMsg is finished

// Create A New Account in First steps
var saveAccount = $resource("/saveAccount")
$scope.createNewAccount=function(){
      var newAccount = new saveAccount();
      newAccount.name = $scope.fullname;
      newAccount.email = $scope.email;
      newAccount.city = $scope.selectedCity;
      newAccount.phone = $scope.phone;
      newAccount.password = $scope.password;
      if($scope.isStudent){
        newAccount.isstudent = true;
      }
      else {
        newAccount.isstudent = false;
      }

      newAccount.isCompleted=false;

      newAccount.$save(
        function(data){
          $scope.userAccount=data;
          //window.location = "/profile.html"; // save and move to profile page
          if($scope.userAccount.isstudent)
          {
            $scope.langPrevBtn=true;
          }
          else {
            $scope.mainSrvPrevBtn=true;
          }
        },
        function(errorMsg){
          $scope.errorMsg="There is a problem in creating your account, Please Try Again."
          $scope.errorModal();
        }
      );
}
// End of $scope.createNewAccount

// Save an account
  var completeAccount = $resource("/completeAccount")
  $scope.finishClick=function(){
      var finishAccount = new completeAccount();
      finishAccount.id=$scope.userAccount.id;
      finishAccount.languages = $scope.langs;

      finishAccount.educations = $scope.educations;

      finishAccount.skills = $scope.mainServices;
      finishAccount.isCompleted=true;

      finishAccount.$save(
        function(data){
          $scope.userAccount=data;
          window.location = "/profile.html"; // save and move to profile page
        },
        function(errorMsg){
          $scope.errorMsg="There is a problem in creating your account, Please Try Again."
          $scope.errorModal();
        }
      );
  }// finish click

  $scope.init();
}
]// scope, resource
);


angular.module('globalHiring').controller('ModalInstanceCtrl', function ($scope, $uibModalInstance, skillProfeciencies, selectedSkill) {

  $scope.skillProfeciencies = skillProfeciencies;
  $scope.selectedSkill=selectedSkill;
  $scope.selected = {
    profeciency: $scope.skillProfeciencies[0]
  };

  $scope.close= function () {
    $uibModalInstance.close($scope.selected);
  };

  $scope.cancel = function () {
    $uibModalInstance.dismiss('cancel');
  };
});


angular.module('globalHiring').controller('ErrormodalInstanceCtrl', function ($scope, $uibModalInstance, errorMsg) {

  $scope.errorMsgToShow=errorMsg;

  $scope.close= function () {
    $uibModalInstance.close();
  };
  $scope.ok= function () {
    $uibModalInstance.close();
  };
});


angular.module('globalHiring').controller('ForceEnglishModalInstanceCtrl', function ($scope, $uibModalInstance,langToAdd,LanguageProfeciencies) {

  $scope.langToAdd = langToAdd;
  $scope.LanguageProfeciencies=LanguageProfeciencies;
  $scope.langToAdd = {
    name:"English",
    profeciency: $scope.LanguageProfeciencies[0].value
  };

  $scope.close= function () {
    $uibModalInstance.close($scope.langToAdd);
  };

  $scope.cancel = function () {
    $uibModalInstance.dismiss('cancel');
  };
});
