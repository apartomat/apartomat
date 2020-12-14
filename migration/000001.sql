CREATE SCHEMA apartomat;

CREATE TABLE apartomat.users (
    id SERIAL PRIMARY KEY,
    email text NOT NULL,
    full_name text NOT NULL,
    is_active boolean NOT NULL,
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