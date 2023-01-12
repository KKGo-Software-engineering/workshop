# workshop

สวัสดีชาวโลก 👋 ยินดีต้อนรับสู่ Software ฎngineering with Go ในส่วนของ Workshop

## 🏁 เริ่มยังไง?

ใน Repository นี้เราใช้ Makefile ในการทำงานได้ ดังนั้นสามารถสั่งรันง่ายๆ ผ่าน `make` ได้เลย

1.เริ่มต้นลองสั่งติดตั้ง dependencies ของ Go มาก่อน

```console
make install
```

2.จากนั้นสั่งรันได้เลย

```console
make dev
```

เมื่อ Server ทำงานได้ควรจะสามารถเรียกจาก [http://localhost:1323](http://localhost:1323)

## 👻 รัน Test ยังไง?

โปรเจกนี้มี 3 ระดับคือ `unit`, `integration` ระดับของ Go และ `end-to-end` ในระดับ Backend รันได้ดังนี้

### 🪛 Unit

```console
make test-unit
```

### ⚙️ Integration

```console
make test-ingegration
```

### 🤖 End-to-End

```console
make test-e2e
```

## 🚀 Deploy ยังไง?

เรา Deploy ด้วย ArgoCD ที่อยู่บน AWS EKS ผ่าน Terraform เพื่อให้ใช้งานได้ ต้องเตรียมของดังนี้

### 🎃 Setup AWS Credential

เพื่อให้ Access AWS ได้เราจะกำหนด Credential เข้าไปให้ Shell ของเราก่อน

```bash
export AWS_ACCESS_KEY_ID=<KEY>
export AWS_SECRET_ACCESS_KEY=<SECRET> 
```

### 🧾 Terraform

1.ติดตั้ง Terraform ให้เรียบร้อย

2.ไปที่ `infra/terraform` จากนั้นรัน

```console
terraform -chdir=infra/terraform apply -var group_name="group-<ID>"
```

### 🛟 Kubernetes

1.ติดตั้ง AWS CLI ให้เรียบร้อย

2.ติดตั้ง Kubernetes CLI ให้เรียบร้อย

3.เนื่องจากเราใช้ AWS EKS เป็น Kubernetes Cluster ดังนั้นเราต้องเอา Kubernetes Context จาก AWS EKS โดยสั่ง

```console
aws eks update-kubeconfig --region ap-southeast-1 --name "eks-group-<ID>"
```

4.ลองสั่ง kubectl

```console
kubectl get ns
```

ถ้าได้ผลลัพธ์ประมาณนี้เป็นอันใช้ได้

```console
NAME              STATUS   AGE
default           Active   3d
kube-system       Active   3d
kube-public       Active   3d
kube-node-lease   Active   3d
```

### 💺 ArgoCD

1.รันคำสั่งสร้าง Namespace

```console
kubectl create namespace argocd
```

2.รันคำสั่งติดตั้ง ArgoCD

```console
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

3.หารหัสผ่านของ `admin` ไว้ก่อน

```console
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

4.Forword Port เพื่อใช้งาน [http://localhost:8080](http://localhost:8080)

```console
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

### 💣 ใช้ AWS เสร็จแล้วอย่าลืม Destroy ทิ้งน๊า

```console
terraform -chdir=infra/terraform destroy -var group_name="group-<ID>"
```
