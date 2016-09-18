var globalHiring = angular.module('globalHiring',['ngResource','ngAnimate', 'ui.bootstrap','selectize']);
globalHiring.controller('stepsController',['$scope', '$resource','$uibModal','$location','$parse', function($scope, $resource,$uibModal,$location,$parse, $http, $window, $log) {
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
                $scope.stepEduHide=false;
                $scope.eduBarImage="images/step_2_bar.png";
                $scope.eduPrevBtn=true;
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

  //$scope.langPrevBtn=false;
  $scope.eduPrevBtn=false;
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
  $scope.srvShouldSelect=false;

  $scope.fullname='';
  $scope.termsOfServiceShow= false;
  $scope.checkboxValue=false;
  $scope.password='';
  $scope.email='';
  $scope.isRedFNPlacehoder=false;
  $scope.isRedEmailPlacehoder=false;
  $scope.isRedPWPlacehoder=false;
  $scope.isRedLangPlacehoder=false;

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

  $scope.langPlaceholder="English";
  $scope.requiedStyleLang="";
  $scope.langProfErrorShow=false;

  $scope.skillModalProfErrorShow=false;

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
      "background-repeat":"no-repeat",
      "color":"red"
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
        $scope.isRedFNPlacehoder=true;
        $scope.requiedStyleFullname={
          "border-color": "red",
          "border-style": "solid",
          "border-width" : "1px"
        }
        $scope.fullnamePlaceholder="Fullname is required";
      }
      else {
        $scope.isRedFNPlacehoder=false;
        $scope.requiedStyleFullname='';
        $scope.fullnamePlaceholder="";
      }

      if($scope.email=='' )
      {
        $scope.isRedEmailPlacehoder=true;
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
        $scope.isRedEmailPlacehoder=false;
      }

      if( $scope.selectedCity=="" )
      {
          $scope.requiedSelectedCity={
          "border-color": "red",
          "border-style": "solid",
          "border-width" : "1px",
          "background-image":"url('../images/combo-red.png')",
          "background-color":"#f9f9f9",
          "background-repeat":"no-repeat",
          "color":"red"
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
        $scope.isRedPWPlacehoder=true;
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
        $scope.isRedPWPlacehoder=false;
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
        $scope.stepEduHide=false;
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
    isValidationComplete= $scope.validateLang();
    if (isValidationComplete &&(($scope.langs.length==0 && $scope.langToAdd.name!='')||($scope.langs.length>0 && !$scope.checkIfHasLang())))
    {
      $scope.addMoreLangClick($scope.langToAdd);
    }

    if(isValidationComplete && $scope.langs.length>0)
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
    if(isValidationComplete && $scope.langs.length==0)
    {
      $scope.errorMsg="NO ENGLISH"
      $scope.forceEnglishModal();
    }

    if(isValidationComplete && $scope.langs.length!=0)
    {
      $scope.langProfErrorShow=false;
      $scope.langPlaceholder="English";
      $scope.requiedStyleLang="";

    //$scope.stepLangHide=true;
    //if($scope.isStudent){
    //  $scope.stepEduHide=false;
    //  $scope.eduBTNLabel="Next Step";
    //  $scope.eduBarImage="images/step_3_bar.png";
  //  $scope.finishClick();
  //  }
  //  else{
    //  $scope.stepEduHide=false;
    //  $scope.eduBTNLabel="Finish";
    //  $scope.eduBarImage="images/step_5_bar.png";
    $scope.finishClick();

  //  }
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
      $scope.mainSRVBarImage="images/step_3_bar.png";
      $scope.getMainServices();
    }
    else{
      //$scope.stepSkillHide=false;
      $scope.stepLangHide=false;
      $scope.langBarImage="images/step_5_bar.png";
      //$scope.finishClick();
    }

  };
  $scope.mainSRVNextClick = function(){
    serviceIsSelected=false;
    for(i=0;i<$scope.mainServices.length;i++)
    {
      if($scope.mainServices[i].isselected)
        {serviceIsSelected=true;}

    }
    if(!serviceIsSelected){$scope.srvShouldSelect=true;}
    else{
      $scope.srvShouldSelect=false;
      $scope.stepMainSRVHide=true;
      $scope.setErrorLabel();
      if($scope.isStudent){
        $scope.stepSkillHide=false;
        //$scope.skillsBTNLabel="Finish";
        $scope.skillBarImage="images/step_4_bar.png";
      }
      else{
        $scope.stepSkillHide=false;
      //  $scope.skillsBTNLabel="Next Step";
        $scope.skillBarImage="images/step_3_bar.png";
      }
    }
  };

  $scope.skillNextClick=function(){
    if(!$scope.chkIFNotSelected())
    {
    if($scope.isStudent){
      //$scope.finishClick();
      $scope.stepSkillHide=true;
      $scope.stepLangHide=false;
      $scope.langBarImage="images/step_5_bar.png";
    }
    else{
      $scope.stepSkillHide=true;
      $scope.stepEduHide=false;
      $scope.langEduImage="images/step_4_bar.png";

    }
  }
  }

  // Previous buttons click
  $scope.langPreClick=function(){
    $scope.stepLangHide=true;
    if($scope.isStudent){
      $scope.stepSkillHide=false;
      $scope.skillBarImage="images/step_4_bar.png"
    }
    else {
      $scope.stepEduHide=false;
      $scope.eduBarImage="images/step_4_bar.png";
    }
  }
  $scope.eduPreClick=function(){
    $scope.stepEduHide=true;
    //$scope.stepLangHide=false;
    if($scope.isStudent){
      //$scope.langBarImage="images/step_2_bar.png";
      $scope.step1Hide=false;
    }
    else {
      $scope.stepSkillHide=false;
      $scope.skillBarImage="images/step_3_bar.png";
    }
  }
  $scope.mainSRVPreClick=function(){
    $scope.stepMainSRVHide=true;
    if($scope.isStudent){
      $scope.stepEduHide=false;
      $scope.eduBarImage="images/step_2_bar.png";
    }
    else {
      $scope.step1Hide=false;
    }
  }
  $scope.skillPreClick=function(){
    $scope.stepMainSRVHide=false;
    $scope.stepSkillHide=true;
    if($scope.isStudent){
      $scope.mainSRVBarImage="images/step_3_bar.png";
    }
    else {
      $scope.mainSRVBarImage="images/step_2_bar.png";
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

    $scope.cityConfig = {
        valueField: 'id',
        labelField: 'name',
        searchField: 'name',
        closeAfterSelect:true,
        sortField: 'text',
        maxItems: 1,
        maxOptions: 5,
    };
    var degrees = $resource("/degrees")
    degrees.query(
      function(data){
        $scope.degrees = data;
      } // function(data)
    ) // service.query


  //Add more Language
  $scope.LanguageProfeciencies = [
    {text: "Elementary", value: "E", desc: "I am only able to understand written text in it."},
    {text: "Basic", value: "B", desc: "I am only able to communicate with it through written communication."},
    {text: "Conversational", value: "C", desc: "I know it well enough to verbally discuss project details with a client."},
    {text: "Fluent", value: "F", desc: "I have complete command of this language with perfect grammer."},
    {text: "Native", value: "N", desc: "I have complete command of this language, including breadth of vocabulary, idioms and colloquialisms."}
  ]
  $scope.langs = [];
  $scope.langToAdd={
    name:'',
    profeciency:''
  };

  $scope.checkIfHasLang=function(){
    if($scope.langToAdd.name!='')
    {
      haslang=false;
      for(i=0;i<$scope.langs.length;i++)
      {
        if($scope.langs[i].name==$scope.langToAdd.name)
        {
          haslang=true;
        }
      }
      return haslang;
    }
    else{return true;}
  }
  $scope.validateLang=function(){
    isValidationComplete=false;
    if($scope.langToAdd.name==''&& $scope.langToAdd.profeciency!='')
    {
      $scope.isRedLangPlacehoder=true;
      $scope.langPlaceholder="This field is required";
      $scope.requiedStyleLang={
      "border-color": "red",
      "border-style": "solid",
      "border-width" : "1px"
      }
    }
    else if($scope.langToAdd.name!=''&& $scope.langToAdd.profeciency==''){
      $scope.isRedLangPlacehoder=false;
      $scope.langProfErrorShow=true;
    }
    else if ($scope.langToAdd.name!=''&& $scope.langToAdd.profeciency!=''){
      $scope.isRedLangPlacehoder=false;
      $scope.langProfErrorShow=false;
      $scope.langPlaceholder="English";
      $scope.requiedStyleLang="";
      isValidationComplete=true;
    }
    else if($scope.langToAdd.name==''&& $scope.langToAdd.profeciency==''){ isValidationComplete=true;}
    return isValidationComplete;
  }


  var index =0;
  $scope.addMoreLangClick=function(langToAdd){
    isValidationComplete=$scope.validateLang();
    if(langToAdd.name!="English"){
      if(langToAdd.profeciency==''){$scope.langProfErrorShow=true;}
      if(langToAdd.name==''){
        $scope.isRedLangPlacehoder=true;
        $scope.langPlaceholder="This field is required";
        $scope.requiedStyleLang={
        "border-color": "red",
        "border-style": "solid",
        "border-width" : "1px"
        }
      }
    }
    if(isValidationComplete && langToAdd.name!='')
    {
      $scope.langs.push(angular.copy(langToAdd));
      $scope.langToAdd.name='';
      $scope.langToAdd.profeciency='';
      $scope.langProfErrorShow=false;
      $scope.langPlaceholder="English";
      $scope.requiedStyleLang="";
      $scope.isRedLangPlacehoder=false;
    }
  }

$scope.removeLanguage=function(lang){
  var removeLangIndex = $scope.langs.indexOf(lang);
  if (index > -1) {
    $scope.langs.splice(removeLangIndex, 1);
  }
}

  //Force Englisg Modal

  $scope.animationsEnabled = true;
  $scope.forceEnglishModal = function () {

    var modalInstance = $uibModal.open({
      animation: $scope.animationsEnabled,
      templateUrl: 'forceEnglishModalContent.html',
      controller: 'ForceEnglishModalInstanceCtrl',
      windowClass: 'app-lang-modal-window',
      backdrop  : 'static',
      keyboard  : false,
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
      if (serviceSelect.isselected) return serviceSelect.selectimageurl;
      else return serviceSelect.unselectimageurl;
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
      windowClass: 'app-skill-modal-window',
      backdrop  : 'static',
      keyboard  : false,
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
    else {
      selectedSkill.profeciency="";
    }
  }

$scope.setErrorLabel=function(){
 for(i=0;i<$scope.mainServices.length;i++)
  {
    if($scope.mainServices[i].isselected)
      {
        var showSrvlbl='errorshow'+$scope.mainServices[i].id;
        $parse(showSrvlbl).assign($scope, false);
      }
  }
}

$scope.setChkedErrorLabel=function(mainSRV,value){
  if(mainSRV.isselected)
    {
      var showSrvlbl='errorshow'+ mainSRV.id;
      $parse(showSrvlbl).assign($scope, value);
    }

}

$scope.chkIFNotSelected=function(){
  skillNotSelected=false;
  for(i=0;i<$scope.mainServices.length;i++)
   {
     if($scope.mainServices[i].isselected)
     {
       isSelected=false;
       for(j=0;j<$scope.mainServices[i].skills.length;j++){
         if($scope.mainServices[i].skills[j].profeciency!=""){
           isSelected=true;
         }
       }
       if(!isSelected){
         $scope.setChkedErrorLabel($scope.mainServices[i],true);
         skillNotSelected=true;
       }
       else{$scope.setChkedErrorLabel($scope.mainServices[i],false);}
     }
   }
   return skillNotSelected;
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
    //  newAccount.city = $scope.selectedCity;
      newAccount.city = $scope.cities[0];
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
            //$scope.langPrevBtn=true;
            $scope.eduPrevBtn=true;
            $scope.eduBarImage="images/step_2_bar.png";

          }
          else {
            $scope.mainSrvPrevBtn=true;
            $scope.mainSRVBarImage="images/step_2_bar.png"
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

  $scope.langModalProfErrorShow=false;
  $scope.langToAdd = langToAdd;
  $scope.LanguageProfeciencies=LanguageProfeciencies;
  $scope.langToAdd = {
    name:"English",
    //profeciency: $scope.LanguageProfeciencies[0].value
    profeciency:''
  };

  $scope.close= function () {
    if($scope.langToAdd.profeciency!=''){
    $uibModalInstance.close($scope.langToAdd);
    }
    else {
        $scope.langModalProfErrorShow=true;
    }
  };

  $scope.cancel = function () {
    $uibModalInstance.dismiss('cancel');
  };
});
