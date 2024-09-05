CREATE TABLE IF NOT EXISTS articles(
  "slug" TEXT PRIMARY KEY NOT NULL,
  "title" TEXT UNIQUE NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" INTEGER NOT NULL,
  "updated_at" INTEGER NOT NULL
);

CREATE VIRTUAL TABLE IF NOT EXISTS articles_fts USING fts5 (
  slug,
  title,
  content,
  created_at,
  updated_at,
  content='articles'
);

INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('introduction-to-golang', 'Introduction to Golang', 'Golang is a statically typed, compiled programming language...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('getting-started-with-elixir', 'Getting Started with Elixir', 'Elixir is a dynamic, functional language designed for building scalable and maintainable applications...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('websockets-in-real-time-apps', 'WebSockets in Real-Time Apps', 'WebSockets provide a full-duplex communication channel over a single TCP connection...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('building-multiplayer-games', 'Building Multiplayer Games', 'When building multiplayer games, managing client inputs and server updates efficiently is crucial...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('golang-raylib-integration', 'Golang and Raylib Integration', 'Using Golang with Raylib provides a simple way to create 2D and 3D games...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('fiat-payments-in-crypto', 'Fiat Payments in the Crypto World', 'Handling fiat payments in the crypto world involves dealing with deposits, withdrawals, and more...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('implementing-batching-in-websockets', 'Implementing Batching in WebSockets', 'Batching is an essential technique for optimizing network performance in real-time applications...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('optimizing-golang-loops', 'Optimizing Golang Loops', 'Optimizing loops in Golang can lead to significant performance improvements...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('working-with-sqlite-in-golang', 'Working with SQLite in Golang', 'SQLite is a popular choice for small to medium-sized applications that need a lightweight database...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('setting-up-ci-with-github-actions', 'Setting Up CI with GitHub Actions', 'Continuous Integration (CI) is a critical practice in modern software development...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('building-an-onboarding-app', 'Building an Onboarding App', 'An onboarding app can streamline the process of bringing new employees into a company...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('error-handling-in-golang', 'Error Handling in Golang', 'Error handling is a crucial aspect of writing robust Go applications...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('working-with-webhooks-in-go', 'Working with Webhooks in Go', 'Webhooks are a powerful way to integrate services and build event-driven applications...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('server-driven-ui-design', 'Server-Driven UI Design', 'Server-driven UI design allows developers to manage the user interface from the server...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('building-a-budget-app', 'Building a Budget App', 'Building a budget app requires careful attention to data parsing and financial calculations...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('high-value-dating-material', 'Becoming High-Value Dating Material', 'Self-improvement is key to becoming high-value dating material...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('customizing-neovim-with-lua', 'Customizing Neovim with Lua', 'Neovim provides a powerful API for customization using Lua...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('building-2d-games-with-raylib', 'Building 2D Games with Raylib', 'Raylib is a simple and easy-to-use library for creating 2D and 3D games...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('handling-flags-with-math-big-int', 'Handling Flags with math/big.Int in Go', 'Handling large sets of flags in Go can be efficiently done using the math/big.Int package...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('introduction-to-box2d', 'Introduction to Box2D', 'Box2D is a popular physics engine used in 2D games...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('server-context-over-http-in-go', 'Server Context over HTTP in Go', 'Managing context across HTTP requests in Go is crucial for building scalable microservices...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('transcripts-in-golang', 'Generating Transcripts in Golang', 'Generating transcripts in Golang can be done using various techniques...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('balancing-work-and-personal-life', 'Balancing Work and Personal Life', 'Achieving a balance between work and personal life is essential for long-term happiness...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('advanced-golang-techniques', 'Advanced Golang Techniques', 'Exploring advanced techniques in Golang can help you become a more proficient developer...', strftime('%s', 'now'), strftime('%s', 'now'));
INSERT OR IGNORE INTO "articles" ("slug", "title", "content", "created_at", "updated_at") VALUES ('sophisticated-metalhead-outfits', 'Sophisticated Metalhead Outfits', 'Finding outfits that show your metal roots while being sophisticated is a unique challenge...', strftime('%s', 'now'), strftime('%s', 'now'));

INSERT INTO articles_fts (slug, title, content, created_at, updated_at) SELECT slug, title, content, created_at, updated_at FROM articles;
