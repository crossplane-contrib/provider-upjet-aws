FROM BASEIMAGE
RUN apk --no-cache add ca-certificates bash

ARG ARCH
ARG TINI_VERSION
ENV USER_ID=1001

# Setup Terraform environment

## Provider-dependent configuration
ARG CROSSPLANE_PROVIDER_VERSION
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME
ARG TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX
## End of - Provider-dependent configuration

ENV PLUGIN_DIR /terraform/provider-mirror/registry.terraform.io/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/linux_${ARCH}
ENV TF_CLI_CONFIG_FILE /terraform/.terraformrc
ENV TF_FORK 0

RUN mkdir -p ${PLUGIN_DIR}

ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip /tmp
ADD https://releases.hashicorp.com/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}/${TERRAFORM_PROVIDER_VERSION}/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_linux_${ARCH}.zip /tmp
ADD terraformrc.hcl ${TF_CLI_CONFIG_FILE}

RUN unzip /tmp/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip -d /usr/local/bin \
  && chmod +x /usr/local/bin/terraform \
  && rm /tmp/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip \
  && unzip /tmp/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_linux_${ARCH}.zip -d ${PLUGIN_DIR} \
  && chmod +x ${PLUGIN_DIR}/* \
  && rm /tmp/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_linux_${ARCH}.zip \
  && chown -R ${USER_ID}:${USER_ID} /terraform
# End of - Setup Terraform environment

ADD provider /usr/local/bin/crossplane-provider

# Provider controller needs these environment variable at runtime
ENV TERRAFORM_VERSION ${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE ${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION ${TERRAFORM_PROVIDER_VERSION}
# ENV TERRAFORM_NATIVE_PROVIDER_PATH ${PLUGIN_DIR}/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_v${TERRAFORM_PROVIDER_VERSION}_x5
ENV TF_APPEND_USER_AGENT crossplane-provider-aws/${CROSSPLANE_PROVIDER_VERSION}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["crossplane-provider"]
