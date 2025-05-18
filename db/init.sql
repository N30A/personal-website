CREATE TABLE technologies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE,
    body TEXT NOT NULL,
    image_name TEXT NULL,
    project_url TEXT NULL
);

CREATE TABLE project_technologies (
    project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
    technology_id INTEGER REFERENCES technologies(id) ON DELETE CASCADE,
    PRIMARY KEY (project_id, technology_id)
);

INSERT INTO technologies (name) VALUES
  -- Språk
  ('Go'), ('SQL'), ('HTML'), ('CSS'), ('C#'), ('Python'), ('JavaScript'),
  ('Java'), ('HTMX'), ('TypeScript'), ('Rust'), ('Kotlin'), ('Swift'),
  ('PHP'), ('Ruby'), ('Perl'), ('Bash'), ('Scala'), ('Flutter'),
  ('Dart'), ('Elixir'), ('Haskell'), ('R'), ('Assembly'),
  ('C'), ('C++'), ('Objective-C'), ('F#'), ('Groovy'), ('Lua'),

  -- Ramverk & bibliotek
  ('.NET'), ('ASP.NET'), ('React'), ('Vue'), ('Angular'), ('Node.js'),
  ('Django'), ('Flask'), ('Spring'), ('Express'), ('Rails'), ('Laravel'),
  ('Gin'), ('Fiber'), ('Echo'), ('sqlx'), ('GORM'), ('Dapper'),
  ('Entity Framework'), ('Pandas'), ('NumPy'), ('TensorFlow'),
  ('Keras'), ('Bootstrap'), ('Tailwind CSS'), ('jQuery'),

  -- Databaser
  ('PostgreSQL'), ('MySQL'), ('SQLite'), ('MongoDB'),
  ('Redis'), ('MariaDB'), ('SQL Server'), ('Oracle'),
  ('Cassandra'), ('Elasticsearch'), ('Firebase'), ('DynamoDB'),

  -- Verktyg / övrigt
  ('Docker'), ('Kubernetes'), ('Git'), ('Linux'), ('Nginx'), ('GraphQL'),
  ('gRPC'), ('WebAssembly'), ('Ansible'), ('Terraform'), ('Jenkins'),
  ('CircleCI'), ('Prometheus'), ('Grafana');
