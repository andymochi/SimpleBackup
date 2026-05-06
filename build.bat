@echo off
chcp 65001 >nul
echo ========================================
echo    Simple Backup 多语言版编译工具 v2.0
echo ========================================
echo.

:: 检查 Go 是否安装
where go >nul 2>nul
if errorlevel 1 (
    echo ❌ 错误: 未找到 Go 编译器
    echo 请安装 Go: https://golang.org/dl/
    pause
    exit /b 1
)

:: 创建构建目录
if not exist "build" mkdir build
if not exist "build\windows" mkdir build\windows
if not exist "build\linux" mkdir build\linux
if not exist "build\macos" mkdir build\macos

set VERSION=2.0.0
set APP_NAME=simplebackup

echo 版本: %VERSION%
echo 正在编译各平台版本...
echo.

:: ===== 编译 Windows 版本 =====
echo [Windows]
echo.

:: Windows 64位
echo  编译 64位版本...
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-s -w -X main.Version=%VERSION%" -o build\windows\%APP_NAME%-windows-amd64.exe
if errorlevel 1 (
    echo  ❌ 编译失败
) else (
    echo  ✅ 成功: %APP_NAME%-windows-amd64.exe
    echo  ✅ 多语言支持: 简体中文/繁体中文/英文
)

:: Windows 32位
echo.
echo  编译 32位版本...
set GOOS=windows
set GOARCH=386
go build -ldflags "-s -w -X main.Version=%VERSION%" -o build\windows\%APP_NAME%-windows-386.exe
if errorlevel 1 (
    echo  ⚠️  编译失败 (跳过)
) else (
    echo  ✅ 成功: %APP_NAME%-windows-386.exe
)

:: ===== 编译 Linux 版本 =====
echo.
echo [Linux]
echo.

:: Linux 64位
echo  编译 64位版本...
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-s -w -X main.Version=%VERSION%" -o build\linux\%APP_NAME%-linux-amd64
if errorlevel 1 (
    echo  ⚠️  编译失败 (跳过)
) else (
    echo  ✅ 成功: %APP_NAME%-linux-amd64
    echo  ✅ 多语言支持: 简体中文/繁体中文/英文
)

:: Linux ARM64 (树莓派等)
echo.
echo  编译 ARM64版本...
set GOOS=linux
set GOARCH=arm64
go build -ldflags "-s -w -X main.Version=%VERSION%" -o build\linux\%APP_NAME%-linux-arm64
if errorlevel 1 (
    echo  ⚠️  编译失败 (跳过)
) else (
    echo  ✅ 成功: %APP_NAME%-linux-arm64
)

:: ===== 编译 macOS 版本 =====
echo.
echo [macOS]
echo.

:: macOS Intel
echo  编译 Intel版本...
set GOOS=darwin
set GOARCH=amd64
go build -ldflags "-s -w -X main.Version=%VERSION%" -o build\macos\%APP_NAME%-macos-amd64
if errorlevel 1 (
    echo  ⚠️  编译失败 (跳过)
) else (
    echo  ✅ 成功: %APP_NAME%-macos-amd64
    echo  ✅ 多语言支持: 简体中文/繁体中文/英文
)

:: macOS Apple Silicon
echo.
echo  编译 Apple Silicon版本...
set GOOS=darwin
set GOARCH=arm64
go build -ldflags "-s -w -X main.Version=%VERSION%" -o build\macos\%APP_NAME%-macos-arm64
if errorlevel 1 (
    echo  ⚠️  编译失败 (跳过)
) else (
    echo  ✅ 成功: %APP_NAME%-macos-arm64
)

:: 创建 README 文件
echo.
echo 创建说明文件...
(
echo # Simple Backup v%VERSION%
echo.
echo ## 功能特性
echo - ✅ 增量备份（智能检测文件变化）
echo - ✅ 自动多语言支持（简体中文/繁体中文/英文）
echo - ✅ 时间戳文件列表（自动生成在目标目录）
echo - ✅ 磁盘空间检查与提醒
echo - ✅ 跨平台支持（Windows/macOS/Linux）
echo - ✅ 文件大小过滤
echo - ✅ 目录排除功能
echo.
echo ## 使用方法
echo.
echo ### Windows
echo 直接运行对应的exe文件，程序会自动检测系统语言。
echo.
echo ### Linux/macOS
echo chmod +x %APP_NAME%-linux-amd64
echo ./%APP_NAME%-linux-amd64 "源目录" "目标目录"
echo.
echo ## 文件列表
echo 备份完成后会在目标目录生成格式为 backup_YYYYMMDD_HHMMSS.txt 的文件列表。
) > build\README.txt

:: 创建压缩包
echo.
echo 正在创建发布包...
cd build

:: Windows 包
where 7z >nul 2>nul
if %errorlevel% equ 0 (
    7z a -tzip %APP_NAME%-v%VERSION%-windows.zip windows\* README.txt >nul
    if exist %APP_NAME%-v%VERSION%-windows.zip (
        echo ✅ 创建: %APP_NAME%-v%VERSION%-windows.zip
    )
) else (
    echo ⚠️  跳过Windows压缩包（需要7zip）
)

:: Linux 包
tar czf %APP_NAME%-v%VERSION%-linux.tar.gz linux/* README.txt 2>nul
if exist %APP_NAME%-v%VERSION%-linux.tar.gz (
    echo ✅ 创建: %APP_NAME%-v%VERSION%-linux.tar.gz
) else (
    echo ⚠️  创建Linux包失败
)

:: macOS 包
tar czf %APP_NAME%-v%VERSION%-macos.tar.gz macos/* README.txt 2>nul
if exist %APP_NAME%-v%VERSION%-macos.tar.gz (
    echo ✅ 创建: %APP_NAME%-v%VERSION%-macos.tar.gz
) else (
    echo ⚠️  创建macOS包失败
)

cd ..

echo.
echo ========================================
echo 编译完成!
echo.
echo 输出目录结构:
echo build\
echo ├── windows\    (Windows可执行文件)
echo ├── linux\      (Linux可执行文件) 
echo ├── macos\      (macOS可执行文件)
echo ├── README.txt  (使用说明)
echo └── *.zip/*.tar.gz (发布包)
echo.
echo 功能说明:
echo 1. 自动检测系统语言（简中/繁中/英文）
echo 2. 文件列表自动生成在目标目录
echo 3. 格式: backup_年月日_时分秒.txt
echo ========================================
echo.
dir build\windows\
echo.
pause