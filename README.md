## 仕様
事前にkubernetesクラスタに接続し、namespaceを指定した上でmain.goを実行すると下記のような出力が得られる。

```
b8da [admin app bastion front]
p5wz [admin app img]
qlr2 [front img]
```

各nodeと、nodeに属するpod（Runningのみ）の種類をそれぞれ、文字列と配列で表す。

```
node名 [pod　pod...]
```
