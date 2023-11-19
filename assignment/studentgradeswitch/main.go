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
		switch{
		case studentGrade >= 90 && studentGrade <=100:
			grade = "A"
		case studentGrade>= 80 && studentGrade <= 89 :
			grade = "B"
		case studentGrade >= 70 && studentGrade <= 79 :
			grade = "C"
		case studentGrade >= 60 && studentGrade <= 69 :
			grade = "D"
		case studentGrade >= 0 && studentGrade <= 59 :
			grade = "F"
		 default:
            fmt.Println("Invalid grade. Please enter a grade between 0 and 100.")
            return
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