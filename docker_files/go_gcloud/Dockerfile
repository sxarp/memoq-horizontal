# version: 2.0
ARG GO_VERSION=1.11


FROM circleci/golang:${GO_VERSION}

RUN curl https://sdk.cloud.google.com | bash

ENV PATH "/home/circleci/google-cloud-sdk/bin:${PATH}"

# install kubectl
RUN sudo apt-get update && sudo apt-get install -y apt-transport-https
RUN curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add - \
&& sudo touch /etc/apt/sources.list.d/kubernetes.list \
&& echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee -a /etc/apt/sources.list.d/kubernetes.list \
&& sudo apt-get update && sudo apt-get install -y kubectl

# install dep
RUN DEP_VERSION=0.5.0; sudo curl -L -s "https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64" -o /bin/dep; sudo chmod +x /bin/dep
