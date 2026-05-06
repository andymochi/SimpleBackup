# Simple Backup v2.0.0

## Functional features
- ✅ Incremental backup, the first time all will be backed up, and only those updated files will be backed up each time in the future
- ✅ Automatic multilingual support (Chinese Simplified/Traditional Chinese/English)
- ✅ Cross-platform support (Windows/macOS/Linux)
- ✅ A backup record is generated for each backup, making it easy to find the file name of the backup

## How to use

### Windows
Run the corresponding exe file directly, and the program will automatically detect the system language.

### Linux/macOS
chmod +x simplebackup-linux-amd64
./simplebackup-linux-amd64 "source directory" "destination directory"

### File List
After the backup is complete, a list of files in the format backup_YYYYMMDD_HHMMSS.txt is generated in the target directory.

## 功能特性
- ✅ 增量备份，第一次会全部备份，以后每次只备份那些更新的文件
- ✅ 自动多语言支持（简体中文/繁体中文/英文）
- ✅ 跨平台支持（Windows/macOS/Linux）
- ✅ 每次备份会生成备份记录，便于查找备份的文件名称

## 使用方法

### Windows
直接运行对应的exe文件，程序会自动检测系统语言。

### Linux/macOS
chmod +x simplebackup-linux-amd64
./simplebackup-linux-amd64 "源目录" "目标目录"

### 文件列表
备份完成后会在目标目录生成格式为 backup_YYYYMMDD_HHMMSS.txt 的文件列表。


## English usage
simplebackup-windows-amd64.exe

Simple Incremental Backup v2.0

Author: Andy Mo(Like football, live in Beijing China)
Usage: simplebackup-windows-amd64.exe [options] <source_dir> <dest_dir>

Options:
  -v    Show verbose output
  -f    Force backup all files
  -list string
        File list output path
  -exclude string
        Directories to exclude, comma separated (default ".git,node_modules,temp")
  -min-size int
        Minimum file size (bytes), skip smaller files
  -max-size int
        Maximum file size (bytes), skip larger files

Examples:
  simplebackup-windows-amd64.exe "C:\My Documents" "D:\Backup\Docs"
  simplebackup-windows-amd64.exe -v "C:\Photos" "E:\Backup\Photos"
  simplebackup-windows-amd64.exe -exclude=".git,node_modules" "C:\Projects" "F:\Backup\Projects"
  simplebackup-windows-amd64.exe -list=backup_list.txt -min-size=1024 "C:\Music" "G:\Backup"


## 中文用法
./simplebackup-macos-amd64 

简易增量备份工具 v2.0

作者：莫迟 (喜欢足球，来自中国北京)

用法: ./simplebackup-macos-amd64 [选项] <源目录> <目标目录>


选项:

  -v 显示详细日志

  -f 强制重新备份所有文件

  -list string

文件列表输出路径

  -exclude string

要排除的目录，用逗号分隔 (default ".git,node_modules,temp")

  -min-size int

最小文件大小（字节），小于此值的文件不备份

  -max-size int

最大文件大小（字节），大于此值的文件不备份


示例:

  ./simplebackup-macos-amd64 "C:\我的文档" "D:\备份\文档"

  ./simplebackup-macos-amd64 -v "C:\照片" "E:\备份\照片"

  ./simplebackup-macos-amd64 -exclude=".git,node_modules" "C:\项目" "F:\备份\项目"

  ./simplebackup-macos-amd64 -list=备份列表.txt -min-size=1024 "C:\音乐" "G:\备份"


## Contact me 联系我
有问题，请关注微信公众号：太酷学习笔记，并发私信，谢谢！祝备份顺利！