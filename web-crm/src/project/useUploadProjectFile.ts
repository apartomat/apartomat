import { FetchResult, useApolloClient } from "@apollo/client"
import { useUploadProjectFileMutation, UploadProjectFileMutation, UploadProjectFileMutationResult, ProjectFile, Forbidden, AlreadyExists, ServerError, ProjectFileType } from "../api/types.d"

export type UploadFn = ({projectId, type, file}: { projectId: number, type: ProjectFileType, file: File}) => Promise<FetchResult<UploadProjectFileMutation>>

export { ProjectFileType } from "../api/types.d";

export enum State {
    INITIAL = "INITIAL",
    UPLOADING = "UPLOADING",
    DONE = "DONE",
    FAILED = "FAILED"
}

export type Result = 
    Initial |
    Uploading |
    Done |
    Failed

type Initial = {
    state: State.INITIAL
}

type Uploading = {
    state: State.UPLOADING
}

type Done = {
    state: State.DONE
    file: Pick<ProjectFile, 'id' | 'url'>
}

type Failed = {
    state: State.FAILED
    error: Error | (
        { __typename: 'Forbidden' }
        & Pick<Forbidden, 'message'>
      ) | (
        { __typename: 'AlreadyExists' }
        & Pick<AlreadyExists, 'message'>
      ) | (
        { __typename: 'ServerError' }
        & Pick<ServerError, 'message'>
      )
}

export function useUploadProjectFile(): [
    UploadFn,
    UploadProjectFileMutationResult,
    Result
] {
    const client = useApolloClient(); 
    const [upload, result] = useUploadProjectFileMutation({ client, errorPolicy: 'all' });

    const state = new Proxy<Result>({ state: State.INITIAL }, {
        get: (target: Result, name: string) => {
            switch (name) {
                case "state":
                    if (!result.called) {
                        return State.INITIAL
                    } else if (result.called && result.loading) {
                        return State.UPLOADING
                    } else if (result.called && !result.loading &&
                        (result.error || (result.data && result.data?.uploadProjectFile.__typename !== "ProjectFileUploaded"))
                    ) {
                        return State.FAILED
                    } else if (result.called && !result.loading && result.data && result.data?.uploadProjectFile.__typename === "ProjectFileUploaded") {
                        return State.DONE
                    }

                    return State.INITIAL

                case "file":
                    if (result.data?.uploadProjectFile.__typename === "ProjectFileUploaded") {
                        return result.data.uploadProjectFile.file
                    }

                    return undefined

                case "error":
                    if (result.data && result.data?.uploadProjectFile.__typename !== "ProjectFileUploaded") {
                        return result.data.uploadProjectFile
                    }
                    
                    return result.error

                default:
                    return undefined
            }
        }
    })

    return [({ projectId, type, file }) => upload({ variables: { input: { projectId, type, file } } }), result, state as Result];
}

export default useUploadProjectFile;