#!/bin/bash

cat << EOF | kubectl apply --namespace kafka -f -
apiVersion: kafka.banzaicloud.io/v1alpha1
kind: KafkaTopic
metadata:
  name: my-topic
spec:
  clusterRef:
    name: kafka
  name: my-topic
  partitions: 1
  replicationFactor: 1
EOF