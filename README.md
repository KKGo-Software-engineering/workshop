# workshop

สวัสดีชาวโลก 👋 ยินดีต้อนรับสู่ Software Engineering with Go ในส่วนของ Workshop

### STEP0: Fork Repository

1. Fork repository นี้ไปที่ Github ของตัวเอง
1. 1 ทีม 1 Fork 1 คนเท่านั้นพอ *เราจะเพิ่มเพื่อนใน Step ถัดๆไปไม่ต้องกังวัล*

### STEP1: 🎃 Setup AWS Credential

เพื่อให้ Access AWS ได้เราจะกำหนด Credential เข้าไปให้ Shell ของเราก่อน หรือ set ใน .bashrc, .zshrc ก็ได้

```bash
export AWS_ACCESS_KEY_ID=<KEY>
export AWS_SECRET_ACCESS_KEY=<SECRET>
```
### STEP2: 🧾 Terraform

1.ติดตั้ง Terraform ให้เรียบร้อย

1.ไปที่ `infra/terraform` จากนั้นรัน

2. สั่งเริ่มต้น Terraform

```console
terraform init
```

3. สั่งสร้าง Terraform สร้าง Resource ใน AWS

```console
terraform apply -var group_name="group-<ID>"
```

4. ตอบ `yes` กด enter แล้วรอไปกดกินข้าวก่อน

```console
Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes
```

## STEP3: 🚀 Deploy ยังไง?

เรา Deploy ด้วย ArgoCD ที่อยู่บน AWS EKS ผ่าน Terraform เพื่อให้ใช้งานได้ ต้องเตรียมของดังนี้

### STEP3.1: 🍻 Setup CI/CD

หลังจากที่ Fork repository ไปแล้ว จะให้ CI/CD ทำงานได้ จะต้องเซ็ตค่า `WORKFLOW_TOKEN` ด้วย โดยเอามาจาก PAT ของ Github จากนั้นเอาไปใส่ที่ Secret
1. สร้าง Personal Access Token ของ Github 
	1. ไปที่ [Personal Access Token](https://github.com/settings/tokens)
	1. กด `Generate new token (classic)`  *เอา classic นะ*
	1. ตั้งชื่อ Note เป็น ชื่อกลุ่มตัวเอง เช่น `group-1`
	1. กดเลือกทีละอัน ทั้งหมด ทุกอัน แล้วกด `Generate token`
	1. เสร็จแล้ว copy ไว้ *มันจะไม่แสดงอีกแล้ว copy ได้แค่ครั้งเดียว ห้ามทำหาย*
1. set ค่า `WORKFLOW_TOKEN` ใน Github ของทีม
	1. กลับไปที่ Repository workshop ของเรา ที่เรา fork มา (e.g. https://github.com/<your-account>/workshop)
	1. ไปที่ Settings > Secrets and variables > Actions 
	1. กด `New repository secret`
	1. ใส่ชื่อ Name เป็น `WORKFLOW_TOKEN` แล้วใส่ค่า Personal Access Token ที่เรา copy ไว้ใน ใส่เข้าไปในช่อง `Secret`
	1. กด `Add secret`
	1. จบแล้ว

References:

- [Generate PAT](https://www.youtube.com/watch?v=jW7tbvHSChg)
- [Create GHA Secret](https://www.youtube.com/watch?v=IuT0Ua7V4xA)

1. แก้ไข CI/CD ให้เป็นชื่อ repository ของทีม
	1. find and replace `<your-github-account>` ให้เป็นชื่อ github account ของคุณคนที่ fork มา
	1. git add -> git commit -> git push

### STEP3.2: เพิ่มสมาชิกใน Github

1. ทำการเพิ่มสมาชิกใน Github ของทีมเพื่อให้สามารถเข้าถึง Repository ได้
	1. ไปที่ Settings > Collaborators and teams > Manage access
		- กด `Add People`
		- เลือก Role เป็น `Admin` ทุกคนเลย


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
terraform destroy -var group_name="group-<ID>"
```


## 🏁 Development เริ่มยังไง?

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