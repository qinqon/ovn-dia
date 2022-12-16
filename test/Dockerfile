FROM quay.io/fedora/fedora:37

RUN dnf install -y net-tools util-linux ovn ovn-central ovn-host openvswitch procps-ng && \
    dnf clean all

COPY populate.sh /

RUN mkdir -p /var/run/openvswitch
RUN touch /etc/default/openvswitch
