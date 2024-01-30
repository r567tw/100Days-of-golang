# 使用官方 Go 映像作為構建環境
FROM golang:1.19 as builder

# 設置工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下載依賴項
RUN go mod download

# 複製源代碼
COPY . .

# 構建應用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用 alpine 映像以最小化最終映像的大小
FROM alpine:latest

# 為應用創建一個目錄
WORKDIR /root/

# 從構建階段映像中複製構建的應用
COPY --from=builder /app/main .

# 執行應用
CMD ["./main"]
