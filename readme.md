# Data (un)marshaling in go

Apa itu marshal dan unmarshaling di go. Pada dasarnya ini adalah proses encoding dan decoding.
Misalnya, ketika kita ingin menkonversi struct di go menjadi json, maka kita bisa melakukan marshaling.
sebaliknya, ketika kita ingin menerima masukan json maka kita dapat melakukan unmarshaling

## Bagaimana cara encoding dan decoding berkerja

Encoding sebenernya kita memanggil fungsi yang mengembalikan data array byte `[]byte()`.
Decoding itu sebaliknya kita memberikan array byte sebagai parameter dan object kosong,
yang nantinya akan di isi.

Isi dari array byte bisa dari manapun, misal dari http request, dari file, apapun bisa dimasukan.

## Custom encoding dan decoding

Bisa gak sih ngekustom encoding dan decoding? bisa
Pada dasarnya kita bisa nge extend dari interface MarshalJSON()

## Gimana caranya mengandle tipe data yang berbeda beda

Golang itu static type, ketika kita set type data misal ke int maka jika dimasukan string akan error
Lalu bagaimana, jika request yang kita terima itu bisa berbeda beda. Kita bisa liat lagi macam macam type di go
ada salah satu type yang bernama [interface{}](https://tour.golang.org/methods/14) tipe ini bisa menerima semuanya
Tapi akan ada kerja extra untuk memparsingnya