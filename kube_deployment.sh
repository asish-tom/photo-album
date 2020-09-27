rm -rf  web-deployment.yaml \
db-claim0-persistentvolumeclaim.yaml \
db-service.yaml \
web-service.yaml \
db-deployment.yaml \
zookeeper-deployment.yaml \
zookeeper-claim0-persistentvolumeclaim.yaml \
zookeeper-claim2-persistentvolumeclaim.yaml \
zookeeper-claim1-persistentvolumeclaim.yaml \
zookeeper-service.yaml \
kafka-broker-service.yaml \
kafka-broker-deployment.yaml \
kafka-broker-claim2-persistentvolumeclaim.yaml \
kafka-broker-claim1-persistentvolumeclaim.yaml \
kafka-broker-claim0-persistentvolumeclaim.yaml \

kompose convert

kubectl create -f \
web-service.yaml,\
web-deployment.yaml,\
db-service.yaml,\
db-deployment.yaml,\
zookeeper-deployment.yaml,\
zookeeper-claim0-persistentvolumeclaim.yaml,\
zookeeper-claim2-persistentvolumeclaim.yaml,\
zookeeper-claim1-persistentvolumeclaim.yaml,\
zookeeper-service.yaml,\
kafka-broker-service.yaml,\
kafka-broker-deployment.yaml,\
kafka-broker-claim2-persistentvolumeclaim.yaml,\
kafka-broker-claim1-persistentvolumeclaim.yaml,\
kafka-broker-claim0-persistentvolumeclaim.yaml\

kubectl get po -A
# kubectl port-forward  web-75f4894c47-fph5h  9090:9090