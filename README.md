#### dummyProject
This is Foo Bar project implementing a pipeline using Concourse CI 

###### This pipeline uses many parties, there's no why regarding to what i used. As this is not production app. I wanted to do a full build and deployment for multi components using different approaches that can introduce me to  different possible features in concourse ci while am learning, also to integrate other tools like ansible and to see how can i utilize docker-compose for dummy simple app, the pipeline is not for production it is AS is.

#### Technogloies :
 * Local concourse ci where i used the docker version provided by [official documentation] 
 * Docker, and docker-compose for both app and concourese ci
 * golang
 * ansible
 * consul
 * nginx
 * registrator
 
 #### related repos:
 * [ansible-concourse-docker]
 * [docker-nginx]
 
 
 #### This project having  another part called [ansible-concourse-docker]. 
 The ansible part, is a job for building and preparing ansible docker image with necessary access and playbooks.
 The image will be pushed to local registry, while some other images like wikiapp will be pushed to my https://hub.docker.com 
 Another the wikiserver-test pipeline will use the ansible image to do the release.
 
 #### What is going on ?
 
 
 ##### set pipeline:
 
 ```
 export registry_address=$(ifconfig|awk '/(inet)/ && !/(inet6)/&& !/127.0.0.1/{print $2}')
 
 fly  -t lite set-pipeline -p wikiserver-test -c ci/wikiserver-pipline.yml  -l .credentials.yml --var="reg_address=$registry_address"

.credentials.yml will have the following content


-p: this is the pipeline name: wikiserver-test
-c: this is the pipeline config ci/wikiserver-pipline.yml
-l: .credentials.yml or --load-vars-from .credentials.yml  stored your credentials and private keys if needed. example of content: 
PASS: USER
USER: USER
ansible_user: USER

I did not store it on github, it is within concourseci configs, if needed use vault.
 ```

##### inside the pipeline's resources:

```
# wikiserver-pipline.yml
we use follwong resources for the main piepline git and docker-image 2 bellow are samples.
resources:
- name: wiki-server
  type: git
  source:
    uri: https://gitlab.com/mdab0ur/dummywikiapp.git
    password: {{PASS}}
    username: {{USER}}
    branch: master

- name: wikigoapp-docker
  type: docker-image
  source:
    repository: mdab0ur/wikigoapp
    username: {{USER}}
    password: {{PASS}}
{{USER}} and {{PASS}} will be loaded from .credentials.yml

You can use github ssh and load priavte keys using --var :

--var="ssh_container_key=$(cat you/key/file)"
Adding they keys to .credentials.yml did not work for me, as am using concourse ci v3.9 but using --var worked.

```

##### Inside the pipeline jobs:

```
jobs:
- name: wikiserver-test
  plan:
    - get: wiki-server
      trigger: true
    - task: test-wikiserver
      file: wiki-server/ci/tasks/test-task.yml
    - task: publish-testing-results-wikiserver
      file: wiki-server/ci/tasks/test-publish.yml
      
With every git push on the resource wiki-server which reperesnts dummywikiapp.git  the job  will be triggered. 
2 tasks will happen:
* code testing.
* test results publish. it is fake!.


- name: build-wiki
  plan:
    - get: wiki-server
      trigger: true
      passed: [wikiserver-test]
    - task: task-build
      file: wiki-server/ci/tasks/task-build.yml
    - put: wikigoapp-docker
      params: {build: "wiki-build"}


If the wikiserver-test job passed. code build starts. 
which will produce wikigoapp-docker into mdab0ur/wikigoapp my repostry on dockerhub.
you can use your own repo, do not depend on mine.
Also you can set it up as local registry we'll see ansible docker image loaded into local registry.



- name: job-deploy-app
  plan:
    - get: nginx-server
      trigger: true
      passed: [docker-nginx-build]
    - get: wiki-server
      trigger: true
      passed: [build-wiki]
    - task: run-docker-ansible
      config:
        platform: linux
        image_resource:
            type: docker-image
            source:
              repository: ((reg_address)):5000/ansible
              insecure_registries: ["((reg_address)):5000"]
        params:
           AnsibleUSER: {{ansible_user}}
           password: {{PASS}} #this is  dockerhub paasword
           username: {{USER}} #this is  dockerhub user
        run:
         path: /bin/bash
         args:
           - -c
           - |
              ansible-playbook -v -i /tmp/inventory /tmp/playbook.yml -u $AnsibleUSER --private-key /tmp/tsa_host_key --sudo --extra-vars "remoteUser=$USER user=$username pass=$password" 


Last joib after making  sure your compoenents images were build successfuly,
then run-docker-ansible reg_address was passed as :  

--var="reg_address=$(ifconfig|awk '/(inet)/ && !/(inet6)/&& !/127.0.0.1/{print $2}')"

It will start deploying reffering  to our:
https://github.com/mohammad-dabour/ansible-concourse-docker.git
```


##### The pipeline will look something like this:

![alt text](https://raw.githubusercontent.com/mohammad-dabour/dummyProject/master/Screen%20Shot%202018-03-28%20at%202.49.42%20PM.png)

 
 
 ##### What's going on with [ansible-pipeline.yml]  ?
 
 ###### set pipeline:
 
 ```
 export registry_address=$(ifconfig|awk '/(inet)/ && !/(inet6)/&& !/127.0.0.1/{print $2}')
 fly -t lite set-pipeline -p ansible-deployer -c ci/ansible-pipeline.yml  -l .credentials.yml --var="ssh_container_key=$(cat path/to/key/file)" --var "reg_address=$registry_address"
 
 reg_address: is your local dockerhub registry
 ```
 
 
 ```
 
 
 resources:
- name: ansible-deployer
  type: git
  source:
    uri: https://github.com/mohammad-dabour/ansible-concourse-docker.git
    password: {{PASS}}
    username: {{USER}}
    branch: master


- name: ansible-docker
  type: docker-image
  source:
    repository: ((reg_address)):5000/ansible
    insecure_registries: ["((reg_address)):5000"]
    
## Above insecure_registries: ((reg_address)) will be taken from --var "reg_address=$(ifconfig|awk '/(inet)/ && !/(inet6)/&& !/127.0.0.1/{print $2}')"" during the setup.

Also note: $(ifconfig|awk '/(inet)/ && !/(inet6)/&& !/127.0.0.1/{print $2}')" was used on mac os, on gcloud instance i used  awk '/(instance)/{print $1}' /etc/hosts.

jobs:
- name: build-deployer
  plan:
    - get: ansible-deployer
      trigger: true
    - put: ansible-docker
      params: 
         build: ansible-deployer
         build_args:
            SSH_KEY: {{ssh_container_key}}
            USER: {{docker_user}}
This will build the ansible image with passing necessery keys localy also the image will be in local registry or private repo.
```

 
 
 ##### The Ansible pipeline will look something like this:
![alt text]( https://raw.githubusercontent.com/mohammad-dabour/dummyProject/master/Screen%20Shot%202018-03-28%20at%203.02.16%20PM.png)

 I have read concourse official docoumentation mainly  however in addition to the official docoumentation  i think [this one was helpful too]
 
 [official documentation]: https://concourse-ci.org/docker-repository.html
 [this one was helpful too]: https://github.com/JeffDeCola/hello-go
 [ansible-concourse-docker]: https://github.com/mohammad-dabour/ansible-concourse-docker.git
 [ansible-pipeline.yml]: https://github.com/mohammad-dabour/ansible-concourse-docker/blob/master/ci/ansible-pipeline.yml
 [docker-nginx]: https://github.com/mohammad-dabour/docker-nginx
