package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"tamiyochi-backend/entity"
// 	"time"

// 	"github.com/gocarina/gocsv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// Entry defines both the CSV layout and database schema
type Entry struct {
    // gorm.Model

    Field1 string `csv:"title"`
    Field2 string `csv:"synopsis"`
    Field3 string `csv:"start_date"`
    Field4 string `csv:"score"`
    Field5 string `csv:"scored_by"`
    Field6 string `csv:"members"`
    Field7 string `csv:"main_picture"`
    Field8 string `csv:"serializations"`
    Field9 int    `csv:"manga_id"`
    Field10 string `csv:"genres"`
    Field11 string `csv:"authors"`
}

// type Iot struct {
//     Id              int             `json:"id"`
//     FirstName       string          `json:"first_name"`
//     LastName        string          `json:"last_name"` // RawMessage here! (not a string)
//     Role            string          `json:"role"`
// }
 
func main() {
// 	komentar := []string{
// 		"Manga ini memiliki plot yang sangat menarik dan penuh dengan twist yang tak terduga! Saya selalu tidak bisa berhenti membacanya setiap kali ada rilis baru. Setiap karakter memiliki latar belakang yang kompleks dan menarik, dan saya sangat terhubung dengan perjuangan mereka. Grafiknya luar biasa indah, dengan setiap panel yang dirancang dengan hati-hati dan detail yang luar biasa. Alur ceritanya terasa alami dan sangat memikat, tetapi terkadang juga memberikan kejutan dan momen yang membuat saya tercengang. Manga ini benar-benar luar biasa dan saya merekomendasikannya kepada semua penggemar manga!",
// 		"Grafiknya benar-benar menakjubkan! Setiap halaman seperti karya seni yang hidup. Saya terpesona dengan detailnya, mulai dari desain karakter hingga latar belakang yang penuh dengan tekstur. Penggambaran adegan pertempuran sangat dinamis dan penuh dengan aksi yang epik. Plotnya juga sangat menarik dan terus memperluas dunia yang dibangun oleh manga ini. Saya tidak bisa cukup memuji betapa bagusnya kualitas grafisnya.",
// 		"Saya sangat terhubung dengan karakter utamanya. Mereka memiliki kepribadian yang kompleks dan mendalam, dan perkembangan mereka sepanjang cerita benar-benar menarik untuk diikuti. Saya merasa seolah-olah saya benar-benar mengenal mereka dan hidup bersama mereka. Tidak hanya karakter utama, tetapi juga karakter pendukung memiliki keunikan mereka sendiri dan berkontribusi pada dinamika cerita. Saya sangat investasi emosional dengan manga ini dan selalu ingin tahu apa yang akan terjadi selanjutnya.",
// 		"Alur ceritanya sangat menarik, tetapi juga penuh dengan twist dan plot twist yang tak terduga. Saya suka bagaimana setiap elemen cerita saling terkait dan mengungkapkan misteri yang lebih besar. Meskipun ada beberapa bagian yang agak lambat, secara keseluruhan ceritanya sangat memikat. Tidak jarang saya menemukan diri saya terjebak dalam teka-teki yang disajikan oleh manga ini, dan itulah yang membuatnya begitu menghibur. Meskipun ada sedikit kebingungan di sepanjang jalan, semua terhubung dengan baik di akhir.",
// 		"Manga ini adalah salah satu favorit saya sepanjang masa! Saya terpesona dengan cara cerita ini menggabungkan elemen fantasi dan realitas dengan indahnya. Setiap halaman terasa seperti petualangan yang menegangkan, dan saya tidak pernah bosan melihatnya. Saya merekomendasikannya kepada semua orang yang mencari pengalaman membaca yang unik dan mendebarkan. Saya tidak sabar untuk melihat apa yang ditawarkan manga ini di masa depan!",
// 		"Endingnya sangat memuaskan! Semua benang plot terikat dengan baik dan setiap pertanyaan terjawab dengan baik. Saya sangat menghargai cara cerita ini berakhir dengan cara yang memuaskan dan tidak meninggalkan kekosongan. Tidak ada yang lebih menyenangkan daripada melihat semua elemen cerita berpadu menjadi satu kesatuan yang utuh. Manga ini memiliki salah satu ending terbaik yang pernah saya baca.",
// 		"Tokoh antagonisnya benar-benar menakutkan! Setiap kali muncul di halaman, saya merinding dan merasa gugup. Penulis dengan cerdik membangun karakter antagonis yang kuat dan misterius, dan setiap pertemuan dengan tokoh ini selalu meninggalkan kesan yang mendalam. Saya menghargai betapa kompleksnya motivasi dan tujuan antagonis ini, dan itu membuat pertarungan antara dia dan karakter utama semakin menegangkan.",
// 		"Saya tidak sabar menunggu kelanjutan cerita manga ini! Setiap bab mengakhiri dengan cliffhanger yang menarik, dan saya selalu ingin tahu apa yang akan terjadi selanjutnya. Manga ini memiliki kemampuan yang luar biasa untuk membuat saya terus kembali untuk bab berikutnya. Saya sangat terkesan dengan kemampuan penulis untuk menjaga ketegangan cerita dan membuat saya penasaran. Tidak diragukan lagi, saya akan terus mengikuti manga ini dengan antusiasme yang tinggi.",
// 		"Pertempuran aksi dalam manga ini begitu epik dan spektakuler! Setiap adegan pertarungan penuh dengan kecepatan dan kekuatan yang luar biasa. Saya selalu terkagum-kagum melihat bagaimana ilustrator menghidupkan gerakan dan aksi karakter-karakter ini di halaman-halaman manga. Setiap adegan benar-benar menegangkan dan saya tidak pernah bosan melihatnya. Jika Anda menyukai pertempuran yang intens, manga ini adalah pilihan yang tepat.",
// 		"Karakter-karakter dalam manga ini begitu kompleks dan memiliki perkembangan yang kuat. Mereka menghadapi berbagai tantangan dan rintangan yang menguji kepribadian mereka. Saya mengagumi bagaimana penulis berhasil menggambarkan perubahan karakter dengan cara yang realistis dan terasa alami. Ini membuat cerita menjadi lebih dalam dan menarik. Saya merasa seolah-olah saya tumbuh bersama karakter-karakter ini, dan itu adalah pengalaman yang sangat berharga.",
// 		"Manga ini memiliki dunia yang sangat kaya dan mendalam. Setiap detail dalam dunia yang dibangun oleh penulis terasa terpikirkan dengan baik. Mulai dari kebudayaan hingga sistem kekuasaan, semuanya memiliki fondasi yang kuat. Saya selalu menikmati menjelajahi dunia ini dan menemukan nuansa dan keunikan yang berbeda-beda. Dunia ini benar-benar hidup dan itu menambahkan dimensi baru pada cerita.",
// 		"Dialog dalam manga ini sangat tajam dan cerdas. Setiap percakapan antara karakter terasa alami dan terasa memiliki tujuan yang jelas. Saya menghargai betapa baiknya penulis dalam menggambarkan kepribadian dan suara unik setiap karakter melalui dialog mereka. Ini membuat cerita menjadi lebih hidup dan membuat saya terhubung dengan karakter secara lebih mendalam.",
// 		"Manga ini penuh dengan kejutan dan twist yang tak terduga. Setiap kali saya pikir saya tahu apa yang akan terjadi, cerita ini selalu berhasil mengubah arah dan membuat saya tercengang. Saya sangat menghargai kejelian penulis dalam mengatur plot dan menjaga ketegangan cerita. Tidak ada yang lebih menarik daripada membaca manga ini dan menemukan bahwa segalanya tidak seperti yang terlihat.",
// 		"Tokoh protagonisnya adalah salah satu karakter yang paling inspiratif yang pernah saya temui dalam manga. Saya terpesona dengan dedikasi, kekuatan, dan semangatnya yang menginspirasi. Saya merasa termotivasi untuk menghadapi tantangan saya sendiri setelah membaca perjuangannya. Saya yakin bahwa karakter ini akan menjadi panutan bagi banyak pembaca lainnya.",
// 		"Salah satu hal yang membuat manga ini begitu istimewa adalah hubungan yang kuat antara karakter-karakternya. Mereka memiliki ikatan yang mendalam, baik itu persahabatan, cinta, atau kekeluargaan. Saya merasa emosi yang mereka alami dan bisa merasakan ikatan mereka dengan cara yang sangat kuat. Hubungan ini memberikan dimensi yang lebih dalam pada cerita dan membuat saya benar-benar terhubung dengan karakter-karakter tersebut.",
// 		"Kejutan-kejutan dalam manga ini benar-benar tak terduga! Saya selalu terkejut dengan plot twist yang penulis masukkan ke dalam cerita. Momen-momen itu membuat saya terpukau dan membuat saya ingin terus membaca untuk mengetahui apa yang akan terjadi selanjutnya. Kejutan-kejutan ini menjaga cerita tetap segar dan menghibur sepanjang perjalanan.",
// 		"Cerita dalam manga ini benar-benar menggugah emosi. Saya tertawa, menangis, marah, dan merasa terinspirasi saat membacanya. Penulis memiliki cara yang luar biasa dalam menggambarkan emosi dan membuat pembaca terhubung dengan cerita secara emosional. Saya merasa terbawa dalam perasaan dan pengalaman karakter-karakter ini, dan itu adalah tanda kualitas yang luar biasa.",
// 		"Manga ini memiliki twist yang luar biasa di setiap akhir babnya. Saya selalu terkejut dengan arah cerita yang diambil dan tidak pernah bisa menebak apa yang akan terjadi selanjutnya. Twist ini membuat cerita menjadi lebih menarik dan memastikan bahwa saya tidak pernah merasa bosan saat membacanya. Saya mengagumi kejelian penulis dalam membangun ketegangan dan mengungkapkan kejutan-kejutan yang menggetarkan.",
// 		"Tokoh-tokoh dalam manga ini memiliki desain karakter yang unik dan menarik. Setiap karakter memiliki ciri khasnya sendiri, baik itu dalam penampilan maupun kepribadian. Saya selalu menantikan penampilan karakter baru dan menyukai cara mereka ditampilkan dalam manga ini. Desain karakter yang kuat menjadi tambahan yang luar biasa pada kualitas visual manga ini.",
// 		"Cerita dalam manga ini tidak hanya menawarkan aksi dan pertarungan yang seru, tetapi juga menghadirkan momen-momen yang sangat emosional. Saya tersentuh dengan cara penulis menggambarkan hubungan antara karakter-karakter dan menghadapi konflik emosional mereka. Momen-momen emosional ini membuat cerita lebih mendalam dan memberikan dampak yang kuat.",
// 		"Manga ini berhasil menggabungkan genre yang berbeda dengan harmonis. Saya sangat menghargai cara penulis menggabungkan elemen-elemen fantasi, aksi, dan drama dengan baik. Tidak ada pertentangan antara genre-genre ini, malah mereka saling melengkapi dan membuat cerita menjadi lebih kaya dan menarik. Jika Anda menyukai campuran genre, manga ini adalah pilihan yang sempurna.",
// 		"Manga ini memiliki pesan yang kuat dan menginspirasi. Ceritanya mampu menyampaikan nilai-nilai yang mendalam dan mengajarkan pembaca tentang keberanian, persahabatan, pengorbanan, dan banyak lagi. Saya merasa terinspirasi setelah membaca manga ini dan merasa bahwa cerita ini memberikan pengaruh positif pada hidup saya. Pesan yang disampaikan dalam manga ini tidak akan pernah saya lupakan.",
// 		"Seni dalam manga ini adalah karya seni yang luar biasa! Setiap panel penuh dengan detail dan ekspresi yang kuat. Desain karakter yang indah dan latar belakang yang menakjubkan membuat setiap halaman menjadi sepotong seni yang memikat. Saya sangat menghargai keindahan visual yang disajikan oleh ilustrator dalam manga ini. Seni ini benar-benar meningkatkan pengalaman membaca saya.",
// 		"Cerita dalam manga ini memiliki tingkat ketegangan yang sangat tinggi. Setiap bab penuh dengan aksi dan kejutan yang membuat saya tegang dan tidak bisa berhenti membacanya. Saya sering menemukan diri saya menahan napas saat membaca adegan-adegan kritis. Ketegangan cerita ini membuat saya terus terjaga dan mengikuti setiap perkembangan dengan antusiasme.",
// 		"Manga ini memiliki karakter-karakter wanita yang kuat dan inspiratif. Mereka adalah tokoh-tokoh yang tidak hanya memiliki kekuatan fisik, tetapi juga kekuatan mental dan emosional. Saya merasa terinspirasi oleh keberanian, ketangguhan, dan determinasi mereka. Karakter-karakter wanita ini membuktikan bahwa perempuan dapat menjadi pahlawan sejati dan memiliki peran yang signifikan dalam cerita.",
// 	}
    

// //     // Open the CSV file for reading
//     file, err := os.Open("manga_mal.csv")
//     if err != nil {
//         panic(err)
//     }
//     defer file.Close()

//     var entries []Entry
//     err = gocsv.Unmarshal(file, &entries)
//     if err != nil {
//         panic(err)
//     }

//     dsn := fmt.Sprintf("host=localhost user=postgres password=rencist dbname=tamiyochi port=5432 TimeZone=Asia/Jakarta")
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(db)
// 		panic(err)
// 	}

//     for _, entry := range entries {
//         for i := 0; i < rand.Intn(20 - 5) + 5; i++ {
//             users := entity.User{}
//             db.Raw("SELECT * FROM users ORDER BY RANDOM() LIMIT 1").Scan(&users)
//             komentar := entity.Komentar{
//                 ID: uuid.New(),
//                 Isi: komentar[rand.Intn(24)],
//                 UserID: users.ID,
//                 SeriID: entry.Field9,
//             }
//             db.Create(&komentar)
//         }
//     }

    // nicknames := make([]string, 100)
    // emails := make([]string, 100)
    // passwords := make([]string, 100)
	// for i := 0; i < 100; i++ {
	// 	nicknames[i], emails[i], passwords[i] = generateUniqueNickname()
    //     kabupaten := entity.Kabupaten{}
    //     db.Raw("SELECT * FROM kabupatens ORDER BY RANDOM() LIMIT 1").Scan(&kabupaten)
    //     user := entity.User{
    //         ID: uuid.New(),
    //         Nama: nicknames[i],
    //         Email: emails[i],
    //         Password: passwords[i],
    //         NoTelp: "085" + strconv.Itoa(rand.Intn(999999999 - 100000000) + 100000000),
    //         KabupatenID: kabupaten.ID,
    //         Alamat: kecamatan[rand.Intn(99)],
    //         Peran: "user",
    //     }
    //     db.Create(&user)
    //     // fmt.Println(nicknames[i], " | ", emails[i], " | ", passwords[i], " | ", "085" + strconv.Itoa(rand.Intn(999999999 - 100000000) + 100000000), " | ", kecamatan[rand.Intn(99)], " | ", kabupaten.ID);
	// }

    // for _, entry := range entries {
    //     komentarCount := rand.Intn(14)
    //     fmt.Println(komentarCount, " | ", entry.Field9)
    // }
    
//     arrHarga := [8]int{18000, 20000, 15000, 24000, 22000, 16000, 25000, 30000}
//     idCount := 1;
//     for _, entry := range entries {
//         var manga entity.Manga

//         rand.Seed(time.Now().UnixNano())
//         mangaCount := 1 + rand.Intn(8-1+1)
//         for i := 1; i < mangaCount+1; i++ {
//             manga.ID = idCount
//             idCount++

//             manga.SeriID = entry.Field9
//             manga.Volume = i

//             rand.Seed(time.Now().UnixNano())
//             volumeCount := 1 + rand.Intn(5-1+1)
//             manga.JumlahTersedia = volumeCount

//             rand.Seed(time.Now().UnixNano())
//             hargaCount := 0 + rand.Intn(7-0+1)
//             manga.HargaSewa = arrHarga[hargaCount]

//             db.Create(&manga)
//         }
//     }

//     // for _, entry := range entries {
//     //     // mangaID := entries[0].Field9
//     //     genre := entry.Field10
//     //     genreArray := strings.Split(genre, ",")
//     //     for _, lmao := range genreArray {
//     //         var genre entity.Genre
//     //         if string(lmao[1]) == "]" {
//     //             rand.Seed(time.Now().UnixNano())
//     //             n := 1 + rand.Intn(225-1+1)
//     //             genre.ID = n
//     //         } else {
//     //             lmao2 := lmao[2:]
//     //             lmao3 := lmao2[:len(lmao2)-1]
//     //             lmaoFinal := strings.Replace(lmao3, "'", "", 1)

//     //             db.Where("nama = ?", lmaoFinal).First(&genre)
//     //         }
//     //         var seriGenre = entity.SeriGenre{
//     //             SeriID: entry.Field9,
//     //             GenreID: genre.ID,
//     //         }
//     //         db.Create(&seriGenre)
//     //     }

//     //     // rand.Seed(time.Now().UnixNano())
//     //     // n := 1 + rand.Intn(225-1+1)
//     //     // if seri.PenerbitID == 31 {
//     //     //     seri.PenerbitID = n
//     //     // }

//     //     // var seriPenerbit = entity.P{
//     //     //     SeriID: seri.ID,
//     //     //     PenerbitID: seri.PenerbitID,
//     //     // }
//     //     // db.Create(&seri)
//     // }

//     // for _, entry := range entries {
//     //     var penerbit = entity.Penerbit{}
//     //     db.Where("nama = ?", strings.Replace(strings.Replace(strings.Replace(entry.Field8, "[", "", 1), "]", "", 1), "'", "", 2)).First(&penerbit)
//     //     var seri = entity.Seri{
//     //         ID: entry.Field9,
//     //         Judul: entry.Field1,
//     //         Sinopsis: entry.Field2,
//     //         TahunTerbit: entry.Field3,
//     //         Skor: entry.Field4,
//     //         TotalPenilai: entry.Field5,
//     //         TotalPembaca: entry.Field6,
//     //         Foto: entry.Field7,
//     //         PenerbitID: penerbit.ID,
//     //     }

//     //     rand.Seed(time.Now().UnixNano())
//     //     n := 1 + rand.Intn(225-1+1)
//     //     if seri.PenerbitID == 31 {
//     //         seri.PenerbitID = n
//     //     }
//     //     db.Create(&seri)
//     // }

//     // fmt.Println(entries[0].Field1 + " | " + entries[0].Field2 + " | " + entries[0].Field3 + " | " + entries[0].Field4 + " | " + entries[0].Field5 + " | " + entries[0].Field6 + " | " + entries[0].Field7 + " | " + entries[0].Field8)
//     // for _, entry := range entries {

//     //     // if strings.Replace(strings.Replace(strings.Replace(entry.Field8, "[", "", 1), "]", "", 1), "'", "", 2) == "" {
//     //     //     fmt.Println(strings.Replace(strings.Replace(strings.Replace(entry.Field8, "[", "", 1), "]", "", 1), "'", "", 2))
//     //     // }

//     // }

//     // var iot Iot
//     // // fmt.Println(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[0].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 10), "},")[1], " ", "`", 1))
//     // // for i := 0; i < 2; i++ {
//     // err = json.Unmarshal([]byte(`{"id": 1868, "first_name": "Kentarou", "last_name": "Miura", "role": "Story & Art"}`), &iot)
//     // if err != nil {
//     //     panic(err)
//     // }
//     // // }
//     // // Context is []byte, so you can keep it as string in DB
//     // fmt.Println(iot.Id)

//     // genre := [17]string{"Action",
//     // "Adventure",
//     // "AwardWinning",
//     // "Drama",
//     // "Fantasy",
//     // "Horror",
//     // "Supernatural",
//     // "Mystery",
//     // "SliceOfLife",
//     // "Psychological",
//     // "Comedy",
//     // "Suspense",
//     // "Sports",
//     // "Sci-Fi",
//     // "Ecchi",
//     // "Romance",
//     // "GirlsLove"}

//     // for q := 0; q < 2; q++ {
//     //     for i := 0; i < strings.Count(entries[q].Field1, ",") + 1; i++ {
//     //         fmt.Println(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Split(entries[q].Field1, ",")[i],"[","", 1), """, "", 10), "]", "", 1), " ", "", 1) )
//     //     }
//     //     fmt.Println(" |||| ")
//     //     // fmt.Println(strings.Replace(strings.Replace(strings.Replace(strings.Split(entry.Field1, ",")[0],"[","", 1), "'", "", 2), "]", "", 1))
//     //     // fmt.Println(entry.Field1)
//     // }

// 	// dsn := fmt.Sprintf("host=localhost user=postgres password=rencist dbname=tamiyochi port=5432 TimeZone=Asia/Jakarta")
// 	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	panic(err)
// 	// }
//     // count := 1
//     // for _, entry := range entries {
//         // var penerbitEntity entity.Penerbit
//         // dbRresult := db.Where("nama = ?", strings.Replace(strings.Replace(strings.Replace(entry.Field1, "[", "", 1), "]", "", 1), "'", "", 2)).First(&penerbitEntity)
//     //     if errors.Is(dbRresult.Error, gorm.ErrRecordNotFound) {
//     //         penerbit := entity.Penerbit{
//     //             ID: count,
//     //             Nama: strings.Replace(strings.Replace(strings.Replace(entry.Field1, "[", "", 1), "]", "", 1), "'", "", 2),
//     //         }
//     //         count++
//     //         db.Create(&penerbit)
//     //     }
//         // var penerbitEntity entity.Penerbit
// 	    // db.Find(&penerbitEntity)
//         // db.Where("nama = ?", strings.Replace(strings.Replace(strings.Replace(entry.Field1, "[", "", 1), "]", "", 1), "'", "", 2)).Take(&penerbitEntity)
//         // // if err != nil {
//         //     if penerbitEntity.Nama != strings.Replace(strings.Replace(strings.Replace(entry.Field1, "[", "", 1), "]", "", 1), "'", "", 2) {
//         //         penerbit := entity.Penerbit{
//         //             ID: count,
//         //             Nama: strings.Replace(strings.Replace(strings.Replace(entry.Field1, "[", "", 1), "]", "", 1), "'", "", 2),
//         //         }
//         //         count++
//         //         db.Create(&penerbit)
//         //     }
//         // } else {
//             // fmt.Println(err)
//         // }

//         // -------------------------------------------------------------------------------------------------
//         // var penulis entity.Penulis
//         // count1 := 1
//         // for _, entry := range entries {
//         //     count := strings.Count(entry.Field11, "{")
//         //     if count > 1 {
//         //         var penulisSeri entity.PenulisSeri
//         //         var iot Iot
//         //         err = json.Unmarshal([]byte(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "", 1), "]", "", 1), "'", "\"", 100), "},")[1], " ", "", 1)), &iot)
//         //         if err != nil {
//         //             panic(err)
//         //         }
//         //         // fmt.Println(iot.LastName)
//         //         penulis.ID = iot.Id - 1865
//         //         penulis.NamaDepan = iot.FirstName
//         //         penulis.NamaBelakang = iot.LastName
//         //         penulis.Peran = iot.Role

//         //         penulisSeri.ID = count1
//         //         penulisSeri.PenulisID = penulis.ID
//         //         penulisSeri.SeriID = entry.Field9
//         //         // fmt.Println(count1, "|", entry.Field9, "|", penulis.ID)
//         //         db.Create(&penulisSeri)
//         //         count1++

//         //         // db.Create(&penulis)

//         //         err = json.Unmarshal([]byte(strings.Split(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "", 1), "]", "", 1), "'", "\"", 100), ", {")[0]), &iot)
//         //         if err != nil {
//         //             panic(err)
//         //         }
//         //         // fmt.Println(iot.LastName)
//         //         penulis.ID = iot.Id - 1865
//         //         penulis.NamaDepan = iot.FirstName
//         //         penulis.NamaBelakang = iot.LastName
//         //         penulis.Peran = iot.Role

//         //         penulisSeri.ID = count1
//         //         penulisSeri.PenulisID = penulis.ID
//         //         penulisSeri.SeriID = entry.Field9
//         //         // fmt.Println(count1, "|", entry.Field9, "|", penulis.ID)
//         //         db.Create(&penulisSeri)
//         //         count1++

//         //         // fmt.Println(penulis.NamaDepan + " | " + penulis.NamaBelakang + " | " + penulis.Peran)

//         //         // fmt.Println(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "`", 1), "]", "`", 1), "'", "\"", 100), "},")[1], " ", "`", 1))
//         //         // fmt.Println(strings.Split(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "`", 1), "]", "`", 1), "'", "\"", 100), ", {")[0] + "`")
//         //     } else {
//         //         var penulisSeri entity.PenulisSeri
//         //         var iot Iot
//         //         err = json.Unmarshal([]byte(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "", 1), "]", "", 1), "'", "\"", 100)), &iot)
//         //         if err != nil {
//         //             panic(err)
//         //         }
//         //         // fmt.Println(iot.Id)
//         //         penulis.ID = iot.Id - 1865
//         //         penulis.NamaDepan = iot.FirstName
//         //         penulis.NamaBelakang = iot.LastName
//         //         penulis.Peran = iot.Role

//         //         penulisSeri.ID = count1
//         //         penulisSeri.PenulisID = penulis.ID
//         //         penulisSeri.SeriID = entry.Field9
//         //         // fmt.Println(count1, "|", entry.Field9, "|", penulis.ID)
//         //         db.Create(&penulisSeri)
//         //         count1++

//         //         // db.Create(&penulis)
//         //         // fmt.Println(penulis.NamaDepan + " | " + penulis.NamaBelakang + " | " + penulis.Peran)

//         //         // fmt.Println(strings.Replace(strings.Replace(strings.Replace(entry.Field11, "[", "`", 1), "]", "`", 1), "'", "\"", 100))
//         //     }
//         // }
//         // -------------------------------------------------------------------------------------------------

//         // count := strings.Count(entries[i].Field1, "{")
//         // if count > 1 {
//         //     var iot Iot
//         //     err = json.Unmarshal([]byte(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "", 1), "]", "", 1), "'", "\"", 100), "},")[1], " ", "", 1)), &iot)
//         //     if err != nil {
//         //         panic(err)
//         //     }
//         //     // fmt.Println(iot.LastName)
//         //     penulis.NamaDepan = iot.FirstName
//         //     penulis.NamaBelakang = iot.LastName
//         //     penulis.Peran = iot.Role
//         //     fmt.Println(penulis.NamaDepan + " | " + penulis.NamaBelakang + " | " + penulis.Peran)

//         //     err = json.Unmarshal([]byte(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "", 1), "]", "", 1), "'", "\"", 100), ", {")[0]), &iot)
//         //     if err != nil {
//         //         panic(err)
//         //     }
//         //     // fmt.Println(iot.LastName)
//         //     penulis.NamaDepan = iot.FirstName
//         //     penulis.NamaBelakang = iot.LastName
//         //     penulis.Peran = iot.Role
//         //     fmt.Println(penulis.NamaDepan + " | " + penulis.NamaBelakang + " | " + penulis.Peran)

//         //     // fmt.Println(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100), "},")[1], " ", "`", 1))
//         //     // fmt.Println(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100), ", {")[0] + "`")
//         // } else {
//         //     var iot Iot
//         //     err = json.Unmarshal([]byte(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "", 1), "]", "", 1), "'", "\"", 100)), &iot)
//         //     if err != nil {
//         //         panic(err)
//         //     }
//         //     // fmt.Println(iot.Id)
//         //     penulis.NamaDepan = iot.FirstName
//         //     penulis.NamaBelakang = iot.LastName
//         //     penulis.Peran = iot.Role
//         //     fmt.Println(penulis.NamaDepan + " | " + penulis.NamaBelakang + " | " + penulis.Peran)

//         //     // fmt.Println(strings.Replace(strings.Replace(strings.Replace(entries[i].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100))
//         // }

//         // fmt.Println(strings.Replace(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[0].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100), "},")[1], " ", "`", 1))
//         // fmt.Println(strings.Split(strings.Replace(strings.Replace(strings.Replace(entries[0].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100), ", {")[0] + "`")
//         // fmt.Println(strings.Replace(strings.Replace(strings.Replace(entries[1].Field1, "[", "`", 1), "]", "`", 1), "'", "\"", 100))

    }

    // func generateUniqueNickname() (string, string, string) {
    //     adjectives := []string{"Brave", "Mighty", "Swift", "Fierce", "Cunning", "Wise", "Noble", "Daring", "Valiant", "Epic"}
    //     nouns := []string{"Warrior", "Hero", "Mage", "Knight", "Rogue", "Sorcerer", "Paladin", "Archer", "Assassin", "Druid"}
    
    //     adjective := adjectives[rand.Intn(len(adjectives))]
    //     noun := nouns[rand.Intn(len(nouns))]
    //     return fmt.Sprintf("%s %s", adjective, noun), fmt.Sprintf("%s%s@gmail.com", strings.ToLower(adjective), strings.ToLower(noun)), fmt.Sprintf("%s_%s", strings.ToLower(adjective), strings.ToLower(noun))
        
    // }

    
    // kecamatan := [100]string{
	// 	"Gambir",
	// 	"Tanah Abang",
	// 	"Menteng",
	// 	"Senen",
	// 	"Cempaka Putih",
	// 	"Johar Baru",
	// 	"Kemayoran",
	// 	"Sawah Besar",
	// 	"Gunung Sahari",
	// 	"Kemuning",
	// 	"Pasar Rebo",
	// 	"Ciracas",
	// 	"Cipayung",
	// 	"Makasar",
	// 	"Kramat Jati",
	// 	"Jatinegara",
	// 	"Duren Sawit",
	// 	"Cakung",
	// 	"Pulo Gadung",
	// 	"Matraman",
	// 	"Penjaringan",
	// 	"Tanjung Priok",
	// 	"Koja",
	// 	"Kelapa Gading",
	// 	"Cilincing",
	// 	"Pademangan",
	// 	"Kebayoran Lama",
	// 	"Kebayoran Baru",
	// 	"Pancoran",
	// 	"Jagakarsa",
	// 	"Mampang Prapatan",
	// 	"Pasar Minggu",
	// 	"Tebet",
	// 	"Setiabudi",
	// 	"Kebon Jeruk",
	// 	"Palmerah",
	// 	"Grogol Petamburan",
	// 	"Tambora",
	// 	"Taman Sari",
	// 	"Cengkareng",
	// 	"Kali Deres",
	// 	"Kalideres",
	// 	"Kebon Jeruk",
	// 	"Kembangan",
	// 	"Kebayoran Lama",
	// 	 "Pesanggrahan",
	// 	"Cengkareng",
	// 	"Kebayoran Baru",
	// 	"Grogol Petamburan",
	// 	"Tambora",
	// 	"Kebon Jeruk",
	// 	"Taman Sari",
	// 	"Tangerang",
	// 	"Serpong",
	// 	"Pamulang",
	// 	"Ciputat",
	// 	"Ciledug",
	// 	"Karawaci",
	// 	"Cikokol",
	// 	"Neglasari",
	// 	"Benda",
	// 	"Batuceper",
	// 	"Karang Tengah",
	// 	"Pinang",
	// 	"Larangan",
	// 	"Cipondoh",
	// 	"Jatiuwung",
	// 	"Cibodas",
	// 	"Parung",
	// 	"Rumpin",
	// 	"Ciseeng",
	// 	"Sukabumi",
	// 	"Gunung Sindur",
	// 	"Cibinong",
	// 	"Bojonggede",
	// 	"Leuwiliang",
	// 	"Ciampea",
	// 	"Tajur Halang",
	// 	"Parung Panjang",
	// 	"Jasinga",
	// 	"Sukaraja",
	// 	"Bogor Barat",
	// 	"Bogor Selatan",
	// 	"Cileungsi",
	// 	"Cibinong",
	// 	"Gunung Putri",
	// 	"Citeureup",
	// 	"Cileungsi",
	// 	"Cibinong",
	// 	"Parung",
	// 	"Sukabumi",
	// 	"Gunung Sindur",
	// 	"Cibinong",
	// 	"Bojong Gede",
	// 	"Leuwiliang",
	// 	"Ciampea",
	// 	"Tenjo",
	// 	"Parung Panjang",
	// 	"Cileungsi",
	// 	"Gunung Putri",
	// }