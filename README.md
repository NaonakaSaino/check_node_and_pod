## 仕様
事前にkubernetesクラスタに接続し、namespaceを指定した上でmain.goを実行すると下記のような出力が得られる。

```
qmes [admin front img]
qlr2 [admin app app bastion cronjob cronjob cronjob cronjob cronjob cronjob front img]
b8da [cronjob]
p5wz [cronjob cronjob]
```

各nodeと、nodeに属するpodの種類をそれぞれ、文字列と配列で表す。
