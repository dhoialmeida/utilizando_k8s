# Primeira etapa
# Usa a imagem do golang para buildar o programa
FROM golang:1.13.8-alpine3.11 as builder

ADD ./src /workspace
WORKDIR /workspace
# -ldflags "-w" remove informações de debug do executável
# -extldflags "-static" garante que o executável seja linkado estáticamente
RUN CGO_ENABLED=0 go build -ldflags '-w -extldflags "-static"' main.go

# Segunda etapa
# Cria uma imagem mínima copiando o executável gerado na primeira etapa
FROM scratch

COPY --from=builder /workspace/main /workspace/main
ENTRYPOINT [ "/workspace/main" ]
