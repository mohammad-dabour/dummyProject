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
 
 
 #### This pipeline having  another part called [ansible-concourse-docker]. 
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


 
 
 
 
 
 I have read many concourse ci tutorials and docouments however in addition to the official docoumentation i think [this one was helpful too]
 
 [official documentation]: https://concourse-ci.org/docker-repository.html
 [this one was helpful too]: https://github.com/JeffDeCola/hello-go
 [ansible-concourse-docker]: https://github.com/mohammad-dabour/ansible-concourse-docker.git
