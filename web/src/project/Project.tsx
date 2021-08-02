import React, { Fragment, ChangeEvent, useState } from "react";
import { useParams } from "react-router-dom";
import { fileURLToPath } from "url";

import { useProject, ProjectFilesList } from "./useProject";
import { useUploadProjectFile } from "./useUploadProjectFile";

interface RouteParams {
    id: string
};

export function Project () {
    let { id } = useParams<RouteParams>();

    const { data, loading, error } = useProject(parseInt(id, 10));

    if (loading) {
        return (
            <div>
                <p>Loading project...</p>
            </div>
        );
    }

    if (error) {
        return (
            <div>
                <h1>Error</h1>
                <p>Can't get project: {error}</p>
            </div>
        );
    }

    switch (data?.project.__typename) {
        case "Project":
            const { project } = data;
            return (
                <Fragment>
                    <h2>{project.title}</h2>
                    <Files list={project.files.list}/>
                    <Upload projectId={project.id} />
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
                    <p>{data?.project.__typename}</p>
                </div>
            );
    }
}

export default Project;

function Files({ list }: { list: ProjectFilesList }) {
    switch (list.__typename) {
        case "ProjectFilesList":
            return (
                <ul>
                    {list.items.map(file => <li key={file.url}>
                        <a href={file.url}>{file.url}</a>
                    </li>)}
                </ul>
            )
        default:
            return <div>n/a</div>
    }
}

function Upload({ projectId }: {projectId: number}) {
    const [ file, setFile ] = useState<File | null>(null)
    const [ upload, { loading, error } ] = useUploadProjectFile()

    const onChange = (event: ChangeEvent<HTMLInputElement>) => {
        if (event.target.files) {
            setFile(event.target.files[0])
        }
    }

    const handleSubmit = (event: React.FormEvent) => {
        console.log({ file, loading, error });
        if (file && !loading) {
            upload({ projectId, file })
        }

        event.preventDefault();
    }

    return (
        <div>
            {error ? <p>{error.message}</p> : null}
            {loading ? null : <input type="file" name="file" onChange={onChange}/>}
            {loading ? <p>Upload file...</p> : <button disabled={!file} onClick={handleSubmit}>Upload</button>}
        </div>
    )
}