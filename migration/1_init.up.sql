create schema apartomat;

create table apartomat.users (
    id char(21) primary key,
    email text not null,
    full_name text not null,
    is_active boolean not null,
    use_gravatar boolean not null,
    created_at timestamp with time zone not null default now(),
    modified_at timestamp with time zone not null default now(),
    constraint users_email_ukey unique (email)
);

create table apartomat.workspaces (
    id char(21) primary key,
    name text not null,
    is_active boolean not null,
    created_at timestamp with time zone not null default now(),
    modified_at timestamp with time zone not null default now(),
    user_id char(21) not null,
    constraint workspaces_user_id_fkey foreign key (user_id) references apartomat.users on delete cascade
);

create table apartomat.workspace_users (
    id char(21) primary key,
    user_id char(21) not null,
    role text not null,
    created_at timestamp with time zone not null default now(),
    modified_at timestamp with time zone not null default now(),
    workspace_id char(21) not null,
    constraint workspace_users_workspace_id_fkey foreign key (workspace_id) references apartomat.workspaces on delete cascade,
    constraint workspace_users_user_id_fkey foreign key (user_id) references apartomat.users on delete cascade
);

create table apartomat.projects (
    id char(21) primary key,
    name text not null,
    status text not null,
    start_at timestamp with time zone,
    end_at timestamp with time zone,
    created_at timestamp with time zone not null default now(),
    modified_at timestamp with time zone not null default now(),
    workspace_id char(21) not null,
    constraint projects_workspace_id_fkey foreign key (workspace_id) references apartomat.workspaces on delete cascade
);

create table apartomat.project_files (
    id char(21) primary key,
    name text not null,
    type text not null,
    mime_type text not null,
    url text not null,
    created_at timestamp with time zone not null default now(),
    modified_at timestamp with time zone not null default now(),
    project_id char(21) not null,
    constraint project_files_project_id_fkey foreign key (project_id) references apartomat.projects on delete cascade
);

create table apartomat.contacts (
    id char(21) primary key,
    full_name text not null,
    photo text not null default '',
    details jsonb,
    created_at timestamp with time zone not null default now(),
    modified_at timestamp with time zone not null default now(),
    project_id char(21) not null,
    constraint contacts_project_id_fkey foreign key (project_id) references apartomat.projects on delete cascade
);

create table apartomat.houses (
    id char(21) primary key,
    city text not null,
    address text not null,
    housing_complex text not null,
    created_at timestamp with time zone not null default now(),
    modified_at timestamp with time zone not null default now(),
    project_id char(21) not null,
    constraint houses_project_id_fkey foreign key (project_id) references apartomat.projects on delete cascade
);

create table apartomat.rooms (
     id char(21) primary key,
     name text not null,
     square real,
     level integer,
     created_at timestamp with time zone not null default now(),
     modified_at timestamp with time zone not null default now(),
     house_id char(21) not null,
     constraint rooms_house_id_fkey foreign key (house_id) references apartomat.houses on delete cascade
);

create table apartomat.visualizations (
     id char(21) primary key,
     name text not null,
     description text not null,
     version integer not null default 0,
     created_at timestamp with time zone not null default now(),
     modified_at timestamp with time zone not null default now(),
     project_id char(21) not null,
     project_file_id char(21) not null,
     room_id char(21),
     constraint visualizations_project_id_fkey foreign key (project_id) references apartomat.projects on delete cascade,
     constraint visualizations_project_file_id_fkey foreign key (project_file_id) references apartomat.project_files on delete cascade,
     constraint visualizations_room_id_fkey foreign key (room_id) references apartomat.rooms
);