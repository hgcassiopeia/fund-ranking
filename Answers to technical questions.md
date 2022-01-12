## Technical questions

1. ใช้เวลาทำแบบทดสอบไปเท่าไร ถ้ามีเวลามากกว่านี้จะทำอะไรเพิ่ม ถ้าใช้เวลาน้อยในการทำโจทย์สามารถใช้โอกาสนี้ในการอธิบายได้ว่าอยากเพิ่มอะไร หรือแก้ไขในส่วนไหน
- ใช้เวลาประมาณ 9 ชม.
- อยากจะ refactor code ใหม่
- จัดการ project structure ใหม่ จัด domain ใหม่
- ปรับปรุง flow ของ CLI Application ให้ UX ดีขึ้น ใช้งานง่ายขึ้น
  
2. อะไรคือ feature ที่นำเข้ามาใช้ในการพัฒนา application นี้ กรุณาแนบ code snippet มาด้วยว่าใช้อย่างไร ในส่วนไหน
- CLI Application นี้ใช้ Package Cobra ในการสร้าง Project Scaffolding `github.com/spf13/cobra`
- Application นี้มี feature หลักคือแสดง list ของกองทุนตามช่วงเวลาที่ผู้ใช้งานส่ง flag `--time` เข้ามา โดยส่วนที่ทำการรับ flag คือ
```go
time, _ := cmd.Flags().GetString("time")
```
- เมื่อรับ flag เข้ามาก็จะทำการตรวจสอบความถูกต้องค่าของ flag ที่รับเข้ามาที่ส่วนนี้
```go
correctTime := contains([]string{"1D", "1W", "1M", "1Y"}, strings.ToUpper(time))
if correctTime {
    time := strings.ToUpper(time)
    getFundByRange(time)
} else {
    log.Printf("Your input: %v is incorrect type please input time contain 1D, 1W, 1M, 1Y", time)
}
```
- เมื่อตรวจสอบแล้วว่า flag ที่รับมาถูกต้องก็ทำการส่ง `--time` ไปที่ function `getFundByRange` เพื่อดึงข้อมูลจาก API และแสดงข้อมูลกองทุนออกมาในรูปแบบของ table โดยใช้ `github.com/alexeyco/simpletable` ในการแสดงข้อมูลกองทุนแบบ Table
```go
table := simpletable.New()
table.Header = &simpletable.Header{
    Cells: []*simpletable.Cell{
        {Align: simpletable.AlignCenter, Text: "ชื่อกองทุน"},
        {Align: simpletable.AlignCenter, Text: "อันดับของกองทุน"},
        {Align: simpletable.AlignCenter, Text: "เวลาที่ข้อมูลถูกอัพเดต"},
        {Align: simpletable.AlignCenter, Text: "ผลตอบแทน"},
        {Align: simpletable.AlignCenter, Text: "ราคา"},
    },
}

for _, row := range fund {
    t, _ := time.Parse("2006-01-02T00:00:00.000Z", row.NavDate)
    r := []*simpletable.Cell{
        {Align: simpletable.AlignRight, Text: row.ThailandFundCode},
        {Align: simpletable.AlignRight, Text: fmt.Sprintf("%.2f", row.NavReturn)},
        {Align: simpletable.AlignCenter, Text: t.Format("02 Jan 2006")},
        {Align: simpletable.AlignRight, Text: fmt.Sprintf("%.2f", row.AvgReturn)},
        {Align: simpletable.AlignRight, Text: fmt.Sprintf("%.2f", row.Nav)},
    }
    table.Body.Cells = append(table.Body.Cells, r)
}
table.SetStyle(simpletable.StyleDefault)
```

3. เราจะสามารถติดตาม performance issue บน production ได้อย่างไร เคยมีประสบการณ์ด้านนี้ไหม
- เคยมีประสบการณ์ใช้ redmine ในการทำ issue tracking บน production แต่การ track จะเป็นการให้ user เข้าไปทดสอบ Manual แล้วแจ้ง issue มา ส่วนที่เป็นพวก Elastic APM เคยอ่านแต่บทความ ยังไม่เคยมีประสบการณ์ใช้งาน

4. อยากปรับปรุง FINNOMENA APIs ที่ใช้ในการพัฒนา ในส่วนไหนให้ดียิ่งขึ้น
- เพิ่ม Swagger API documentation เพื่อให้ง่ายต่อการนำ API ไปพัฒนา