import React, { Fragment } from "react"
import { useParams } from "react-router-dom"

import { Clipboard } from "./Clipboard/Clipboard"

import useSpecScreen from "./useSpecScreen"

type RouteParams = {
    projectId: string
}

export default function Spec() {
    let { projectId } = useParams<RouteParams>();

    const { data, loading, error } = useSpecScreen(projectId);

    if (loading) {
        return (
            <div>
                <p>Loading spec...</p>
            </div>
        );
    }

    if (error) {
        return (
            <div>
                <h1>Error</h1>
                <p>Can't get spec: {error}</p>
            </div>
        );
    }

    const handleOnAdd = (...args: []) => {
        console.log("add", { ...args })
    }

    switch (data?.screen.screen.project.__typename) {
        case "Project":
            const { project } = data?.screen.screen;
            return (
                <Fragment>
                    <h2>{project.title}</h2>
                    <Clipboard onAdd={handleOnAdd}/>
                </Fragment>
            );
        case "NotFound":
            return (
                <div>
                    <h1>Error</h1>
                    <p>project not found</p>
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
                    <p>{data?.screen.__typename}</p>
                </div>
            );
    }
}