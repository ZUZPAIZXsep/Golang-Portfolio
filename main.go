package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type PageData struct {
	Title          string
	ActivePage     string
	Education      []Education
	WorkExperience []WorkExperience
	Projects       []Project
}

type Education struct {
	Degree      string
	Faculty     sql.NullString
	Major       string
	Institution string
	StartYear   int
	EndYear     int
	GPAX        float64
}

type WorkExperience struct {
	ID               int
	Company          string
	Position         string
	StartDate        time.Time
	EndDate          time.Time
	Responsibilities []string
}

type Project struct {
	ID           int
	Name         string
	Description  string
	Bullets      []string
	Technologies []string
}

var db *sql.DB

// homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/index.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:      "	Home | Kantharakorn",
		ActivePage: "home",
	}

	err = tmpl.ExecuteTemplate(w, "base", data)

}

// aboutpage
func aboutHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/about.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:      "About | Kantharakorn",
		ActivePage: "about",
	}

	tmpl.ExecuteTemplate(w, "base", data)
}

// educationpage
func educationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/education.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	educationList, err := getEducation()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:      "Education | Kantharakorn",
		ActivePage: "education",
		Education:  educationList,
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// fetch data from educataion table
func getEducation() ([]Education, error) {
	rows, err := db.Query(`
		SELECT degree, faculty, major, institution, start_year, end_year, gpax
		FROM education
		ORDER BY start_year DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var educationList []Education

	for rows.Next() {
		var education Education

		err := rows.Scan(
			&education.Degree,
			&education.Faculty,
			&education.Major,
			&education.Institution,
			&education.StartYear,
			&education.EndYear,
			&education.GPAX,
		)
		if err != nil {
			return nil, err
		}

		educationList = append(educationList, education)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return educationList, nil
}

// experiencepage
func experienceHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/experience.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	experienceList, err := getWorkExperience()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:          "Work Experience | Kantharakorn",
		ActivePage:     "experience",
		WorkExperience: experienceList,
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// fetch data from experience table
func getWorkExperience() ([]WorkExperience, error) {
	rows, err := db.Query(`
        SELECT
            w.id,
            w.company,
            w.position,
            w.start_date,
            w.end_date,
            r.description
        FROM work_experience w
        JOIN work_experience_responsibilities r
            ON w.id = r.experience_id
        ORDER BY
            w.start_date ASC,
            r.sort_order
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experienceList []WorkExperience

	for rows.Next() {
		var (
			id          int
			company     string
			position    string
			startDate   time.Time
			endDate     time.Time
			description string
		)

		err := rows.Scan(
			&id,
			&company,
			&position,
			&startDate,
			&endDate,
			&description,
		)
		if err != nil {
			return nil, err
		}

		if len(experienceList) == 0 ||
			experienceList[len(experienceList)-1].ID != id {

			experienceList = append(experienceList, WorkExperience{
				ID:               id,
				Company:          company,
				Position:         position,
				StartDate:        startDate,
				EndDate:          endDate,
				Responsibilities: []string{},
			})
		}

		last := &experienceList[len(experienceList)-1]

		last.Responsibilities =
			append(last.Responsibilities, description)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return experienceList, nil
}

// Projects page
func projectsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/projects.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	projectList, err := getProjects()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:      "Projects | Kantharakorn",
		ActivePage: "projects",
		Projects:   projectList,
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// query data from database projects
func getProjects() ([]Project, error) {

	// Get projects
	rows, err := db.Query(`
        SELECT
            id,
            name,
            description
        FROM projects
        ORDER BY id
    `)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projectList []Project

	for rows.Next() {

		var project Project

		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
		)

		if err != nil {
			return nil, err
		}

		project.Bullets = []string{}
		project.Technologies = []string{}

		projectList = append(projectList, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Get project bullets
	for i := range projectList {

		rows, err := db.Query(`
            SELECT
                description
            FROM project_bullets
            WHERE project_id = $1
            ORDER BY sort_order
        `, projectList[i].ID)

		if err != nil {
			return nil, err
		}

		for rows.Next() {

			var bullet string

			err := rows.Scan(&bullet)

			if err != nil {
				rows.Close()
				return nil, err
			}

			projectList[i].Bullets =
				append(projectList[i].Bullets, bullet)
		}

		if err := rows.Err(); err != nil {
			rows.Close()
			return nil, err
		}

		rows.Close()
	}

	// Get project technologies
	for i := range projectList {

		rows, err := db.Query(`
            SELECT
                technology
            FROM project_technologies
            WHERE project_id = $1
            ORDER BY sort_order
        `, projectList[i].ID)

		if err != nil {
			return nil, err
		}

		for rows.Next() {

			var technology string

			err := rows.Scan(&technology)

			if err != nil {
				rows.Close()
				return nil, err
			}

			projectList[i].Technologies =
				append(projectList[i].Technologies, technology)
		}

		if err := rows.Err(); err != nil {
			rows.Close()
			return nil, err
		}

		rows.Close()
	}

	return projectList, nil
}

func main() {

	//conect PostgreSQL
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=portfolio_db sslmode=disable"

	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/education", educationHandler)
	http.HandleFunc("/experience", experienceHandler)
	http.HandleFunc("/projects", projectsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Connected to PostgreSQL")
	println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
