properties([
    buildDiscarder(logRotator(daysToKeepStr: '365', numToKeepStr: '350')),
    parameters([
        string(name: 'GIT_SHA', description: 'Git SHA to checkout.'),
        string(name: 'PROMETHEUS_BRANCH', description: 'Branch to checkout. <b>Will not use if GIT_SHA is set.</b>'),
        booleanParam(name: 'FINAL_CUT', defaultValue: false, description: 'Automation parameter, do not touch')
    ])
])

library 'jenkins-pipeline-library@main'

boolean FINAL_CUT = params.FINAL_CUT
String gitSha = params.GIT_SHA
String gitShaShort = ""
String prometheusBranch = params.PROMETHEUS_BRANCH

if (gitSha == "") {
    if (prometheusBranch == "") {
        error("This build requires either a git sha or a branch. Please check upstream job!")
    }
    println "Git SHA not provided. Using branch ${prometheusBranch} as param for checkout."
}

node(label: "centos7_fast_devserver") {
// node(label: "dev3-177") {
    notify(slackChannel: '#api-team-build-notification') {

        cleanWs()

        // Set build description
        if (gitSha == "") {
            currentBuild.description = "Source Path: pubsubplus-prometheus-exporter.git/${prometheusBranch}"
        } else {
            currentBuild.description = "Source Path: pubsubplus-prometheus-exporter.git@${gitSha}"
        }

        // Checkout source code using either <GIT_SHA> or <PROMETHEUS_BRANCH>
        stage("Checkout") {

            String branchName = gitSha
            if (branchName == "") {
                branchName = "*/${prometheusBranch}"
            }

            println "Checkout out from branch ${branchName}"
            try {
                def checkoutResults = checkout([$class: 'GitSCM',
                    branches: [[name: branchName]],
                    doGenerateSubmoduleConfigurations: false,
                    extensions: [],
                    submoduleCfg: [],
                    userRemoteConfigs: [[credentialsId: 're-github-bot-1', url: 'git@github.com:SolaceDev/pubsubplus-prometheus-exporter.git']]
                ])

                if (gitSha == "") {
                    gitSha = checkoutResults.GIT_COMMIT
                }
                gitShaShort = gitSha.substring(0,5)
            } catch (e) {
                println "Error while checking out pubsubplus-prometheus-exporter repository."
                error(e)
            }
        }

        // Replace -SNAPSHOT depending on the following:
        //      Feature branch     --> version-SNAPSHOT
        //      Main branch        --> version-${GIT_SHA}
        //      Release branch     --> version-RC
        //      Release branch+FC  --> version
        //
        //      Release pattern currently matches: 1.0.0
        String releasePattern = /[0-9].[0-9].[0-9]$/
        boolean isReleaseBranch = false
        stage("Version and Save Docker Image") {
            //get version path from version.go: version ex:1.0.0
            sh "cd /opt/cvsdirs/loadbuild/jenkins/slave/workspace/prometheus-exporter-build"
            String version = sh(returnStdout:true, script:"cd goversion && cat version.go | grep \"const version\" | sed 's/const version = \"\\(.*\\)\"/\\1/'").trim()

            String uniqueVersion = gitShaShort

            //build docker image of pubsubplus-prometheus-exporter project
            sh "docker build -t apps-jenkins:18888/pubsubplus-prometheus-exporter:${version}-${uniqueVersion} ."

            //save docker image as tar file
            sh "docker save apps-jenkins:18888/pubsubplus-prometheus-exporter:${version}-${uniqueVersion} | gzip > ./pubsubplus-prometheus-exporter_${version}-${uniqueVersion}.tar.gz"

            //make a new directory to store the tar file
            sh "mkdir -p /home/public/RND/loads/pubsubplus-prometheus-exporter/${version}"

            //move the tar file to the new directory
            sh "mv pubsubplus-prometheus-exporter_${version}-${uniqueVersion}.tar.gz /home/public/RND/loads/pubsubplus-prometheus-exporter/${version}"
        }

    }
}