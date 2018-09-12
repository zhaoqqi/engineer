### wget使用方法

#### 断点续传
```bash
wget -c -t 0 -O new_name.tar.gz https://download.libsodium.org/libsodium/releases/LATEST.tar.gz
```
-c   
支持断点续传   

-t 0   
重试次数，0代表不限次数   

-O   
把下载的文件重命名为 new_name.tar.gz   


