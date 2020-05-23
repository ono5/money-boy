# Bootstrap

* navbar-expand  
トグルボタンの表示・非表示に関わるclass

* ml-auto  
ボタンを左側に寄せる

# ボタンをちょっと上にあげるテクニック
transformのtranslateYを使うホバーじにボタン位置を少しずらすことが可能

```
.current:hover {
    background-color: #f66436;
    border-radius: 7px;
    box-shadow: 2px 5px 10px #111;
    transform: translateY(-1px);
}
```

* align-items-center  
要素内でitemを中央にできる
