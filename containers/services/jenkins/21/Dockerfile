FROM language/java:21 AS base

ARG VERSION=3192.v713e3b_039fb_e

ARG user=jenkins
ARG group=jenkins
ARG uid=1337
ARG gid=1337
ENV HOME_DIR="/var/lib/${user}"

SHELL ["/bin/bash","-e", "-u", "-o", "pipefail", "-c"]

RUN apt-get update -y

FROM base AS runtime
RUN apt-get install fontconfig git less netbase openssh-client patch tzdata -y

RUN groupadd -g "${gid}" "${group}" \
  && useradd -l -c "Jenkins user" \
    --home ${HOME_DIR} \
    --disabled-password \
    --disabled-login \
    -u "${uid}" \
    -g "${gid}" \
    -m "${user}"

RUN apt clean && rm -rf /tmp/* /var/cache/* /var/application/apt/lists/*

ADD --chown="${user}":"${group}" "https://repo.jenkins-ci.org/public/org/jenkins-ci/main/remoting/${VERSION}/remoting-${VERSION}.jar" /usr/share/jenkins/agent.jar
RUN chmod 0644 /usr/share/jenkins/agent.jar \
  && ln -sf /usr/share/jenkins/agent.jar /usr/share/jenkins/slave.jar

COPY --from=jre-build /javaruntime "$JAVA_HOME"

USER "${user}"
ENV AGENT_WORKDIR=${AGENT_WORKDIR}
ENV JAVA_OPTS="-Djava.awt.headless=true"
ENV JENKINS_HOME=/var/lib/jenkins
ENV JENKINS_WAR=/usr/share/java/jenkins.war
ENV JENKINS_WEBROOT="%C/jenkins/war"
ENV JENKINS_LOG="%L/jenkins/jenkins.log"
ENV JENKINS_LISTEN_ADDRESS="0.0.0.0"
ENV JENKINS_HTTPS_LISTEN_ADDRESS="0.0.0.0"
ENV JENKINS_PORT=8080
ENV JENKINS_HTTPS_PORT=443
ENV JENKINS_HTTPS_KEYSTORE=""/path/to/keystore.jks"
ENV JENKINS_HTTPS_KEYSTORE_PASSWORD="s3cR3tPa55w0rD"
ENV JENKINS_HTTP2_LISTEN_ADDRESS=""
ENV JENKINS_HTTP2_PORT=""
ENV JENKINS_DEBUG_LEVEL=5
ENV JENKINS_ENABLE_ACCESS_LOG="false"
ENV JENKINS_PREFIX="/jenkins"
ENV JENKINS_OPTS=""
ENV LANG C.UTF-8
ENV PATH="${JAVA_HOME}/bin:${PATH}"
ARG AGENT_WORKDIR=/home/"${user}"/agent
ENV TZ=Americas/Chicago


RUN mkdir /home/${user}/.jenkins && mkdir -p "${AGENT_WORKDIR}"

VOLUME ${HOME_DIR}/.jenkins
VOLUME "${AGENT_WORKDIR}"
WORKDIR ${HOME_DIR}
ENV user=${user}
LABEL \
  org.opencontainers.image.title="Jenkins Agent" \
  org.opencontainers.image.version="${VERSION}" \
  org.opencontainers.image.url="https://www.jenkins.io/" \
  org.opencontainers.image.source="https://github.com/jenkinsci/docker-agent" \
  org.opencontainers.image.licenses="MIT"
