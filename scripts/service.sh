#!/bin/bash

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 应用信息
APP_NAME="min-dsp"
APP_BIN="bin/${APP_NAME}"
PID_FILE="bin/${APP_NAME}.pid"
CONFIG_FILE="config/config.yaml"
LOG_FILE="logs/app.log"

# 检查必要的目录
check_directories() {
    local dirs=("bin" "logs" "config")
    for dir in "${dirs[@]}"; do
        if [ ! -d "$dir" ]; then
            mkdir -p "$dir"
            echo -e "${YELLOW}创建目录: $dir${NC}"
        fi
    done
}

# 检查环境
check_environment() {
    # 检查 Go 环境
    if ! command -v go &> /dev/null; then
        echo -e "${RED}错误: 未找到 Go 环境${NC}"
        exit 1
    fi

    # 检查配置文件
    if [ ! -f "$CONFIG_FILE" ]; then
        echo -e "${RED}错误: 配置文件不存在 ($CONFIG_FILE)${NC}"
        exit 1
    fi

    # 检查必要的目录
    check_directories
}

# 获取服务状态
get_status() {
    if [ -f "$PID_FILE" ]; then
        local pid=$(cat "$PID_FILE")
        if ps -p "$pid" > /dev/null 2>&1; then
            echo -e "${GREEN}服务运行中 (PID: $pid)${NC}"
            return 0
        else
            echo -e "${YELLOW}服务未运行 (PID文件存在但进程不存在)${NC}"
            rm -f "$PID_FILE"
            return 1
        fi
    else
        echo -e "${YELLOW}服务未运行${NC}"
        return 1
    fi
}

# 启动服务
start_service() {
    echo -e "${YELLOW}正在启动 $APP_NAME 服务...${NC}"
    
    # 检查环境
    check_environment

    # 检查服务是否已运行
    if get_status > /dev/null; then
        echo -e "${RED}错误: 服务已经在运行${NC}"
        exit 1
    fi

    # 编译项目
    echo -e "${YELLOW}正在编译项目...${NC}"
    if ! go build -o "$APP_BIN"; then
        echo -e "${RED}错误: 编译失败${NC}"
        exit 1
    fi

    # 启动服务
    echo -e "${YELLOW}正在启动服务...${NC}"
    nohup "$APP_BIN" > "$LOG_FILE" 2>&1 &
    local pid=$!
    
    # 保存PID
    echo $pid > "$PID_FILE"
    
    # 等待服务启动
    sleep 2
    if ps -p "$pid" > /dev/null 2>&1; then
        echo -e "${GREEN}服务已启动 (PID: $pid)${NC}"
        echo -e "${GREEN}日志文件: $LOG_FILE${NC}"
    else
        echo -e "${RED}错误: 服务启动失败${NC}"
        rm -f "$PID_FILE"
        exit 1
    fi
}

# 停止服务
stop_service() {
    echo -e "${YELLOW}正在停止 $APP_NAME 服务...${NC}"
    
    if [ -f "$PID_FILE" ]; then
        local pid=$(cat "$PID_FILE")
        if ps -p "$pid" > /dev/null 2>&1; then
            echo -e "${YELLOW}正在停止进程 (PID: $pid)...${NC}"
            kill "$pid"
            
            # 等待进程结束
            local count=0
            while ps -p "$pid" > /dev/null 2>&1; do
                sleep 1
                count=$((count + 1))
                if [ $count -ge 10 ]; then
                    echo -e "${RED}服务未能正常停止，强制终止...${NC}"
                    kill -9 "$pid"
                    break
                fi
            done
        fi
        rm -f "$PID_FILE"
        echo -e "${GREEN}服务已停止${NC}"
    else
        echo -e "${YELLOW}服务未运行${NC}"
    fi
}

# 重启服务
restart_service() {
    stop_service
    sleep 2
    start_service
}

# 显示帮助信息
show_help() {
    echo "用法: $0 {start|stop|restart|status}"
    echo
    echo "命令:"
    echo "  start   - 启动服务"
    echo "  stop    - 停止服务"
    echo "  restart - 重启服务"
    echo "  status  - 查看服务状态"
    echo "  help    - 显示此帮助信息"
}

# 主函数
main() {
    case "$1" in
        start)
            start_service
            ;;
        stop)
            stop_service
            ;;
        restart)
            restart_service
            ;;
        status)
            get_status
            ;;
        help|--help|-h)
            show_help
            ;;
        *)
            show_help
            exit 1
            ;;
    esac
}

# 执行主函数
main "$@" 