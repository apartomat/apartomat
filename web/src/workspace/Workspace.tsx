import React, { Fragment } from "react";
import { useParams } from "react-router-dom";

import UserAvatar from "./UserAvatar";
import { useAuthContext } from "../common/context/auth/useAuthContext";

import Avatar from "../common/ui/Avatar";
import AvatarGroup from "../common/ui/AvatarGroup";

import { useWorkspace, WorkspaceUsersResult, WorkspaceProjectsListResult } from "./useWorkspace";
import { useCreateProject } from "./useCreateProject";
import { useState } from "react";

interface RouteParams {
    id: string
};

export function Workspace () {
    const { user } = useAuthContext();
    let { id } = useParams<RouteParams>();
    const { data, loading, error } = useWorkspace(parseInt(id, 10));

    if (loading) {
        return (
            <div>
                <p>Loading workspace...</p>
            </div>
        );
    }

    if (error) {
        return (
            <div>
                <h1>Error</h1>
                <p>Can't get workspace: {error}</p>
            </div>
        );
    }

    switch (data?.workspace.__typename) {
        case "Workspace":
            const { workspace } = data;
            return (
                <Fragment>
                    <nav className="navbar">
                        <div className="navbar__logo">apartomat</div>
                        <div className="navbar__user">
                            <UserAvatar user={user} className="header-user" />
                        </div>
                    </nav>
                    
                    <h2>{workspace.name}</h2>
                    <WorkspaceUsers users={workspace.users} />
                    <Projects projects={workspace.projects.list} />
                    <CreateProject workspaceId={workspace.id}/>
                </Fragment>
            );
        case "NotFound":
            return (
                <div>
                    <h1>Error</h1>
                    <p>Workspace not found</p>
                </div>
            );
        case "Forbidden":
            return (
                <div>
                    <h1>Error</h1>
                    <p>Access is denied</p>
                </div>
            );
        default:
            return (
                <div>
                    <h1>Error</h1>
                    <p>{data?.workspace.__typename}</p>
                </div>
            );
    }
}

function WorkspaceUsers({ users }: {users: WorkspaceUsersResult}) {
    switch (users.__typename) {
        case "WorkspaceUsers":
            return (
                <AvatarGroup>
                    {users.items.map(user => <Avatar key={user.id} src={user.profile.gravatar.url} alt={user.profile.email}/>)}
                </AvatarGroup>
            )
        default:
            return <div>n/a</div>
    }
}

function Projects({ projects }: { projects: WorkspaceProjectsListResult }) {
    switch (projects.__typename) {
        case "WorkspaceProjectsList":
            return (
                <ul>
                    {projects.items.map(project => <li key={project.id}>
                        <a href={`/p/${project.id}`}>{project.name}</a>
                    </li>)}
                </ul>
            )
        default:
            return <div>n/a</div>
    }
}

function CreateProject({ workspaceId }: { workspaceId: number }) {
    const [ title, setTitle ] = useState("")
    const [ create, { loading, error } ] = useCreateProject()

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault()
        create({ workspaceId, title: title })
    }

    const handleChangeTitle = (event: React.FormEvent<HTMLInputElement>) => {
        setTitle(event.currentTarget.value)
    }

    if (loading) {
        return (
            <div>Createing project...</div>
        )
    }

    return (
        <form onSubmit={handleSubmit}>
            {error ? <p>{error.message}</p> : null}
            <input type="text" onChange={handleChangeTitle} />
            <input type="submit"/>
        </form>
    )
}

export default Workspace;