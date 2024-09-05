package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "embed"

	_ "modernc.org/sqlite"
)

var (
	//go:embed schema.sql
	schema string
)

var searchQuery = `
	SELECT
		articles_fts.slug AS slug,
		highlight(articles_fts, 1, '<span class="text-lg font-bold underline bg-green-200">', '</span>') AS title, 
		highlight(articles_fts, 2, '<span class="text-lg font-bold underline bg-green-200">', '</span>') AS content
	FROM articles_fts
		JOIN articles ON articles_fts.slug = articles.slug
	WHERE articles_fts MATCH ?;
`

var listQuery = `
	SELECT slug, title, content FROM articles;
`

var lookupBySlugQuery = `
	SELECT slug, title, content FROM articles WHERE slug = ?;
`

var insertQuery = `
	INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES (?, ?, ?, strftime('%s', 'now'), strftime('%s', 'now'));
`

func migrate(dbc *sql.DB) error {
	_, err := dbc.Exec(schema)
	return err
}

func insert(dbc *sql.DB, slug, title, content string) error {
	_, err := dbc.Exec(insertQuery, slug, title, content)
	return err
}

func searchForArticles(dbc *sql.DB, q string) (string, error) {
	rows, err := dbc.Query(searchQuery, q)
	if err != nil {
		return "", err
	}

	cnt := ""
	for rows.Next() {
		slug := ""
		title := ""
		content := ""
		if err := rows.Scan(&slug, &title, &content); err != nil {
			return "", err
		}

		title = fmt.Sprintf(`<a class="hover:text-green-500" href="/blog/%s">
			<h1>
				%s
			</h1>
		</a>`, slug, title)
		content = fmt.Sprintf("<p>%s</p>", content)
		entry := fmt.Sprintf(`<div class="p-5 border border-slate-600 bg-slate-200">%s%s</div>`, title, content)

		cnt += entry
	}

	return cnt, nil
}

func listArticles(dbc *sql.DB) (string, error) {
	rows, err := dbc.Query(listQuery)
	if err != nil {
		return "", err
	}

	cnt := ""
	for rows.Next() {
		slug := ""
		title := ""
		content := ""
		if err := rows.Scan(&slug, &title, &content); err != nil {
			return "", err
		}

		title = fmt.Sprintf(`<a class="hover:text-green-500" href="/blog/%s">
			<h1>
				%s
			</h1>
		</a>`, slug, title)
		content = fmt.Sprintf("<p>%s</p>", content)
		entry := fmt.Sprintf(`<div class="p-5 border border-slate-600 bg-slate-200">%s%s</div>`, title, content)

		cnt += entry
	}

	return cnt, nil
}

func lookupBySlug(dbc *sql.DB, slug string) (string, error) {
	slugs := ""
	title := ""
	content := ""

	err := dbc.QueryRow(lookupBySlugQuery, slug).Scan(&slugs, &title, &content)
	if err != nil {
		return "", err
	}

	html := `<h1 class="text-xl px-10 mt-10"> %s </h1>
			<div class="px-10">
				%s
			</div>
	`

	html = fmt.Sprintf(html, title, content)
	return html, nil
}

func main() {
	dbc, err := sql.Open("sqlite", "file.db")
	if err != nil {
		panic("fucked")
	}
	defer dbc.Close()

	if err := migrate(dbc); err != nil {
		panic(err)
	}
	c := NewSearchCache[string]()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		html := `
			<!doctype html>
			<html>
				<head>
				  <meta charset="UTF-8">
				  <meta name="viewport" content="width=device-width, initial-scale=1.0">
				  <script src="https://cdn.tailwindcss.com"></script>
				  <script src="https://unpkg.com/htmx.org@2.0.2"></script>
				</head>
				<body class="h-screen flex flex-col overflow-hidden" hx-boost="true">
					<header class="bg-white">
					  <div class="mx-auto max-w-screen-xl px-4 py-8 sm:px-6 sm:py-12 lg:px-8">
						<div class="flex flex-col items-start gap-4 md:flex-row md:items-center md:justify-between">
						  <div>
							<h1 class="text-2xl font-bold text-gray-900 sm:text-3xl">Blog Posts</h1>

							<p class="mt-1.5 text-sm text-gray-500">
							  Lorem ipsum dolor, sit amet consectetur adipisicing elit. Iure, recusandae.
							</p>
						  </div>

						  <div class="flex items-center gap-4">
							<button
							  class="inline-flex items-center justify-center gap-1.5 rounded border border-gray-200 bg-white px-5 py-3 text-gray-900 transition hover:text-gray-700 focus:outline-none focus:ring"
							  type="button"
							>
							  <span class="text-sm font-medium"> View Website </span>

							  <svg
								xmlns="http://www.w3.org/2000/svg"
								class="size-4"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
								stroke-width="2"
							  >
								<path
								  stroke-linecap="round"
								  stroke-linejoin="round"
								  d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
								/>
							  </svg>
							</button>

							<button
							  class="inline-block rounded bg-indigo-600 px-5 py-3 text-sm font-medium text-white transition hover:bg-indigo-700 focus:outline-none focus:ring"
							  type="button"
							>
							  Create Post
							</button>
						  </div>
						</div>
					  </div>
					</header>
					<h1 class="text-xl px-10 mt-10"> Blogs </h1>
					<div class="p-10">
						<form class="flex flex-col space-y-2" hx-post="/search" hx-target="#results" hx-swap="innerHTML">
							<input class="w-full p-5 border border-black rounded" id="query" name="query">
							<button class="py-2 px-5 bg-blue-500 text-white" type="submit">Submit</button>
						</form>
					</div>
					<div id="results" class="flex flex-col space-y-5 px-10 overflow-y-auto">
						%s
					</div>
				</body>
			</html>
		`
		articleHTML, err := listArticles(dbc)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		html = fmt.Sprintf(html, articleHTML)

		w.Write([]byte(html))
	})

	mux.HandleFunc("GET /blog/{slug}", func(w http.ResponseWriter, r *http.Request) {
		html := `
			<body class="h-screen flex flex-col overflow-hidden" hx-boost="true">
				<header class="bg-white">
				  <div class="mx-auto max-w-screen-xl px-4 py-8 sm:px-6 sm:py-12 lg:px-8">
					<div class="flex flex-col items-start gap-4 md:flex-row md:items-center md:justify-between">
					  <div>
						<h1 class="text-2xl font-bold text-gray-900 sm:text-3xl">Blog Posts</h1>

						<p class="mt-1.5 text-sm text-gray-500">
						  Lorem ipsum dolor, sit amet consectetur adipisicing elit. Iure, recusandae.
						</p>
					  </div>

					  <div class="flex items-center gap-4">
						<button
						  class="inline-flex items-center justify-center gap-1.5 rounded border border-gray-200 bg-white px-5 py-3 text-gray-900 transition hover:text-gray-700 focus:outline-none focus:ring"
						  type="button"
						>
						  <span class="text-sm font-medium"> View Website </span>

						  <svg
							xmlns="http://www.w3.org/2000/svg"
							class="size-4"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
							stroke-width="2"
						  >
							<path
							  stroke-linecap="round"
							  stroke-linejoin="round"
							  d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
							/>
						  </svg>
						</button>

						<button
						  class="inline-block rounded bg-indigo-600 px-5 py-3 text-sm font-medium text-white transition hover:bg-indigo-700 focus:outline-none focus:ring"
						  type="button"
						>
						  Create Post
						</button>
					  </div>
					</div>
				  </div>
				</header>
				<div class="p-10">
					%s
				<div>
			</body>
		`

		slug := r.PathValue("slug")
		if slug == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		result, err := lookupBySlug(dbc, slug)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		html = fmt.Sprintf(html, result)
		w.Write([]byte(html))
	})

	mux.HandleFunc("POST /search", func(w http.ResponseWriter, r *http.Request) {
		query := r.FormValue("query")
		if query == "" {
			cnt, err := listArticles(dbc)
			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Write([]byte(cnt))
			return
		}

		v, err := c.Get(query)
		if err == nil {
			fmt.Println("From cache")
			w.Write([]byte(*v))
			return
		}

		cnt, err := searchForArticles(dbc, query)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := c.Set(query, cnt); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(cnt))
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}
