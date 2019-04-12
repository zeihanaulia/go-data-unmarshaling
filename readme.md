# Data (un)marshaling in go

Apa itu marshal dan unmarshaling di go. Pada dasarnya ini adalah proses encoding dan decoding.
Misalnya, ketika kita ingin menkonversi struct di go menjadi json, maka kita bisa melakukan marshaling.
sebaliknya, ketika kita ingin menerima masukan json maka kita dapat melakukan unmarshaling

## Bagaimana cara encoding dan decoding berkerja

Encoding sebenernya kita memanggil fungsi yang mengembalikan data array byte `[]byte()`.
Decoding itu sebaliknya kita memberikan array byte sebagai parameter dan object kosong,
yang nantinya akan di isi.

Isi dari array byte bisa dari manapun, misal dari http request, dari file, apapun bisa dimasukan.