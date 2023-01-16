

## Soal 3

Terdapat Pesan ERROR ```Error response from daemon: dockerfile parse error line 17: unknown instruction: LISTEN
jadi EXPOSE ```
Hal ini terjadi dikarenakan sintaks `LISTEN`

Berikut untuk perbaikan dari DOCKERFILE

```
FROM golang
ADD . /go/src/github.com/telkomdev/indihome/backend
WORKDIR /go/src/github.com/telkomdev/indihome
RUN go get github.com/tools/godep
RUN godep restore
RUN go install github.com/telkomdev/indihome
ENTRYPOINT /go/bin/indihome
EXPOSE 80

```

## Soal 4 
Tujuan microservice adalah untuk memudahkan dalam maintanance ketika satu fitur microservice mati atau bermasalah tidak menganggu yang lainnya

## Soal 5
index bekerja pada db
cara kerjanya yaitu seperti pada hash map dimana jika kita mengakses index tersebut maka akan mengambil key dari index yang sudah di deklarasikan tersebut jika tidak maka db akan meloop dulu dan mencari data yang dimaksud
