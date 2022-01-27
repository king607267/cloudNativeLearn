#!/bin/bash

wget https://raw.githubusercontent.com/cncamp/101/master/module4/envoy.yaml
kubectl create configmap envoy-config --from-file=envoy.yaml

wget https://raw.githubusercontent.com/cncamp/101/master/module4/envoy-deploy.yaml
kubectl create -f envoy-deploy.yaml

podIP=$(kubectl get po -oyaml --selector run=envoy | grep -e '^    podIP:' | awk '{print $2}')

curl "$podIP":10000

#修改端口到10001
kubectl edit cm envoy-config

podName=$(kubectl get po -oyaml --selector run=envoy | grep -e '^    name: ' | awk '{print $2}')

#重启pod
kubectl scale deploy "$podName" --replicas=0

kubectl scale deploy "$podName" --replicas=1

curl "$podIP":10001

kubectl delete deployment envoy







