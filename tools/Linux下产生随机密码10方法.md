### Linux下产生随机密码10方法

1. 使用SHA来哈希日期，输出头32个字符
```bash
date +%s | sha256sum | base64 | head -c 32 ; echo
```    

2. 命令使用内嵌的/dev/urandom，只输出字符，结果取头32个
```bash
< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32};echo
```

3. 使用系统自带的openssl的随机特点来产生随机密码
```bash
openssl rand -base64 32
```

4.     
```bash
tr -cd ‘[:alnum:]‘ < /dev/urandom | fold -w30 | head -n1
```

5. 通过过滤字符命令，输出随机密码
```bash
strings /dev/urandom | grep -o ‘[[:alnum:]]’ | head -n 30 | tr -d ‘\n’; echo
```

6.     
```bash
< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c6
```

7. 使用命令dd的强大功能
```bash
dd if=/dev/urandom bs=1 count=32 2>/dev/null | base64 -w 0 | rev | cut -b 2- | rev
```

8.     
```bash
</dev/urandom  tr -dc ’12345!@#$%qwertQWERTasdfgASDFGzxcvbZXCVB’ | head -c8; echo “”
```

9. 使用randpw随时产生随机密码，可以把它放到~/.bashrc文件里面
```bash
randpw(){ < /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-16};echo;}
```

10. 最简洁的方式
```bash
date | md5sum
```    

[原文链接](http://os.51cto.com/art/201102/246360.htm, "Linux下产生随机密码10方法")   