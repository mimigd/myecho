# 使用Golang的官方鏡像作為基礎鏡像
FROM golang:1.20.7-alpine

# 設定工作目錄
WORKDIR /app

# 複製go.mod和go.sum文件到工作目錄
COPY go.mod go.sum ./

# 下載依賴
RUN go mod download

# 複製專案所有文件到工作目錄
COPY . .

# 編譯專案
RUN go build -o main .

# 設定啟動命令
CMD ["./main"]
