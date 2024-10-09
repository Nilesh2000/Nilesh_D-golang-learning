package grading

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close() // Close file at the end of the program

	// Read CSV data
	reader := csv.NewReader(f)

	// Read first line
	if _, err := reader.Read(); err != nil {
		panic(err)
	}

	students := []student{}
	data, err := reader.ReadAll()

	for _, each := range data {
		test1, _ := strconv.ParseInt(each[3], 10, 64)
		test2, _ := strconv.ParseInt(each[4], 10, 64)
		test3, _ := strconv.ParseInt(each[5], 10, 64)
		test4, _ := strconv.ParseInt(each[6], 10, 64)
		st := student{
			firstName:  each[0],
			lastName:   each[1],
			university: each[2],
			test1Score: int(test1),
			test2Score: int(test2),
			test3Score: int(test3),
			test4Score: int(test4),
		}
		students = append(students, st)
	}

	return students
}

func determineGrade(score float32) Grade {
	switch {
	case score >= 70:
		return A
	case score >= 50 && score < 70:
		return B
	case score >= 35 && score < 50:
		return C
	default:
		return F
	}
}

func calculateGrade(students []student) []studentStat {
	studentStats := []studentStat{}

	for _, st := range students {
		finalScore := float32(st.test1Score+st.test2Score+st.test3Score+st.test4Score) / 4
		grade := determineGrade(finalScore)
		sts := studentStat{
			student:    st,
			finalScore: finalScore,
			grade:      grade,
		}
		studentStats = append(studentStats, sts)
	}
	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	topper := studentStat{}
	var topperScore float32

	for _, ss := range gradedStudents {
		if ss.finalScore > topperScore {
			topper = ss
			topperScore = ss.finalScore
		}
	}
	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	topperPerUniversity := make(map[string]studentStat)

	for _, ss := range gs {
		val, ok := topperPerUniversity[ss.university]
		if !ok {
			topperPerUniversity[ss.university] = ss
		} else {
			if ss.finalScore > val.finalScore {
				topperPerUniversity[ss.university] = ss
			}
		}
	}
	return topperPerUniversity
}
