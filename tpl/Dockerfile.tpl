# Static go service
FROM alpine:latest

ADD bin/* /opt/###__PROJ_ORG__###/
WORKDIR /opt/###__PROJ_ORG__###
EXPOSE 9900
CMD [ "/opt/###__PROJ_ORG__###/###__PROJ_NAME__###" ]
