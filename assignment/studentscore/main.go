package main


import "fmt"

var myStudentScore[] string

func main()  {
	StudentGrade()



	
}


func StudentGrade (){
		
	fmt.Println("Enter the total number of student in your class: ")
	var studentTotalNumber int
	fmt.Scan(& studentTotalNumber)
	
	for i := 1; i <= studentTotalNumber; i++{
		fmt.Println("Enter the Grade of Each student in the class:")
		var studentGrade int 
		
		if _, err := fmt.Scan(& studentGrade); err != nil{
			fmt.Println("Invalid input. Please enter a valid integer.")
            return
		}
		var grade string

		if(studentGrade >= 90 && studentGrade <=100){
			grade = "A"
		}else if (studentGrade>= 80 && studentGrade <= 89){
			grade = "B"
		}else if (studentGrade >= 70 && studentGrade <= 79){
			grade = "C"
		}else if (studentGrade >= 60 && studentGrade <= 69){
			grade = "D"
		}else if (studentGrade >= 0 && studentGrade <= 59){
			grade = "F"
		}
		myStudentScore = append(myStudentScore, grade)

	}
	generateResult()
}

func generateResult()  {
	for index, result := range myStudentScore{
		fmt.Printf("Student %d: %s\n", index+1, result)
	}
}