package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"personal-web/connection"
	"personal-web/middleware"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	connection.DatabaseConnect()
	route := mux.NewRouter()

	route.PathPrefix("/public").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	route.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")

	route.HandleFunc("/register-form", registerForm).Methods("GET")
	route.HandleFunc("/send-register-data", sendRegisterData).Methods("POST")
	route.HandleFunc("/login-form", loginForm).Methods("GET")
	route.HandleFunc("/process-login-data", processLoginData).Methods("POST")
	route.HandleFunc("/process-logout", processLogout).Methods("GET")

	route.HandleFunc("/project-detail/{id}", projectDetail).Methods("GET")

	route.HandleFunc("/add-project-form", addProjectForm).Methods("GET")
	route.HandleFunc("/send-add-project-data", middleware.UploadFile(sendAddProjectData)).Methods("POST")

	route.HandleFunc("/edit-project-form/{id}", editProjectForm).Methods("GET")
	route.HandleFunc("/send-edit-project-data/{id}", middleware.UploadFile(sendEditProjectData)).Methods("POST")

	route.HandleFunc("/delete-project/{id}", deleteProject).Methods("GET")

	fmt.Println("Server running on localhost:8000")
	http.ListenAndServe("localhost:8000", route)
}

type userStruct struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type sessionDataStruct struct {
	IsLogin      bool
	UserId       int
	UserName     string
	FlashMessage string
}

var sessionData = sessionDataStruct{}

type projectDataStruc struct {
	Id              int
	ProjectName     string
	StartDate       time.Time
	EndDate         time.Time
	StartDateFormat string
	EndDateFormat   string
	Duration        string
	Description     string
	Technologies    []string
	Image           string
	IsLogin         bool
	FlashMessage    string
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("views/home.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return

	} else {
		sessionData = sessionDataStruct{}

		var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")

		if session.Values["IsLogin"] != true {
			sessionData.IsLogin = false
		} else {
			getFlashMessage := session.Flashes("Message")
			session.Save(r, w)
			var buildFlashMessage []string
			if len(getFlashMessage) > 0 {
				for _, fMLetter := range getFlashMessage {
					buildFlashMessage = append(buildFlashMessage, fMLetter.(string))
				}
			}

			sessionData.IsLogin = session.Values["IsLogin"].(bool)
			sessionData.UserId = session.Values["Id"].(int)
			sessionData.UserName = session.Values["Name"].(string)
			sessionData.FlashMessage = strings.Join(buildFlashMessage, "")
		}

		if sessionData.IsLogin != true {
			var projectData []projectDataStruc
			allDataFrom_db_project, _ := connection.Conn.Query(context.Background(), "SELECT id, project_name, start_date, end_date, duration, description, technologies, image FROM db_project")
			for allDataFrom_db_project.Next() {
				selectedProjectData := projectDataStruc{}
				err := allDataFrom_db_project.Scan(&selectedProjectData.Id, &selectedProjectData.ProjectName, &selectedProjectData.StartDate, &selectedProjectData.EndDate, &selectedProjectData.Duration, &selectedProjectData.Description, &selectedProjectData.Technologies, &selectedProjectData.Image)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				projectData = append(projectData, selectedProjectData)
			}

			response := map[string]interface{}{
				"SessionData": sessionData,
				"ProjectData": projectData,
			}

			w.WriteHeader(http.StatusOK)
			tmpl.Execute(w, response)

		} else {
			var projectData []projectDataStruc
			allDataFrom_db_project, _ := connection.Conn.Query(context.Background(), "SELECT db_project.id, project_name, start_date, end_date, duration, description, technologies, image FROM db_project LEFT JOIN db_user ON db_project.author_id = db_user.id WHERE db_user.id = $1 ORDER BY id DESC", sessionData.UserId)
			for allDataFrom_db_project.Next() {
				selectedProjectData := projectDataStruc{}
				err := allDataFrom_db_project.Scan(&selectedProjectData.Id, &selectedProjectData.ProjectName, &selectedProjectData.StartDate, &selectedProjectData.EndDate, &selectedProjectData.Duration, &selectedProjectData.Description, &selectedProjectData.Technologies, &selectedProjectData.Image)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				selectedProjectData.IsLogin = sessionData.IsLogin

				projectData = append(projectData, selectedProjectData)
			}

			response := map[string]interface{}{
				"SessionData": sessionData,
				"ProjectData": projectData,
			}

			w.WriteHeader(http.StatusOK)
			tmpl.Execute(w, response)
		}
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	session, _ := store.Get(r, "SESSION_KEY")

	if session.Values["IsLogin"] != true {
		sessionData.IsLogin = false
	} else {
		sessionData.IsLogin = session.Values["IsLogin"].(bool)
		sessionData.UserName = session.Values["Name"].(string)
	}
	response := map[string]interface{}{
		"SessionData": sessionData,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, response)
}

func registerForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/register.html")

	if tmpl == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message : " + err.Error()))
	} else {
		var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")

		if session.Values["IsLogin"] != true {
			sessionData.IsLogin = false
		} else {
			sessionData.IsLogin = session.Values["IsLogin"].(bool)
			sessionData.UserName = session.Values["Name"].(string)
		}
		response := map[string]interface{}{
			"SessionData": sessionData,
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, response)
	}
}

func sendRegisterData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	name := r.PostForm.Get("input-name")
	email := r.PostForm.Get("input-email")
	password := r.PostForm.Get("input-password")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO db_user(name, email, password) VALUES ($1, $2, $3)", name, email, hashedPassword)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	http.Redirect(w, r, "/login-form", http.StatusMovedPermanently)
}

func loginForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/login.html")

	if tmpl == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message : " + err.Error()))
	} else {
		store := sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")

		if session.Values["IsLogin"] != true {
			sessionData.IsLogin = false
		} else {
			sessionData.IsLogin = session.Values["IsLogin"].(bool)
			sessionData.UserName = session.Values["Name"].(string)
		}

		getFlashMessage := session.Flashes("Message")
		session.Save(r, w) //untuk mereset flash message dari session browser
		var buildFlashMessage []string
		if len(getFlashMessage) > 0 {
			for _, fMLetter := range getFlashMessage {
				buildFlashMessage = append(buildFlashMessage, fMLetter.(string))
			}
		}
		sessionData.FlashMessage = strings.Join(buildFlashMessage, "")

		response := map[string]interface{}{
			"SessionData": sessionData,
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, response)
	}
}

func processLoginData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := r.PostForm.Get("input-email")
	password := r.PostForm.Get("input-password")

	user := userStruct{}

	errEmail := connection.Conn.QueryRow(context.Background(), "SELECT * FROM db_user WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if errEmail != nil {
		store := sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")
		session.AddFlash("Email belum terdaftar!", "Message")
		session.Save(r, w)

		http.Redirect(w, r, "/login-form", http.StatusMovedPermanently)

	} else if errPassword != nil {
		store := sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")
		session.AddFlash("Password Salah!", "Message")
		session.Save(r, w)

		http.Redirect(w, r, "/login-form", http.StatusMovedPermanently)

	} else {
		var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")
		session.Values["Name"] = user.Name
		session.Values["Email"] = user.Email
		session.Values["Id"] = user.Id
		session.Values["IsLogin"] = true
		session.Options.MaxAge = 10800 // Detik

		session.AddFlash("Login succesfully... Now you can create, show, edit, and delete your Projects", "Message")
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func processLogout(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	session, _ := store.Get(r, "SESSION_KEY")
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
	} else {
		var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")

		if session.Values["IsLogin"] != true {
			sessionData.IsLogin = false
		} else {
			sessionData.IsLogin = session.Values["IsLogin"].(bool)
			sessionData.UserName = session.Values["Name"].(string)
		}

		where_id, _ := strconv.Atoi(mux.Vars(r)["id"])

		selectedProjectData := projectDataStruc{}

		err = connection.Conn.QueryRow(context.Background(), "SELECT id, project_name, start_date, end_date, duration, description, technologies, image FROM db_project WHERE id=$1", where_id).
			Scan(&selectedProjectData.Id, &selectedProjectData.ProjectName, &selectedProjectData.StartDate, &selectedProjectData.EndDate, &selectedProjectData.Duration, &selectedProjectData.Description, &selectedProjectData.Technologies, &selectedProjectData.Image)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		selectedProjectData.StartDateFormat = selectedProjectData.StartDate.Format("2006-01-02")
		selectedProjectData.EndDateFormat = selectedProjectData.EndDate.Format("2006-01-02")

		response := map[string]interface{}{
			"SessionData":         sessionData,
			"selectedProjectData": selectedProjectData,
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, response)
	}
}

func addProjectForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if tmpl == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message : " + err.Error()))
	} else {
		var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")

		if session.Values["IsLogin"] != true {
			sessionData.IsLogin = false
		} else {
			sessionData.IsLogin = session.Values["IsLogin"].(bool)
			sessionData.UserId = session.Values["Id"].(int)
			sessionData.UserName = session.Values["Name"].(string)
		}
		response := map[string]interface{}{
			"SessionData": sessionData,
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, response)
	}
}

func sendAddProjectData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	} else {
		projectName := r.PostForm.Get("project-name")
		startDate := r.PostForm.Get("start-date")
		endDate := r.PostForm.Get("end-date")
		var duration string
		description := r.PostForm.Get("description")
		technologies := []string{r.PostForm.Get("node"), r.PostForm.Get("react"), r.PostForm.Get("vue"), r.PostForm.Get("typescript")}
		dataContext := r.Context().Value("dataFile")
		image := dataContext.(string)
		// image := r.PostForm.Get("project-image")

		layoutDate := "2006-01-02"
		startDateParse, _ := time.Parse(layoutDate, startDate)
		endDateParse, _ := time.Parse(layoutDate, endDate)

		hour := 1
		day := hour * 24
		week := hour * 24 * 7
		month := hour * 24 * 30
		year := hour * 24 * 365

		differHour := endDateParse.Sub(startDateParse).Hours()
		var differHours int = int(differHour)
		// fmt.Println(differHours)
		days := differHours / day
		weeks := differHours / week
		months := differHours / month
		years := differHours / year

		if differHours < week {
			duration = strconv.Itoa(int(days)) + " Days"
		} else if differHours < month {
			duration = strconv.Itoa(int(weeks)) + " Weeks"
		} else if differHours < year {
			duration = strconv.Itoa(int(months)) + " Months"
		} else if differHours > year {
			duration = strconv.Itoa(int(years)) + " Years"
		}

		_, err = connection.Conn.Exec(context.Background(), "INSERT INTO db_project(author_id, project_name, start_date, end_date, duration, description, technologies, image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", sessionData.UserId, projectName, startDate, endDate, duration, description, technologies, image)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func editProjectForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/edit-project.html")

	if tmpl == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message : " + err.Error()))
	} else {
		var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")

		if session.Values["IsLogin"] != true {
			sessionData.IsLogin = false
		} else {
			sessionData.IsLogin = session.Values["IsLogin"].(bool)
			sessionData.UserName = session.Values["Name"].(string)
		}

		where_id, _ := strconv.Atoi(mux.Vars(r)["id"])
		selectedProjectData := projectDataStruc{}
		err = connection.Conn.QueryRow(context.Background(), "SELECT id, project_name, start_date, end_date, duration, description, technologies, image FROM db_project WHERE id=$1", where_id).
			Scan(&selectedProjectData.Id, &selectedProjectData.ProjectName, &selectedProjectData.StartDate, &selectedProjectData.EndDate, &selectedProjectData.Duration, &selectedProjectData.Description, &selectedProjectData.Technologies, &selectedProjectData.Image)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		selectedProjectData.StartDateFormat = selectedProjectData.StartDate.Format("2006-01-02")
		selectedProjectData.EndDateFormat = selectedProjectData.EndDate.Format("2006-01-02")

		response := map[string]interface{}{
			"SessionData":         sessionData,
			"selectedProjectData": selectedProjectData,
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, response)
	}
}

func sendEditProjectData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	} else {
		where_id, _ := strconv.Atoi(mux.Vars(r)["id"])

		projectName := r.PostForm.Get("project-name")
		startDate := r.PostForm.Get("start-date")
		endDate := r.PostForm.Get("end-date")
		var duration string
		description := r.PostForm.Get("description")
		technologies := []string{r.PostForm.Get("node"), r.PostForm.Get("react"), r.PostForm.Get("vue"), r.PostForm.Get("typescript")}
		dataContext := r.Context().Value("dataFile")
		image := dataContext.(string)
		// image := r.PostForm.Get("project-image")

		layoutDate := "2006-01-02"
		startDateParse, _ := time.Parse(layoutDate, startDate)
		endDateParse, _ := time.Parse(layoutDate, endDate)

		hour := 1
		day := hour * 24
		week := hour * 24 * 7
		month := hour * 24 * 30
		year := hour * 24 * 365

		differHour := endDateParse.Sub(startDateParse).Hours()
		var differHours int = int(differHour)
		// fmt.Println(differHours)
		days := differHours / day
		weeks := differHours / week
		months := differHours / month
		years := differHours / year

		if differHours < week {
			duration = strconv.Itoa(int(days)) + " Days"
		} else if differHours < month {
			duration = strconv.Itoa(int(weeks)) + " Weeks"
		} else if differHours < year {
			duration = strconv.Itoa(int(months)) + " Months"
		} else if differHours > year {
			duration = strconv.Itoa(int(years)) + " Years"
		}

		_, err = connection.Conn.Exec(context.Background(), "UPDATE db_project SET project_name=$1, start_date=$2, end_date=$3, duration=$4, description=$5, technologies=$6, image=$7 WHERE id=$8",
			projectName, startDate, endDate, duration, description, technologies, image, where_id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	where_id, _ := strconv.Atoi(mux.Vars(r)["id"])
	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM db_project WHERE id=$1", where_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
