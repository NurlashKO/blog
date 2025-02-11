# Required because my GCP Host is running on legacy iptables
update-alternatives --set iptables /usr/sbin/iptables-legacy

# https://hub.docker.com/layers/openvpn/openvpn-as/latest/images/sha256-be5d1babf17a0f0bdea232d037a8e9865c664ca895f7a203d6822e415ef42808
/usr/local/openvpn_as/scripts/openvpnas --nodaemon --pidfile=/ovpn/tmp/openvpn.pid
