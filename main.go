package main

import (
	"fmt"
	"gobasic/test"
)




func main() {
	// helloGo()
	// datatype()
	// outputDataType()
	// operator()
	// scanf()
	// condition()
	// arrayType()
	// slices()
	// maps()
	// loop()
	// breakcontinue()
	// ranges()
	// function()
	// returnmutivalue()
	// variabicfunction()
	// structs()
	packages()


}




func helloGo() {
	fmt.Println("Hello Go Programming")
}

func datatype() {
	fmt.Println("bool = ค่าทางตรรกศาสตร์ true/false ค่าเริ่มต้น false")
	fmt.Println("int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr = ตัวเลขที่ไม่มีทศนิยม ค่าเริ่มต้น 0")
	fmt.Println("float32 float64 = ตัวเลขทศนิยม ค่าเริ่มต้น 0")
	fmt.Println("string = ชุดข้อความ ค่าเริ่มต้น double quote")
	fmt.Println("")

	fmt.Println("Manual Type Declaration")
	var name string = "jetsadawwts"
	var age int = 26
	var score float32 = 99.9
	var isPass bool = true
	fmt.Println("My name is", name)
	fmt.Println("Age =", age)
	fmt.Println("Score =", score)
	fmt.Println("Pass Exam =", isPass)
	fmt.Println("")
	fmt.Println("Type Interence")
	names := "jetsadawwts"
	ages := 26
	scores := 99.9
	isPasss := true
	fmt.Println("My name is", names)
	fmt.Println("Age =", ages)
	fmt.Println("Score =", scores)
	fmt.Println("Pass Exam =", isPasss)
	fmt.Println("Type Interence")
	fmt.Println("")
	fmt.Println("ค่าคงที่")
	const firstname string = "Jetsadawwts"
	fmt.Println("My name is", firstname)
}

func outputDataType() {
	name := "Jetsadawwts"
	fmt.Printf("My name is %v \n", name) //value
	fmt.Printf("Data type = %T \n", name)
}

func operator() {
	fmt.Println("+ บวก \n")
	fmt.Println("- ลบ\n")
	fmt.Println("* คูณ\n")
	fmt.Println("/ หาร\n")
	fmt.Println("% หารเอาเศษ\n")

	// var number1 int = 10
	// var number2 int = 3
	// var number1, number2 = 10, 3
	number1, number2 := 10, 3
	fmt.Println("ผลบวก =", number1+number2)
	fmt.Println("ผลลบ =", number1-number2)
	fmt.Println("ผลคูณ =", number1*number2)
	fmt.Println("ผลหาร =", number1/number2)
	fmt.Println("หารเอาเศษ = \n", number1%number2)
	fmt.Println("== เท่ากับ \n")
	fmt.Println("!= ไม่เท่ากับ\n")
	fmt.Println("> มากกว่า\n")
	fmt.Println("< น้อยกว่า\n")
	fmt.Println(">= มากกว่าเท่ากับ\n")
	fmt.Println("<= น้อยกว่าเท่ากับ\n")
	fmt.Println(number1, "เท่ากับ", number2, "หรือไม่ =", number1 == number2)
	fmt.Println(number1, "ไม่เท่ากับ", number2, "หรือไม่ =", number1 != number2)
	fmt.Println(number1, "มากกว่า", number2, "หรือไม่ =", number1 > number2)
	fmt.Println(number1, "น้อยกว่า", number2, "หรือไม่ =", number1 < number2)
	fmt.Println(number1, "มากกว่าหรือเท่ากับ", number2, "=", number1 >= number2)
	fmt.Println(number1, "น้อยกว่าหรือเท่ากับ", number2, "=", number1 <= number2)

}

func scanf() {
	fmt.Println("String Format \n")
	fmt.Println("string %s \n")
	fmt.Println("interger %d \n")
	fmt.Println("floating point %f \n")
	var name, score = "", 0

	fmt.Print("ป้อนชื่อนักเรียนเเละคะเเนนสอบ = ")
	fmt.Scanf("%s", &name)
	fmt.Scanf("%d", &score)

	fmt.Println("สวัสดี =", name)
	fmt.Println("คะเเนนสอบ + คะเเนนจิตพิสัย =", score+10)
}

func condition() {
	fmt.Println("if else stature \n")
	// var score int
	// fmt.Print("คะเเนนสอบ = ")
	// fmt.Scanf("%d", &score)
	// if score >= 50 {
	// 	fmt.Println("สอบผ่าน")
	// } else {
	// 	fmt.Println("สอบไม่ผ่าน")
	// }
	var number int
	fmt.Print("ป้อนตัวเลข = ")
	fmt.Scanf("%d", &number)
	// if number % 2 == 0 {
	// 	fmt.Println("เลขคู่")
	// } else {
	// 	fmt.Println("เลขคี่")
	// }

	// if number == 1 {
	// 	fmt.Println("เปิดบัญชีใหม่")
	// } else if number == 2 {
	// 	fmt.Println("ฝากเงิน-ถอนเงิน")
	// } else {
	// 	fmt.Println("ข้อมูลไม่ถูกต้อง")
	// }

	fmt.Println("switch case \n")
	switch number {
	case 1:
		fmt.Println("เปิดบัญชีใหม่")

	case 2:
		fmt.Println("ฝากเงิน-ถอนเงิน")
	default:
		fmt.Println("ข้อมูลไม่ถูกต้อง")
	}

}

func arrayType() {
	fmt.Println("Array \n")

	//สร้าง Array
	//var numbers[3]int
	//กำหนดค่า
	// numbers[0] = 100
	// numbers[1] = 200
	// numbers[2] = 300

	// var numbers[3]int = [3]int{100, 200, 300}
	//  numbers := [3]int{100, 200, 300}
	numbers := [...]int{100, 200, 300}
	//len ขนาดของ Array
	size := len(numbers)

	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(size)
}

func slices() {
	fmt.Println("Slices Dynamic Size Array \n")
	numbers := []int{100, 200}
	// fmt.Println(numbers)
	// numbers[0] = 1000
	// numbers[1] = 2000

	//เพิ่มค่า Array
	numbers = append(numbers, 300)
	fmt.Println(numbers[:2])
	fmt.Println(numbers[1:])
	fmt.Println(numbers)
}

func maps() {
	fmt.Println("Map \n")
	// country := map[string]string{"TH": "ประเทศไทย", "EN": "อังกฤษ"}
	country := map[string]string{}
	country["TH"] = "ประเทศไทย"
	country["EN"] = "ประเทศอังกฤษ"
	

	// value, check := country["EN"]
    value, isKey := country["EN"]
	if isKey  {
		fmt.Println(value)
	} else {
		fmt.Println("ไม่พบช้อมูล")
	}


}


func loop() {
	for count:= 1; count <= 3; count++ {
		fmt.Println("Hello Go Programming")
	}

	for count:= 3; count >= 1; count-- {
		fmt.Println(count)
	}

	for count:= 1; count <= 3; count++ {
		fmt.Println(count)
	}
}


func breakcontinue() {
	for count:= 1; count <= 5; count++ {
		fmt.Println(count)
		if(count == 3) {
			break
		}
	}
	for count:= 1; count <= 5; count++ {
		
		if(count == 3) {
			continue
		}
		fmt.Println(count)
	}

	fmt.Println("End Program")
}

func ranges() {
	numbers:= []int{10,20,30,40,50,60,70,80,90,100}
	language:= map[string]string{"TH":"Thai", "EN":"English"}

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Index = ",i,"Value = ",numbers[i])
	// }

	for index, value := range numbers {
		fmt.Println("Index = ",index,"Value = ",value)
	}


	for key, value := range language {
		fmt.Println("key = ",key,"Value = ",value)
	}
}

func function() {
	function1()
	function2("Hello Go Programming")
	function3(50,100)
	fmt.Println("ค่าจัดส่ง = ",function4())
	fmt.Println("ยอดรวม = ",function5(500, 500))
}

func function1() {
	fmt.Println("Hello Go Programming")
}

func function2(name string) {
	fmt.Println(name)
}

func function3(number1, number2 int) {
	fmt.Println("ยอดรวม = ",number1+number2)
}

func function4() int {
	return 50
}

func function5(number1, number2 int) int {
	total:= number1 + number2
	return total
}	

func returnmutivalue() {
	result, check := returnmutivalue1(100, 200)
	fmt.Println("ผลรวม =",result)
	fmt.Println("number =",check)
}

func returnmutivalue1(number1, number2 int) (int, string) {
	total := number1 + number2
	status := ""
	if(total%2 == 0) {
		status = "เลขคู่"
	} else {
		status = "เลขคี่"
	}
	return total, status
}

func variabicfunction() {
	fmt.Println(variabicfunction1(10,20,30,40,50,60))
}	

func variabicfunction1(numbers...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

type structss1 struct {
	name string
	price float64
	category string
	discount int
}

func structs() {
	product1 := structss1{name: "ปากกา", price: 50.5, category: "เครื่องเขียน", discount: 10}
	fmt.Println(product1)
}

func packages() {
	fmt.Println(test.Add(100,100))
	fmt.Println(test.Subtract(100,200))
}






//Go Programming Comment

/*
	Go Programming Comment
*/
