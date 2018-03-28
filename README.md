#### dummyProject
This is Foo Bar project implementing a pipeline using Concourse CI 

###### This pipeline uses many parties, there's no why regarding to what i used. As this is not production app. I wanted to do a full deployment for multi components using different approaches that can introduce me to  different possible features in concourse ci while am learning, the pipeline is not for production it is AS is.

#### Technogloies :
 * Local concourse ci where i used the docker version provided by [official documentation] 
 * Docker, and docker-compose for both app and concourese ci
 * golang
 * ansible
 * consul
 * nginx
 * registrator
 
 
 #### This project having  another part called [ansible-concourse-docker]. 
 The ansible part, is a job for building and preparing ansible docker image with necessary access and playbooks.
 The image will be pushed to local registry, while some other images like wikiapp will be pushed to my https://hub.docker.com 
 Another the wikiserver-test pipeline will use the ansible image to do the release.
 
 #### What is going on ?
 
 ```
 fly  -t lite set-pipeline -p wikiserver-test -c ci/wikiserver-pipline.yml  -l .credentials.yml --var="reg_address=$(ifconfig|awk '/(inet)/ && !/(inet6)/&& !/127.0.0.1/{print $2}')"

.credentials.yml will have the following content


-p: this is the pipeline name: wikiserver-test
-c: this is the pipeline config ci/wikiserver-pipline.yml
-l: .credentials.yml or --load-vars-from .credentials.yml  stored your credentials and private keys if needed. example of content: 
PASS: USER
USER: USER
docker_user: USER

I did not store it on github, it is within concourseci configs, if needed use vault.
 ```


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

 
 
 
 
 
 I have read many concourse ci tutorials and docouments however in addition to the official docoumentation i think [this one was helpful too]
 
 [official documentation]: https://concourse-ci.org/docker-repository.html
 [this one was helpful too]: https://github.com/JeffDeCola/hello-go
 [ansible-concourse-docker]: https://github.com/mohammad-dabour/ansible-concourse-docker.git
