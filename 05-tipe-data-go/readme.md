# Tipe Data

Go memiliki beberapa jenis tipe data, diantaranya adalah tipe data string, numerik, boolean, nil.

## String
String adalah tipe data yang nilainya diapit oleh tanda petik dua (") atau dengan tanda backtik (`) dan berupa karakter atau kumpulan karakter. Perhatikan contoh berikut:

```go
package main
import "fmt"

func main() {
    var name = "Asrul harahap"
    var alamat = `
    Jl. Tanah Kusir 2
    Kec. Kebayoran Lama Selatan
    Kota Jakarta Selatan
    DKI Jakarta - 12240
    `

    fmt.Println(name)
    fmt.Println(alamat)
}
```
> N/b: Jika menggunakan tanda (`) maka hasilnya sesuai dengan penulisan.

## Numerik Non-Desimal
Tipe data numerik non-desimal Go ada beberapa jenis. Secara umum ada 2 tipe data kategori ini yang perlu diketahui.

- uint, tipe data untuk bilangan cacah (bilangan positif).
- int, tipe data untuk bilangan bulat (bilangan negatif dan positif).

Kedua tipe data di atas kemudian dibagi lagi menjadi beberapa jenis, dengan pembagian berdasarkan lebar cakupan nilainya, detailnya bisa dilihat di tabel berikut.

Tipe data | Range bilangan
:---------:|------------
uint8 |0 ↔ 255
uint16 | 0 ↔ 65535
uint32 | 0 ↔ 4294967295
uint64 | 0 ↔ 18446744073709551615
uint | bisa berupa uint32 atau uint64 (tergantung nilai)
byte | sama dengan uint8
int8 | -128 ↔ 127
int16 | -32768 ↔ 32767
int32 | -2147483648 ↔ 2147483647
int64 | -9223372036854775808 ↔ 9223372036854775807
int | bisa berupa int32 atau int64 (tergantung nilai)
rune | bisa berupa int32

Pemilihan tipe data yang baik adalah cara optimasi penggunaan memori, gunakanlah tipedata sesuai kebutuhan.

```go
package main
import "fmt"

func main() {
	var bil1 uint8 = 101
	var bil2 = -2147483648

	fmt.Println("bilangan positif", bil1)
	fmt.Println("bilangan negatif", bil2)
}
```

Variabel bil1 bertipe uint8 dengan nilai awal 101. Sedangkan variabel bil2 dideklarasikan dengan nilai awal -2147483648. Compiler go akan menentukan tipe data variabel bil2  sebagai int32 karena angka tersebut ada pada range int32.

## Numerik Desimal
Tipe data numerik desimal yang perlu diketahui ada 2, float32 dan float64. Perbedaan kedua tipe data tersebut berada range cakupan nilai desimal yang bisa ditampung. 

## bool (Boolean)
Tipe data bool adalah tipedata spesial yang hanya berisikan 2 jenis nilai, `true` dan `false`. 

```go
package main
import "fmt"

func main() {
    var benar bool = true
    fmt.Println("Benar ? ", benar)
}
```

## Tipe nil
Variabel yang isi nilainya `nil` berarti memiliki nilai kosong. Secara default semua tipe data memiliki nilai zero.

- nilai zero dari string adalah "" (string kosong).
- nilai zero dari bool adalah false.
- nilai zero dari tipe numerik non-desimal adalah 0.
- nilai zero dari tipe numerik desimal adalah 0.0.

Nilai zero berbeda dengan nil. Nil adalah nilai kosong, benar-benar kosong. nil tidak bisa digunakan pada tipe data yang sudah dibahas di atas. Ada beberapa tipe data yang bisa di-set nilainya dengan nil.