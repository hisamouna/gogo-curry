pipeline {
 agent any
 
 stages{
  stage("git pull"){
   steps {
    dir(gogocurry) {
     checkout scm: [$class: 'GitSCM', 
      userRemoteConfigs: [[url: 'https://github.com/hikarusasa/gogo-curry.git']], 
      branches: [[name: "refs/tags/${TAG}"]]], changelog: false, poll: false
    }
   }
  }

  stage("run") {
   steps {
    sh ''' 
    echo "Hello My World"

    ls -trl
	
    '''
  }
 }
 
  stage("clean") {
   steps {
    deleteDir()
  }
 }
}
}
