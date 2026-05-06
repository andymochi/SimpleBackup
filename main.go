package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// ========== 多语言支持 ==========
type Language struct {
	AppName              string
	Author               string
	Usage                string
	Examples             string
	SourceDir            string
	DestDir              string
	Options              string
	Example1             string
	Example2             string
	Example3             string
	Example4             string
	FlagVerbose          string
	FlagForce            string
	FlagList             string
	FlagExclude          string
	FlagMinSize          string
	FlagMaxSize          string
	ErrorSourceNotFound  string
	ErrorDestCreate      string
	ErrorNoSpace         string
	WarningDiskSpace     string
	PromptContinue       string
	BackupStarted        string
	CurrentTime          string
	SystemInfo           string
	DiskSpaceCheck       string
	DiskSpaceAvailable   string
	DiskSpaceNeeded      string
	DiskSpaceWarning     string
	DiskSpaceEnough      string
	BackupCancelled      string
	BackupComplete       string
	BackupFailed         string
	StatsTotalFiles      string
	StatsNewFiles        string
	StatsUpdatedFiles    string
	StatsSkippedFiles    string
	StatsFailedFiles     string
	StatsBackupSize      string
	StatsBackupTime      string
	StatsFileList        string
	StatsAllSuccess      string
	StatsSomeFailed      string
	StatusStarting       string
	StatusNewFile        string
	StatusModified       string
	StatusForced         string
	StatusSkipped        string
	StatusSkippedSize    string
	StatusSkippedError   string
	StatusCopying        string
	StatusCopyDone       string
	StatusNoChange       string
	StatusTargetError    string
	StatusHashError      string
	StatusCreateDirError string
	StatusCopyError      string
	StatusDiskFull       string
	Suggestions          string
	Suggestion1          string
	Suggestion2          string
	Suggestion3          string
	Suggestion4          string
	Suggestion5          string
}

var (
	// 简体中文
	langZH = Language{
		AppName:              "简易增量备份工具 v2.0",
		Author:               "莫迟 (喜欢足球，来自中国北京)",
		Usage:                "用法: %s [选项] <源目录> <目标目录>\n\n",
		Examples:             "示例:\n",
		SourceDir:            "源目录",
		DestDir:              "目标目录",
		Options:              "选项:\n",
		Example1:             "  %s \"C:\\我的文档\" \"D:\\备份\\文档\"\n",
		Example2:             "  %s -v \"C:\\照片\" \"E:\\备份\\照片\"\n",
		Example3:             "  %s -exclude=\".git,node_modules\" \"C:\\项目\" \"F:\\备份\\项目\"\n",
		Example4:             "  %s -list=备份列表.txt -min-size=1024 \"C:\\音乐\" \"G:\\备份\"\n",
		FlagVerbose:          "显示详细日志",
		FlagForce:            "强制重新备份所有文件",
		FlagList:             "文件列表输出路径",
		FlagExclude:          "要排除的目录，用逗号分隔",
		FlagMinSize:          "最小文件大小（字节），小于此值的文件不备份",
		FlagMaxSize:          "最大文件大小（字节），大于此值的文件不备份",
		ErrorSourceNotFound:  "❌ 错误: 源目录不存在 - %s\n",
		ErrorDestCreate:      "❌ 错误: 无法创建目标目录 - %v\n",
		ErrorNoSpace:         "❌ 错误: 磁盘空间不足",
		WarningDiskSpace:     "⚠️  警告: 可能需要 %.2f GB，但只有 %.2f GB 可用空间\n",
		PromptContinue:       "是否继续?",
		BackupStarted:        "🚀 开始备份: %s → %s\n",
		CurrentTime:          "📅 时间: %s\n",
		SystemInfo:           "💻 系统: %s/%s\n",
		DiskSpaceCheck:       "检查磁盘空间...\n",
		DiskSpaceAvailable:   "可用空间: %.2f GB",
		DiskSpaceNeeded:      "需要空间: %.2f GB",
		DiskSpaceWarning:     "⚠️  磁盘空间检查: %v\n",
		DiskSpaceEnough:      "✓ 磁盘空间: 需要 %.2f GB，可用 %.2f GB\n",
		BackupCancelled:      "备份已取消",
		BackupComplete:       "✅ 备份完成!",
		BackupFailed:         "❌ 备份失败: %v\n",
		StatsTotalFiles:      "📊 处理文件总数: %d",
		StatsNewFiles:        "🆕 新增文件: %d",
		StatsUpdatedFiles:    "🔄 更新文件: %d",
		StatsSkippedFiles:    "⏭️  跳过文件: %d",
		StatsFailedFiles:     "❌ 失败文件: %d",
		StatsBackupSize:      "💽 备份数据量: %s",
		StatsBackupTime:      "⏱️  备份耗时: %v",
		StatsFileList:        "📄 文件列表: %s",
		StatsAllSuccess:      "✅ 所有文件备份成功!",
		StatsSomeFailed:      "⚠️  有 %d 个文件备份失败，请检查日志",
		StatusStarting:       "开始备份: %s",
		StatusNewFile:        "新文件",
		StatusModified:       "文件已修改",
		StatusForced:         "强制备份",
		StatusSkipped:        "跳过",
		StatusSkippedSize:    "跳过大小限制",
		StatusSkippedError:   "跳过错误文件",
		StatusCopying:        "复制中",
		StatusCopyDone:       "复制完成",
		StatusNoChange:       "无变化",
		StatusTargetError:    "目标文件错误",
		StatusHashError:      "哈希计算失败",
		StatusCreateDirError: "创建目录失败",
		StatusCopyError:      "复制失败",
		StatusDiskFull:       "磁盘空间不足",
		Suggestions:          "建议操作:",
		Suggestion1:          "  1. 清理目标磁盘",
		Suggestion2:          "  2. 更换更大容量的存储设备",
		Suggestion3:          "  3. 使用 -exclude 参数排除大文件或目录",
		Suggestion4:          "  4. 备份到其他位置",
		Suggestion5:          "  5. 分批次备份不同目录",
	}

	// 英文 (默认)
	langEN = Language{
		AppName:              "Simple Incremental Backup v2.0",
		Author:               "Andy Mo(Like football, live in Beijing China)",
		Usage:                "Usage: %s [options] <source_dir> <dest_dir>\n\n",
		Examples:             "Examples:\n",
		SourceDir:            "Source directory",
		DestDir:              "Destination directory",
		Options:              "Options:\n",
		Example1:             "  %s \"C:\\My Documents\" \"D:\\Backup\\Docs\"\n",
		Example2:             "  %s -v \"C:\\Photos\" \"E:\\Backup\\Photos\"\n",
		Example3:             "  %s -exclude=\".git,node_modules\" \"C:\\Projects\" \"F:\\Backup\\Projects\"\n",
		Example4:             "  %s -list=backup_list.txt -min-size=1024 \"C:\\Music\" \"G:\\Backup\"\n",
		FlagVerbose:          "Show verbose output",
		FlagForce:            "Force backup all files",
		FlagList:             "File list output path",
		FlagExclude:          "Directories to exclude, comma separated",
		FlagMinSize:          "Minimum file size (bytes), skip smaller files",
		FlagMaxSize:          "Maximum file size (bytes), skip larger files",
		ErrorSourceNotFound:  "❌ Error: Source directory not found - %s\n",
		ErrorDestCreate:      "❌ Error: Cannot create destination directory - %v\n",
		ErrorNoSpace:         "❌ Error: Disk space insufficient",
		WarningDiskSpace:     "⚠️  Warning: May need %.2f GB, but only %.2f GB available\n",
		PromptContinue:       "Continue?",
		BackupStarted:        "🚀 Starting backup: %s → %s\n",
		CurrentTime:          "📅 Time: %s\n",
		SystemInfo:           "💻 System: %s/%s\n",
		DiskSpaceCheck:       "Checking disk space...\n",
		DiskSpaceAvailable:   "Available: %.2f GB",
		DiskSpaceNeeded:      "Needed: %.2f GB",
		DiskSpaceWarning:     "⚠️  Disk space check: %v\n",
		DiskSpaceEnough:      "✓ Disk space: Need %.2f GB, Available %.2f GB\n",
		BackupCancelled:      "Backup cancelled",
		BackupComplete:       "✅ Backup Complete!",
		BackupFailed:         "❌ Backup failed: %v\n",
		StatsTotalFiles:      "📊 Total files processed: %d",
		StatsNewFiles:        "🆕 New files: %d",
		StatsUpdatedFiles:    "🔄 Updated files: %d",
		StatsSkippedFiles:    "⏭️  Skipped files: %d",
		StatsFailedFiles:     "❌ Failed files: %d",
		StatsBackupSize:      "💽 Backup size: %s",
		StatsBackupTime:      "⏱️  Backup time: %v",
		StatsFileList:        "📄 File list: %s",
		StatsAllSuccess:      "✅ All files backed up successfully!",
		StatsSomeFailed:      "⚠️  %d files failed, check the log",
		StatusStarting:       "Starting backup: %s",
		StatusNewFile:        "New file",
		StatusModified:       "File modified",
		StatusForced:         "Forced backup",
		StatusSkipped:        "Skipped",
		StatusSkippedSize:    "Skipped due to size",
		StatusSkippedError:   "Skipped due to error",
		StatusCopying:        "Copying",
		StatusCopyDone:       "Copy done",
		StatusNoChange:       "No change",
		StatusTargetError:    "Target file error",
		StatusHashError:      "Hash calculation failed",
		StatusCreateDirError: "Create directory failed",
		StatusCopyError:      "Copy failed",
		StatusDiskFull:       "Disk full",
		Suggestions:          "Suggestions:",
		Suggestion1:          "  1. Clean up destination disk",
		Suggestion2:          "  2. Use larger storage device",
		Suggestion3:          "  3. Use -exclude to skip large files/dirs",
		Suggestion4:          "  4. Backup to another location",
		Suggestion5:          "  5. Backup in batches",
	}

	// 当前语言
	lang *Language
)

// 检测系统语言
func detectLanguage() *Language {
	// 通过环境变量检测
	envVars := []string{
		"LANGUAGE",
		"LANG",
		"LC_ALL",
		"LC_MESSAGES",
	}

	for _, env := range envVars {
		if lang := os.Getenv(env); lang != "" {
			if strings.Contains(strings.ToLower(lang), "zh_cn") ||
				strings.Contains(strings.ToLower(lang), "zh-cn") ||
				strings.Contains(strings.ToLower(lang), "zh.hans") {
				return &langZH
			}
		}
	}

	// 默认使用英文
	return &langEN
}

// ========== 变量定义 ==========
var (
	verbose     = flag.Bool("v", false, "")
	force       = flag.Bool("f", false, "")
	listFile    = flag.String("list", "", "自动生成时间戳文件列表，留空则自动生成")
	excludeDirs = flag.String("exclude", ".git,node_modules,temp", "")
	minSize     = flag.Int64("min-size", 0, "")
	maxSize     = flag.Int64("max-size", 0, "")
)

// ========== 主函数 ==========
func main() {
	// 检测并设置语言
	lang = detectLanguage()

	// 设置flag的使用说明
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, lang.AppName+"\n")
		fmt.Fprintf(os.Stderr, lang.Author+"\n")
		fmt.Fprintf(os.Stderr, lang.Usage, os.Args[0])
		fmt.Fprintf(os.Stderr, lang.Options)

		// 重新定义flag的帮助文本
		fmt.Fprintf(os.Stderr, "  -v\t%s\n", lang.FlagVerbose)
		fmt.Fprintf(os.Stderr, "  -f\t%s\n", lang.FlagForce)
		fmt.Fprintf(os.Stderr, "  -list string\n\t%s\n", lang.FlagList)
		fmt.Fprintf(os.Stderr, "  -exclude string\n\t%s (default \".git,node_modules,temp\")\n", lang.FlagExclude)
		fmt.Fprintf(os.Stderr, "  -min-size int\n\t%s\n", lang.FlagMinSize)
		fmt.Fprintf(os.Stderr, "  -max-size int\n\t%s\n", lang.FlagMaxSize)

		fmt.Fprintf(os.Stderr, "\n"+lang.Examples)
		fmt.Fprintf(os.Stderr, lang.Example1, os.Args[0])
		fmt.Fprintf(os.Stderr, lang.Example2, os.Args[0])
		fmt.Fprintf(os.Stderr, lang.Example3, os.Args[0])
		fmt.Fprintf(os.Stderr, lang.Example4, os.Args[0])
	}

	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	sourceDir := filepath.Clean(flag.Arg(0))
	destDir := filepath.Clean(flag.Arg(1))

	// 检查源目录
	if !isDirExists(sourceDir) {
		fmt.Printf(lang.ErrorSourceNotFound, sourceDir)
		os.Exit(1)
	}

	// 检查是否备份到自身
	if strings.HasPrefix(destDir, sourceDir) && destDir != sourceDir {
		fmt.Println("❌ Error: Cannot backup to a subdirectory of source")
		fmt.Printf("   Source: %s\n", sourceDir)
		fmt.Printf("   Dest: %s\n", destDir)
		os.Exit(1)
	}

	// 确保目标目录存在
	if err := os.MkdirAll(destDir, 0755); err != nil {
		fmt.Printf(lang.ErrorDestCreate, err)
		os.Exit(1)
	}

	// 生成时间戳文件列表名称
	if *listFile == "" {
		timestamp := time.Now().Format("backup_20060102_150405")
		*listFile = filepath.Join(destDir, timestamp+".txt")
	}

	fmt.Printf(lang.BackupStarted, sourceDir, destDir)
	fmt.Printf(lang.CurrentTime, time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf(lang.SystemInfo, runtime.GOOS, runtime.GOARCH)
	fmt.Println(strings.Repeat("=", 60))

	startTime := time.Now()

	// 执行备份
	stats, err := backupDirectory(sourceDir, destDir)
	if err != nil {
		fmt.Printf(lang.BackupFailed, err)
		os.Exit(1)
	}

	// 输出统计信息
	elapsed := time.Since(startTime)
	stats.BackupTime = elapsed

	printStats(stats, sourceDir, destDir)
}

// ========== 备份主函数 ==========
func backupDirectory(source, dest string) (*BackupStats, error) {
	stats := &BackupStats{}
	excludeList := parseExcludeList(*excludeDirs)

	// 创建文件列表
	listFile, err := os.Create(*listFile)
	if err != nil {
		return nil, fmt.Errorf("Cannot create file list: %v", err)
	}
	defer listFile.Close()

	writer := bufio.NewWriter(listFile)
	defer writer.Flush()

	// 写入文件列表头部
	writeFileListHeader(writer, source, dest)

	// 遍历源目录
	err = filepath.Walk(source, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			stats.FailedFiles++
			if *verbose {
				fmt.Printf("⚠️  %s: %s (%v)\n", lang.StatusSkippedError, filePath, err)
			}
			return nil
		}

		// 跳过目录
		if info.IsDir() {
			// 检查是否在排除列表中
			if isExcluded(filePath, source, excludeList) {
				if *verbose {
					fmt.Printf("%s: %s\n", lang.StatusSkipped, filePath)
				}
				return filepath.SkipDir
			}
			return nil
		}

		// 检查文件大小限制
		if (*minSize > 0 && info.Size() < *minSize) || (*maxSize > 0 && info.Size() > *maxSize) {
			stats.SkippedFiles++
			if *verbose {
				fmt.Printf("%s: %s (%d bytes)\n", lang.StatusSkippedSize, filePath, info.Size())
			}
			return nil
		}

		stats.TotalFiles++
		stats.TotalSize += info.Size()

		// 计算相对路径
		relPath, err := filepath.Rel(source, filePath)
		if err != nil {
			return err
		}

		// 计算目标路径
		destPath := filepath.Join(dest, relPath)
		destDir := filepath.Dir(destPath)

		// 计算文件哈希
		fileHash, err := calculateFileHash(filePath)
		if err != nil {
			stats.FailedFiles++
			writeFileListEntry(writer, relPath, info, fileHash, lang.StatusHashError)
			if *verbose {
				fmt.Printf("⚠️  %s: %s (%v)\n", lang.StatusHashError, relPath, err)
			}
			return nil
		}

		// 检查是否需要备份
		needBackup, reason := needToBackup(filePath, destPath, fileHash, info.ModTime())
		if !needBackup && !*force {
			stats.SkippedFiles++
			writeFileListEntry(writer, relPath, info, fileHash, reason)
			if *verbose {
				fmt.Printf("%s: %s (%s)\n", lang.StatusSkipped, relPath, reason)
			}
			return nil
		}

		// 创建目标目录
		if err := os.MkdirAll(destDir, 0755); err != nil {
			stats.FailedFiles++
			writeFileListEntry(writer, relPath, info, fileHash, lang.StatusCreateDirError)
			if *verbose {
				fmt.Printf("❌ %s: %s (%v)\n", lang.StatusCreateDirError, destDir, err)
			}
			return nil
		}

		// 复制文件
		if err := copyFileWithProgress(filePath, destPath, info.Size()); err != nil {
			stats.FailedFiles++
			status := lang.StatusCopyError
			if strings.Contains(err.Error(), lang.StatusDiskFull) {
				status = lang.StatusDiskFull
			}
			writeFileListEntry(writer, relPath, info, fileHash, status)

			if strings.Contains(err.Error(), lang.StatusDiskFull) {
				return fmt.Errorf("❌ %s: %v", lang.StatusDiskFull, err)
			}

			if *verbose {
				fmt.Printf("❌ %s: %s → %s (%v)\n", lang.StatusCopyError, relPath, destPath, err)
			}
			return nil
		}

		// 更新统计
		if reason == lang.StatusNewFile {
			stats.NewFiles++
		} else {
			stats.UpdatedFiles++
		}

		// 记录到文件列表
		writeFileListEntry(writer, relPath, info, fileHash, lang.StatusCopyDone)

		// 显示进度
		if *verbose {
			fmt.Printf("%s: %s → %s\n", reason, relPath, destPath)
		} else {
			// 简洁进度显示
			totalOps := stats.NewFiles + stats.UpdatedFiles + stats.FailedFiles
			if totalOps%10 == 0 {
				fmt.Print(".")
			}
			if totalOps%200 == 0 {
				fmt.Printf(" (%d)\n", totalOps)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return stats, nil
}

// 写入文件列表头部
func writeFileListHeader(writer *bufio.Writer, source, dest string) {
	writer.WriteString("# " + lang.AppName + "\n")
	writer.WriteString(fmt.Sprintf("# Generated: %s\n", time.Now().Format("2006-01-02 15:04:05")))
	writer.WriteString(fmt.Sprintf("# Source: %s\n", source))
	writer.WriteString(fmt.Sprintf("# Destination: %s\n", dest))
	writer.WriteString("# Language: " + getSystemLanguage() + "\n")
	writer.WriteString("#\n")
	writer.WriteString("# Format: Relative Path | Size (bytes) | Modified Time | MD5 Hash | Status\n")
	writer.WriteString("#" + strings.Repeat("-", 100) + "\n")
}

// 获取系统语言
func getSystemLanguage() string {
	envVars := []string{"LANGUAGE", "LANG", "LC_ALL"}
	for _, env := range envVars {
		if lang := os.Getenv(env); lang != "" {
			return lang
		}
	}
	return "en"
}

// 写入文件列表条目
func writeFileListEntry(writer *bufio.Writer, relPath string, info os.FileInfo, hash, status string) {
	writer.WriteString(fmt.Sprintf("%s | %d | %s | %s | %s\n",
		relPath, info.Size(), info.ModTime().Format("2006-01-02 15:04:05"), hash, status))
}

// ========== 辅助函数 ==========
func calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func needToBackup(source, dest, sourceHash string, sourceModTime time.Time) (bool, string) {
	if *force {
		return true, lang.StatusForced
	}

	destInfo, err := os.Stat(dest)
	if os.IsNotExist(err) {
		return true, lang.StatusNewFile
	}
	if err != nil {
		return true, lang.StatusTargetError
	}

	// 检查修改时间和大小
	if sourceModTime.After(destInfo.ModTime()) || destInfo.Size() <= 0 {
		destHash, err := calculateFileHash(dest)
		if err != nil || destHash != sourceHash {
			return true, lang.StatusModified
		}
	}

	return false, lang.StatusNoChange
}

func copyFileWithProgress(src, dst string, size int64) error {
	source, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Open source failed: %v", err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return checkDiskError(err, "Create destination failed")
	}
	defer destination.Close()

	// 设置缓冲区
	buf := make([]byte, 64*1024) // 64KB
	var total int64
	start := time.Now()

	for {
		n, err := source.Read(buf)
		if n > 0 {
			wn, werr := destination.Write(buf[:n])
			if werr != nil {
				return checkDiskError(werr, "Write failed")
			}
			total += int64(wn)

			// 显示大文件的复制进度
			if size > 10*1024*1024 && *verbose && time.Since(start) > time.Second {
				percent := float64(total) / float64(size) * 100
				fmt.Printf("\r%s: %.1f%% (%.1f/%.1f MB)",
					lang.StatusCopying, percent,
					float64(total)/(1024*1024), float64(size)/(1024*1024))
				start = time.Now()
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return checkDiskError(err, "Read failed")
		}
	}

	if size > 10*1024*1024 && *verbose {
		fmt.Println()
	}

	return destination.Sync()
}

func checkDiskError(err error, context string) error {
	errStr := strings.ToLower(err.Error())

	// 磁盘错误关键词
	diskErrors := []string{
		"no space",
		"disk full",
		"enospc",
		"not enough space",
		"quota exceeded",
		"insufficient disk space",
		"磁盘空间不足",
	}

	for _, keyword := range diskErrors {
		if strings.Contains(errStr, keyword) {
			return fmt.Errorf("%s: %v (%s)", context, err, lang.StatusDiskFull)
		}
	}

	return fmt.Errorf("%s: %v", context, err)
}

func isDirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func parseExcludeList(excludeStr string) []string {
	excludeList := strings.Split(excludeStr, ",")
	for i := range excludeList {
		excludeList[i] = strings.TrimSpace(excludeList[i])
	}
	return excludeList
}

func isExcluded(path, base string, excludeList []string) bool {
	relPath, err := filepath.Rel(base, path)
	if err != nil {
		return false
	}

	for _, exclude := range excludeList {
		if exclude == "" {
			continue
		}
		// 检查目录名
		if filepath.Base(path) == exclude {
			return true
		}
		// 检查相对路径
		if strings.Contains(relPath, exclude) {
			return true
		}
	}
	return false
}

func askForConfirmation(prompt string) bool {
	fmt.Printf("%s [y/N]: ", prompt)

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes"
}

func printStats(stats *BackupStats, source, dest string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println(lang.BackupComplete)
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("📁 %s: %s\n", lang.SourceDir, source)
	fmt.Printf("💾 %s: %s\n", lang.DestDir, dest)
	fmt.Printf(lang.StatsTotalFiles+"\n", stats.TotalFiles)
	fmt.Printf(lang.StatsNewFiles+"\n", stats.NewFiles)
	fmt.Printf(lang.StatsUpdatedFiles+"\n", stats.UpdatedFiles)
	fmt.Printf(lang.StatsSkippedFiles+"\n", stats.SkippedFiles)
	fmt.Printf(lang.StatsFailedFiles+"\n", stats.FailedFiles)

	if stats.TotalSize > 0 {
		fmt.Printf(lang.StatsBackupSize+"\n", formatFileSize(uint64(stats.TotalSize)))
	}

	fmt.Printf(lang.StatsBackupTime+"\n", stats.BackupTime.Round(time.Second))

	if stats.FailedFiles == 0 {
		fmt.Printf(lang.StatsFileList+"\n", *listFile)
		fmt.Println(lang.StatsAllSuccess)
	} else {
		fmt.Printf(lang.StatsSomeFailed+"\n", stats.FailedFiles)
	}

	fmt.Println(strings.Repeat("=", 60))
}

func formatFileSize(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// 备份统计结构
type BackupStats struct {
	TotalFiles   int
	NewFiles     int
	UpdatedFiles int
	SkippedFiles int
	FailedFiles  int
	TotalSize    int64
	BackupTime   time.Duration
}
