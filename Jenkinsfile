pipeline {
 agent any
 
 stages{
  stage("git pull"){
   steps {
    dir(gogocurry) {
     git url: "https://github.com/hikarusasa/gogo-curry.git", branch: 'master'
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
