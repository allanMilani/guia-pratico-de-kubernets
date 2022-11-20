# Kubernetes

## O que é

“Kubernetes (K8s) é um produto Open Source utilizado para automatizar a implantação, o dimensionamento e o gerenciamento de aplicativos em conteiner.”

## De onde veio

Desenvolvida pela Google

- Borg
- Omega
- Kubernetes

## Pontos importantes

- Kubernetes é disponibilizado através de um conjunto de APIs
- Normalmente acessamos a API usando a CLI: **kubectl**
- Tudo é baseado em estado. Você configura o estado de cada objeto, dependendo do estado o kubernets toma uma ação
- Kubernetes Master
    - Kube-apiserver
    - Kube-controller-manager
    - Kube-scheduler
- Outros Node
    - Kubelet
    - Kubeproxy
- ~/.kube/config → Arquivos de configurações de credenciais de cluster de plataformas como Azure, AWS, Google Cloud, podendo mexer em um de cada vez alterando apenas o contexto
- Replicasets realizam a criação de um gerenciador de pods onde é possivel informar as imagens e a quantidade de replicas necessárias, no entanto quanto há a atualização de uma replicaset os pods não são atualizados a não ser que delete os pods para que eles sejam recriados
- Deployment → ReplicaSet → Pod
- Usando deployment é possivel contornar o erro do replicaset, pois ele deleta os pods e recria-os automaticamente
- *Service* é a porta de entrada da aplicação

## Dinâmica “superficial”

**Cluster**: Conjunto de máquinas (Nodes). Cada máquina possui uma quantidade de vCPU e Memória

**Pods**: Unidade que contém os containers provisionados. O Pod representa os processos rodando no cluster ****

## Deployment

*as replicas são utilizadas para caso um container caia o outro assume o lugar, além de balancear as requisições

*caso um cluster caia o kubernets provisiona um novo automaticamente

---

## Configurações

- Rodando local: instalar o [Kind](https://kind.sigs.k8s.io/)
- No windows com WSL2 basta habilitar a funcinalidade pela propria interface do docker desktop em outros SO é necessário instalar manualmente o kubectl
- VScode Extensões:
    - Kubernetes

---

## Kind

```bash
kind create cluster
```

Cria um novo cluster

- - -*config:* especifica um arquivo de configuração
    - *kind create cluster - -config=<ARQUIVO_COM_O_CAMINHO>*
- *- -name:* define um nome na criação do cluster
    - *kind create cluster - -name=<NOME_DO_ARQUIVO>*

```docker
kind get clusters
```

- Exibe uma lista dos clusters criados

```docker
kind delete clusters [NOME DO CLUSTER]
```

- Deleta um cluster expecifico

Configurando o Kind ([https://kind.sigs.k8s.io/docs/user/configuration/](https://kind.sigs.k8s.io/docs/user/configuration/))

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: [NOME DO CLUSTER]

nodes: 
	- role:[FUNÇÃO DO NODE control-pane | worker]
```

---

## Comandos

```docker
kubectl cluster-info --context [CONTEXTO]
```

- Seleciona o contexto de operação do client do Kubernetes (aponta o kubctl para o contexto)

```jsx
kubectl exec -it [NOME DO POD] -- bash
```

- Acessa o terminal do pod em execução

```jsx
kubectl logs [NOME DO POD]
```

- Captura o erro do pod

```yaml
kubctl top pod [NOME DO POD]
```

- Exibe o consumo de recurso do pod

```docker
kubectl get
```

Exibe uma lista do complemento passado

- *nodes:* Exibe a lista de Nodes
    - *kubectl get nodes*
- *pods:* Exibe uma lista de Pods
    - *kubectl get pods*
- *replicasets:* Exibe uma lista dos replicasets criados
    - *kubectl get replicasets*
- *deployments:* Exibe uma lista dos deployments
    - *kubectl get deployments*
- *service: Exibe uma lista de serviços*
    - *kubectl get service*
- namespaces: Exibe uma lista de namespaces
    - **********************kubectl get namespaces**********************

```docker
kubectl config 
```

Exibe uma lista de configurações

- *get-clusters*: exibe uma lista de clusters definidos no arquivo de configurações
    - *kubectl config get-clusters*
- *use-context:* define um contexto pré-existente no arquivo de configurações
    - *kubectl config use-context <NOME_DO_CONTEXTO>*
- 

```yaml
kubectl apply
```

Aplica um arquivo de configuração:

- *-f*: especifica o arquivo que será aplicado
    - *kubectl apply -f [ARQUIVO]*
- 

```yaml
kubectl port-foward [POD/SERVICE] [PORTA LOCAL]:[PORTA DO POD]
```

Redireciona postas locais para o Pod

```yaml
kubectl delete 
```

Deleta o complemento do comando

- *pod*: Deleta um Pod
    - *kubectl delete pod [NOME_DO_POD]*
- *replicaset:* Deleta um Replicaset
    - *kubectl delete replicaset [NOME_DO_REPLICASET]*
- *service:* Deleta um serviço
    - *kubectl delete service [NOME_DO_SERVICO]*

```yaml
kubectl describe
```

Detalhas as informações de um item

- *pods:* Detalhas todas as informações do pod
    - *kubectl describe pod [NOME_DO_POD]*

```yaml
kubectl rollout
```

Gerencia uma implantação

- *history:* Exibe uma lista de versões do objeto
    - *kubectl rollout history deployment [NOME_DO_OBJETO]*
- *undo:* volta ultima versão rodada
    - *kubectl rollout undo deployment [NOME_DO_OBJETO]*
- *- -to-revision*: volta para uma versão expecifica
    - *kubectl rollout undo deployment [NOME_DO_OBJETO] - -to-revision=[REVISION]*

```yaml
kubectl proxy
```

Acessa a rede do kubernetss gerando um proxy da maquina com o kubernetes

- - -*port:* Expecifica a porta de acesso do proxy
    - *kubectl proxy - -port=[PORTA_DA_MAQUINA]*
- 

Criando um POD com yml - (Um container por POD)

```yaml
apiVersion: [VERSÂO DA API]
kind: [TIPO D
O OBJETO (POD)]
metadata: 
	name: [NOME DO OBJETO]
	labels:
		app: [NOME DA LABEL]
spec: [EXPECIFICAÇÃO DO OBJETO]
	containers:
	- name: [NOME DO CONTAINER]
	  image: [IMAGEM UTILIZADA]
		volumeMounts: [MONTA UM VOLUME NO CONTAINER]
        - mountPath: "[PASTA DENTRO DO CONTAINER]"
          name: [NOME DO VOLUME CRIADO]
----------------------------- Aplicando variaveis de Ambiente ----------------------
		env: 
			- name: [NOME DA VARIAVEL DE AMBIENTE]
				value: [VALOR DA VARIAVEL DE AMBIENTE]
------------------------------- OU -------------------------------------
		envFrom:
            - configMapRef:
                name: [NOME DO CONFIGMAP]
						--------------------- Chaves Secretas --------------------------
						- secretRef:
								name: [NOME DO CONFIGMAP]
----------------------------- Criando um template para gerar replicas ------------------
	selector: 
		matchLabels:
			app: [LABEL]
	replicas: [QUANTIDADE DE REPLICAS]
	template: 
		metadata:
			name: [NOME DO OBJETO]
			labels: 
				app: [NOME DA LABEL]
			spec:
				containers:
					- name: [NOME DO CONTAINER]
					  image: [IMAGEM UTILIZADA]
----------------------------------------------------------------------------------------
	selector: [RESPONSAVEL POR FILTRAR PODS DO SERVIÇO]
		app: [LABEL]
	type: [TIPO DO SERVIÇO]
	ports:
		- name: [NOME DA PORTA]
		  port: [PORTA DO SERVIÇO]
			targetPort: [PORTA DO CONTAINER]
			protocol: [PORTOCOLO]
			
	
```

## Configmap

Um arquivo de configurações responsável por mapear variáveis de ambiente e afins

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: [NOME DO CONFIGMAP]
data:
  [KEY]: [VALUE]
```

**Injeção de ConfigMap na aplicação**

Pega os dados do configmap e cria um volume 

```yaml
volumes:
  - name: [NOME DO VOLUME]
    configMap:
	    name: [NOME DO CONFIGMAP]
      items:
        - key: [CHAVE DO CONFIG MAP]
          path: [DESTINO]
```

**Chaves secretas**

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: goserver-secret
type: Opaque
data:
  USER: [VALOR EM BASE64]
  PASSWORD: [VALOR EM BASE64]
```

## Healthcheck

Checa a saúde da aplicação e realiza ações para manter tudo ok

Configuração de checkagem com Liveness (Recria o pod)

```yaml
containers:
    - name: [NOME DO CONTAINER]
    image: "[IMAGE]"
    livenessProbe:
      httpGet:
        path: [ROTA DA APLICACAO]
        port: [PORTA DA APLICACAO]
			periodSeconds: [PERIDO PARA REALIZAR A CHECAGEM]
			failureThreshold: [QUANTIDADE DE FALHAS]
			timeoutSeconds: [TEMPO DE ESPERA]
			successThreshold: [QUANTIDADE DE TESTES]
```

Configuração de checagem com Readness (Checa se a aplicação já está no ar para mandar trafego)

```yaml
containers:
    - name: [NOME DO CONTAINER]
    image: "[IMAGE]"
    readinessProbe:
      httpGet:
        path: [ROTA DA APLICACAO]
        port: [PORTA DA APLICACAO]
			periodSeconds: [PERIDO PARA REALIZAR A CHECAGEM]
			failureThreshold: [QUANTIDADE DE FALHAS]
			timeoutSeconds: [TEMPO DE ESPERA]
			successThreshold: [QUANTIDADE DE TESTES]
			initialDelaySeconds: [TEMPO DE ESPERA PARA INICIAR A CHECAGEM]
```

Configuração de checagem do StartupProbe (Checa se a aplicação subiu para que inicie os processos de Liveness e Readness)

```yaml
containers:
    - name: [NOME DO CONTAINER]
    image: "[IMAGE]"
    startupProbe:
      httpGet:
        path: [ROTA DA APLICACAO]
        port: [PORTA DA APLICACAO]
			periodSeconds: [PERIDO PARA REALIZAR A CHECAGEM]
			failureThreshold: [QUANTIDADE DE FALHAS]
			timeoutSeconds: [TEMPO DE ESPERA]
			successThreshold: [QUANTIDADE DE TESTES]
			initialDelaySeconds: [TEMPO DE ESPERA PARA INICIAR A CHECAGEM]
```

## Metrics-server

Coleta as métricas em tempo real de cada parte do kubernetes

[https://github.com/kubernetes-sigs/metrics-server](https://github.com/kubernetes-sigs/metrics-server)

Configurações de quantidade máxima de recurso do container (Caso não exista quantidade suficiente de recursos o Pod ficará como pendente)

***Sempre evitar que a soma dos limites ultrapassem a quantidade máxima dos servidor***

```yaml
containers:
    - name: [NOME DO CONTAINER]
	    image: "[IMAGE]"
      resources:
        requests: [MINIMO NECESSARIO]
          cpu: [QUANTIDADE DE CPU]
          memory: [QUANTIDADE DE MEMORIA]
				limits: [MAXIMO DE RECURSOS]
					cpu: [QUATIDADE DE CPU]
					memory: [QUANTIDADE DE MEMORIA]
```

Configuração do HPA

```yaml
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  name: [NOME DO SERVIÇO]
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: [QUEM SERA MONITORADO]
  minReplicas: [QUANTIDADE MINIMA DE REPLICAS]
  maxReplicas: [QUANTIDADE MAXIMA DE REPLICAS]
  targetCPUUtilizationPercentage: [PERCENTUAL DE CPU PARA INICIAR A ESCALAÇÂO]
```

Comando para teste de estresse

```yaml
kubectl run -it fortio \
	--rm --image=fortio/fortio \
	-- load -qps 800 \
	-t 120s \
	-c 70 "http://goserver-service/healthz"
```

## Volumes

Stateless → não armazena estado

Statful → 

Persistent Volume aloca de maneira estática o tamanho

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: [NOME DO SERVICO]
spec:
  capacity:
    storage: [TAMANHO DO STORAGE]
  accessModes:
    - ReadWriteOnce #Ação dentro do mesmo Node
```

Persistent Volume Claim aloca de maneira dinâmica o espaço de armazenamento

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: goserver-pvc
spec:
  resources:
    requests:
      storage: 5Gi
  accessModes:
    - ReadWriteOnce #Ação dentro do mesmo Node
```

Criando o volume

```yaml
volumes:
  - name: [NOME DO VOLUME]
    persistentVolumeClam:
      claimName: [IDENTIFICADOR DO PERSISTENT VOLUME]
```

Configurando um statefulset, responsável por gerenciar os pods (ordem de up e down)

Ex: Mysql

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mysql
spec:
  replicas: 4
  serviceName: mysql-h
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
        - name: mysql
          image: mysql
          env:
            - name: MYSQL_ROOT_PASSWORD
              value: root
				  volumeMounts:
             - mountPath: /var/lib/mysql
               name: mysql-volume
---------------------- Gera um volume de maneira dinamica ------------------------
 volumeClaimTemplates:
   - metadata:
       name: mysql-volume
     spec:
       accessModes:
         - ReadWriteOnce
       resources:
         requests:
           storage: 5Gi
```

**Headless Service**

Serviço de apontamento de DNS

```yaml
apiVersion: apps/v1
kind: Service
metadata:
  name: [SERVICE NAME IGUAL AO DO ARQUIVO DE STATEFULSET]
spec:
  selector:
    app: [NOME DO SERVICO]
  ports:
    - port: [PORTA]
  clusterIP: None
```

## Ingress

Ponto único de entrada na aplicação

[https://kubernetes.io/docs/concepts/services-networking/ingress/](https://kubernetes.io/docs/concepts/services-networking/ingress/)

[https://kubernetes.github.io/ingress-nginx/](https://kubernetes.github.io/ingress-nginx/)

## Cert-manager

[https://cert-manager.io/docs/tutorials/getting-started-with-cert-manager-on-google-kubernetes-engine-using-lets-encrypt-for-ingress-ssl/](https://cert-manager.io/docs/tutorials/getting-started-with-cert-manager-on-google-kubernetes-engine-using-lets-encrypt-for-ingress-ssl/)

## Namespace

Criando um namespace

```yaml
kubectl create namespace [NOME DO NAMESPACE]
```

Adicionando um namespace

- No paramtro spec adicionar namespace
- No momento de aplicar o manifesto utilizar -ns=[NOME DO NAMESPACE]

Visualizando informações do cluster

```yaml
kubectl config view
```

Criando contextos

```yaml
kubectl config set-context dev --namespace=[NAMESPACE] --cluster=[CLUSTER NAME] --user=[USERNAME]
```

Setando contexto

```yaml
kubectl config use-context [NOME DO CONTEXTO]
```
