FROM circleci/golang:1.11

RUN curl https://sdk.cloud.google.com | bash

ENV PATH "/home/circleci/google-cloud-sdk/bin:${PATH}"

# install kubectl
RUN sudo apt-get update && sudo apt-get install -y apt-transport-https
RUN curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add - \
&& sudo touch /etc/apt/sources.list.d/kubernetes.list \
&& echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee -a /etc/apt/sources.list.d/kubernetes.list \
&& sudo apt-get update && sudo apt-get install -y kubectl
