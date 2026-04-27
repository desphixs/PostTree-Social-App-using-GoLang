# 🚀 Go-Social: Learning & Building Plan (Django to Go)

This roadmap tracks our journey from Django expertise to Go mastery by building a minimal social media platform.

## 🏁 Phase 0: Foundations & Concepts
- [x] **What is Go?** Understanding compilation vs interpretation.
- [x] **The "Project" in Go:** Modules vs Virtualenvs.
- [x] **Vocabulary Sync:** Mapping DRF terms to Go terms.

## 🏗️ Phase 1: Project Scaffolding (The "Startproject")
- [x] Initialize Go Module (`go mod init`).
- [x] Create entry point (`main.go`).
- [x] Run the first "Server" (Hello World).

## 🚦 Phase 2: Web Server & Routing (The "URLconf")
- [ ] Setting up a Router (e.g., Gin or Fiber).
- [ ] Creating Handlers (DRF APIViews equivalent).
- [ ] Path parameters & JSON responses.

## 💾 Phase 3: Structs & Data (The "Models")
- [ ] Writing Structs (Defining the Shape of Data).
- [ ] JSON Tags (The "Serializers" equivalent).
- [ ] Type Safety: Why you can't just pass any dictionary.

## 🗄️ Phase 4: Database & ORM (The "Persistence")
- [ ] Connecting GORM (The Django-like ORM).
- [ ] Defining Relationships (`ForeignKey`, `ManyToMany`).
- [ ] Auto-Migrations.

## 📝 Phase 5: Feature Implementation
- [ ] **Authentication:** JWT vs Sessions (DRF Auth).
- [ ] **Posts:** CRUD for text posts.
- [ ] **Likes/Interactions:** Many-to-Many logic.
- [ ] **Replies/Retweets:** Nested and shared logic.

## ⚙️ Phase 6: Production Polish
- [ ] Middlewares (Auth, Logging, CORS).
- [ ] Dependency Injection (Clean code).
- [ ] Goroutines (Async tasks/Celery equivalent).

---

### 📚 Django-to-Go Rosetta Stone

| Django | Go | Description |
| :--- | :--- | :--- |
| `python manage.py runserver` | `go run main.go` | Running the dev server |
| `pip install` | `go get` | Adding dependencies |
| `venv/requirements.txt` | `go.mod` | Dependency management |
| `models.py` | `models/structs.go` | Defining data structures |
| `serializers.py` | `Struct Tags` | JSON conversion |
| `views.py` | `handlers/controllers` | Business logic |
| `urls.py` | `router.go` | Path routing |
| `settings.py` | `.env` or `config.go` | Configuration |