CREATE SCHEMA apartomat;

CREATE TABLE apartomat.users (
    id SERIAL PRIMARY KEY,
    email text NOT NULL,
    full_name text NOT NULL,
    is_active boolean NOT NULL,
    use_gravatar boolean NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    modified_at timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT users_email_ukey UNIQUE (email)
);

CREATE TABLE apartomat.workspaces (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    is_active boolean NOT NULL,
    user_id INT NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    modified_at timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT workspaces_user_id_fkey FOREIGN KEY (user_id) REFERENCES apartomat.users ON DELETE CASCADE
);

CREATE TABLE apartomat.workspace_users (
    id SERIAL PRIMARY KEY,
    workspace_id integer NOT NULL,
    user_id integer NOT NULL,
    role text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    modified_at timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT workspace_users_workspace_id_fkey FOREIGN KEY (workspace_id) REFERENCES apartomat.workspaces ON DELETE CASCADE,
    CONSTRAINT workspace_users_user_id_fkey FOREIGN KEY (user_id) REFERENCES apartomat.users ON DELETE CASCADE
);

CREATE TABLE apartomat.projects (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    is_active boolean NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    modified_at timestamp with time zone NOT NULL DEFAULT now(),
    workspace_id INT NOT NULL,
    CONSTRAINT projects_workspace_id_fkey FOREIGN KEY (workspace_id) REFERENCES apartomat.workspaces ON DELETE CASCADE
);

CREATE TABLE apartomat.project_files (
    id SERIAL PRIMARY KEY,
    project_id integer NOT NULL,
    name text NOT NULL,
    type text NOT NULL,
    url text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    modified_at timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT project_files_project_id_fkey FOREIGN KEY (project_id) REFERENCES apartomat.projects ON DELETE CASCADE,
    CONSTRAINT project_files_ukey UNIQUE (project_id, name)
);